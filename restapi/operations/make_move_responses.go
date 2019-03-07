// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "98point6DropToken/models"
)

// MakeMoveOKCode is the HTTP code returned for type MakeMoveOK
const MakeMoveOKCode int = 200

/*MakeMoveOK OK

swagger:response makeMoveOK
*/
type MakeMoveOK struct {

	/*
	  In: Body
	*/
	Payload *models.MoveOutput `json:"body,omitempty"`
}

// NewMakeMoveOK creates MakeMoveOK with default headers values
func NewMakeMoveOK() *MakeMoveOK {

	return &MakeMoveOK{}
}

// WithPayload adds the payload to the make move o k response
func (o *MakeMoveOK) WithPayload(payload *models.MoveOutput) *MakeMoveOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the make move o k response
func (o *MakeMoveOK) SetPayload(payload *models.MoveOutput) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *MakeMoveOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// MakeMoveBadRequestCode is the HTTP code returned for type MakeMoveBadRequest
const MakeMoveBadRequestCode int = 400

/*MakeMoveBadRequest Malformed Request. Illegal move.

swagger:response makeMoveBadRequest
*/
type MakeMoveBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.MalformedRequest `json:"body,omitempty"`
}

// NewMakeMoveBadRequest creates MakeMoveBadRequest with default headers values
func NewMakeMoveBadRequest() *MakeMoveBadRequest {

	return &MakeMoveBadRequest{}
}

// WithPayload adds the payload to the make move bad request response
func (o *MakeMoveBadRequest) WithPayload(payload *models.MalformedRequest) *MakeMoveBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the make move bad request response
func (o *MakeMoveBadRequest) SetPayload(payload *models.MalformedRequest) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *MakeMoveBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// MakeMoveNotFoundCode is the HTTP code returned for type MakeMoveNotFound
const MakeMoveNotFoundCode int = 404

/*MakeMoveNotFound Game/player not found

swagger:response makeMoveNotFound
*/
type MakeMoveNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.MalformedRequest `json:"body,omitempty"`
}

// NewMakeMoveNotFound creates MakeMoveNotFound with default headers values
func NewMakeMoveNotFound() *MakeMoveNotFound {

	return &MakeMoveNotFound{}
}

// WithPayload adds the payload to the make move not found response
func (o *MakeMoveNotFound) WithPayload(payload *models.MalformedRequest) *MakeMoveNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the make move not found response
func (o *MakeMoveNotFound) SetPayload(payload *models.MalformedRequest) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *MakeMoveNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// MakeMoveConflictCode is the HTTP code returned for type MakeMoveConflict
const MakeMoveConflictCode int = 409

/*MakeMoveConflict Played out of turn

swagger:response makeMoveConflict
*/
type MakeMoveConflict struct {

	/*
	  In: Body
	*/
	Payload *models.MalformedRequest `json:"body,omitempty"`
}

// NewMakeMoveConflict creates MakeMoveConflict with default headers values
func NewMakeMoveConflict() *MakeMoveConflict {

	return &MakeMoveConflict{}
}

// WithPayload adds the payload to the make move conflict response
func (o *MakeMoveConflict) WithPayload(payload *models.MalformedRequest) *MakeMoveConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the make move conflict response
func (o *MakeMoveConflict) SetPayload(payload *models.MalformedRequest) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *MakeMoveConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}