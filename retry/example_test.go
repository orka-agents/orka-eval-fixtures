package retry_test

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/orka-agents/orka-eval-fixtures/retry"
)

func ExampleDo() {
	calls := 0
	err := retry.Do(context.Background(), retry.Policy{MaxAttempts: 3, BaseDelay: time.Millisecond, MaxDelay: 4 * time.Millisecond},
		func(context.Context) error {
			calls++
			if calls < 3 {
				return errors.New("not yet")
			}
			return nil
		})
	fmt.Println(calls, err)
	// Output: 3 <nil>
}
