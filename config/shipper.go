package cargo

import (
	"context"
)

type Shipper interface {
	GetContext() context.Context
	SetContext(ctx context.Context)

	GetServiceType() string
	SetServiceType(serviceType string)

	GetEndpoints() []Endpoint
	GetEndpoint(title string) Endpoint
	AddEndpoint(title, production, development string)
	DelEndpoint(title string) []Endpoint

	GetMode() string
	SetMode(mode string)

	GetDomestic() bool
	SetDomestic(isDomestic bool)

	GetCurrentCredential() string
	SetCurrentCredential(currentCredential string)

	GetCurrentEndpoint() string
	SetCurrentEndpoint(currentEndpoint string)

	GetCredentials() []Credential
	GetCredential(title string) Credential
	AddCredential(title, username, password string)
	DelCredential(title string) []Credential

	CreateCargo(data map[string]any) (map[string]any, error)

	WhereIsTheCargo(tracking string) (map[string]any, error)

	CancelCargo(tracking string) (map[string]any, error)
}
