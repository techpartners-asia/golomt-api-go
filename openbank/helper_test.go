package openbank

import (
	"encoding/json"
	"errors"
	"strings"
	"testing"

	"github.com/techpartners-asia/golomt-api-go/openbank/model"
)

func TestParseEncryptedResponseAcceptsPlainJSON(t *testing.T) {
	decrypt := func(string) (string, error) { return "", errors.New("must not be called") }

	got, err := parseEncryptedResponse[*model.ServiceListResp]([]byte(` {"clientId":"c","state":"s","scope":"x"}`), decrypt)
	if err != nil || got == nil || got.State != "s" || got.Scope != "x" {
		t.Fatalf("got %+v, err %v", got, err)
	}
}

func TestParseEncryptedResponseDecryptsBase64(t *testing.T) {
	decrypt := func(s string) (string, error) { return `{"state":"s"}`, nil }

	got, err := parseEncryptedResponse[*model.ServiceListResp]([]byte("QUJD"), decrypt)
	if err != nil || got.State != "s" {
		t.Fatalf("got %+v, err %v", got, err)
	}
}

func TestParseEncryptedResponseErrorShowsBody(t *testing.T) {
	decrypt := func(string) (string, error) { return "", errors.New("illegal base64 data at input byte 0") }

	_, err := parseEncryptedResponse[*model.ServiceListResp]([]byte("Forbidden"), decrypt)
	if err == nil || !strings.Contains(err.Error(), "Forbidden") {
		t.Fatalf("err = %v, want it to include the body", err)
	}
}

func TestServiceListReqIsFlat(t *testing.T) {
	b, err := json.Marshal(model.ServiceListReq{RegisterNo: "1234", Services: []string{"ACCTBALINQ"}, Code: "ACCTBALINQ"})
	if err != nil {
		t.Fatal(err)
	}
	want := `{"registerNo":"1234","services":["ACCTBALINQ"],"code":"ACCTBALINQ"}`
	if string(b) != want {
		t.Fatalf("got %s, want %s", b, want)
	}
}
