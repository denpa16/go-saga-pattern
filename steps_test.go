package saga

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddStep(t *testing.T) {
	testSagaName := "testSaga"
	testStepName := "testStep"

	executeFunc := func(ctx context.Context) error {
		return nil
	}

	compensateFunc := func(ctx context.Context) error {
		return nil
	}

	saga := NewSaga(testSagaName)
	saga.AddStep(
		testStepName,
		executeFunc,
		compensateFunc,
	)

	assert.Equal(t, len(saga.steps), 1)
	assert.Equal(t, saga.steps[0].Name, testStepName)
	assert.Equal(t, saga.steps[0].Compensate(context.Background()), compensateFunc(context.Background()))
	assert.Equal(t, saga.steps[0].Execute(context.Background()), executeFunc(context.Background()))
}
