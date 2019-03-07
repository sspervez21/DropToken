// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "98point6DropToken/models"
)

// GetAllGamesOKCode is the HTTP code returned for type GetAllGamesOK
const GetAllGamesOKCode int = 200

/*GetAllGamesOK OK

swagger:response getAllGamesOK
*/
type GetAllGamesOK struct {

	/*
	  In: Body
	*/
	Payload *models.AllGames `json:"body,omitempty"`
}

// NewGetAllGamesOK creates GetAllGamesOK with default headers values
func NewGetAllGamesOK() *GetAllGamesOK {

	return &GetAllGamesOK{}
}

// WithPayload adds the payload to the get all games o k response
func (o *GetAllGamesOK) WithPayload(payload *models.AllGames) *GetAllGamesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get all games o k response
func (o *GetAllGamesOK) SetPayload(payload *models.AllGames) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAllGamesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
