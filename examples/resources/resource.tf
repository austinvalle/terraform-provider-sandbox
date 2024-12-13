terraform {
  required_providers {
    examplecloud = {
      source = "austinvalle/sandbox"
    }
  }
}

provider "examplecloud" {
  set_namespace_from_token = false
}

resource "examplecloud_sdkv2_thing" "test" {
  trigger = "a"
}
