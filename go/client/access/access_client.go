// Code generated by go-swagger; DO NOT EDIT.

package access

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new access API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for access API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteAccessApikeyNameorid(params *DeleteAccessApikeyNameoridParams, opts ...ClientOption) (*DeleteAccessApikeyNameoridOK, error)

	GetAccessApikey(params *GetAccessApikeyParams, opts ...ClientOption) (*GetAccessApikeyOK, error)

	GetAccessApikeyNameorid(params *GetAccessApikeyNameoridParams, opts ...ClientOption) (*GetAccessApikeyNameoridOK, error)

	PostAccessApikey(params *PostAccessApikeyParams, opts ...ClientOption) (*PostAccessApikeyOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
DeleteAccessApikeyNameorid Deletes an API Key given its `name` or `id`. If the user's session is tied to an organization, the API Key must be for that organization.
*/
func (a *Client) DeleteAccessApikeyNameorid(params *DeleteAccessApikeyNameoridParams, opts ...ClientOption) (*DeleteAccessApikeyNameoridOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteAccessApikeyNameoridParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteAccessApikeyNameorid",
		Method:             "DELETE",
		PathPattern:        "/access/apikey/{nameorid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteAccessApikeyNameoridReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteAccessApikeyNameoridOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteAccessApikeyNameorid: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetAccessApikey Lists all the API Keys for the authenticated user. If the user's session is tied to an organization, only the API Keys for that organization will be returned.
*/
func (a *Client) GetAccessApikey(params *GetAccessApikeyParams, opts ...ClientOption) (*GetAccessApikeyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAccessApikeyParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetAccessApikey",
		Method:             "GET",
		PathPattern:        "/access/apikey",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAccessApikeyReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetAccessApikeyOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetAccessApikey: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetAccessApikeyNameorid Gets information about a specific API Key given its `name` or `id`. If the user's session is tied to an organization, the API Key must be for that organization.
*/
func (a *Client) GetAccessApikeyNameorid(params *GetAccessApikeyNameoridParams, opts ...ClientOption) (*GetAccessApikeyNameoridOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAccessApikeyNameoridParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetAccessApikeyNameorid",
		Method:             "GET",
		PathPattern:        "/access/apikey/{nameorid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAccessApikeyNameoridReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetAccessApikeyNameoridOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetAccessApikeyNameorid: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PostAccessApikey Creates a new API Key with the given `name` scoped to all the organizations the user is a member or owner of. If the user's session is already tied to an organization, the new API Key will be scoped to that organization.
*/
func (a *Client) PostAccessApikey(params *PostAccessApikeyParams, opts ...ClientOption) (*PostAccessApikeyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostAccessApikeyParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostAccessApikey",
		Method:             "POST",
		PathPattern:        "/access/apikey",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostAccessApikeyReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostAccessApikeyOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostAccessApikey: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
