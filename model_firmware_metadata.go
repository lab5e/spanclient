/*
 * The Span API
 *
 * API for device, collection, output and firmware management
 *
 * API version: 4.0.8 freckled-fawn
 * Contact: dev@lab5e.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package spanclient
// FirmwareMetadata Metadata about firmware on devices.
type FirmwareMetadata struct {
	CurrentFirmwareId string `json:"currentFirmwareId,omitempty"`
	TargetFirmwareId string `json:"targetFirmwareId,omitempty"`
	// Last reported firmware version.
	FirmwareVersion string `json:"firmwareVersion,omitempty"`
	SerialNumber string `json:"serialNumber,omitempty"`
	ModelNumber string `json:"modelNumber,omitempty"`
	Manufacturer string `json:"manufacturer,omitempty"`
	// State of the firmware.
	State string `json:"state,omitempty"`
	StateMessage string `json:"stateMessage,omitempty"`
}
