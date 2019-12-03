package test

import (
	"fmt"
	"github.com/rh-messaging/shipshape/pkg/framework"
	"github.com/rh-messaging/shipshape/pkg/framework/ginkgowrapper"
	"github.com/onsi/ginkgo"
	"github.com/onsi/ginkgo/config"
	"github.com/onsi/ginkgo/reporters"
	"github.com/onsi/gomega"
	"k8s.io/klog"
	"os"
	"path"
	"testing"
)


// CommonMainFunc must be called by each test suite in order to parse
// command line arguments.
func CommonMainFunc(m *testing.M) {
	framework.HandleFlags()
	gomega.RegisterFailHandler(ginkgowrapper.Fail)
	os.Exit(m.Run())
}

// RunSpecs will use Ginkgo to run your test specs, and eventually setup
// report dir accordingly. The uniqueId is used to help composing the generated
// JUnit file name (when --report-dir is specified when running your tests).
func RunSpecs(t *testing.T, uniqueId string, description string) {
	// If any ginkgoReporter has been defined, use them.
	if framework.TestContext.ReportDir != "" {
		ginkgo.RunSpecsWithDefaultAndCustomReporters(t, description, generateReporter(uniqueId))
	} else {
		ginkgo.RunSpecs(t, description)
	}
}

// generateReporter returns a slice of ginkgo.Reporter if reportDir has been provided
func generateReporter(uniqueId string) []ginkgo.Reporter {
	var ginkgoReporters []ginkgo.Reporter

	// If report dir specified, create it
	if framework.TestContext.ReportDir != "" {
		if err := os.MkdirAll(framework.TestContext.ReportDir, 0755); err != nil {
			klog.Errorf("Failed creating report directory: %v", err)
		} else {
			ginkgoReporters = append(ginkgoReporters, reporters.NewJUnitReporter(
				path.Join(framework.TestContext.ReportDir,
					fmt.Sprintf("junit_%v%s%02d.xml",
						framework.TestContext.ReportPrefix,
						uniqueId,
						config.GinkgoConfig.ParallelNode))))
		}
	}

	return ginkgoReporters
}

// Before suite validation setup (happens only once per test suite)
var _ = ginkgo.SynchronizedBeforeSuite(func() []byte {
	// Unique initialization (node 1 only)
	return nil
}, func(data []byte) {
	// Initialization for each parallel node
}, 10)

// After suite validation teardown (happens only once per test suite)
var _ = ginkgo.SynchronizedAfterSuite(func() {
	// All nodes tear down
}, func() {
	// Node1 only tear down
	framework.RunCleanupActions(framework.AfterEach)
	framework.RunCleanupActions(framework.AfterSuite)
}, 10)
