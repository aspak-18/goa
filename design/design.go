package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = Service("weather", func() {
	Description("The weather service returns current weather data for a given city.")
	Method("show", func() {
		Payload(func() {
			Attribute("city", String, "Name of the city")
			Required("city")
		})
		Result(Weather)
		HTTP(func() {
			GET("/{city}")
			Response(StatusOK)
		})
	})
})

var Weather = Type("Weather", func() {
	Attribute("temp", Float64, "Temperature in Celsius")
	Attribute("wind_speed", Float64, "Wind speed in meters per second")
	Attribute("humidity", Float64, "Humidity in percent")
	Required("temp", "wind_speed", "humidity")
})
