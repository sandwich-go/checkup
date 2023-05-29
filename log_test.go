package checkup

import (
	"fmt"
	"testing"
)

func TestLogError(t *testing.T) {
	New()
	LogError(fmt.Errorf("test error"))
}
