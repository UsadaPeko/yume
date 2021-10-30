package domain

import "errors"

type Plan struct {}

func (p Plan) PutFund(amount uint32) error {
	// TODO: 추후에 구현될 내용.
	return errors.New("Plan이 아직 설정되지 않아서 자금을 계산할 수 없습니다.")
}
