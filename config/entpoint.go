package cargo

type Endpoint struct {
	title       string // some cargo companies may have more than one endpoint. (shipment, reference, vs, vs)
	production  string // live url
	development string // test url
	isActive    bool   // is active
}

func (e *Endpoint) GetTitle() string {
	return e.title
}

// SetTitle sets the title of the endpoint.
func (e *Endpoint) SetTitle(title string) {
	e.title = title
}

func (e *Endpoint) GetProduction() string {
	return e.production
}

// SetProduction sets the production endpoint URL.
func (e *Endpoint) SetProduction(production string) {
	e.production = production
}

func (e *Endpoint) GetDevelopment() string {
	return e.development
}

// SetDevelopment sets the development endpoint URL.
func (e *Endpoint) SetDevelopment(development string) {
	e.development = development
}

// GetActiveUrl Returns Url according to Mode value
func (e *Endpoint) GetActiveUrl(mode Mode) string {
	if mode == DEVELOPMENT {
		return e.development
	}
	if mode == PRODUCTION {
		return e.development
	}
	return e.development
}

// GetActive retrieves the isActive (True or False) status of the Endpoint.
func (e *Endpoint) GetActive() bool {
	return e.isActive
}

// SetActive sets the isActive (True or False) status of the Endpoint.
func (e *Endpoint) SetActive(isActive bool) {
	e.isActive = isActive
}
