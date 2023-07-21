// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 Authors of kdoctor-io
// SPDX-License-Identifier: Apache-2.0

package echo

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewGetParams creates a new GetParams object
//
// There are no default values defined in the spec.
func NewGetParams() GetParams {

	return GetParams{}
}

// GetParams contains all the bound params for the get operation
// typically these are obtained from a http.Request
//
// swagger:parameters Get
type GetParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*delay some second return response
	  In: query
	*/
	Delay *int64
	/*task name
	  In: query
	*/
	Task *string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetParams() beforehand.
func (o *GetParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qDelay, qhkDelay, _ := qs.GetOK("delay")
	if err := o.bindDelay(qDelay, qhkDelay, route.Formats); err != nil {
		res = append(res, err)
	}

	qTask, qhkTask, _ := qs.GetOK("task")
	if err := o.bindTask(qTask, qhkTask, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindDelay binds and validates parameter Delay from query.
func (o *GetParams) bindDelay(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("delay", "query", "int64", raw)
	}
	o.Delay = &value

	return nil
}

// bindTask binds and validates parameter Task from query.
func (o *GetParams) bindTask(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}
	o.Task = &raw

	return nil
}
