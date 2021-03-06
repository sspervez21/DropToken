// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "98point6DropToken/models"
)

// GetGameOKCode is the HTTP code returned for type GetGameOK
const GetGameOKCode int = 200

/*GetGameOK OK

swagger:response getGameOK
*/
type GetGameOK struct {

	/*
	  In: Body
	*/
	Payload *models.GameState `json:"body,omitempty"`
}

// NewGetGameOK creates GetGameOK with default headers values
func NewGetGameOK() *GetGameOK {

	return &GetGameOK{}
}

// WithPayload adds the payload to the get game o k response
func (o *GetGameOK) WithPayload(payload *models.GameState) *GetGameOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get game o k response
func (o *GetGameOK) SetPayload(payload *models.GameState) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetGameOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetGameBadRequestCode is the HTTP code returned for type GetGameBadRequest
const GetGameBadRequestCode int = 400

/*GetGameBadRequest Malformed Request

swagger:response getGameBadRequest
*/
type GetGameBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.MalformedRequest `json:"body,omitempty"`
}

// NewGetGameBadRequest creates GetGameBadRequest with default headers values
func NewGetGameBadRequest() *GetGameBadRequest {

	return &GetGameBadRequest{}
}

// WithPayload adds the payload to the get game bad request response
func (o *GetGameBadRequest) WithPayload(payload *models.MalformedRequest) *GetGameBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get game bad request response
func (o *GetGameBadRequest) SetPayload(payload *models.MalformedRequest) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetGameBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetGameNotFoundCode is the HTTP code returned for type GetGameNotFound
const GetGameNotFoundCode int = 404

/*GetGameNotFound Game not found

swagger:response getGameNotFound
*/
type GetGameNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.MalformedRequest `json:"body,omitempty"`
}

// NewGetGameNotFound creates GetGameNotFound with default headers values
func NewGetGameNotFound() *GetGameNotFound {

	return &GetGameNotFound{}
}

// WithPayload adds the payload to the get game not found response
func (o *GetGameNotFound) WithPayload(payload *models.MalformedRequest) *GetGameNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get game not found response
func (o *GetGameNotFound) SetPayload(payload *models.MalformedRequest) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetGameNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
