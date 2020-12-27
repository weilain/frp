package basic

import (
	"fmt"

	"github.com/fatedier/frp/test/e2e/framework"
	"github.com/fatedier/frp/test/e2e/framework/consts"

	. "github.com/onsi/ginkgo"
)

var _ = Describe("[Feature: Server Manager]", func() {
	f := framework.NewDefaultFramework()

	It("Ports Whitelist", func() {
		// TODO
		serverConf := consts.DefaultServerConfig
		clientConf := consts.DefaultClientConfig

		clientConf += fmt.Sprintf(`
			[tcp]
			type = tcp
			local_port = {{ .%s }}
			remote_port = {{ .%s }}
			`, framework.TCPEchoServerPort, framework.GenPortName("TCP"))

		f.RunProcesses([]string{serverConf}, []string{clientConf})

		framework.ExpectTCPRequest(f.UsedPorts[framework.GenPortName("TCP")], []byte(consts.TestString), []byte(consts.TestString), connTimeout)

	})
})
