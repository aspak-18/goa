// Code generated by goa v3.11.0, DO NOT EDIT.
//
// weather service
//
// Command:
// $ goa gen example/design

package weather

import (
	"context"
)

// The weather service returns current weather data for a given city.
type Service interface {
	// Show implements show.
	Show(context.Context, *ShowPayload) (res *Weather, err error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "weather"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [1]string{"show"}

// ShowPayload is the payload type of the weather service show method.
type ShowPayload struct {
	// Name of the city
	City string
}

// Weather is the result type of the weather service show method.
type Weather struct {
	// Temperature in Celsius
	Temp float64
	// Wind speed in meters per second
	WindSpeed float64
	// Humidity in percent
	Humidity float64
}
