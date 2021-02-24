terraform {
  required_providers {
    hivelocity = {
      versions = ["0.0.1"]
      source = "hivelocity.net/prod/hivelocity"
    }
  }
}

// Find a plan with 16GB of memory in Tampa.
data "hivelocity_product" "tampa_product" {
  first = true

  filter {
    name   = "product_memory"
    values = ["64GB"]
  }

  filter {
    name   = "data_center"
    values = ["TPA1"]
  }

  filter {
    name   = "stock"
    values = ["limited", "available"]
  }
}

// Provision your device with CentOS 7.
resource "hivelocity_bare_metal_device" "tampa_server" {
    product_id    = "${data.hivelocity_product.tampa_product.product_id}"
    os_name       = "CentOS 7.x"
    location_name = "${data.hivelocity_product.tampa_product.data_center}"
    hostname      = "hivelocity.terraform.test"
    tags          = ["hello", "world"]
    script        = file("${path.module}/cloud_init_example.yaml")
}
