/*
 * Hivelocity API for Partners
 *
 * Interact with Hivelocity
 *
 * API version: 2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type ContactCreate struct {
	Active int32 `json:"active"`
	Email string `json:"email"`
	Password string `json:"password,omitempty"`
	Description string `json:"description,omitempty"`
	Phone string `json:"phone,omitempty"`
	FullName string `json:"fullName"`
}
