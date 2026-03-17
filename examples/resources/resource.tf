terraform {
  required_providers {
    examplecloud = {
      source = "austinvalle/sandbox"
    }
  }
}

resource "examplecloud_thing" "test" {
  config_attr = "hello world"
  ip_address  = "127.0.0."
}
