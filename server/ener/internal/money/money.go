package money

type Money struct{
	amount uint32
}

func (m Money) Minus(m2 Money) {
	m.amount -= m2.amount
}

func (m Money) Plus(m2 Money) {
	m.amount += m2.amount
}

func New(amount uint32) Money {
	return Money{amount}
}

func Zero() Money {
	return Money{amount: 0}
}
