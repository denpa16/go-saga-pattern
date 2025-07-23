# Package go-saga-pattern

## Библиотека, с помощью которой можно реализовать паттер Saga для распределенных систем

## Где можно использовать:
### когда необходимо откатить состояния систем в цепочке вызовов, если одна из вернула ошибку

## Примеры использования

```shell
func(s *Service) ProcessSomeAction() error {
    saga := NewSaga("service - method: ProcessSomeAction")
	var sagaState struct {
	  actionFirstResultDo int
	  actionSecondResultDo string
	  actionThirdResultDo bool
	}
	
    saga.AddStep("actionFirst",
        func(ctx context.Context) error {
            actionFirstResultDo, err := s.actionFirstDo()
            if err != nil {
              return err
            }
            sagaState.actionFirstResult = actionFirstResult
            return nil
        },
        func(ctx context.Context) error {
            _, err := s.actionFirstCompensate()
            if err != nil {
              return err
            }
            return nil
        },
    }
    
    saga.AddStep("actionSecond",
        func(ctx context.Context) error {
            actionSecondResultDo, err := s.actionSecondDo()
            if err != nil {
              return err
            }
            sagaState.actionSecondResult = actionSecondResult
            return nil
        },
        func(ctx context.Context) error {
            _, err := s.actionSecondCompensate()
            if err != nil {
              return err
            }
            return nil
        },
    }
    
    saga.AddStep("actionThird",
        func(ctx context.Context) error {
            actionThirdResultDo, err := s.actionThirdDo()
            if err != nil {
              return err
            }
            sagaState.actionThirdResult = actionThirdResult
            return nil
        },
        func(ctx context.Context) error {
            _, err := s.actionThirdCompensate()
            if err != nil {
              return err
            }
            return nil
        },
    }
    
    err = saga.Run(ctx)
    if err != nil {
        return nil
    }
    return nil
}
```