terraform {
  required_providers {
    example = {
      source = "austinvalle/sandbox"
    }
  }
}

resource "example_widget" "example" {
  existing_attribute = "hello there!"
}
