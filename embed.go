package main

import "fmt"

type Payer interface {
	Pay(int) error
}

type Ringer interface {
	Ring(string) error
}
type NFCPhone interface {
	Ringer
	Payer
}
type Phone struct {
	Money int
	PhoneID string
}

func (p *Phone) Pay(amount int) error {
	if p.Money < amount {
		return fmt.Errorf("Not enought money on account")
	}
	p.Money -= amount
	return nil
}

func (p *Phone) Ring(number string) error {
	if number == "" {
		return fmt.Errorf("Please enter number")
	}
	return nil
}

func PayForMetWithPhone(phone NFCPhone){
	err := phone.Pay(1)
	if err != nil {
		fmt.Printf("Ошибка при оплате %v \n\n", err)
		return
	}
	fmt.Printf("Турникет открыт через %T\n\n", phone)
}

func main() {
	myPhone := &Phone{Money: 100}
	PayForMetWithPhone(myPhone)

}
