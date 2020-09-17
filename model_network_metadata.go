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
// NetworkMetadata Network metadata for devices.
type NetworkMetadata struct {
	// Allocated IP address.
	AllocatedIp string `json:"allocatedIp,omitempty"`
	AllocatedAt string `json:"allocatedAt,omitempty"`
	// Cell ID for device. This might not be set, depending on the provider in use.
	CellId string `json:"cellId,omitempty"`
}
