package openbank

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/techpartners-asia/golomt-api-go/openbank/model"
)

const testKey, testIV = "0123456789abcdef", "fedcba9876543210"

// fakeBank answers login, then ACCTBALINQ per SPEC 5: an OAuth grant while
// state is empty, the balance once the grant is sent back.
func fakeBank(t *testing.T, grant string, queries *[]string) *httptest.Server {
	t.Helper()
	crypt := openbank{sessionKey: testKey, ivKey: testIV}
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/v1/auth/login" {
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{"token":"T","refreshToken":"R","tokenType":"Bearer","expiresIn":300}`))
			return
		}
		*queries = append(*queries, r.URL.RawQuery)
		reply := grant
		if r.URL.Query().Get("state") == "S1" && r.URL.Query().Get("scope") == "SC1" && r.URL.Query().Get("client_id") == "C1" {
			reply = `{"accountId":"ACC","accountName":"Magic Tech","currency":"MNT","balanceLL":[{"type":"AVAIL","amount":{"value":100000000,"currency":"MNT"}}]}`
		}
		enc, err := crypt.EncryptAESCBC(reply)
		if err != nil {
			t.Fatal(err)
		}
		w.Write([]byte(enc))
	}))
}

func newTestClient(url string) Openbank {
	return New(model.OpenbankInput{Username: "u", Password: "p", IvKey: testIV, SessionKey: testKey, Url: url, ClientID: "CONFIGURED"})
}

func TestAccountBalcInqFollowsOAuthGrant(t *testing.T) {
	var queries []string
	srv := fakeBank(t, `{"clientId":"C1","responseType":"code","state":"S1","scope":"SC1"}`, &queries)
	defer srv.Close()

	got, err := newTestClient(srv.URL).AccountBalcInq(model.AccountBalcInqReq{AccountID: "ACC", RegisterNo: "REG"})
	if err != nil {
		t.Fatal(err)
	}
	if got.AccountName != "Magic Tech" || len(got.BalanceLL) != 1 || got.BalanceLL[0].Amount.Value != 100000000 {
		t.Fatalf("got %+v", got)
	}
	if len(queries) != 2 {
		t.Fatalf("want 2 ACCTBALINQ calls, got %d: %q", len(queries), queries)
	}
	if queries[0] != "client_id=&scope=&state=" {
		t.Errorf("first call must send empty client_id/state/scope, got %q", queries[0])
	}
	if queries[1] != "client_id=C1&scope=SC1&state=S1" {
		t.Errorf("second call must send the grant, got %q", queries[1])
	}
}

func TestAccountBalcInqReturnsConsentURL(t *testing.T) {
	var queries []string
	srv := fakeBank(t, `{"clientId":"C1","responseType":"code","url":"https://bank/consent"}`, &queries)
	defer srv.Close()

	_, err := newTestClient(srv.URL).AccountBalcInq(model.AccountBalcInqReq{AccountID: "ACC", RegisterNo: "REG"})
	if err == nil || !strings.Contains(err.Error(), "https://bank/consent") {
		t.Fatalf("err = %v, want the consent URL", err)
	}
	if len(queries) != 1 {
		t.Fatalf("must not retry without a grant, got %d calls", len(queries))
	}
}
