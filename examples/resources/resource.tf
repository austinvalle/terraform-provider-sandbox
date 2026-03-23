terraform {
  required_providers {
    examplecloud = {
      source = "austinvalle/sandbox"
    }
  }
}

resource "examplecloud_thing" "test" {
  config_attr = "hello"

  # break_data_consistency = true # ru roh
}
