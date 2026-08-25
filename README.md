# golomt-api-go

Go client for Golomt Bank's online services — eCommerce, SocialPay, SocialPay Mini App
and Open Banking (Corporate Gateway).

> **Internal library.** This is a private module maintained by techpartners.asia for use
> in our own projects. It is not published to a public proxy and is not intended for
> outside consumers. Access requires membership in the `techpartners-asia` GitHub org.

## Requirements

- **Go 1.27.0 or newer.** The `go` directive in `go.mod` pins a patch version, so older
  toolchains will refuse to build the module.
- Read access to `github.com/techpartners-asia/golomt-api-go`.
- Credentials issued by Golomt Bank for whichever service you are integrating.

## Installation

Because the repository is private, the Go module proxy cannot fetch it. Configure your
environment once:

```bash
# Bypass the public proxy and checksum database for our org
go env -w GOPRIVATE=github.com/techpartners-asia/*

# Let the toolchain authenticate over SSH instead of HTTPS
git config --global url."git@github.com:".insteadOf "https://github.com/"
```

Then add the dependency as normal:

```bash
go get github.com/techpartners-asia/golomt-api-go@latest
```

If you use HTTPS with a personal access token instead of SSH, put the token in
`~/.netrc` rather than in the URL rewrite:

```
machine github.com login <your-github-username> password <your-personal-access-token>
```

### CI

CI runners need the same access. Grant the job a deploy key or a token with `repo`
scope, export `GOPRIVATE=github.com/techpartners-asia/*`, and configure the rewrite
before `go mod download`.

## Packages

| Package | Constructor | Methods | Purpose |
| --- | --- | --- | --- |
| `ecommerce` | `New(baseUrl, secret, bearerToken string)` | 11 | Card invoices, tokenized payments, settlement, refunds |
| `socialpay` | `New(terminal, secret, endpoint string)` | 6 | QR and phone-number invoices, payment cancellation, settlement |
| `mini_app` | `New(baseUrl, clientId, base64PublicKey string)` | 2 | SocialPay Mini App user info and push notifications |
| `openbank` | `New(input model.OpenbankInput)` | 37 | Corporate Gateway — accounts, transfers, cards, reference data |
| `utils` | — | — | HMAC, RSA, TOTP and JSON value helpers shared by the above |

Each service package exposes an interface (`GolomtEcommerce`, `SocialPay`,
`SocialPayMiniApp`, `Openbank`) and an unexported implementation, so consumers can mock
the interface in tests.

## Usage

### eCommerce

```go
import "github.com/techpartners-asia/golomt-api-go/ecommerce"

client := ecommerce.New(baseUrl, secret, bearerToken)

invoice, err := client.CreateInvoice(ecommerce.CreateInvoiceInput{ /* ... */ })
if err != nil {
    return err
}

// Redirect the customer's browser to the card entry page
url := client.GetInvoiceUrl(ecommerce.GetInvoiceInput{ /* ... */ })
```

Also available: `Inquiry`, `CreateToken`, `CheckToken`, `PayTokenPayment`,
`CheckTokenPayment`, `GetConfirmationUrl`, `GetSettlementDetails`,
`ParsePushNotificationResponse`, and `Refund` (test environment only).

### SocialPay

```go
import "github.com/techpartners-asia/golomt-api-go/socialpay"

client := socialpay.New(terminal, secret, endpoint)

resp, err := client.CreateInvoiceQR(socialpay.InvoiceInput{ /* ... */ })
```

Also available: `CreateInvoicePhone`, `CancelInvoice`, `CheckInvoice`, `CancelPayment`,
`Settlement`.

### SocialPay Mini App

```go
import mini_app "github.com/techpartners-asia/golomt-api-go/mini_app"

client := mini_app.New(baseUrl, clientId, base64PublicKey)

info, err := client.GetUserInfo(token)
```

### Open Banking (Corporate Gateway)

```go
import (
    "github.com/techpartners-asia/golomt-api-go/openbank"
    "github.com/techpartners-asia/golomt-api-go/openbank/model"
)

client := openbank.New(model.OpenbankInput{
    OrganizationName: "...",
    Username:         "...",
    Password:         "...",
    IvKey:            "...",
    SessionKey:       "...",
    Url:              "...",
    RegisterNo:       "...",
    ClientID:         "...",
    XGolomtKey:       "...",
})

balance, err := client.AccountBalcInq(model.AccountBalcInqReq{ /* ... */ })
```

Authentication is handled internally: the client logs in on first use and refreshes the
token when it expires, so callers do not manage sessions.

The interface follows the numbering in Golomt's Corporate Gateway specification —
section 4 is access and OTP, 5 accounts, 6 transfers, 7 reference data, 8 cards.

## Error handling

Every method returns a Go `error`; the client does not panic on gateway failures.

- HTTP 4xx/5xx responses return an error carrying the status code and the raw response
  body, e.g. `golomt socialpay: http 502: <body>`.
- Empty or non-JSON success bodies return an explicit error rather than a nil
  dereference.
- Application-level failures (`header.code != 200`) return the gateway's own error
  description.

Do not match on error strings — the wording is not part of the contract and has changed
between releases.

## Configuration and secrets

All credentials are passed to the constructors at runtime. Nothing is read from the
environment and no defaults are baked into the library. Keep terminal IDs, secrets,
bearer tokens, session keys and RSA keys in your own secret store; never commit them.

## Development

```bash
go build ./...
go vet ./...
gofmt -l .
```

Vulnerability scanning, both of which should report clean before a release:

```bash
govulncheck ./...
grype dir:.
```

> **Never commit build artifacts.** A Delve debug binary was committed in v0.0.15 and
> reported 29 vulnerabilities (1 Critical, 13 High) against its embedded dependency
> graph — none of which existed in the source tree. `.gitignore` covers `__debug_bin*`,
> but that only applies to files that are not already tracked, so a binary committed
> once will keep being scanned until it is explicitly removed.

## Versioning

Tagged `v0.0.x`. The module is below `v1`, so no compatibility is promised between
tags — read the release notes before upgrading. Recent releases have changed the
minimum Go version, error message formats and helper semantics.

## License

The repository currently carries an MIT `LICENSE` file from its initial commit, which
does not reflect how this library is actually distributed. If it is meant to be
proprietary, that file should be replaced with an internal-use notice — a decision for
whoever owns licensing at techpartners.asia.
