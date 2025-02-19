// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package monitoring

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// RemoveAlarmSuppressionRequest wrapper for the RemoveAlarmSuppression operation
type RemoveAlarmSuppressionRequest struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of an alarm.
	AlarmId *string `mandatory:"true" contributesTo:"path" name:"alarmId"`

	// Customer part of the request identifier token. If you need to contact Oracle about a particular
	// request, please provide the complete request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request RemoveAlarmSuppressionRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request RemoveAlarmSuppressionRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request RemoveAlarmSuppressionRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// RemoveAlarmSuppressionResponse wrapper for the RemoveAlarmSuppression operation
type RemoveAlarmSuppressionResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response RemoveAlarmSuppressionResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response RemoveAlarmSuppressionResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
