package domain

type PlanFactory struct {}

func (p PlanFactory) NewEmptyPlan() Plan {
	return Plan{
		defaultFundSource: NewFundSource(),
	}
}
