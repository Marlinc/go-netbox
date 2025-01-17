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

// NewDcimPowerPortTemplatesListParams creates a new DcimPowerPortTemplatesListParams object
// with the default values initialized.
func NewDcimPowerPortTemplatesListParams() *DcimPowerPortTemplatesListParams {
	var ()
	return &DcimPowerPortTemplatesListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDcimPowerPortTemplatesListParamsWithTimeout creates a new DcimPowerPortTemplatesListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDcimPowerPortTemplatesListParamsWithTimeout(timeout time.Duration) *DcimPowerPortTemplatesListParams {
	var ()
	return &DcimPowerPortTemplatesListParams{

		timeout: timeout,
	}
}

// NewDcimPowerPortTemplatesListParamsWithContext creates a new DcimPowerPortTemplatesListParams object
// with the default values initialized, and the ability to set a context for a request
func NewDcimPowerPortTemplatesListParamsWithContext(ctx context.Context) *DcimPowerPortTemplatesListParams {
	var ()
	return &DcimPowerPortTemplatesListParams{

		Context: ctx,
	}
}

// NewDcimPowerPortTemplatesListParamsWithHTTPClient creates a new DcimPowerPortTemplatesListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDcimPowerPortTemplatesListParamsWithHTTPClient(client *http.Client) *DcimPowerPortTemplatesListParams {
	var ()
	return &DcimPowerPortTemplatesListParams{
		HTTPClient: client,
	}
}

/*DcimPowerPortTemplatesListParams contains all the parameters to send to the API endpoint
for the dcim power port templates list operation typically these are written to a http.Request
*/
type DcimPowerPortTemplatesListParams struct {

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

// WithTimeout adds the timeout to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) WithTimeout(timeout time.Duration) *DcimPowerPortTemplatesListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) WithContext(ctx context.Context) *DcimPowerPortTemplatesListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) WithHTTPClient(client *http.Client) *DcimPowerPortTemplatesListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDevicetypeID adds the devicetypeID to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) WithDevicetypeID(devicetypeID *int64) *DcimPowerPortTemplatesListParams {
	o.SetDevicetypeID(devicetypeID)
	return o
}

// SetDevicetypeID adds the devicetypeId to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) SetDevicetypeID(devicetypeID *int64) {
	o.DevicetypeID = devicetypeID
}

// WithLimit adds the limit to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) WithLimit(limit *int64) *DcimPowerPortTemplatesListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithName adds the name to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) WithName(name *string) *DcimPowerPortTemplatesListParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) SetName(name *string) {
	o.Name = name
}

// WithOffset adds the offset to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) WithOffset(offset *int64) *DcimPowerPortTemplatesListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithQ adds the q to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) WithQ(q *string) *DcimPowerPortTemplatesListParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) SetQ(q *string) {
	o.Q = q
}

// WriteToRequest writes these params to a swagger request
func (o *DcimPowerPortTemplatesListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
