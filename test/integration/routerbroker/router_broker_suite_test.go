package routerbroker_test

import (
	"github.com/fgiorgetti/qpid-dispatch-gollum/test"
	"testing"
)

func TestMain(m *testing.M) {
	test.CommonMainFunc(m)
}

// Just to illustrate the structure for this test suite
func TestRouterBroker(t *testing.T) {
	test.RunSpecs(t, "routerbroker","RouterBroker Suite")
}
