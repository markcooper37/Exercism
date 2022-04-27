package account

import "sync"

type Account struct {
	mu    sync.Mutex
	money int64
	open bool
}

func Open(amount int64) *Account {
	if amount < 0 {
		return nil
	}
	return &Account{money: amount, open: true}
}

func (a *Account) Balance() (int64, bool) {
	a.mu.Lock()
	defer a.mu.Unlock()
	if a.open {
		return a.money, true
	} else {
		return 0, false
	}
}

func (a *Account) Deposit(amount int64) (int64, bool) {
	a.mu.Lock()
	defer a.mu.Unlock()
	if !a.open || a.money + amount < 0 {
		return 0, false
	} else {
		a.money += amount
		return a.money, true
	}
}

func (a *Account) Close() (int64, bool) {
	a.mu.Lock()
	defer a.mu.Unlock()
	if a.open {
		a.open = false
		money := a.money
		a.money = 0
		return money, true
	} else {
		return 0, false
	}
}
