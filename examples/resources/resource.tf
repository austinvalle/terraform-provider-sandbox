terraform {
  required_providers {
    example = {
      source = "austinvalle/sandbox"
    }
  }
}

resource "examplecloud_thing" "this" {}


data "examplecloud_thing" "this" {}
