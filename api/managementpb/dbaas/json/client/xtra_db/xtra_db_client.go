// Code generated by go-swagger; DO NOT EDIT.

package xtra_db

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new xtra db API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for xtra db API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	CreateXtraDBCluster(params *CreateXtraDBClusterParams) (*CreateXtraDBClusterOK, error)

	DeleteXtraDBCluster(params *DeleteXtraDBClusterParams) (*DeleteXtraDBClusterOK, error)

	ListXtraDBClusters(params *ListXtraDBClustersParams) (*ListXtraDBClustersOK, error)

	UpdateXtraDBCluster(params *UpdateXtraDBClusterParams) (*UpdateXtraDBClusterOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateXtraDBCluster creates xtra DB cluster creates a new xtra DB cluster
*/
func (a *Client) CreateXtraDBCluster(params *CreateXtraDBClusterParams) (*CreateXtraDBClusterOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateXtraDBClusterParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CreateXtraDBCluster",
		Method:             "POST",
		PathPattern:        "/v1/management/DBaaS/XtraDB/Create",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateXtraDBClusterReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateXtraDBClusterOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreateXtraDBClusterDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  DeleteXtraDBCluster deletes xtra DB cluster deletes xtra DB cluster
*/
func (a *Client) DeleteXtraDBCluster(params *DeleteXtraDBClusterParams) (*DeleteXtraDBClusterOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteXtraDBClusterParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteXtraDBCluster",
		Method:             "POST",
		PathPattern:        "/v1/management/DBaaS/XtraDB/Delete",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteXtraDBClusterReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteXtraDBClusterOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteXtraDBClusterDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ListXtraDBClusters lists xtra DB clusters returns a list of xtra DB clusters
*/
func (a *Client) ListXtraDBClusters(params *ListXtraDBClustersParams) (*ListXtraDBClustersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListXtraDBClustersParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListXtraDBClusters",
		Method:             "POST",
		PathPattern:        "/v1/management/DBaaS/XtraDB/List",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListXtraDBClustersReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListXtraDBClustersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListXtraDBClustersDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  UpdateXtraDBCluster updates xtra DB cluster updates existing xtra DB cluster
*/
func (a *Client) UpdateXtraDBCluster(params *UpdateXtraDBClusterParams) (*UpdateXtraDBClusterOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateXtraDBClusterParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "UpdateXtraDBCluster",
		Method:             "POST",
		PathPattern:        "/v1/management/DBaaS/XtraDB/Update",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpdateXtraDBClusterReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateXtraDBClusterOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UpdateXtraDBClusterDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
