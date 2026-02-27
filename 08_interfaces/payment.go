package _8_interfaces

import "fmt"

type Payment interface {
	pay(amount float64) (string, error)
}

type CreditCardPayment struct {
	cardnumber string
	cvv        int
	exp        string
}

type UPIPayment struct {
	upiid string
}

type NetBankingPayment struct {
	bankName string
}

func (cc CreditCardPayment) pay(amount float64) (string, error) {
	return fmt.Sprintf("Paying %d via via credit card with number %s", amount, cc.cardnumber), nil
}

func (u UPIPayment) pay(amount float64) (string, error) {
	return fmt.Sprintf("Paying %d via via  UPIPayment with UPI Id %s", amount, u.upiid), nil
}

func (n NetBankingPayment) pay(amount float64) (string, error) {
	return fmt.Sprintf("Paying %d via via Netbanking using %s", amount, n.bankName), nil
}

func MakePayment(p Payment, amount float64) (string, error) {
	return p.pay(amount)
}

func PaymentProcessor(pType string, pg PaymentGateway, amount float64) (string, error) {

	var payment Payment
	switch pType {
	case "credit-card":
		payment = CreditCardPayment{cardnumber: "123456", cvv: 100, exp: "12/34"}
	case "upi":
		payment = UPIPayment{upiid: "parshya@upi"}
	case "netbanking":
		payment = NetBankingPayment{bankName: "ICICI Bank"}
	default:
		return "", fmt.Errorf("Unknown payment type: %s", pType)
	}
	return MakePayment(payment, amount)
}
