/*
 * Hivelocity API
 *
 * Interact with Hivelocity
 *
 * API version: 2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type Option struct {
	Name string `json:"name,omitempty"`
	Id int32 `json:"id,omitempty"`
	Currency string `json:"currency,omitempty"`
	MonthlyPrice float32 `json:"monthlyPrice,omitempty"`
	Expressions []string `json:"expressions,omitempty"`
	Tags []string `json:"tags,omitempty"`
}
