package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = API("calc", func() {
	Title("Calculator Service")
	Description("Service for calc numbers")
	Server("calc", func() {
		Host("localhost", func() {
			URI("http://localhost:8000")
		})
	})
})

var _ = Service("calc", func() {
	Description("The calc service performs operations on numbers.")
	Method("multiply", func() {
		Payload(func() {
			Attribute("a", Int, "Left operand")
			Attribute("b", Int, "Right operand")
			Required("a", "b")
		})
		Result(Int)
		HTTP(func() {
			GET("/multiply/{a}/{b}")
		})
	})
	Method("div", func() {
		Error("DivByZero")
		Payload(func() {
			Attribute("a", Int, "Left operand")
			Attribute("b", Int, "Right operand")
			Required("a", "b")
		})
		Result(Int)
		HTTP(func() {
			GET("/div/{a}/{b}")
			Response("DivByZero", StatusBadRequest)
		})
	})
	Files("/openapi.json", "./gen/http/openapi3.json")
})
