terraform {
  required_providers {
    aidial = {
      source = "alexey-ban/aidial"
    }
  }
}

provider "aidial" {
  base_url = var.aidial_base_url
  api_key  = var.aidial_api_key
}
