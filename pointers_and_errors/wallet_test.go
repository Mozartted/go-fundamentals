package pointersanderrors

import (
	"testing"
)

func TestWallet(t *testing.T) {
	wallet := Wallet{}

	wallet.Deposit(10)

	got := wallet.Balance()

	want := Bitcoin(10)

	// Helper functions
	assertBalance := func(t testing.TB, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	assertError := func(t testing.TB, got error, want string) {
		t.Helper()

		if got == nil {
			t.Fatal("didn't get an error but wanted one")
		}

		if got.Error() != want {
			t.Errorf("got %q, want %q", got, want)
		}
	}

	if got != want {
		t.Errorf("got %s want %s", got.String(), want.String())
	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(40)}
		wallet.Deposit(Bitcoin(2))
	})

	t.Run("withdraw the funds", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		value := wallet.Withdraw(Bitcoin(50))
		assertError(t, value, "cannot withdraw, insufficient funds")
		assertBalance(t, wallet, Bitcoin(10))
	})
}
