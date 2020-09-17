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
// CollectionFirmware struct for CollectionFirmware
type CollectionFirmware struct {
	// The current firmware is the firmware that the devices are currently using.
	CurrentFirmwareId string `json:"currentFirmwareId,omitempty"`
	// The target firmware is set to the desired firmware image for the devices in this collection. If the management is set to \"device\" this will only be used if the target firmware isn't set on the device itself.
	TargetFirmwareId string `json:"targetFirmwareId,omitempty"`
	Management CollectionFirmwareFirmwareManagement `json:"management,omitempty"`
}