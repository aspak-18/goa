// Code generated by goa v3.11.0, DO NOT EDIT.
//
// weather client
//
// Command:
// $ goa gen example/design

package weather

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "weather" service client.
type Client struct {
	ShowEndpoint goa.Endpoint
}

// NewClient initializes a "weather" service client given the endpoints.
func NewClient(show goa.Endpoint) *Client {
	return &Client{
		ShowEndpoint: show,
	}
}

// Show calls the "show" endpoint of the "weather" service.
func (c *Client) Show(ctx context.Context, p *ShowPayload) (res *Weather, err error) {
	var ires interface{}
	ires, err = c.ShowEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*Weather), nil
}
