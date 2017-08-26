/*                          _       _
 *__      _____  __ ___   ___  __ _| |_ ___
 *\ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
 * \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
 *  \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
 *
 * Copyright © 2016 Weaviate. All rights reserved.
 * LICENSE: https://github.com/weaviate/weaviate/blob/master/LICENSE
 * AUTHOR: Bob van Luijt (bob@weaviate.com)
 * See www.weaviate.com for details
 * Contact: @weaviate_iot / yourfriends@weaviate.com
 */
   

package things


 

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// WeaviateThingsValidateOKCode is the HTTP code returned for type WeaviateThingsValidateOK
const WeaviateThingsValidateOKCode int = 200

/*WeaviateThingsValidateOK Successful validated.

swagger:response weaviateThingsValidateOK
*/
type WeaviateThingsValidateOK struct {
}

// NewWeaviateThingsValidateOK creates WeaviateThingsValidateOK with default headers values
func NewWeaviateThingsValidateOK() *WeaviateThingsValidateOK {
	return &WeaviateThingsValidateOK{}
}

// WriteResponse to the client
func (o *WeaviateThingsValidateOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
}

// WeaviateThingsValidateUnauthorizedCode is the HTTP code returned for type WeaviateThingsValidateUnauthorized
const WeaviateThingsValidateUnauthorizedCode int = 401

/*WeaviateThingsValidateUnauthorized Unauthorized or invalid credentials.

swagger:response weaviateThingsValidateUnauthorized
*/
type WeaviateThingsValidateUnauthorized struct {
}

// NewWeaviateThingsValidateUnauthorized creates WeaviateThingsValidateUnauthorized with default headers values
func NewWeaviateThingsValidateUnauthorized() *WeaviateThingsValidateUnauthorized {
	return &WeaviateThingsValidateUnauthorized{}
}

// WriteResponse to the client
func (o *WeaviateThingsValidateUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
}

// WeaviateThingsValidateForbiddenCode is the HTTP code returned for type WeaviateThingsValidateForbidden
const WeaviateThingsValidateForbiddenCode int = 403

/*WeaviateThingsValidateForbidden The used API-key has insufficient permissions.

swagger:response weaviateThingsValidateForbidden
*/
type WeaviateThingsValidateForbidden struct {
}

// NewWeaviateThingsValidateForbidden creates WeaviateThingsValidateForbidden with default headers values
func NewWeaviateThingsValidateForbidden() *WeaviateThingsValidateForbidden {
	return &WeaviateThingsValidateForbidden{}
}

// WriteResponse to the client
func (o *WeaviateThingsValidateForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
}

// WeaviateThingsValidateUnprocessableEntityCode is the HTTP code returned for type WeaviateThingsValidateUnprocessableEntity
const WeaviateThingsValidateUnprocessableEntityCode int = 422

/*WeaviateThingsValidateUnprocessableEntity Can not validate, check the body.

swagger:response weaviateThingsValidateUnprocessableEntity
*/
type WeaviateThingsValidateUnprocessableEntity struct {
}

// NewWeaviateThingsValidateUnprocessableEntity creates WeaviateThingsValidateUnprocessableEntity with default headers values
func NewWeaviateThingsValidateUnprocessableEntity() *WeaviateThingsValidateUnprocessableEntity {
	return &WeaviateThingsValidateUnprocessableEntity{}
}

// WriteResponse to the client
func (o *WeaviateThingsValidateUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
}

// WeaviateThingsValidateNotImplementedCode is the HTTP code returned for type WeaviateThingsValidateNotImplemented
const WeaviateThingsValidateNotImplementedCode int = 501

/*WeaviateThingsValidateNotImplemented Not (yet) implemented.

swagger:response weaviateThingsValidateNotImplemented
*/
type WeaviateThingsValidateNotImplemented struct {
}

// NewWeaviateThingsValidateNotImplemented creates WeaviateThingsValidateNotImplemented with default headers values
func NewWeaviateThingsValidateNotImplemented() *WeaviateThingsValidateNotImplemented {
	return &WeaviateThingsValidateNotImplemented{}
}

// WriteResponse to the client
func (o *WeaviateThingsValidateNotImplemented) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(501)
}
