// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2018 The go-netbox Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package circuits

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewCircuitsCircuitTypesListParams creates a new CircuitsCircuitTypesListParams object
// with the default values initialized.
func NewCircuitsCircuitTypesListParams() *CircuitsCircuitTypesListParams {
	var ()
	return &CircuitsCircuitTypesListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCircuitsCircuitTypesListParamsWithTimeout creates a new CircuitsCircuitTypesListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCircuitsCircuitTypesListParamsWithTimeout(timeout time.Duration) *CircuitsCircuitTypesListParams {
	var ()
	return &CircuitsCircuitTypesListParams{

		timeout: timeout,
	}
}

// NewCircuitsCircuitTypesListParamsWithContext creates a new CircuitsCircuitTypesListParams object
// with the default values initialized, and the ability to set a context for a request
func NewCircuitsCircuitTypesListParamsWithContext(ctx context.Context) *CircuitsCircuitTypesListParams {
	var ()
	return &CircuitsCircuitTypesListParams{

		Context: ctx,
	}
}

// NewCircuitsCircuitTypesListParamsWithHTTPClient creates a new CircuitsCircuitTypesListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCircuitsCircuitTypesListParamsWithHTTPClient(client *http.Client) *CircuitsCircuitTypesListParams {
	var ()
	return &CircuitsCircuitTypesListParams{
		HTTPClient: client,
	}
}

/*CircuitsCircuitTypesListParams contains all the parameters to send to the API endpoint
for the circuits circuit types list operation typically these are written to a http.Request
*/
type CircuitsCircuitTypesListParams struct {

	/*Limit
	  Number of results to return per page.

	*/
	Limit *int64
	/*Name*/
	Name *string
	/*Offset
	  The initial index from which to return the results.

	*/
	Offset *int64
	/*Q*/
	Q *string
	/*Slug*/
	Slug *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the circuits circuit types list params
func (o *CircuitsCircuitTypesListParams) WithTimeout(timeout time.Duration) *CircuitsCircuitTypesListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the circuits circuit types list params
func (o *CircuitsCircuitTypesListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the circuits circuit types list params
func (o *CircuitsCircuitTypesListParams) WithContext(ctx context.Context) *CircuitsCircuitTypesListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the circuits circuit types list params
func (o *CircuitsCircuitTypesListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the circuits circuit types list params
func (o *CircuitsCircuitTypesListParams) WithHTTPClient(client *http.Client) *CircuitsCircuitTypesListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the circuits circuit types list params
func (o *CircuitsCircuitTypesListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the circuits circuit types list params
func (o *CircuitsCircuitTypesListParams) WithLimit(limit *int64) *CircuitsCircuitTypesListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the circuits circuit types list params
func (o *CircuitsCircuitTypesListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithName adds the name to the circuits circuit types list params
func (o *CircuitsCircuitTypesListParams) WithName(name *string) *CircuitsCircuitTypesListParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the circuits circuit types list params
func (o *CircuitsCircuitTypesListParams) SetName(name *string) {
	o.Name = name
}

// WithOffset adds the offset to the circuits circuit types list params
func (o *CircuitsCircuitTypesListParams) WithOffset(offset *int64) *CircuitsCircuitTypesListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the circuits circuit types list params
func (o *CircuitsCircuitTypesListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithQ adds the q to the circuits circuit types list params
func (o *CircuitsCircuitTypesListParams) WithQ(q *string) *CircuitsCircuitTypesListParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the circuits circuit types list params
func (o *CircuitsCircuitTypesListParams) SetQ(q *string) {
	o.Q = q
}

// WithSlug adds the slug to the circuits circuit types list params
func (o *CircuitsCircuitTypesListParams) WithSlug(slug *string) *CircuitsCircuitTypesListParams {
	o.SetSlug(slug)
	return o
}

// SetSlug adds the slug to the circuits circuit types list params
func (o *CircuitsCircuitTypesListParams) SetSlug(slug *string) {
	o.Slug = slug
}

// WriteToRequest writes these params to a swagger request
func (o *CircuitsCircuitTypesListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Limit != nil {

		// query param limit
		var qrLimit int64
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	if o.Name != nil {

		// query param name
		var qrName string
		if o.Name != nil {
			qrName = *o.Name
		}
		qName := qrName
		if qName != "" {
			if err := r.SetQueryParam("name", qName); err != nil {
				return err
			}
		}

	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int64
		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt64(qrOffset)
		if qOffset != "" {
			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}

	}

	if o.Q != nil {

		// query param q
		var qrQ string
		if o.Q != nil {
			qrQ = *o.Q
		}
		qQ := qrQ
		if qQ != "" {
			if err := r.SetQueryParam("q", qQ); err != nil {
				return err
			}
		}

	}

	if o.Slug != nil {

		// query param slug
		var qrSlug string
		if o.Slug != nil {
			qrSlug = *o.Slug
		}
		qSlug := qrSlug
		if qSlug != "" {
			if err := r.SetQueryParam("slug", qSlug); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
