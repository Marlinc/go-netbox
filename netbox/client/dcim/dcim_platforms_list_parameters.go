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

// NewDcimPlatformsListParams creates a new DcimPlatformsListParams object
// with the default values initialized.
func NewDcimPlatformsListParams() *DcimPlatformsListParams {
	var ()
	return &DcimPlatformsListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDcimPlatformsListParamsWithTimeout creates a new DcimPlatformsListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDcimPlatformsListParamsWithTimeout(timeout time.Duration) *DcimPlatformsListParams {
	var ()
	return &DcimPlatformsListParams{

		timeout: timeout,
	}
}

// NewDcimPlatformsListParamsWithContext creates a new DcimPlatformsListParams object
// with the default values initialized, and the ability to set a context for a request
func NewDcimPlatformsListParamsWithContext(ctx context.Context) *DcimPlatformsListParams {
	var ()
	return &DcimPlatformsListParams{

		Context: ctx,
	}
}

// NewDcimPlatformsListParamsWithHTTPClient creates a new DcimPlatformsListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDcimPlatformsListParamsWithHTTPClient(client *http.Client) *DcimPlatformsListParams {
	var ()
	return &DcimPlatformsListParams{
		HTTPClient: client,
	}
}

/*DcimPlatformsListParams contains all the parameters to send to the API endpoint
for the dcim platforms list operation typically these are written to a http.Request
*/
type DcimPlatformsListParams struct {

	/*Limit
	  Number of results to return per page.

	*/
	Limit *int64
	/*Manufacturer*/
	Manufacturer *string
	/*ManufacturerID*/
	ManufacturerID *int64
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

// WithTimeout adds the timeout to the dcim platforms list params
func (o *DcimPlatformsListParams) WithTimeout(timeout time.Duration) *DcimPlatformsListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim platforms list params
func (o *DcimPlatformsListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim platforms list params
func (o *DcimPlatformsListParams) WithContext(ctx context.Context) *DcimPlatformsListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim platforms list params
func (o *DcimPlatformsListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim platforms list params
func (o *DcimPlatformsListParams) WithHTTPClient(client *http.Client) *DcimPlatformsListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim platforms list params
func (o *DcimPlatformsListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the dcim platforms list params
func (o *DcimPlatformsListParams) WithLimit(limit *int64) *DcimPlatformsListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the dcim platforms list params
func (o *DcimPlatformsListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithManufacturer adds the manufacturer to the dcim platforms list params
func (o *DcimPlatformsListParams) WithManufacturer(manufacturer *string) *DcimPlatformsListParams {
	o.SetManufacturer(manufacturer)
	return o
}

// SetManufacturer adds the manufacturer to the dcim platforms list params
func (o *DcimPlatformsListParams) SetManufacturer(manufacturer *string) {
	o.Manufacturer = manufacturer
}

// WithManufacturerID adds the manufacturerID to the dcim platforms list params
func (o *DcimPlatformsListParams) WithManufacturerID(manufacturerID *int64) *DcimPlatformsListParams {
	o.SetManufacturerID(manufacturerID)
	return o
}

// SetManufacturerID adds the manufacturerId to the dcim platforms list params
func (o *DcimPlatformsListParams) SetManufacturerID(manufacturerID *int64) {
	o.ManufacturerID = manufacturerID
}

// WithName adds the name to the dcim platforms list params
func (o *DcimPlatformsListParams) WithName(name *string) *DcimPlatformsListParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the dcim platforms list params
func (o *DcimPlatformsListParams) SetName(name *string) {
	o.Name = name
}

// WithOffset adds the offset to the dcim platforms list params
func (o *DcimPlatformsListParams) WithOffset(offset *int64) *DcimPlatformsListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the dcim platforms list params
func (o *DcimPlatformsListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithQ adds the q to the dcim platforms list params
func (o *DcimPlatformsListParams) WithQ(q *string) *DcimPlatformsListParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the dcim platforms list params
func (o *DcimPlatformsListParams) SetQ(q *string) {
	o.Q = q
}

// WithSlug adds the slug to the dcim platforms list params
func (o *DcimPlatformsListParams) WithSlug(slug *string) *DcimPlatformsListParams {
	o.SetSlug(slug)
	return o
}

// SetSlug adds the slug to the dcim platforms list params
func (o *DcimPlatformsListParams) SetSlug(slug *string) {
	o.Slug = slug
}

// WriteToRequest writes these params to a swagger request
func (o *DcimPlatformsListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Manufacturer != nil {

		// query param manufacturer
		var qrManufacturer string
		if o.Manufacturer != nil {
			qrManufacturer = *o.Manufacturer
		}
		qManufacturer := qrManufacturer
		if qManufacturer != "" {
			if err := r.SetQueryParam("manufacturer", qManufacturer); err != nil {
				return err
			}
		}

	}

	if o.ManufacturerID != nil {

		// query param manufacturer_id
		var qrManufacturerID int64
		if o.ManufacturerID != nil {
			qrManufacturerID = *o.ManufacturerID
		}
		qManufacturerID := swag.FormatInt64(qrManufacturerID)
		if qManufacturerID != "" {
			if err := r.SetQueryParam("manufacturer_id", qManufacturerID); err != nil {
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
