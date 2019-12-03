package twointerior_test

import (
	"github.com/fgiorgetti/qpid-dispatch-gollum/test"
	"testing"
)

func TestMain(m *testing.M) {
	test.CommonMainFunc(m)
}

func TestTwoInterior(t *testing.T) {
	test.RunSpecs(t, "twointerior", "TwoInterior Suite")
}
