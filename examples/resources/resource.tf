terraform {
  required_providers {
    sandbox = {
      source = "austinvalle/sandbox"
    }
  }
}
resource "sandbox_thing" "this" {
  example_attr = "hello there!"
}
