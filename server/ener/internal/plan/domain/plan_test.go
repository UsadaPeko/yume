package domain_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"pekora.dev/yume/ener/internal/plan/domain"
)

var _ = Describe("Plan이", func() {
	Describe("생성되는 경우", func() {
		planFactory := domain.PlanFactory{}

		Context("아무런 데이터가 주어지지 않았다면", func() {
			emptyPlan := planFactory.NewEmptyPlan()

			It("기본적인 비어있는 Plan이 만들어진다.", func() {
				Expect(emptyPlan).NotTo(Equal(nil))
			})
			It("자금을 넣을 수 없다", func() {
				err := emptyPlan.PutFund(10)
				Expect(err).ShouldNot(BeNil())
			})
		})

	})
})
