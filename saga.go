package saga

import (
	"context"
	"log"
)

type Saga struct {
	Name  string
	steps []Step
}

func NewSaga(name string) *Saga {
	return &Saga{
		Name: name,
	}
}

func (s *Saga) Run(ctx context.Context) error {
	var executedSteps []int

	for i, step := range s.steps {
		log.Printf("%s: executing saga step: %s", s.Name, step.Name)
		if err := step.Execute(ctx); err != nil {
			log.Printf("%s: failed executing saga step %s: %v", s.Name, step.Name, err)
			for j := len(executedSteps) - 1; j >= 0; j-- {
				stepIdx := executedSteps[j]
				compStep := s.steps[stepIdx]
				log.Printf("%s: compensating saga step: %s", s.Name, compStep.Name)
				if compStep.Compensate != nil {
					if compErr := compStep.Compensate(ctx); compErr != nil {
						log.Printf("%s: failed compensating saga step %s: %v", s.Name, compStep.Name, compErr)
					}
				}
			}
			return err
		}
		executedSteps = append(executedSteps, i)
	}
	return nil
}
