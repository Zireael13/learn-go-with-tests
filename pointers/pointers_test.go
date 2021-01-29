package pointers

import "testing"

func assertBalance(t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertError(t testing.TB, got error, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("Wanted error but didn't get one")
	}

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func assertNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Fatalf("Got error %s", got)
	}
}

func TestWallet(t *testing.T) {

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))

		assertBalance(t, wallet, Bitcoin(10))
	})
	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(10)}

		err := wallet.Withdraw(Bitcoin(3))

		assertBalance(t, wallet, Bitcoin(7))

		assertNoError(t, err)
	})

	t.Run("Withdraw less than 0", func(t *testing.T) {
		startBal := Bitcoin(10)
		wallet := Wallet{balance: startBal}
		withdraw := wallet.Withdraw(Bitcoin(15))
		assertBalance(t, wallet, startBal)

		assertError(t, withdraw, InsufficentFundsErr)

	})
}
