package twointerior

import (
	"github.com/rh-messaging/shipshape/pkg/framework"
	"github.com/rh-messaging/shipshape/pkg/validation/qpiddispatch/management"
	"github.com/onsi/ginkgo"
)

/**
Validates the formed mesh
*/
var _ = ginkgo.Describe("Validates the formed mesh", func() {

	var (
		ctx1 *framework.ContextData
		ctx2 *framework.ContextData
	)

	// Initialize after frameworks have been created
	ginkgo.JustBeforeEach(func() {
		ctx1 = FrameworkQdrOne.GetFirstContext()
		ctx2 = FrameworkQdrTwo.GetFirstContext()
	})

	ginkgo.It("Query routers in the network on each pod", func() {
		management.ValidateRoutersInNetwork(ctx1, QdrOneName, 2)
		management.ValidateRoutersInNetwork(ctx2, QdrTwoName, 2)

	})
})
