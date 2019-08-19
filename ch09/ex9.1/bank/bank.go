package bank

var deposits = make(chan int) // send amount to deposit
var balances = make(chan int) // receive balance
var withdraws = make(chan int)

func Deposit(amount int)  { deposits <- amount }
func Balance() int        { return <-balances }
func Withdraw(amount int) { withdraws <- amount }

func teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		case amount := <-withdraws:
			balance -= amount
		}
	}
}

func init() {
	go teller() // start the monitor goroutine
}
