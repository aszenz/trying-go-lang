package pointers

import (
	"testing"

	"github.com/aszenz/basic-golang/go-test-helpers"
)

func TestWallet(t *testing.T) {
	assertBalance := func(t *testing.T, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		want := Bitcoin(10)
		assertBalance(t, wallet, want)
	})
	t.Run("Withdraw with funds", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(10))
		want := Bitcoin(10)
		assertBalance(t, wallet, want)
		testhelpers.AssertNoError(t, err)
	})
	t.Run("Withdraw more than balance", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{balance: startingBalance}
		got := wallet.Withdraw(Bitcoin(30))
		assertBalance(t, wallet, startingBalance)
		testhelpers.AssertError(t, got, errorInsufficientFunds)
	})
}
