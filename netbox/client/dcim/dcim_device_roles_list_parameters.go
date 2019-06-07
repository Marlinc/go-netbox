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

// NewDcimDeviceRolesListParams creates a new DcimDeviceRolesListParams object
// with the default values initialized.
func NewDcimDeviceRolesListParams() *DcimDeviceRolesListParams {
	var ()
	return &DcimDeviceRolesListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDcimDeviceRolesListParamsWithTimeout creates a new DcimDeviceRolesListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDcimDeviceRolesListParamsWithTimeout(timeout time.Duration) *DcimDeviceRolesListParams {
	var ()
	return &DcimDeviceRolesListParams{

		timeout: timeout,
	}
}

// NewDcimDeviceRolesListParamsWithContext creates a new DcimDeviceRolesListParams object
// with the default values initialized, and the ability to set a context for a request
func NewDcimDeviceRolesListParamsWithContext(ctx context.Context) *DcimDeviceRolesListParams {
	var ()
	return &DcimDeviceRolesListParams{

		Context: ctx,
	}
}

// NewDcimDeviceRolesListParamsWithHTTPClient creates a new DcimDeviceRolesListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDcimDeviceRolesListParamsWithHTTPClient(client *http.Client) *DcimDeviceRolesListParams {
	var ()
	return &DcimDeviceRolesListParams{
		HTTPClient: client,
	}
}

/*DcimDeviceRolesListParams contains all the parameters to send to the API endpoint
for the dcim device roles list operation typically these are written to a http.Request
*/
type DcimDeviceRolesListParams struct {

	/*Color*/
	Color *string
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
	/*VMRole*/
	VMRole *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the dcim device roles list params
func (o *DcimDeviceRolesListParams) WithTimeout(timeout time.Duration) *DcimDeviceRolesListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim device roles list params
func (o *DcimDeviceRolesListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim device roles list params
func (o *DcimDeviceRolesListParams) WithContext(ctx context.Context) *DcimDeviceRolesListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim device roles list params
func (o *DcimDeviceRolesListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim device roles list params
func (o *DcimDeviceRolesListParams) WithHTTPClient(client *http.Client) *DcimDeviceRolesListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim device roles list params
func (o *DcimDeviceRolesListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithColor adds the color to the dcim device roles list params
func (o *DcimDeviceRolesListParams) WithColor(color *string) *DcimDeviceRolesListParams {
	o.SetColor(color)
	return o
}

// SetColor adds the color to the dcim device roles list params
func (o *DcimDeviceRolesListParams) SetColor(color *string) {
	o.Color = color
}

// WithLimit adds the limit to the dcim device roles list params
func (o *DcimDeviceRolesListParams) WithLimit(limit *int64) *DcimDeviceRolesListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the dcim device roles list params
func (o *DcimDeviceRolesListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithName adds the name to the dcim device roles list params
func (o *DcimDeviceRolesListParams) WithName(name *string) *DcimDeviceRolesListParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the dcim device roles list params
func (o *DcimDeviceRolesListParams) SetName(name *string) {
	o.Name = name
}

// WithOffset adds the offset to the dcim device roles list params
func (o *DcimDeviceRolesListParams) WithOffset(offset *int64) *DcimDeviceRolesListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the dcim device roles list params
func (o *DcimDeviceRolesListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithQ adds the q to the dcim device roles list params
func (o *DcimDeviceRolesListParams) WithQ(q *string) *DcimDeviceRolesListParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the dcim device roles list params
func (o *DcimDeviceRolesListParams) SetQ(q *string) {
	o.Q = q
}

// WithSlug adds the slug to the dcim device roles list params
func (o *DcimDeviceRolesListParams) WithSlug(slug *string) *DcimDeviceRolesListParams {
	o.SetSlug(slug)
	return o
}

// SetSlug adds the slug to the dcim device roles list params
func (o *DcimDeviceRolesListParams) SetSlug(slug *string) {
	o.Slug = slug
}

// WithVMRole adds the vMRole to the dcim device roles list params
func (o *DcimDeviceRolesListParams) WithVMRole(vMRole *string) *DcimDeviceRolesListParams {
	o.SetVMRole(vMRole)
	return o
}

// SetVMRole adds the vmRole to the dcim device roles list params
func (o *DcimDeviceRolesListParams) SetVMRole(vMRole *string) {
	o.VMRole = vMRole
}

// WriteToRequest writes these params to a swagger request
func (o *DcimDeviceRolesListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Color != nil {

		// query param color
		var qrColor string
		if o.Color != nil {
			qrColor = *o.Color
		}
		qColor := qrColor
		if qColor != "" {
			if err := r.SetQueryParam("color", qColor); err != nil {
				return err
			}
		}

	}

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

	if o.VMRole != nil {

		// query param vm_role
		var qrVMRole string
		if o.VMRole != nil {
			qrVMRole = *o.VMRole
		}
		qVMRole := qrVMRole
		if qVMRole != "" {
			if err := r.SetQueryParam("vm_role", qVMRole); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
