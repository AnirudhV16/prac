package main

import (
	"fmt"

	"github.com/AnirudhV16/prac/types"
)

func main() {
	//Account1 := types.NewBankAccount("1234", "Anirudh", 1000.00)
	var Account1 types.Account
	Account1 = types.NewBankAccount("1234", "Anirudh", 1000.00)

	Account1.Deposit(50000)
	Account1.GetBalance()

	Account1.Withdraw(500)
	Account1.GetBalance()

	//Notifier

	notifiers := []types.Notify{&types.MailNotifier{Address: "example@gmail.com"},
		&types.SMSNotifier{Phonenumber: "1234567890"},
		&types.PushNotifier{Device: "HFD12347578"},
		&types.WhatsappNotifier{Phonenumber: "1234567890"}}
	NotifyAll(notifiers, "new movie update")

}

func NotifyAll(notifiers []types.Notify, message string) {
	for _, v := range notifiers {
		err := v.Send(message)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println("notified")
	}
	fmt.Println("notified message by all notifiers")
}
