// Code generated by goa v3.11.0, DO NOT EDIT.
//
// weather HTTP client CLI support package
//
// Command:
// $ goa gen example/design

package client

import (
	weather "example/gen/weather"
)

// BuildShowPayload builds the payload for the weather show endpoint from CLI
// flags.
func BuildShowPayload(weatherShowCity string) (*weather.ShowPayload, error) {
	var city string
	{
		city = weatherShowCity
	}
	v := &weather.ShowPayload{}
	v.City = city

	return v, nil
}
