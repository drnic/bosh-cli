package cmd_test

import (
	. "github.com/onsi/ginkgo"
	// . "github.com/onsi/gomega"

	. "github.com/cloudfoundry/bosh-cli/cmd"
	fakecmdconf "github.com/cloudfoundry/bosh-cli/cmd/config/configfakes"
	fakeui "github.com/cloudfoundry/bosh-cli/ui/fakes"
)

var _ = Describe("OAuthTokenCmd", func() {
	var (
		config  *fakecmdconf.FakeConfig
		ui      *fakeui.FakeUI
		command OAuthTokenCmd
	)

	BeforeEach(func() {
		config = &fakecmdconf.FakeConfig{}
		ui = &fakeui.FakeUI{}
		command = NewOAuthTokenCmd("environment", config, ui)
	})

	Describe("Run", func() {
		var (
			updatedConfig *fakecmdconf.FakeConfig
		)

		BeforeEach(func() {
			updatedConfig = &fakecmdconf.FakeConfig{}
			config.UnsetCredentialsReturns(updatedConfig)
		})

		// act := func() error { return command.Run() }

		It("raises error if basic auth", func() {
		})

		It("generates an access token if uaa/client/credentials", func() {
		})

		It("generates new access token if uaa/user/password", func() {
		})

	})
})
