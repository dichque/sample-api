// Code generated by go-swagger; DO NOT EDIT.

package pet

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/dichque/sample-api/models"
)

// CreateCreatedCode is the HTTP code returned for type CreateCreated
const CreateCreatedCode int = 201

/*CreateCreated Pet Created

swagger:response createCreated
*/
type CreateCreated struct {

	/*
	  In: Body
	*/
	Payload *models.Pet `json:"body,omitempty"`
}

// NewCreateCreated creates CreateCreated with default headers values
func NewCreateCreated() *CreateCreated {

	return &CreateCreated{}
}

// WithPayload adds the payload to the create created response
func (o *CreateCreated) WithPayload(payload *models.Pet) *CreateCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create created response
func (o *CreateCreated) SetPayload(payload *models.Pet) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateBadRequestCode is the HTTP code returned for type CreateBadRequest
const CreateBadRequestCode int = 400

/*CreateBadRequest Bad Request

swagger:response createBadRequest
*/
type CreateBadRequest struct {
}

// NewCreateBadRequest creates CreateBadRequest with default headers values
func NewCreateBadRequest() *CreateBadRequest {

	return &CreateBadRequest{}
}

// WriteResponse to the client
func (o *CreateBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}
