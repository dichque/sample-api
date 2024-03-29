// Code generated by go-swagger; DO NOT EDIT.

package pet

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/dichque/sample-api/models"
)

// GetOKCode is the HTTP code returned for type GetOK
const GetOKCode int = 200

/*GetOK Pet get

swagger:response getOK
*/
type GetOK struct {

	/*
	  In: Body
	*/
	Payload *models.Pet `json:"body,omitempty"`
}

// NewGetOK creates GetOK with default headers values
func NewGetOK() *GetOK {

	return &GetOK{}
}

// WithPayload adds the payload to the get o k response
func (o *GetOK) WithPayload(payload *models.Pet) *GetOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get o k response
func (o *GetOK) SetPayload(payload *models.Pet) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetBadRequestCode is the HTTP code returned for type GetBadRequest
const GetBadRequestCode int = 400

/*GetBadRequest Bad Request

swagger:response getBadRequest
*/
type GetBadRequest struct {
}

// NewGetBadRequest creates GetBadRequest with default headers values
func NewGetBadRequest() *GetBadRequest {

	return &GetBadRequest{}
}

// WriteResponse to the client
func (o *GetBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// GetNotFoundCode is the HTTP code returned for type GetNotFound
const GetNotFoundCode int = 404

/*GetNotFound Pet Not Found

swagger:response getNotFound
*/
type GetNotFound struct {
}

// NewGetNotFound creates GetNotFound with default headers values
func NewGetNotFound() *GetNotFound {

	return &GetNotFound{}
}

// WriteResponse to the client
func (o *GetNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}
