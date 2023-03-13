terraform {
  required_providers {
    example = {
      source = "austinvalle/sandbox"
    }
  }
}

resource "example_existing_widget" "example" {
  string_attribute = "hello there!"
}
