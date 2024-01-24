package cargo

import (
	"context"
)

type Cargo struct {
	Ctx         context.Context
	ServiceType string       // Soap or Rest
	Mode        string       // Production or Development
	Credentials []Credential // Production and Development
	Endpoints   []Endpoint   // Some carriers may have more than one endpoint, and has a Production and Development endpoint for each one
}

// GetContext retrieves the context associated with the Cargo.
// It implements the GetContext() method of the Provider interface.
// The context is used to perform operations within the scope of the Cargo.
func (c *Cargo) GetContext() context.Context {
	return c.Ctx
}

// SetContext assigns the specified context to the Cargo.
// It implements the SetContext() method of the Provider interface.
// The context is crucial for performing operations within the scope of the Cargo.
func (c *Cargo) SetContext(ctx context.Context) {
	c.Ctx = ctx
}

// GetServiceType retrieves the mode (Soap or Rest) of the Cargo.
func (c *Cargo) GetServiceType() string {
	return c.ServiceType
}

// SetServiceType sets the mode (Soap or Rest) of the Cargo.
func (c *Cargo) SetServiceType(serviceType string) {
	c.ServiceType = serviceType
}

// GetMode retrieves the mode (Production or Development) of the Cargo.
func (c *Cargo) GetMode() string {
	return c.Mode
}

// SetMode sets the mode (Production or Development) of the Cargo.
func (c *Cargo) SetMode(mode string) {
	c.Mode = mode
}

// GetCredentials retrieves the list of credentials in the Cargo.
func (c *Cargo) GetCredentials() []Credential {
	return c.Credentials
}

// GetCredential retrieves a specific credential by its title.
func (c *Cargo) GetCredential(title string) Credential {
	for _, credential := range c.Credentials {
		if title == credential.Title {
			return credential
		}
	}
	return Credential{}
}

// AddCredential adds a new credential to the Cargo.
func (c *Cargo) AddCredential(title, username, password string) {
	c.Credentials = append(c.Credentials, Credential{Title: title, Username: username, Password: password})
}

// DelCredential deletes a specific credential by its title.
func (c *Cargo) DelCredential(title string) []Credential {
	for index, credential := range c.Credentials {
		if title == credential.Title {
			c.Credentials = append(c.Credentials[:index], c.Credentials[index+1:]...)
		}
	}
	return c.Credentials
}

// GetEndpoints retrieves the list of endpoints in the Cargo.
func (c *Cargo) GetEndpoints() []Endpoint {
	return c.Endpoints
}

// GetEndpoint retrieves a specific endpoint by its title.
func (c *Cargo) GetEndpoint(title string) Endpoint {
	for _, endpoint := range c.Endpoints {
		if title == endpoint.Title {
			return endpoint
		}
	}
	return Endpoint{}
}

// AddEndpoint adds a new endpoint to the Cargo.
func (c *Cargo) AddEndpoint(title, production, development string) {
	c.Endpoints = append(c.Endpoints, Endpoint{Title: title, Production: production, Development: development})
}

// DelEndpoint deletes a specific endpoint by its title.
func (c *Cargo) DelEndpoint(title string) []Endpoint {
	for index, endpoint := range c.Endpoints {
		if title == endpoint.Title {
			c.Endpoints = append(c.Endpoints[:index], c.Endpoints[index+1:]...)
		}
	}
	return c.Endpoints
}
