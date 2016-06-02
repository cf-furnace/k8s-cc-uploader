package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

/*UploadBuildArtifactsCreated Artifacts have been uploaded

swagger:response uploadBuildArtifactsCreated
*/
type UploadBuildArtifactsCreated struct {
}

// NewUploadBuildArtifactsCreated creates UploadBuildArtifactsCreated with default headers values
func NewUploadBuildArtifactsCreated() *UploadBuildArtifactsCreated {
	return &UploadBuildArtifactsCreated{}
}

// WriteResponse to the client
func (o *UploadBuildArtifactsCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
}

/*UploadBuildArtifactsBadRequest Missing or invalid upload URI, or invalid timeout

swagger:response uploadBuildArtifactsBadRequest
*/
type UploadBuildArtifactsBadRequest struct {
}

// NewUploadBuildArtifactsBadRequest creates UploadBuildArtifactsBadRequest with default headers values
func NewUploadBuildArtifactsBadRequest() *UploadBuildArtifactsBadRequest {
	return &UploadBuildArtifactsBadRequest{}
}

// WriteResponse to the client
func (o *UploadBuildArtifactsBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
}

/*UploadBuildArtifactsInternalServerError Artifacts upload failed

swagger:response uploadBuildArtifactsInternalServerError
*/
type UploadBuildArtifactsInternalServerError struct {
}

// NewUploadBuildArtifactsInternalServerError creates UploadBuildArtifactsInternalServerError with default headers values
func NewUploadBuildArtifactsInternalServerError() *UploadBuildArtifactsInternalServerError {
	return &UploadBuildArtifactsInternalServerError{}
}

// WriteResponse to the client
func (o *UploadBuildArtifactsInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
}

/*UploadBuildArtifactsDefault Response code from upload request

swagger:response uploadBuildArtifactsDefault
*/
type UploadBuildArtifactsDefault struct {
	_statusCode int
}

// NewUploadBuildArtifactsDefault creates UploadBuildArtifactsDefault with default headers values
func NewUploadBuildArtifactsDefault(code int) *UploadBuildArtifactsDefault {
	if code <= 0 {
		code = 500
	}

	return &UploadBuildArtifactsDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the upload build artifacts default response
func (o *UploadBuildArtifactsDefault) WithStatusCode(code int) *UploadBuildArtifactsDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the upload build artifacts default response
func (o *UploadBuildArtifactsDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WriteResponse to the client
func (o *UploadBuildArtifactsDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
}
