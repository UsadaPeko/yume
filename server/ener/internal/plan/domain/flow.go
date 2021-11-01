package domain

import "pekora.dev/yume/ener/internal/money"

// Account 는 계좌들이 갖고있는 공통의 인터페이스를 정의합니다.
type Account interface {
	// GetBalance 는 계좌의 잔액을 반환합니다.
	GetBalance() money.Money
	// Remit 은 돈을 송금하는 인터페이스입니다.
	Remit(money.Money)
	// AcceptRemittance 는 송금한 돈을 받는 인터페이스입니다.
	AcceptRemittance(money.Money)
	// IsReal 은 Optional한 상태에서 값이 있는지 없는지 확인합니다.
	IsReal() bool
}

// Flow 는 돈의 흐름을 재현합니다.
type Flow struct {
	upstream   Account
	downstream Account
}

func MakeFlow(upstream Account) Flow {
	return Flow{
		upstream:   upstream,
		downstream: VirtualAccount{},
	}
}

// CanFlow 는 흐름이 실제로 실행될 수 있는지 확인합니다.
func (f Flow) CanFlow() bool {
	if !f.downstream.IsReal() {
		return false
	}
	return true
}

func (f Flow) Flow() {
	if !f.CanFlow() {
		return
	}
	amount := f.upstream.GetBalance()
	f.upstream.Remit(amount)
	f.downstream.AcceptRemittance(amount)
}
