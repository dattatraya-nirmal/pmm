// Code generated by go-swagger; DO NOT EDIT.

package rules

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new rules API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for rules API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	ChangeAlertingRules(params *ChangeAlertingRulesParams) (*ChangeAlertingRulesOK, error)

	ListAlertingRules(params *ListAlertingRulesParams) (*ListAlertingRulesOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  ChangeAlertingRules changes alerting rules changes integrated alerting rule
*/
func (a *Client) ChangeAlertingRules(params *ChangeAlertingRulesParams) (*ChangeAlertingRulesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewChangeAlertingRulesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ChangeAlertingRules",
		Method:             "POST",
		PathPattern:        "/v1/management/ia/Rules/Change",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ChangeAlertingRulesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ChangeAlertingRulesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ChangeAlertingRulesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ListAlertingRules lists alerting rules returns a list of all integrated alerting rules
*/
func (a *Client) ListAlertingRules(params *ListAlertingRulesParams) (*ListAlertingRulesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListAlertingRulesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListAlertingRules",
		Method:             "POST",
		PathPattern:        "/v1/management/ia/Rules/List",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListAlertingRulesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListAlertingRulesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListAlertingRulesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}