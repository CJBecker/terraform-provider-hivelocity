/*
 * Hivelocity API
 *
 * Interact with Hivelocity
 *
 * API version: 2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type IpmiSensor struct {
	Units string `json:"units,omitempty"`
	SensorId string `json:"sensorId,omitempty"`
	Reading float32 `json:"reading,omitempty"`
	Group string `json:"group,omitempty"`
	Status bool `json:"status,omitempty"`
	Name string `json:"name,omitempty"`
}
