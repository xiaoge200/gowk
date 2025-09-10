package gowk

import (
	"encoding/json"

	"github.com/go-resty/resty/v2"
)

type Option func(*GoWk)

type GoWk struct {
	restyClient *resty.Client
}

func NewClient(baseurl string, options ...Option) *GoWk {
	c := GoWk{
		restyClient: resty.New(),
	}
	c.restyClient.SetBaseURL(baseurl)
	// c.restyClient.SetRetryCount(3)
	// c.restyClient.SetTimeout(5 * time.Second)
	// c.restyClient.SetRetryWaitTime(1 * time.Second)
	// c.restyClient.SetRetryMaxWaitTime(10 * time.Second)
	c.restyClient.OnAfterResponse(handleAPIError)

	for _, option := range options {
		option(&c)
	}
	return &c
}

func (g *GoWk) RestyClient() *resty.Client {
	return g.restyClient
}

func handleAPIError(c *resty.Client, resp *resty.Response) error {
	if resp.IsError() {
		if err, ok := resp.Error().(*APIError); ok {
			err.StatusCode = resp.StatusCode()
			return err
		}
		err := &APIError{}
		json.Unmarshal(resp.Body(), err)
		err.StatusCode = resp.StatusCode()
		return err
	}
	return nil
}
