// Code generated by go-swagger; DO NOT EDIT.

package nodes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/percona/pmm/api/json/models"
)

// AddVirtualMachineNodeReader is a Reader for the AddVirtualMachineNode structure.
type AddVirtualMachineNodeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddVirtualMachineNodeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewAddVirtualMachineNodeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAddVirtualMachineNodeOK creates a AddVirtualMachineNodeOK with default headers values
func NewAddVirtualMachineNodeOK() *AddVirtualMachineNodeOK {
	return &AddVirtualMachineNodeOK{}
}

/*AddVirtualMachineNodeOK handles this case with default header values.

(empty)
*/
type AddVirtualMachineNodeOK struct {
	Payload *models.InventoryAddVirtualMachineNodeResponse
}

func (o *AddVirtualMachineNodeOK) Error() string {
	return fmt.Sprintf("[POST /v0/inventory/Nodes/AddVirtualMachine][%d] addVirtualMachineNodeOK  %+v", 200, o.Payload)
}

func (o *AddVirtualMachineNodeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InventoryAddVirtualMachineNodeResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}