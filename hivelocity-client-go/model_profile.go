/*
 * Hivelocity API
 *
 * Interact with Hivelocity
 *
 * API version: 2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type Profile struct {
	IsClient bool `json:"is_client,omitempty"`
	Phone string `json:"phone,omitempty"`
	Company *interface{} `json:"company,omitempty"`
	Country *interface{} `json:"country,omitempty"`
	Zip *interface{} `json:"zip,omitempty"`
	Last string `json:"last,omitempty"`
	Id int32 `json:"id,omitempty"`
	Login string `json:"login,omitempty"`
	First string `json:"first,omitempty"`
	Address *interface{} `json:"address,omitempty"`
	Fax *interface{} `json:"fax,omitempty"`
	Created *interface{} `json:"created,omitempty"`
	FullName *interface{} `json:"full_name,omitempty"`
	Email string `json:"email,omitempty"`
	City *interface{} `json:"city,omitempty"`
	MetaData *interface{} `json:"meta_data,omitempty"`
	State *interface{} `json:"state,omitempty"`
}
