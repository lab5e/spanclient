/*
 * The Span API
 *
 * API for device, collection, output and firmware management
 *
 * API version: 3.0.0
 * Contact: dev@lab5e.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package spanclient
// MultiSendMessageResponse Broadcast message result. The errors array contains the list of errors ocurred when sending a message.
type MultiSendMessageResponse struct {
	Errors []MessageSendResult `json:"errors,omitempty"`
	Sent int32 `json:"sent,omitempty"`
	Failed int32 `json:"failed,omitempty"`
}
