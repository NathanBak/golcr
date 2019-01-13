package golcr

import (
	"encoding/json"
	"fmt"
	"net/http"

	. "github.com/NathanBak/golcr/pkg/model"
)

type lcrClient struct {
	client *http.Client
	token  string
}

func (c *lcrClient) GetOrgs() (*[]Org, error) {
	orgs := &[]Org{}
	body, err := c.get("https://lcr.lds.org/services/orgs/sub-org-name-hierarchy?lang=eng")
	if err != nil {
		return orgs, err
	}

	err = json.Unmarshal(body, orgs)
	return orgs, err
}

func (c *lcrClient) GetSubOrgWithCallings(subOrgID int64) (*Org, error) {
	orgs := []Org{}

	url := fmt.Sprintf("https://lcr.lds.org/services/orgs/sub-orgs-with-callings?lang=eng&subOrgId=%v", subOrgID)
	body, err := c.get(url)
	if err != nil {
		return nil, err
	}

	//log.Println(string(body))

	err = json.Unmarshal(body, &orgs)
	if err != nil {
		return nil, err
	}

	if len(orgs) == 0 {
		return nil, fmt.Errorf("unable to find sub org %v", subOrgID)
	}

	if len(orgs) > 1 {
		return nil, fmt.Errorf("found multiple matches for sub org %v", subOrgID)
	}

	return &orgs[0], err
}

func (c *lcrClient) GetMembers(unitNumber int64) (*[]Member, error) {
	url := fmt.Sprintf("https://lcr.lds.org/services/umlu/report/member-list?lang=eng&unitNumber=%v", unitNumber)
	members := &[]Member{}
	body, err := c.get(url)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, members)
	return members, err
}
