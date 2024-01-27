package cargo

// Credential information is different for live and test environments.
type Credential struct {
	title    string // Production or Development vs vs
	username string
	password string
	isActive bool // is active
}

// GetTitle retrieves the title associated with the Credential.
func (c *Credential) GetTitle() string {
	return c.title
}

// SetTitle sets the title associated with the Credential.
func (c *Credential) SetTitle(title string) {
	c.title = title
}

// GetUsername retrieves the username associated with the Credential.
func (c *Credential) GetUsername() string {
	return c.username
}

// SetUsername sets the username associated with the Credential.
func (c *Credential) SetUsername(username string) {
	c.username = username
}

// GetPassword retrieves the password associated with the Credential.
func (c *Credential) GetPassword() string {
	return c.password
}

// SetPassword sets the password associated with the Credential.
func (c *Credential) SetPassword(password string) {
	c.password = password
}

// GetActive retrieves the isActive (True or False) status of the Credential.
func (c *Credential) GetActive() bool {
	return c.isActive
}

// SetActive sets the isActive (True or False) status of the Credential.
func (c *Credential) SetActive(isActive bool) {
	c.isActive = isActive
}
