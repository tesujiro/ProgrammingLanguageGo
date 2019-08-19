package bank_test

import (
	"fmt"
	"testing"

	"github.com/tesujiro/TheGoProgrammingLanguage/ch09/ex9.1/bank"
)

func TestBank(t *testing.T) {
	done := make(chan struct{})

	// Alice
	go func() {
		bank.Deposit(200)
		fmt.Println("=", bank.Balance())
		done <- struct{}{}
	}()

	// Bob
	go func() {
		bank.Deposit(100)
		done <- struct{}{}
	}()

	// Wait for both transactions.
	<-done
	<-done

	if got, want := bank.Balance(), 300; got != want {
		t.Errorf("Balance = %d, want %d", got, want)
	}

	// Alice
	go func() {
		if got, want := bank.Withdraw(100), true; got != want {
			t.Errorf("Withdraw(100) = %v, want %v", got, want)
		}
		fmt.Println("=", bank.Balance())
		done <- struct{}{}
	}()

	// Bob
	go func() {
		if got, want := bank.Withdraw(100), true; got != want {
			t.Errorf("Withdraw(100) = %v, want %v", got, want)
		}
		done <- struct{}{}
	}()

	// Wait for both transactions.
	<-done
	<-done

	if got, want := bank.Balance(), 100; got != want {
		t.Errorf("Balance = %d, want %d", got, want)
	}
	if got, want := bank.Withdraw(500), false; got != want {
		t.Errorf("Withdraw(500) = %v, want %v", got, want)
	}
	if got, want := bank.Balance(), 100; got != want {
		t.Errorf("Balance = %d, want %d", got, want)
	}
}
