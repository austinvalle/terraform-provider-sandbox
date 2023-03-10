terraform {
  required_providers {
    example = {
      source = "austinvalle/sandbox"
    }
  }
}

resource "example_widget" "example" {
  new_attribute = "hello there!"
}
