package domain

import (
	"pekora.dev/yume/ener/internal/money"
)

type Plan struct {
	defaultFundSource FundSource
}

func (p Plan) PutFund(amount money.Money) {
	p.defaultFundSource.Put(amount)
}

type FundSource struct {
	id      string
	account DepositAccount
}

func NewFundSource() FundSource {
	return FundSource{
		id: "uuid",
		account: NewDepositAccount(),
	}
}

func (f FundSource) GetBalance() money.Money {
	return f.account.GetBalance()
}

func (f FundSource) Remit(m money.Money) {
	f.account.Remit(m)
}

func (f FundSource) AcceptRemittance(m money.Money) {
	f.account.AcceptRemittance(m)
}

func (f FundSource) Put(m money.Money) {
	f.AcceptRemittance(m)
}
