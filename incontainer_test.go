package incontainer_test

import (
	"os"
	"testing"

	"github.com/calmh/incontainer"
)

func TestIncontainer(t *testing.T) {
	result := incontainer.Detect()
	if os.Getenv("IN_DOCKER_TEST") == "" {
		t.Skipf("not running in test harness (result: %v)", result)
	}
	if result != true {
		t.Errorf("expected incontainer.Detect() to return true, got %v", result)
	}
}
