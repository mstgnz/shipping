package cargo

import (
	"context"
)

type Shipper interface {
	GetContext() context.Context
	SetContext(ctx context.Context)

	GetEndpoint() Endpoint
	SetEndpoint(production, development string)

	GetMode() string
	SetMode(mode string)

	GetUsername() string
	SetUsername(username string)

	GetPassword() string
	SetPassword(password string)

	CreateCargo() error

	WhereIsTheCargo() error

	CancelCargo() error
}

type Endpoint struct {
	Production  string
	Development string
}

type Cargo struct {
	Ctx      context.Context
	Endpoint Endpoint
	Mode     string // Production or Development
	Username string
	Password string
}

// GetContext retrieves the context associated with the OAuth2 provider.
// It implements the GetContext() method of the Provider interface.
// The context is used to perform operations within the scope of the provider.
func (c *Cargo) GetContext() context.Context {
	return c.Ctx
}

// SetContext assigns the specified context to the OAuth2 provider.
// It implements the SetContext() method of the Provider interface.
// The context is crucial for performing operations within the scope of the provider.
func (c *Cargo) SetContext(ctx context.Context) {
	c.Ctx = ctx
}

// GetEndpoint retrieves the human-readable name of the OAuth2 provider.
func (c *Cargo) GetEndpoint() Endpoint {
	return c.Endpoint
}

// SetEndpoint sets the human-readable name of the OAuth2 provider.
func (c *Cargo) SetEndpoint(production, development string) {
	c.Endpoint = Endpoint{Production: production, Development: development}
}

// GetMode retrieves the human-readable name of the OAuth2 provider.
func (c *Cargo) GetMode() string {
	return c.Mode
}

// SetMode sets the human-readable name of the OAuth2 provider.
func (c *Cargo) SetMode(mode string) {
	c.Mode = mode
}

// GetUsername retrieves the human-readable name of the OAuth2 provider.
func (c *Cargo) GetUsername() string {
	return c.Username
}

// SetUsername sets the human-readable name of the OAuth2 provider.
func (c *Cargo) SetUsername(username string) {
	c.Username = username
}

// GetPassword retrieves the human-readable name of the OAuth2 provider.
func (c *Cargo) GetPassword() string {
	return c.Password
}

// SetPassword sets the human-readable name of the OAuth2 provider.
func (c *Cargo) SetPassword(password string) {
	c.Password = password
}
