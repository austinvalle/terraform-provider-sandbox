terraform {
  required_providers {
    examplecloud = {
      source = "austinvalle/sandbox"
    }
  }
}

resource "examplecloud_thing" "test" {
  case_insensitive_string = "test" // modifying the case of this string shouldn't update the resource

  # case_insensitive_string = "TEST"
}
