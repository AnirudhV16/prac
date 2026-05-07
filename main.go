package main

import (
	"fmt"

	"github.com/AnirudhV16/prac/types"
	"github.com/AnirudhV16/prac/use"
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

	//discount
	Pcustomer := types.Customer{
		Name:               "Anirudh",
		DiscountCalculator: types.PremiumDiscount{},
	}
	Rcustomer := types.Customer{
		Name:               "Anirudh",
		DiscountCalculator: types.RegularDiscount{},
	}
	Scustomer := types.Customer{
		Name:               "Anirudh",
		DiscountCalculator: types.StudentDiscount{},
	}

	customers := []types.Customer{Pcustomer, Rcustomer, Scustomer}
	for _, c := range customers {
		fmt.Println(c.CalculateDiscount(1000))
	}

	//L - principle
	circle, err := types.NewCircle(12)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(circle.Area())
	}
	sphere, err := types.NewSphere(14)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(sphere.Volume())
	}

	//D principle
	sms := types.SMSService{}
	email := types.EmailService{}

	//notifier
	services := []types.MessageService{sms, email}
	for _, v := range services {
		notifier := types.NewOrderNotifier(v)
		notifier.NotifyOrder("Hiiiiiiiiiiiiiiiiiiiiii")
	}

	//factory
	eemail, err := types.NewNotification("email")
	smss, err := types.NewNotification("sms")
	pushh, err := types.NewNotification("push")

	notifierss := []types.Notification{eemail, smss, pushh}
	for _, v := range notifierss {
		v.Notify("Helloooooo!")
	}

	//builder
	pizza1, err := types.NewPizzaBuilder("medium", "???").WithSauce("????").WithCheese("*****").WithToppings("----", "88888", "44444", "0000").Build()
	if err != nil {
		fmt.Println(err)
	}
	pizza2, err := types.NewPizzaBuilder("large", "???").WithSauce("????").WithCheese("*****").WithToppings("----").Build()
	if err != nil {
		fmt.Println(err)
	}
	pizza1.Describe()
	pizza2.Describe()

	//abstract
	mysqlfactory := types.NewDatabaseClient(&types.MySQLFactory{})
	mysqlfactory.Run("Hi.......")

	postgresfactory := types.NewDatabaseClient(&types.PostgresFactory{})
	postgresfactory.Run("Hello.......")

	//decorator
	s := "  Hello World  "

	var holder types.TextProcessor = types.NewPlainTextProcessor()
	holder = types.NewUpperCaseDecorator(holder)
	holder = types.NewTrimDecorator(holder)
	holder = types.NewExclamationDecorator(holder)

	fmt.Println(holder.Process(s))

	//fascade
	HomeTheatre := types.HomeTheatreFacade{}
	HomeTheatre.WatchMovie("Raabta")
	HomeTheatre.EndMovie()

	//composite
	smanager := types.NewManager("A", 600, make([]types.Employee, 0))
	dev1 := types.NewDeveloper("B", 500)
	smanager.AddSubordinate(dev1)

	jmanager := types.NewManager("C", 450, make([]types.Employee, 0))
	smanager.AddSubordinate(jmanager)

	dev2 := types.NewDeveloper("D", 400)
	dev3 := types.NewDeveloper("E", 420)
	jmanager.AddSubordinate(dev2)
	jmanager.AddSubordinate(dev3)

	//this should print sum of all the salaries in the tree 2300 + 70 => 2370
	fmt.Println(smanager.GetSalary())

	//observer
	subject := types.NewStockMarket(make([]types.Stocker, 0)) //this is subject
	ob1 := types.NewTradeAlert("AAAAAAA")
	ob2 := types.NewPriceLogger()

	subject.Subscribe(ob1)
	subject.Subscribe(ob2)

	subject.SetPrice("bitcoin", 24234)

	subject.Unsubscribe(ob2)

	subject.SetPrice("dozge", 2000)

	//command

	light := types.NewLight()
	invoker := types.RemoteControl{}
	invoker.PressButton(&types.TurnOnCommand{Light: light})
	invoker.PressButton(&types.TurnOffCommand{Light: light})
	invoker.PressButton(&types.TurnOffCommand{Light: light})
	invoker.PressUndo()
	invoker.PressUndo()

	fmt.Printf("state of light: %d\n", light.GetState())

	//template
	pdf := types.PDFReport{}
	excel := types.ExcelReport{}
	types.Run(pdf)
	types.Run(excel)

	//parking lot
	//made a bike spot
	spot1 := use.NewBasicSpot(use.BikeFeeCalculator{}, "bike")
	//made floor
	floor := use.Floor{}
	//added bike spot in to the floor
	floor.AddSpots(spot1)
	//made a parking and added one floor to it
	parking := use.ParkingLot{Floors: []use.Floor{floor}}
	//made vehicle
	vehicle := use.NewVehicle("XYZ", "AP1234", "bike")
	//park the bike
	ticket, err := parking.Park(*vehicle)
	if err != nil {
		fmt.Println(err)
	} else { //unpark the vehicle
		fees := parking.Unpark(*ticket)
		fmt.Printf("Fees to be paid %f\n", fees)
	}

	//library, 2 books, 1 member -> borrows 2 books

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
