/*
 * Hivelocity API
 *
 * Interact with Hivelocity
 *
 * API version: 2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type BareMetalDeviceCreate struct {
	OsName string `json:"osName"`
	VlanId int32 `json:"vlanId,omitempty"`
	ProductId int32 `json:"productId"`
	LocationName string `json:"locationName"`
	Hostname string `json:"hostname"`
}
