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

// NewDcimConsoleServerPortsListParams creates a new DcimConsoleServerPortsListParams object
// with the default values initialized.
func NewDcimConsoleServerPortsListParams() *DcimConsoleServerPortsListParams {
	var ()
	return &DcimConsoleServerPortsListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDcimConsoleServerPortsListParamsWithTimeout creates a new DcimConsoleServerPortsListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDcimConsoleServerPortsListParamsWithTimeout(timeout time.Duration) *DcimConsoleServerPortsListParams {
	var ()
	return &DcimConsoleServerPortsListParams{

		timeout: timeout,
	}
}

// NewDcimConsoleServerPortsListParamsWithContext creates a new DcimConsoleServerPortsListParams object
// with the default values initialized, and the ability to set a context for a request
func NewDcimConsoleServerPortsListParamsWithContext(ctx context.Context) *DcimConsoleServerPortsListParams {
	var ()
	return &DcimConsoleServerPortsListParams{

		Context: ctx,
	}
}

// NewDcimConsoleServerPortsListParamsWithHTTPClient creates a new DcimConsoleServerPortsListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDcimConsoleServerPortsListParamsWithHTTPClient(client *http.Client) *DcimConsoleServerPortsListParams {
	var ()
	return &DcimConsoleServerPortsListParams{
		HTTPClient: client,
	}
}

/*DcimConsoleServerPortsListParams contains all the parameters to send to the API endpoint
for the dcim console server ports list operation typically these are written to a http.Request
*/
type DcimConsoleServerPortsListParams struct {

	/*Cabled*/
	Cabled *string
	/*ConnectionStatus*/
	ConnectionStatus *string
	/*Device*/
	Device *string
	/*DeviceID*/
	DeviceID *int64
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
	/*Tag*/
	Tag *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the dcim console server ports list params
func (o *DcimConsoleServerPortsListParams) WithTimeout(timeout time.Duration) *DcimConsoleServerPortsListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim console server ports list params
func (o *DcimConsoleServerPortsListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim console server ports list params
func (o *DcimConsoleServerPortsListParams) WithContext(ctx context.Context) *DcimConsoleServerPortsListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim console server ports list params
func (o *DcimConsoleServerPortsListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim console server ports list params
func (o *DcimConsoleServerPortsListParams) WithHTTPClient(client *http.Client) *DcimConsoleServerPortsListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim console server ports list params
func (o *DcimConsoleServerPortsListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCabled adds the cabled to the dcim console server ports list params
func (o *DcimConsoleServerPortsListParams) WithCabled(cabled *string) *DcimConsoleServerPortsListParams {
	o.SetCabled(cabled)
	return o
}

// SetCabled adds the cabled to the dcim console server ports list params
func (o *DcimConsoleServerPortsListParams) SetCabled(cabled *string) {
	o.Cabled = cabled
}

// WithConnectionStatus adds the connectionStatus to the dcim console server ports list params
func (o *DcimConsoleServerPortsListParams) WithConnectionStatus(connectionStatus *string) *DcimConsoleServerPortsListParams {
	o.SetConnectionStatus(connectionStatus)
	return o
}

// SetConnectionStatus adds the connectionStatus to the dcim console server ports list params
func (o *DcimConsoleServerPortsListParams) SetConnectionStatus(connectionStatus *string) {
	o.ConnectionStatus = connectionStatus
}

// WithDevice adds the device to the dcim console server ports list params
func (o *DcimConsoleServerPortsListParams) WithDevice(device *string) *DcimConsoleServerPortsListParams {
	o.SetDevice(device)
	return o
}

// SetDevice adds the device to the dcim console server ports list params
func (o *DcimConsoleServerPortsListParams) SetDevice(device *string) {
	o.Device = device
}

// WithDeviceID adds the deviceID to the dcim console server ports list params
func (o *DcimConsoleServerPortsListParams) WithDeviceID(deviceID *int64) *DcimConsoleServerPortsListParams {
	o.SetDeviceID(deviceID)
	return o
}

// SetDeviceID adds the deviceId to the dcim console server ports list params
func (o *DcimConsoleServerPortsListParams) SetDeviceID(deviceID *int64) {
	o.DeviceID = deviceID
}

// WithLimit adds the limit to the dcim console server ports list params
func (o *DcimConsoleServerPortsListParams) WithLimit(limit *int64) *DcimConsoleServerPortsListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the dcim console server ports list params
func (o *DcimConsoleServerPortsListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithName adds the name to the dcim console server ports list params
func (o *DcimConsoleServerPortsListParams) WithName(name *string) *DcimConsoleServerPortsListParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the dcim console server ports list params
func (o *DcimConsoleServerPortsListParams) SetName(name *string) {
	o.Name = name
}

// WithOffset adds the offset to the dcim console server ports list params
func (o *DcimConsoleServerPortsListParams) WithOffset(offset *int64) *DcimConsoleServerPortsListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the dcim console server ports list params
func (o *DcimConsoleServerPortsListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithQ adds the q to the dcim console server ports list params
func (o *DcimConsoleServerPortsListParams) WithQ(q *string) *DcimConsoleServerPortsListParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the dcim console server ports list params
func (o *DcimConsoleServerPortsListParams) SetQ(q *string) {
	o.Q = q
}

// WithTag adds the tag to the dcim console server ports list params
func (o *DcimConsoleServerPortsListParams) WithTag(tag *string) *DcimConsoleServerPortsListParams {
	o.SetTag(tag)
	return o
}

// SetTag adds the tag to the dcim console server ports list params
func (o *DcimConsoleServerPortsListParams) SetTag(tag *string) {
	o.Tag = tag
}

// WriteToRequest writes these params to a swagger request
func (o *DcimConsoleServerPortsListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Cabled != nil {

		// query param cabled
		var qrCabled string
		if o.Cabled != nil {
			qrCabled = *o.Cabled
		}
		qCabled := qrCabled
		if qCabled != "" {
			if err := r.SetQueryParam("cabled", qCabled); err != nil {
				return err
			}
		}

	}

	if o.ConnectionStatus != nil {

		// query param connection_status
		var qrConnectionStatus string
		if o.ConnectionStatus != nil {
			qrConnectionStatus = *o.ConnectionStatus
		}
		qConnectionStatus := qrConnectionStatus
		if qConnectionStatus != "" {
			if err := r.SetQueryParam("connection_status", qConnectionStatus); err != nil {
				return err
			}
		}

	}

	if o.Device != nil {

		// query param device
		var qrDevice string
		if o.Device != nil {
			qrDevice = *o.Device
		}
		qDevice := qrDevice
		if qDevice != "" {
			if err := r.SetQueryParam("device", qDevice); err != nil {
				return err
			}
		}

	}

	if o.DeviceID != nil {

		// query param device_id
		var qrDeviceID int64
		if o.DeviceID != nil {
			qrDeviceID = *o.DeviceID
		}
		qDeviceID := swag.FormatInt64(qrDeviceID)
		if qDeviceID != "" {
			if err := r.SetQueryParam("device_id", qDeviceID); err != nil {
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

	if o.Tag != nil {

		// query param tag
		var qrTag string
		if o.Tag != nil {
			qrTag = *o.Tag
		}
		qTag := qrTag
		if qTag != "" {
			if err := r.SetQueryParam("tag", qTag); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
