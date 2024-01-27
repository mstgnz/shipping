package cargo

import (
	"context"
)

type ServiceType string

const (
	SOAP ServiceType = "SOAP"
	REST ServiceType = "REST"
)

type Mode string

const (
	PRODUCTION  Mode = "PRODUCTION"
	DEVELOPMENT Mode = "DEVELOPMENT"
)

type Shipper interface {
	GetContext() context.Context
	SetContext(ctx context.Context)

	GetServiceType() ServiceType
	SetServiceType(serviceType ServiceType)

	GetMode() Mode
	SetMode(mode Mode)

	IsDomestic() bool
	SetDomestic(isDomestic bool)

	GetEndpoints() []Endpoint
	GetEndpoint(title string) Endpoint
	AddEndpoint(title, production, development string)
	DelEndpoint(title string) []Endpoint

	GetCredentials() []Credential
	GetCredential(title string) Credential
	AddCredential(title, username, password string)
	DelCredential(title string) []Credential

	GetCurrentEndpointAndCredential() Current

	CreateCargo(data map[string]any) (map[string]any, error)

	WhereIsTheCargo(tracking string) (map[string]any, error)

	CancelCargo(tracking string) (map[string]any, error)
}
