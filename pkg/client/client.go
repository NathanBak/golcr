package golcr

import . "github.com/NathanBak/golcr/pkg/model"

// Client contains functions to interact with LCR.
type Client interface {
	GetOrgs() (*[]Org, error)
	GetSubOrgWithCallings(subOrgID int64) (*Org, error)
	GetMembers(unitNumber int64) (*[]Member, error)
}

// NewClient creates, initializes and returns a new Client.
func NewClient(creds Creds) (Client, error) {
	c := &lcrClient{}

	return c, c.connect(creds)
}
