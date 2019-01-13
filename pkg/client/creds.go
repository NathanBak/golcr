package golcr

// Creds provides the necessary credentials to authenticate with lds.org.
type Creds interface {
	GetUserPass() (username string, password string)
}
