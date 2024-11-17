package pointers

import (
	"testing"
)

func TestWallet(t *testing.T) {

	assertBalance := func(t testing.TB, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}
	assertNoError := func(t testing.TB, got error) {
		t.Helper()

		if got != nil {
			t.Errorf("expected no error but got an error %s", got.Error())
		}

	}

	assertError := func(t testing.TB, got, want error) {
		t.Helper()

		if got == nil {
			t.Errorf("expected an error but did not receive one")
		}
		if want.Error() != got.Error() {
			t.Errorf("got %q but expected %q", got, want)
		}

	}

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(10))
		assertNoError(t, err)
		assertBalance(t, wallet, Bitcoin(10))
	})
	t.Run("withdraw insufficient funds", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(10)}
		err := wallet.Withdraw(Bitcoin(20))
		assertError(t, err, ErrInsuficientFunds)
		assertBalance(t, wallet, Bitcoin(10))
	})

}
