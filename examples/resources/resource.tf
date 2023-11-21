terraform {
  required_providers {
    examplecloud = {
      source = "austinvalle/sandbox"
    }
  }
}
resource "examplecloud_thing" "this" {
  set = [
    {
      name     = "test2"
      bool_two = true
    }
  ]
}
