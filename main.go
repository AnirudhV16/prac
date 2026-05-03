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
	fmt.Println(Account1.GetBalance())

	Account1.Withdraw(500)
	fmt.Println(Account1.GetBalance())

	//Notifier

	notifiers := []types.Notify{&types.MailNotifier{Address: "example@gmail.com"},
		&types.SMSNotifier{PhoneNumber: "1234567890"},
		&types.PushNotifier{Device: "HFD12347578"},
		&types.WhatsappNotifier{PhoneNumber: "1234567890"}}
	NotifyAll(notifiers, "new movie update")

	// Embedding
	var Account2 *types.SavingsAccount
	Base := *types.NewBaseAccount("Anirudh", 1000)
	Account2 = types.NewSavingsAccount(Base, 12)
	Account2.Deposit(10000)
	Account2.ApplyInterest()
	fmt.Println(Account2.GetBalance())

	//Restaurant
	r1 := types.NewRestaurtant("AM to PM", 122234, 57848896, 4, 33)
	r2 := types.NewRestaurtant("Being hungry", 122223452, 578776, 3, 44)
	r3 := types.NewRestaurtant("salt", 765544, 66696, 2, 22)
	r4 := types.NewRestaurtant("sweet Magic", 434432, 99865758, 3, 54)

	ratings := []types.Rateable{&r1.Rating, &r2.Rating, &r3.Rating, &r4.Rating}

	types.PrintTopRated(ratings, 2)

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
