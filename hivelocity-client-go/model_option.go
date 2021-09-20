/*
 * Hivelocity API for Partners
 *
 * Interact with Hivelocity
 *
 * API version: 2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type Option struct {
	Expressions []string `json:"expressions,omitempty"`
	Currency string `json:"currency,omitempty"`
	Name string `json:"name,omitempty"`
	MonthlyPrice float32 `json:"monthlyPrice,omitempty"`
	Tags []string `json:"tags,omitempty"`
	Id int32 `json:"id,omitempty"`
}
