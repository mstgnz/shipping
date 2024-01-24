package cargo

type Credential struct {
	Title    string // Production or Development
	Username string
	Password string
}

// GetTitle retrieves the title associated with the Credential.
func (c *Credential) GetTitle() string {
	return c.Title
}

// SetTitle sets the title associated with the Credential.
func (c *Credential) SetTitle(title string) {
	c.Title = title
}

// GetUsername retrieves the username associated with the Credential.
func (c *Credential) GetUsername() string {
	return c.Username
}

// SetUsername sets the username associated with the Credential.
func (c *Credential) SetUsername(username string) {
	c.Username = username
}

// GetPassword retrieves the password associated with the Credential.
func (c *Credential) GetPassword() string {
	return c.Password
}

// SetPassword sets the password associated with the Credential.
func (c *Credential) SetPassword(password string) {
	c.Password = password
}
