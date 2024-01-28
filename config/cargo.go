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
// Mode specifies the operational mode of the cargo integration, which can be either "production" or "development".
// It defines whether the integration is configured for real-world usage or testing purposes.
//
// isDomestic indicates whether the cargo corresponds to a domestic product (True) or an international one (False).
// It helps in distinguishing between shipments within the country and those abroad.
type Cargo struct {
	ctx         context.Context
	serviceType ServiceType   // Soap or Rest
	mode        Mode          // Production or Development
	isDomestic  bool          // True if the product is domestic, False if it is abroad
	endpoints   []*Endpoint   // List of available endpoints
	credentials []*Credential // List of available credentials
}

// GetContext retrieves the context associated with the Cargo.
// It implements the GetContext() method of the Provider interface.
// The context is used to perform operations within the scope of the Cargo.
func (c *Cargo) GetContext() context.Context {
	return c.ctx
}

// SetContext assigns the specified context to the Cargo.
// It implements the SetContext() method of the Provider interface.
// The context is crucial for performing operations within the scope of the Cargo.
func (c *Cargo) SetContext(ctx context.Context) {
	c.ctx = ctx
}

// GetServiceType retrieves the mode (Soap or Rest) of the Cargo.
func (c *Cargo) GetServiceType() ServiceType {
	return c.serviceType
}

// SetServiceType sets the mode (Soap or Rest) of the Cargo.
func (c *Cargo) SetServiceType(serviceType ServiceType) {
	c.serviceType = serviceType
}

// GetMode retrieves the mode (Production or Development) of the Cargo.
func (c *Cargo) GetMode() Mode {
	return c.mode
}

// SetMode sets the mode (Production or Development) of the Cargo.
func (c *Cargo) SetMode(mode Mode) {
	c.mode = mode
}

// IsDomestic retrieves the isDomestic (True or False) status of the Cargo.
func (c *Cargo) IsDomestic() bool {
	return c.isDomestic
}

// SetDomestic sets the isDomestic (True or False) status of the Cargo.
func (c *Cargo) SetDomestic(isDomestic bool) {
	c.isDomestic = isDomestic
}

// GetEndpoints retrieves the list of endpoints in the Cargo.
func (c *Cargo) GetEndpoints() []*Endpoint {
	return c.endpoints
}

// GetEndpoint retrieves a specific endpoint by its title.
func (c *Cargo) GetEndpoint(title string) *Endpoint {
	for _, endpoint := range c.endpoints {
		if title == endpoint.title {
			return endpoint
		}
	}
	return nil
}

// AddEndpoint adds a new endpoint to the Cargo with the specified title, production, and development URLs.
func (c *Cargo) AddEndpoint(title, production, development string) *Endpoint {
	endpoint := &Endpoint{title: title, production: production, development: development}
	c.endpoints = append(c.endpoints, endpoint)
	return endpoint
}

// DelEndpoint deletes a specific endpoint by its title.
func (c *Cargo) DelEndpoint(title string) []*Endpoint {
	for index, endpoint := range c.endpoints {
		if title == endpoint.title {
			c.endpoints = append(c.endpoints[:index], c.endpoints[index+1:]...)
		}
	}
	return c.endpoints
}

// GetCredentials retrieves the list of credentials in the Cargo.
func (c *Cargo) GetCredentials() []*Credential {
	return c.credentials
}

// GetCredential retrieves a specific credential by its title.
func (c *Cargo) GetCredential(title string) *Credential {
	for _, credential := range c.credentials {
		if title == credential.title {
			return credential
		}
	}
	return nil
}

// AddCredential adds a new credential to the Cargo with the specified title, username, and password.
func (c *Cargo) AddCredential(title, username, password string) *Credential {
	credential := &Credential{title: title, username: username, password: password}
	c.credentials = append(c.credentials, credential)
	return credential
}

// DelCredential deletes a specific credential by its title.
func (c *Cargo) DelCredential(title string) []*Credential {
	for index, credential := range c.credentials {
		if title == credential.title {
			c.credentials = append(c.credentials[:index], c.credentials[index+1:]...)
		}
	}
	return c.credentials
}

// GetCurrentEndpointAndCredential return current endpoint and credential
func (c *Cargo) GetCurrentEndpointAndCredential() Current {
	var current Current
	for _, credential := range c.credentials {
		if credential.isActive {
			current.Credential = credential
		}
	}
	for _, endpoint := range c.endpoints {
		if endpoint.isActive {
			current.Endpoint = endpoint
		}
	}
	return current
}
