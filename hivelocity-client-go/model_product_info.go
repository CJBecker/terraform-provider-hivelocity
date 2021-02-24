/*
 * Hivelocity API
 *
 * Interact with Hivelocity
 *
 * API version: 2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type ProductInfo struct {
	ProductCpuCores string `json:"product_cpu_cores,omitempty"`
	ProductOriginalPrice float32 `json:"product_original_price,omitempty"`
	ProductBandwidth string `json:"product_bandwidth,omitempty"`
	ProductDrive string `json:"product_drive,omitempty"`
	AnnuallyLocationPremium float32 `json:"annually_location_premium,omitempty"`
	ProductDisabledBillingPeriods []string `json:"product_disabled_billing_periods,omitempty"`
	ProductTriennialPrice float32 `json:"product_triennial_price,omitempty"`
	ProductSemiAnnuallyPrice float32 `json:"product_semi_annually_price,omitempty"`
	QuarterlyLocationPremium float32 `json:"quarterly_location_premium,omitempty"`
	ProductId int32 `json:"product_id,omitempty"`
	ProductMemory string `json:"product_memory,omitempty"`
	Edge bool `json:"edge,omitempty"`
	ProductDisplayPrice float32 `json:"product_display_price,omitempty"`
	BiennialLocationPremium float32 `json:"biennial_location_premium,omitempty"`
	ProductQuarterlyPrice float32 `json:"product_quarterly_price,omitempty"`
	TriennialLocationPremium float32 `json:"triennial_location_premium,omitempty"`
	SemiAnnuallyLocationPremium float32 `json:"semi_annually_location_premium,omitempty"`
	ProductName string `json:"product_name,omitempty"`
	ProductMonthlyPrice float32 `json:"product_monthly_price,omitempty"`
	Core bool `json:"core,omitempty"`
	ProductCpu string `json:"product_cpu,omitempty"`
	ProductAnnuallyPrice float32 `json:"product_annually_price,omitempty"`
	ProductHourlyPrice float32 `json:"product_hourly_price,omitempty"`
	ProductOnSale bool `json:"product_on_sale,omitempty"`
	HourlyLocationPremium float32 `json:"hourly_location_premium,omitempty"`
	DataCenter string `json:"data_center,omitempty"`
	Stock string `json:"stock,omitempty"`
	ProductBiennialPrice float32 `json:"product_biennial_price,omitempty"`
	MonthlyLocationPremium float32 `json:"monthly_location_premium,omitempty"`
}
