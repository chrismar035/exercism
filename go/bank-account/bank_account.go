package account

import "sync"

// Account is a bank account that can be accessed in different ways at the same time.
type Account struct {
	closed  bool
	balance int64
	sync.Mutex
}

// Open returns a newly opened account with the provided initial deposit
func Open(initialDeposit int64) *Account {
	if initialDeposit < 0 {
		return nil
	}

	return &Account{balance: initialDeposit}
}

// Close will close the account and return the final balance
func (a *Account) Close() (payout int64, ok bool) {
	a.Lock()
	defer a.Unlock()
	if !a.closed {
		payout = a.balance
		a.closed = true
		ok = true
	}

	return payout, ok
}

// Balance will return the current balance for open accounts
func (a *Account) Balance() (balance int64, ok bool) {
	return a.balance, !a.closed
}

// Deposit will modify the current balance by the given amount and return the new balance.
// Negative provided amounts will ack as a withdraw.
func (a *Account) Deposit(amount int64) (newBalance int64, ok bool) {
	if a.closed {
		return 0, ok
	}

	a.Lock()
	defer a.Unlock()

	if a.balance+amount >= 0 {
		a.balance += amount
		ok = true
	}

	return a.balance, ok
}
