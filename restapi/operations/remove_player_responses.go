// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "98point6DropToken/models"
)

// RemovePlayerAcceptedCode is the HTTP code returned for type RemovePlayerAccepted
const RemovePlayerAcceptedCode int = 202

/*RemovePlayerAccepted OK

swagger:response removePlayerAccepted
*/
type RemovePlayerAccepted struct {
}

// NewRemovePlayerAccepted creates RemovePlayerAccepted with default headers values
func NewRemovePlayerAccepted() *RemovePlayerAccepted {

	return &RemovePlayerAccepted{}
}

// WriteResponse to the client
func (o *RemovePlayerAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(202)
}

// RemovePlayerNotFoundCode is the HTTP code returned for type RemovePlayerNotFound
const RemovePlayerNotFoundCode int = 404

/*RemovePlayerNotFound Game/player not found

swagger:response removePlayerNotFound
*/
type RemovePlayerNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.MalformedRequest `json:"body,omitempty"`
}

// NewRemovePlayerNotFound creates RemovePlayerNotFound with default headers values
func NewRemovePlayerNotFound() *RemovePlayerNotFound {

	return &RemovePlayerNotFound{}
}

// WithPayload adds the payload to the remove player not found response
func (o *RemovePlayerNotFound) WithPayload(payload *models.MalformedRequest) *RemovePlayerNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the remove player not found response
func (o *RemovePlayerNotFound) SetPayload(payload *models.MalformedRequest) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RemovePlayerNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// RemovePlayerGoneCode is the HTTP code returned for type RemovePlayerGone
const RemovePlayerGoneCode int = 410

/*RemovePlayerGone Game is already DONE

swagger:response removePlayerGone
*/
type RemovePlayerGone struct {

	/*
	  In: Body
	*/
	Payload *models.MalformedRequest `json:"body,omitempty"`
}

// NewRemovePlayerGone creates RemovePlayerGone with default headers values
func NewRemovePlayerGone() *RemovePlayerGone {

	return &RemovePlayerGone{}
}

// WithPayload adds the payload to the remove player gone response
func (o *RemovePlayerGone) WithPayload(payload *models.MalformedRequest) *RemovePlayerGone {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the remove player gone response
func (o *RemovePlayerGone) SetPayload(payload *models.MalformedRequest) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RemovePlayerGone) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(410)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}