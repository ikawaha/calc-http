// Code generated by goa v3.7.6, DO NOT EDIT.
//
// calc client
//
// Command:
// $ goa gen calc/design

package calc

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "calc" service client.
type Client struct {
	MultiplyEndpoint goa.Endpoint
	DivEndpoint      goa.Endpoint
}

// NewClient initializes a "calc" service client given the endpoints.
func NewClient(multiply, div goa.Endpoint) *Client {
	return &Client{
		MultiplyEndpoint: multiply,
		DivEndpoint:      div,
	}
}

// Multiply calls the "multiply" endpoint of the "calc" service.
func (c *Client) Multiply(ctx context.Context, p *MultiplyPayload) (res int, err error) {
	var ires interface{}
	ires, err = c.MultiplyEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(int), nil
}

// Div calls the "div" endpoint of the "calc" service.
// Div may return the following errors:
//	- "DivByZero" (type *goa.ServiceError)
//	- error: internal error
func (c *Client) Div(ctx context.Context, p *DivPayload) (res int, err error) {
	var ires interface{}
	ires, err = c.DivEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(int), nil
}
