package _8_interfaces

import "fmt"

type PaymentGateway interface {
	processPayment(amount float64)
}

type RozorPayGateway struct {
	RozarPayDetails
}

type RozarPayDetails struct {
	gatewayUrl string
	apiKey     string
	apiSecret  string
}

type PaypalGateway struct {
	PaypalDetails
}

type PaypalDetails struct {
	gatewayUrl   string
	clientId     string
	clientSecret string
}

func (r RozorPayGateway) processPayment(amount float64) {
	fmt.Println("RozorPayGateway processPayment", r.gatewayUrl)
}

func (p PaypalGateway) processPayment(amount float64) {
	fmt.Println("PaypalGateway processPayment", p.gatewayUrl)
}
