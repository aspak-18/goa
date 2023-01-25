package weatherapi

import (
	"context"
	weather "example/gen/weather"
	"log"
)

// weather service example implementation.
// The example methods log the requests and return zero values.
type weathersrvc struct {
	logger *log.Logger
}

// NewWeather returns the weather service implementation.
func NewWeather(logger *log.Logger) weather.Service {
	return &weathersrvc{logger}
}

// Show implements show.
func (s *weathersrvc) Show(ctx context.Context, p *weather.ShowPayload) (res *weather.Weather, err error) {
	res = &weather.Weather{}
	s.logger.Print("weather.show")
	return
}
