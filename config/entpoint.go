package cargo

type Endpoint struct {
	Title       string // some cargo companies may have more than one endpoint. (shipment, reference, vs, vs)
	Production  string // live url
	Development string // test url
}

func (c *Endpoint) GetTitle() string {
	return c.Title
}

// SetTitle sets the title of the endpoint.
func (c *Endpoint) SetTitle(title string) {
	c.Title = title
}

func (c *Endpoint) GetProduction() string {
	return c.Production
}

// SetProduction sets the production endpoint URL.
func (c *Endpoint) SetProduction(production string) {
	c.Production = production
}

func (c *Endpoint) GetDevelopment() string {
	return c.Development
}

// SetDevelopment sets the development endpoint URL.
func (c *Endpoint) SetDevelopment(development string) {
	c.Development = development
}

// GetActiveUrl Returns Url according to Mode value
func (c *Endpoint) GetActiveUrl(mode string) string {
	if mode == "Development" {
		return c.Development
	}
	return c.Production
}
