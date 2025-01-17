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

package dcim

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

// NewDcimRegionsListParams creates a new DcimRegionsListParams object
// with the default values initialized.
func NewDcimRegionsListParams() *DcimRegionsListParams {
	var ()
	return &DcimRegionsListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDcimRegionsListParamsWithTimeout creates a new DcimRegionsListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDcimRegionsListParamsWithTimeout(timeout time.Duration) *DcimRegionsListParams {
	var ()
	return &DcimRegionsListParams{

		timeout: timeout,
	}
}

// NewDcimRegionsListParamsWithContext creates a new DcimRegionsListParams object
// with the default values initialized, and the ability to set a context for a request
func NewDcimRegionsListParamsWithContext(ctx context.Context) *DcimRegionsListParams {
	var ()
	return &DcimRegionsListParams{

		Context: ctx,
	}
}

// NewDcimRegionsListParamsWithHTTPClient creates a new DcimRegionsListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDcimRegionsListParamsWithHTTPClient(client *http.Client) *DcimRegionsListParams {
	var ()
	return &DcimRegionsListParams{
		HTTPClient: client,
	}
}

/*DcimRegionsListParams contains all the parameters to send to the API endpoint
for the dcim regions list operation typically these are written to a http.Request
*/
type DcimRegionsListParams struct {

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
	/*Parent*/
	Parent *string
	/*ParentID*/
	ParentID *int64
	/*Q*/
	Q *string
	/*Slug*/
	Slug *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the dcim regions list params
func (o *DcimRegionsListParams) WithTimeout(timeout time.Duration) *DcimRegionsListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim regions list params
func (o *DcimRegionsListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim regions list params
func (o *DcimRegionsListParams) WithContext(ctx context.Context) *DcimRegionsListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim regions list params
func (o *DcimRegionsListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim regions list params
func (o *DcimRegionsListParams) WithHTTPClient(client *http.Client) *DcimRegionsListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim regions list params
func (o *DcimRegionsListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the dcim regions list params
func (o *DcimRegionsListParams) WithLimit(limit *int64) *DcimRegionsListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the dcim regions list params
func (o *DcimRegionsListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithName adds the name to the dcim regions list params
func (o *DcimRegionsListParams) WithName(name *string) *DcimRegionsListParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the dcim regions list params
func (o *DcimRegionsListParams) SetName(name *string) {
	o.Name = name
}

// WithOffset adds the offset to the dcim regions list params
func (o *DcimRegionsListParams) WithOffset(offset *int64) *DcimRegionsListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the dcim regions list params
func (o *DcimRegionsListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithParent adds the parent to the dcim regions list params
func (o *DcimRegionsListParams) WithParent(parent *string) *DcimRegionsListParams {
	o.SetParent(parent)
	return o
}

// SetParent adds the parent to the dcim regions list params
func (o *DcimRegionsListParams) SetParent(parent *string) {
	o.Parent = parent
}

// WithParentID adds the parentID to the dcim regions list params
func (o *DcimRegionsListParams) WithParentID(parentID *int64) *DcimRegionsListParams {
	o.SetParentID(parentID)
	return o
}

// SetParentID adds the parentId to the dcim regions list params
func (o *DcimRegionsListParams) SetParentID(parentID *int64) {
	o.ParentID = parentID
}

// WithQ adds the q to the dcim regions list params
func (o *DcimRegionsListParams) WithQ(q *string) *DcimRegionsListParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the dcim regions list params
func (o *DcimRegionsListParams) SetQ(q *string) {
	o.Q = q
}

// WithSlug adds the slug to the dcim regions list params
func (o *DcimRegionsListParams) WithSlug(slug *string) *DcimRegionsListParams {
	o.SetSlug(slug)
	return o
}

// SetSlug adds the slug to the dcim regions list params
func (o *DcimRegionsListParams) SetSlug(slug *string) {
	o.Slug = slug
}

// WriteToRequest writes these params to a swagger request
func (o *DcimRegionsListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Parent != nil {

		// query param parent
		var qrParent string
		if o.Parent != nil {
			qrParent = *o.Parent
		}
		qParent := qrParent
		if qParent != "" {
			if err := r.SetQueryParam("parent", qParent); err != nil {
				return err
			}
		}

	}

	if o.ParentID != nil {

		// query param parent_id
		var qrParentID int64
		if o.ParentID != nil {
			qrParentID = *o.ParentID
		}
		qParentID := swag.FormatInt64(qrParentID)
		if qParentID != "" {
			if err := r.SetQueryParam("parent_id", qParentID); err != nil {
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
