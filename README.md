# avmsgraph-sdk-go

**`avmsgraph-sdk-go`** is an opinionated Go SDK for a subset of the [Microsoft Graph API](https://learn.microsoft.com/en-us/graph/overview)'s [Golang SDK](https://github.com/microsoftgraph/msgraph-sdk-go) generated using [Kiota](https://github.com/microsoft/kiota) and maintained by [Aurva](https://aurva.io).

This SDK simplifies integration with Microsoft Graph for common directory and role management use cases in Go applications and makes the application significantly lighter.

---

## ✨ Features

- Built using [Kiota](https://github.com/microsoft/kiota) for consistency and performance
- Targets selected Microsoft Graph paths:
  - `/users`
  - `/applications`
  - `/groups`
  - `/directoryObjects/getByIds`
  - `/directoryRoles`
  - `/directoryRoleTemplates`
  - `/roleManagement/directory/roleDefinitions`
- Backing store support for advanced state management

---

## 🛠 Installation

Install using `go get`:

```bash
go get github.com/aurva-io/aurva/avmsgraph-sdk-go
```

---

## Generating the SDK

```sh
./kiota generate --language Go --openapi ./openapi.yaml --backing-store --output ./avmsgraph-sdk-go --namespace-name "github.com/aurva-io/avmsgraph-sdk-go" --include-path "/users"  --include-path "/applications" --include-path "/directoryObjects/getByIds"  --include-path "/groups"  --include-path "/directoryRoles" --include-path "/directoryRoleTemplates"  --include-path "/roleManagement/directory/roleDefinitions"
```


## Connection 

```go

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	avmsgraphsdk "github.com/aurva-io/avmsgraph-sdk-go"
	azurekiota "github.com/microsoft/kiota-authentication-azure-go"
	bundle "github.com/microsoft/kiota-bundle-go"

func main() {
	cred, err := azidentity.NewClientSecretCredential("a.tenantID", "a.clientID", "a.clientSecret", nil)
	if err != nil {
		return "", fmt.Errorf("failed to get Azure credentials: %w", err)
	}

	authProvider, err := azurekiota.NewAzureIdentityAuthenticationProvider(cred)

	if err != nil {
		fmt.Printf("Error creating auth provider: %v\n", err)
	}

	adapter, err := bundle.NewDefaultRequestAdapter(authProvider)

	if err != nil {
		fmt.Printf("Error creating request adapter: %v\n", err)
	}

	client := avmsgraphsdk.NewApiClient(adapter, nil)
}

```