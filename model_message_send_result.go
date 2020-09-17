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
// MessageSendResult struct for MessageSendResult
type MessageSendResult struct {
	DeviceId string `json:"deviceId,omitempty"`
	Message string `json:"message,omitempty"`
}