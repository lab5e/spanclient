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
// ListFirmwareResponse struct for ListFirmwareResponse
type ListFirmwareResponse struct {
	Images []Firmware `json:"images,omitempty"`
}
