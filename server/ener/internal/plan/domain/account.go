package domain

import "pekora.dev/yume/ener/internal/money"

// SavingAccount 는 저금용 계좌입니다.
type SavingAccount struct {
	id      string
	balance money.Money

	next Flow
}

// DepositAccount 는 예금용 계좌입니다. 사용자가 돈을
// 빼내어 쉽게 사용할 수 있습니다.
type DepositAccount struct {
	id      string
	balance money.Money

	next Flow
}

func NewDepositAccount() DepositAccount {
	account := DepositAccount{
		id: "id",
		balance: money.Zero(),
	}
	account.next = MakeFlow(account)
	return account
}

func (d DepositAccount) GetBalance() money.Money {
	return d.balance
}

func (d DepositAccount) Remit(m money.Money) {
	d.balance.Minus(m)
}

func (d DepositAccount) AcceptRemittance(m money.Money) {
	d.balance.Plus(m)
	if !d.next.CanFlow() {
		return
	}
	d.next.Flow()
}

func (d DepositAccount) IsReal() bool {
	return true
}

type VirtualAccount struct{}

func (v VirtualAccount) GetBalance() money.Money {
	panic("implement me")
}

func (v VirtualAccount) Remit(m money.Money) {
	panic("implement me")
}

func (v VirtualAccount) AcceptRemittance(m money.Money) {
	panic("implement me")
}

func (v VirtualAccount) IsReal() bool {
	return false
}
