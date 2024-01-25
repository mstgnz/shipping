package cargo

import (
	"context"
)

// Cargo struct represents a structure designed to manage multiple endpoints and credentials in cargo integrations.
// It facilitates the storage of information about various available credentials and endpoints, providing the flexibility to switch between them.
// The fields CurrentCredential and CurrentEndpoint specifically store information about the currently active credential and endpoint.
// This enables you to use the active ones in your cargo requests. You can modify these active settings using the provided set methods.
//
// ServiceType represents the type of service utilized in cargo integrations, which can be either "Soap" or "Rest".
// It indicates the communication protocol employed for data exchange.
//
// Mode specifies the operational mode of the cargo integration, which can be either "Production" or "Development".
// It defines whether the integration is configured for real-world usage or testing purposes.
//
// IsDomestic indicates whether the cargo corresponds to a domestic product (True) or an international one (False).
// It helps in distinguishing between shipments within the country and those abroad.
type Cargo struct {
	Ctx               context.Context
	ServiceType       string       // Soap or Rest
	Mode              string       // Production or Development
	IsDomestic        bool         // True if the product is domestic, False if it is abroad
	CurrentCredential string       // Current Credential
	CurrentEndpoint   string       // Current Endpoint
	Credentials       []Credential // List of available credentials
	Endpoints         []Endpoint   // List of available endpoints
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

// GetDomestic retrieves the IsDomestic (True or False) status of the Cargo.
func (c *Cargo) GetDomestic() bool {
	return c.IsDomestic
}

// SetDomestic sets the IsDomestic (True or False) status of the Cargo.
func (c *Cargo) SetDomestic(isDomestic bool) {
	c.IsDomestic = isDomestic
}

// GetCurrentCredential get active credential title
func (c *Cargo) GetCurrentCredential() string {
	return c.CurrentCredential
}

// SetCurrentCredential set active credential title
func (c *Cargo) SetCurrentCredential(currentCredential string) {
	c.CurrentCredential = currentCredential
}

// GetCurrentEndpoint get active endpoint title
func (c *Cargo) GetCurrentEndpoint() string {
	return c.CurrentEndpoint
}

// SetCurrentEndpoint set active endpoint title
func (c *Cargo) SetCurrentEndpoint(currentEndpoint string) {
	c.CurrentEndpoint = currentEndpoint
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

// AddCredential adds a new credential to the Cargo with the specified title, username, and password.
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

// AddEndpoint adds a new endpoint to the Cargo with the specified title, production, and development URLs.
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
