package fakes

import (
	bmagentclient "github.com/cloudfoundry/bosh-micro-cli/agentclient"
)

type FakeAgentClient struct {
	PingResponses   []pingResponse
	PingCalledCount int

	StopCalled bool
	stopErr    error

	ApplyApplySpec bmagentclient.ApplySpec
	ApplyErr       error
}

type pingResponse struct {
	response string
	err      error
}

func NewFakeAgentClient() *FakeAgentClient {
	return &FakeAgentClient{}
}

func (c *FakeAgentClient) Ping() (string, error) {
	c.PingCalledCount++

	if len(c.PingResponses) > 0 {
		response := c.PingResponses[0]
		c.PingResponses = c.PingResponses[1:]
		return response.response, response.err
	}

	return "", nil
}

func (c *FakeAgentClient) Stop() error {
	c.StopCalled = true
	return c.stopErr
}

func (c *FakeAgentClient) Apply(applySpec bmagentclient.ApplySpec) error {
	c.ApplyApplySpec = applySpec

	return c.ApplyErr
}

func (c *FakeAgentClient) SetPingBehavior(response string, err error) {
	c.PingResponses = append(c.PingResponses, pingResponse{
		response: response,
		err:      err,
	})
}

func (c *FakeAgentClient) SetStopBehavior(err error) {
	c.stopErr = err
}