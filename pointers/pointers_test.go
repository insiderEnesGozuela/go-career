package pointers

import "testing"

func TestWallet(t *testing.T) {

	assertBalance := func(t *testing.T, wallet Wallet, want Bitcoin) {
		got := wallet.balance

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	assertError := func(t *testing.T, got error, want string) {
		t.Helper()
		if got == nil {
			t.Fatal("wanted an error but didn't get one")
		}

		if got.Error() != want {
			t.Errorf("got %q, want %q", got, want)
		}
	}

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(1))

		want := Bitcoin(1)

		assertBalance(t, wallet, want)
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: 2}
		wallet.Withdraw(Bitcoin(1))

		want := Bitcoin(1)

		assertBalance(t, wallet, want)
	})

	t.Run("withdraw not enough", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{balance: startingBalance}
		err := wallet.Withdraw(Bitcoin(100))

		assertError(t, err, "oh no")
		assertBalance(t, wallet, startingBalance)
	})

}
