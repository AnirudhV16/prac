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
	
	var notifiers []types.Notify
	
	func NotifyAll(notifiers []types.Notify, message string){
		for _, v := range notifiers {
			err := v.Send(message string)
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Println("notified")
		}
		fmt.Println("notified message by all notifiers")
	}
}
