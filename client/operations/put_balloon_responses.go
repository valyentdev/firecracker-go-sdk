// Code generated by go-swagger; DO NOT EDIT.

// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
// 	http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/valyentdev/firecracker-go-sdk/client/models"
)

// PutBalloonReader is a Reader for the PutBalloon structure.
type PutBalloonReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutBalloonReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPutBalloonNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutBalloonBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewPutBalloonDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutBalloonNoContent creates a PutBalloonNoContent with default headers values
func NewPutBalloonNoContent() *PutBalloonNoContent {
	return &PutBalloonNoContent{}
}

/*PutBalloonNoContent handles this case with default header values.

Balloon device created/updated
*/
type PutBalloonNoContent struct {
}

func (o *PutBalloonNoContent) Error() string {
	return fmt.Sprintf("[PUT /balloon][%d] putBalloonNoContent ", 204)
}

func (o *PutBalloonNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutBalloonBadRequest creates a PutBalloonBadRequest with default headers values
func NewPutBalloonBadRequest() *PutBalloonBadRequest {
	return &PutBalloonBadRequest{}
}

/*PutBalloonBadRequest handles this case with default header values.

Balloon device cannot be created/updated due to bad input
*/
type PutBalloonBadRequest struct {
	Payload *models.Error
}

func (o *PutBalloonBadRequest) Error() string {
	return fmt.Sprintf("[PUT /balloon][%d] putBalloonBadRequest  %+v", 400, o.Payload)
}

func (o *PutBalloonBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutBalloonBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutBalloonDefault creates a PutBalloonDefault with default headers values
func NewPutBalloonDefault(code int) *PutBalloonDefault {
	return &PutBalloonDefault{
		_statusCode: code,
	}
}

/*PutBalloonDefault handles this case with default header values.

Internal server error
*/
type PutBalloonDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the put balloon default response
func (o *PutBalloonDefault) Code() int {
	return o._statusCode
}

func (o *PutBalloonDefault) Error() string {
	return fmt.Sprintf("[PUT /balloon][%d] putBalloon default  %+v", o._statusCode, o.Payload)
}

func (o *PutBalloonDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutBalloonDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
