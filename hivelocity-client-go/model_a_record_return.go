/*
 * Hivelocity API for Partners
 *
 * Interact with Hivelocity
 *
 * API version: 2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type ARecordReturn struct {
	DomainId int32 `json:"domainId"`
	Type_ string `json:"type"`
	Ttl int32 `json:"ttl"`
	Addresses []RecordValue `json:"addresses"`
	Name string `json:"name"`
}
