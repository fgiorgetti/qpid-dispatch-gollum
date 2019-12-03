package oneinterior_test

import (
	"github.com/fgiorgetti/qpid-dispatch-gollum/test"
	"testing"
)

func TestMain(m *testing.M) {
	test.CommonMainFunc(m)
}

func TestOneInterior(t *testing.T) {
	test.RunSpecs(t, "oneinterior", "OneInterior Suite")
}
