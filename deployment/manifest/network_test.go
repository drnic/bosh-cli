package manifest_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/cloudfoundry/bosh-micro-cli/deployment/manifest"

	bmproperty "github.com/cloudfoundry/bosh-micro-cli/common/property"
)

var _ = Describe("Network", func() {
	var (
		network Network
	)

	BeforeEach(func() {
		network = Network{
			Name: "fake-name",
			Type: Dynamic,
			CloudProperties: bmproperty.Map{
				"subnet": bmproperty.Map{
					"name": "sg-1234",
				},
			},
			IP:      "1.2.3.4",
			Netmask: "255.255.255.0",
			Gateway: "1.2.3.1",
			DNS:     []string{"1.1.1.1"},
		}
	})

	Describe("Interface", func() {
		It("returns an interface that can be used to connect to the network", func() {
			Expect(network.Interface()).To(Equal(bmproperty.Map{
				"type":    "dynamic",
				"ip":      "1.2.3.4",
				"netmask": "255.255.255.0",
				"gateway": "1.2.3.1",
				"dns":     []string{"1.1.1.1"},
				"cloud_properties": bmproperty.Map{
					"subnet": bmproperty.Map{
						"name": "sg-1234",
					},
				},
			}))
		})

		It("leaves out missing optional fields", func() {
			network.Netmask = ""
			network.Gateway = ""
			network.DNS = []string{}

			Expect(network.Interface()).To(Equal(bmproperty.Map{
				"type": "dynamic",
				"ip":   "1.2.3.4",
				"cloud_properties": bmproperty.Map{
					"subnet": bmproperty.Map{
						"name": "sg-1234",
					},
				},
			}))
		})
	})
})