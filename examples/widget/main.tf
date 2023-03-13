terraform {
  required_providers {
    example = {
      source = "austinvalle/sandbox"
    }
  }
}

resource "example_widget" "example" {
  string_attribute = "hello there!"
}
