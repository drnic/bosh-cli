package cmd

import (
	cmdconf "github.com/cloudfoundry/bosh-cli/cmd/config"
	"github.com/cloudfoundry/bosh-cli/director"
	biui "github.com/cloudfoundry/bosh-cli/ui"
)

type OAuthTokenCmd struct {
	sessionFactory func(cmdconf.Config) Session
	config         cmdconf.Config
	ui             biui.UI
}

func NewOAuthTokenCmd(sessionFactory func(cmdconf.Config) Session,
	config cmdconf.Config,
	ui biui.UI) OAuthTokenCmd {
	return OAuthTokenCmd{sessionFactory: sessionFactory, config: config, ui: ui}
}

func (c OAuthTokenCmd) Run() error {
	dir, err := c.sessionFactory(c.config).Director()
	if err != nil {
		return err
	}
	switch d := dir.(type) {
	case director.DirectorImpl:
		c.ui.PrintLinef("%#v\n", d)
	}
	return nil
}
