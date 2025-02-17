---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "hivelocity_product Data Source - terraform-provider-hivelocity"
subcategory: ""
description: |-
  
---

# hivelocity_product (Data Source)



## Example Usage

```terraform
terraform {
  required_providers {
    hivelocity = {
      versions = ["0.1.0"]
      source   = "hivelocity/hivelocity"
    }
  }
}

data "hivelocity_product" "tampa" {
  filter {
    name   = "data_center"
    values = ["TPA1"]
  }
}

output "tampa_product" {
  value = data.hivelocity_product.tampa
}

data "hivelocity_product" "all" {}

output "first_product" {
  value = data.hivelocity_product.all
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `filter` (Block Set) (see [below for nested schema](#nestedblock--filter))
- `first` (Boolean)

### Read-Only

- `annually_location_premium` (Number)
- `biennial_location_premium` (Number)
- `core` (Boolean)
- `data_center` (String)
- `edge` (Boolean)
- `hourly_location_premium` (Number)
- `id` (String) The ID of this resource.
- `monthly_location_premium` (Number)
- `product_annually_price` (Number)
- `product_bandwidth` (String)
- `product_biennial_price` (Number)
- `product_cpu` (String)
- `product_cpu_cores` (String)
- `product_disabled_billing_periods` (List of String)
- `product_display_price` (Number)
- `product_drive` (String)
- `product_gpu` (String)
- `product_hourly_price` (Number)
- `product_id` (Number)
- `product_memory` (String)
- `product_monthly_price` (Number)
- `product_name` (String)
- `product_on_sale` (Boolean)
- `product_original_price` (Number)
- `product_quarterly_price` (Number)
- `product_semi_annually_price` (Number)
- `product_triennial_price` (Number)
- `quarterly_location_premium` (Number)
- `semi_annually_location_premium` (Number)
- `stock` (String)
- `triennial_location_premium` (Number)

<a id="nestedblock--filter"></a>
### Nested Schema for `filter`

Required:

- `name` (String)
- `values` (List of String)


