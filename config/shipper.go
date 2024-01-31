package cargo

import (
	"context"
)

// ShippingData parameter placeholder
type ShippingData interface{}

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

// Shipper contains a common description of all cargo.
type Shipper interface {
	GetContext() context.Context
	SetContext(ctx context.Context)

	GetServiceType() ServiceType
	SetServiceType(serviceType ServiceType)

	GetMode() Mode
	SetMode(mode Mode)

	IsDomestic() bool
	SetDomestic(isDomestic bool)

	GetEndpoints() []*Endpoint
	GetEndpoint(title string) *Endpoint
	AddEndpoint(title, production, development string) *Endpoint
	DelEndpoint(title string) []*Endpoint

	GetCredentials() []*Credential
	GetCredential(title string) *Credential
	AddCredential(title, username, password string) *Credential
	DelCredential(title string) []*Credential

	GetCurrent() Current

	CreateCargo(data ShippingData) (*Response, error)

	WhereIsTheCargo(data ShippingData) (*Response, error)

	CancelCargo(data ShippingData) (*Response, error)
}
