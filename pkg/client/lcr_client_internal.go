package golcr

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strconv"
	"strings"
)

const authURL = "https://signin.lds.org/login.html"
const currentUserURL = "https://www.lds.org/htvt/services/v1/user/currentUser"

func (c *lcrClient) connect(creds Creds) error {
	cookieJar, _ := cookiejar.New(nil)

	c.client = &http.Client{Jar: cookieJar}

	user, pass := creds.GetUserPass()

	data := url.Values{}
	data.Set("username", user)
	data.Add("password", pass)

	r, err := http.NewRequest("POST", authURL, strings.NewReader(data.Encode()))
	if err != nil {
		return err
	}

	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	res, err := c.client.Do(r)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if strings.Contains(string(body), "<title>Sign in</title>") {
		return errors.New("unable to log in with credentials")
	}

	setCookie := res.Header["Set-Cookie"]
	val := strings.Split(setCookie[0], ";")
	c.token = val[0]

	return nil
}

func (c *lcrClient) get(URL string) ([]byte, error) {

	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("cookie", c.token)

	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code %v", res.StatusCode)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
