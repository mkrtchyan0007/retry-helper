package retry

import (
	"errors"
	"testing"
	"time"
)

func TestDo(t *testing.T) {
	count := 0
	err := Do(3, time.Millisecond, func() error {
		count++
		if count < 2 {
			return errors.New("fail")
		}
		return nil
	})

	if err != nil || count != 2 {
		t.Fatal("retry failed")
	}
}
