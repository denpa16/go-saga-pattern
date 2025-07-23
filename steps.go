package saga

import "context"

type Step struct {
	Name       string
	Execute    func(ctx context.Context) error
	Compensate func(ctx context.Context) error
}

func (s *Saga) AddStep(name string, execute, compensate func(ctx context.Context) error) {
	s.steps = append(s.steps, Step{
		Name:       name,
		Execute:    execute,
		Compensate: compensate,
	})
}
