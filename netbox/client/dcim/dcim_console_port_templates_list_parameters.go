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

// NewDcimConsolePortTemplatesListParams creates a new DcimConsolePortTemplatesListParams object
// with the default values initialized.
func NewDcimConsolePortTemplatesListParams() *DcimConsolePortTemplatesListParams {
	var ()
	return &DcimConsolePortTemplatesListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDcimConsolePortTemplatesListParamsWithTimeout creates a new DcimConsolePortTemplatesListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDcimConsolePortTemplatesListParamsWithTimeout(timeout time.Duration) *DcimConsolePortTemplatesListParams {
	var ()
	return &DcimConsolePortTemplatesListParams{

		timeout: timeout,
	}
}

// NewDcimConsolePortTemplatesListParamsWithContext creates a new DcimConsolePortTemplatesListParams object
// with the default values initialized, and the ability to set a context for a request
func NewDcimConsolePortTemplatesListParamsWithContext(ctx context.Context) *DcimConsolePortTemplatesListParams {
	var ()
	return &DcimConsolePortTemplatesListParams{

		Context: ctx,
	}
}

// NewDcimConsolePortTemplatesListParamsWithHTTPClient creates a new DcimConsolePortTemplatesListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDcimConsolePortTemplatesListParamsWithHTTPClient(client *http.Client) *DcimConsolePortTemplatesListParams {
	var ()
	return &DcimConsolePortTemplatesListParams{
		HTTPClient: client,
	}
}

/*DcimConsolePortTemplatesListParams contains all the parameters to send to the API endpoint
for the dcim console port templates list operation typically these are written to a http.Request
*/
type DcimConsolePortTemplatesListParams struct {

	/*DevicetypeID*/
	DevicetypeID *int64
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

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the dcim console port templates list params
func (o *DcimConsolePortTemplatesListParams) WithTimeout(timeout time.Duration) *DcimConsolePortTemplatesListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim console port templates list params
func (o *DcimConsolePortTemplatesListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim console port templates list params
func (o *DcimConsolePortTemplatesListParams) WithContext(ctx context.Context) *DcimConsolePortTemplatesListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim console port templates list params
func (o *DcimConsolePortTemplatesListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim console port templates list params
func (o *DcimConsolePortTemplatesListParams) WithHTTPClient(client *http.Client) *DcimConsolePortTemplatesListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim console port templates list params
func (o *DcimConsolePortTemplatesListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDevicetypeID adds the devicetypeID to the dcim console port templates list params
func (o *DcimConsolePortTemplatesListParams) WithDevicetypeID(devicetypeID *int64) *DcimConsolePortTemplatesListParams {
	o.SetDevicetypeID(devicetypeID)
	return o
}

// SetDevicetypeID adds the devicetypeId to the dcim console port templates list params
func (o *DcimConsolePortTemplatesListParams) SetDevicetypeID(devicetypeID *int64) {
	o.DevicetypeID = devicetypeID
}

// WithLimit adds the limit to the dcim console port templates list params
func (o *DcimConsolePortTemplatesListParams) WithLimit(limit *int64) *DcimConsolePortTemplatesListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the dcim console port templates list params
func (o *DcimConsolePortTemplatesListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithName adds the name to the dcim console port templates list params
func (o *DcimConsolePortTemplatesListParams) WithName(name *string) *DcimConsolePortTemplatesListParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the dcim console port templates list params
func (o *DcimConsolePortTemplatesListParams) SetName(name *string) {
	o.Name = name
}

// WithOffset adds the offset to the dcim console port templates list params
func (o *DcimConsolePortTemplatesListParams) WithOffset(offset *int64) *DcimConsolePortTemplatesListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the dcim console port templates list params
func (o *DcimConsolePortTemplatesListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithQ adds the q to the dcim console port templates list params
func (o *DcimConsolePortTemplatesListParams) WithQ(q *string) *DcimConsolePortTemplatesListParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the dcim console port templates list params
func (o *DcimConsolePortTemplatesListParams) SetQ(q *string) {
	o.Q = q
}

// WriteToRequest writes these params to a swagger request
func (o *DcimConsolePortTemplatesListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.DevicetypeID != nil {

		// query param devicetype_id
		var qrDevicetypeID int64
		if o.DevicetypeID != nil {
			qrDevicetypeID = *o.DevicetypeID
		}
		qDevicetypeID := swag.FormatInt64(qrDevicetypeID)
		if qDevicetypeID != "" {
			if err := r.SetQueryParam("devicetype_id", qDevicetypeID); err != nil {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
