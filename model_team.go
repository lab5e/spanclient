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
// Team struct for Team
type Team struct {
	TeamId string `json:"teamId,omitempty"`
	Tags map[string]string `json:"tags,omitempty"`
	Members []Member `json:"members,omitempty"`
}
