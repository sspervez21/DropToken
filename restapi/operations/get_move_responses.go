// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "98point6DropToken/models"
)

// GetMoveOKCode is the HTTP code returned for type GetMoveOK
const GetMoveOKCode int = 200

/*GetMoveOK OK

swagger:response getMoveOK
*/
type GetMoveOK struct {

	/*
	  In: Body
	*/
	Payload *models.Move `json:"body,omitempty"`
}

// NewGetMoveOK creates GetMoveOK with default headers values
func NewGetMoveOK() *GetMoveOK {

	return &GetMoveOK{}
}

// WithPayload adds the payload to the get move o k response
func (o *GetMoveOK) WithPayload(payload *models.Move) *GetMoveOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get move o k response
func (o *GetMoveOK) SetPayload(payload *models.Move) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetMoveOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetMoveBadRequestCode is the HTTP code returned for type GetMoveBadRequest
const GetMoveBadRequestCode int = 400

/*GetMoveBadRequest Malformed Request.

swagger:response getMoveBadRequest
*/
type GetMoveBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.MalformedRequest `json:"body,omitempty"`
}

// NewGetMoveBadRequest creates GetMoveBadRequest with default headers values
func NewGetMoveBadRequest() *GetMoveBadRequest {

	return &GetMoveBadRequest{}
}

// WithPayload adds the payload to the get move bad request response
func (o *GetMoveBadRequest) WithPayload(payload *models.MalformedRequest) *GetMoveBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get move bad request response
func (o *GetMoveBadRequest) SetPayload(payload *models.MalformedRequest) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetMoveBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetMoveNotFoundCode is the HTTP code returned for type GetMoveNotFound
const GetMoveNotFoundCode int = 404

/*GetMoveNotFound Game/Move not found

swagger:response getMoveNotFound
*/
type GetMoveNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.MalformedRequest `json:"body,omitempty"`
}

// NewGetMoveNotFound creates GetMoveNotFound with default headers values
func NewGetMoveNotFound() *GetMoveNotFound {

	return &GetMoveNotFound{}
}

// WithPayload adds the payload to the get move not found response
func (o *GetMoveNotFound) WithPayload(payload *models.MalformedRequest) *GetMoveNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get move not found response
func (o *GetMoveNotFound) SetPayload(payload *models.MalformedRequest) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetMoveNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}