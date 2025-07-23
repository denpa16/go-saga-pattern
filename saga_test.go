package saga

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewSaga(t *testing.T) {
	testSagaName := "testSaga"

	testSaga := Saga{Name: testSagaName}
	createdSaga := NewSaga(testSagaName)

	assert.Equal(t, testSaga.Name, createdSaga.Name)
}

func TestRun(t *testing.T) {
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

	err := saga.Run(context.Background())
	assert.Equal(t, err, nil)
}

func TestRunWithExecuteError(t *testing.T) {
	testFirstSagaStepName := "testFirstSagaStepName"
	testSecondSagaStepName := "testSecondSagaStepName"
	testStepName := "testStep"

	saga := NewSaga(testFirstSagaStepName)

	firstExecuteFunc := func(ctx context.Context) error {
		return nil
	}
	firstCompensateFunc := func(ctx context.Context) error {
		return nil
	}

	saga.AddStep(
		testStepName,
		firstExecuteFunc,
		firstCompensateFunc,
	)

	executeErr := errors.New("execute error")

	secondExecuteFunc := func(ctx context.Context) error {
		return executeErr
	}
	secondCompensateFunc := func(ctx context.Context) error {
		return nil
	}

	saga.AddStep(
		testSecondSagaStepName,
		secondExecuteFunc,
		secondCompensateFunc,
	)

	err := saga.Run(context.Background())
	assert.Equal(t, err, executeErr)
}

func TestRunWithCompensateError(t *testing.T) {
	testFirstSagaStepName := "testFirstSagaStepName"
	testSecondSagaStepName := "testSecondSagaStepName"
	testStepName := "testStep"

	saga := NewSaga(testFirstSagaStepName)

	compensateErr := errors.New("compensate error")

	firstExecuteFunc := func(ctx context.Context) error {
		return nil
	}
	firstCompensateFunc := func(ctx context.Context) error {
		return compensateErr
	}

	saga.AddStep(
		testStepName,
		firstExecuteFunc,
		firstCompensateFunc,
	)

	executeErr := errors.New("execute error")

	secondExecuteFunc := func(ctx context.Context) error {
		return executeErr
	}
	secondCompensateFunc := func(ctx context.Context) error {
		return nil
	}

	saga.AddStep(
		testSecondSagaStepName,
		secondExecuteFunc,
		secondCompensateFunc,
	)

	err := saga.Run(context.Background())
	assert.Equal(t, err, executeErr)
}
