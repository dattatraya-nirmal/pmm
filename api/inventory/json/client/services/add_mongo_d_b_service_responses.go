// Code generated by go-swagger; DO NOT EDIT.

package services

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// AddMongoDBServiceReader is a Reader for the AddMongoDBService structure.
type AddMongoDBServiceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddMongoDBServiceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewAddMongoDBServiceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAddMongoDBServiceOK creates a AddMongoDBServiceOK with default headers values
func NewAddMongoDBServiceOK() *AddMongoDBServiceOK {
	return &AddMongoDBServiceOK{}
}

/*AddMongoDBServiceOK handles this case with default header values.

A successful response.
*/
type AddMongoDBServiceOK struct {
	Payload *AddMongoDBServiceOKBody
}

func (o *AddMongoDBServiceOK) Error() string {
	return fmt.Sprintf("[POST /v1/inventory/Services/AddMongoDB][%d] addMongoDBServiceOK  %+v", 200, o.Payload)
}

func (o *AddMongoDBServiceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(AddMongoDBServiceOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*AddMongoDBServiceBody add mongo d b service body
swagger:model AddMongoDBServiceBody
*/
type AddMongoDBServiceBody struct {

	// Custom user-assigned labels. Keys must start with "_".
	CustomLabels map[string]string `json:"custom_labels,omitempty"`

	// Node identifier where this instance runs. Required.
	NodeID string `json:"node_id,omitempty"`

	// Unique across all Services user-defined name. Required.
	ServiceName string `json:"service_name,omitempty"`
}

// Validate validates this add mongo d b service body
func (o *AddMongoDBServiceBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AddMongoDBServiceBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddMongoDBServiceBody) UnmarshalBinary(b []byte) error {
	var res AddMongoDBServiceBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddMongoDBServiceOKBody add mongo d b service o k body
swagger:model AddMongoDBServiceOKBody
*/
type AddMongoDBServiceOKBody struct {

	// mongodb
	Mongodb *AddMongoDBServiceOKBodyMongodb `json:"mongodb,omitempty"`
}

// Validate validates this add mongo d b service o k body
func (o *AddMongoDBServiceOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMongodb(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AddMongoDBServiceOKBody) validateMongodb(formats strfmt.Registry) error {

	if swag.IsZero(o.Mongodb) { // not required
		return nil
	}

	if o.Mongodb != nil {
		if err := o.Mongodb.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("addMongoDBServiceOK" + "." + "mongodb")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *AddMongoDBServiceOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddMongoDBServiceOKBody) UnmarshalBinary(b []byte) error {
	var res AddMongoDBServiceOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddMongoDBServiceOKBodyMongodb MongoDBService represents a generic MongoDB instance.
swagger:model AddMongoDBServiceOKBodyMongodb
*/
type AddMongoDBServiceOKBodyMongodb struct {

	// Custom user-assigned labels. Keys starts with "_".
	CustomLabels map[string]string `json:"custom_labels,omitempty"`

	// Node identifier where this instance runs.
	NodeID string `json:"node_id,omitempty"`

	// Unique randomly generated instance identifier.
	ServiceID string `json:"service_id,omitempty"`

	// Unique across all Services user-defined name.
	ServiceName string `json:"service_name,omitempty"`
}

// Validate validates this add mongo d b service o k body mongodb
func (o *AddMongoDBServiceOKBodyMongodb) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AddMongoDBServiceOKBodyMongodb) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddMongoDBServiceOKBodyMongodb) UnmarshalBinary(b []byte) error {
	var res AddMongoDBServiceOKBodyMongodb
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}