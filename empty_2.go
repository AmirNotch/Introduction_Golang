package main

import "fmt"

type Wallet struct {
	Cash int
}
type Payer interface {
	Pay(int) error
}

func (w *Wallet) Pay(amount int) error {
	 if w.Cash < amount{
		 return fmt.Errorf("Not enought money")
	 }
	 w.Cash -= amount
	 return nil
}

func Buy(in interface{}){
	var p Payer
	var ok bool
	if p, ok = in.(Payer); !ok {
		fmt.Printf("%T не является платёжным средством ", in)
		return
	}

	err := p.Pay(10)
	if err != nil {
		fmt.Printf("Ошибка при оплате %T %v\n\n", p, err)
		return
	}
	fmt.Printf("Спасибо за покупку через %T\n\n", p)
}

func Change(in interface{}) {
	var changing Payer
	var nice bool
	if changing,nice = in.(Payer); !nice {
		fmt.Printf("%T не является подленным", )
		return
	}

	erring := changing.Pay(20)
	if erring != nil {
		fmt.Printf("Ошибка при обмене %T %v\n\n", changing, erring)
		return
	}
	fmt.Printf("Спасибо за хороший обмен %T\n\n", changing)

}
func main() {

	myWallet := &Wallet{Cash: 100}
	Buy(myWallet)
	Buy([]int{1,2,3,4})
	Buy(3.14)
	Change(myWallet)
	Change([]int{1,2,3,4})
	Change(3.14)

}
