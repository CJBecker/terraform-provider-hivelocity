/*
 * Hivelocity API
 *
 * Interact with Hivelocity
 *
 * API version: 2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type IpmiInfo struct {
	IpmiVersion string `json:"ipmiVersion,omitempty"`
	IpmbEventReceiver string `json:"ipmbEventReceiver,omitempty"`
	AuxFirmwareRevInfo string `json:"auxFirmwareRevInfo,omitempty"`
	ProductId string `json:"productId,omitempty"`
	SensorDevice string `json:"sensorDevice,omitempty"`
	DeviceId string `json:"deviceId,omitempty"`
	IpmbEventGenerator string `json:"ipmbEventGenerator,omitempty"`
	SdrRepositoryDevice string `json:"sdrRepositoryDevice,omitempty"`
	Bridge string `json:"bridge,omitempty"`
	ManufacturerId string `json:"manufacturerId,omitempty"`
	DeviceAvailable string `json:"deviceAvailable,omitempty"`
	SelDevice string `json:"selDevice,omitempty"`
	FruInventoryDevice string `json:"fruInventoryDevice,omitempty"`
	DeviceSDRs string `json:"deviceSDRs,omitempty"`
	DeviceRevision string `json:"deviceRevision,omitempty"`
	FirmwareRevision string `json:"firmwareRevision,omitempty"`
	ChassisDevice string `json:"chassisDevice,omitempty"`
}
