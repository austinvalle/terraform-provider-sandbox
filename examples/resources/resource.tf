terraform {
  required_providers {
    examplecloud = {
      source = "austinvalle/sandbox"
    }
  }
}

resource "random_string" "str" {
  length = 5
}

provider "examplecloud" {
  # Unknown until after the first apply!
  api_key = random_string.str.result
}

# Will defer until the provider configuration is fully known!
resource "examplecloud_sdkv2_thing" "test" {
  name = data.examplecloud_sdkv2_thing.test.name
  size = 99
}

# Will defer until the provider configuration is fully known!
data "examplecloud_sdkv2_thing" "test" {}
