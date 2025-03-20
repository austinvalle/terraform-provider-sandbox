terraform {
  required_providers {
    examplecloud = {
      source = "austinvalle/sandbox"
    }
  }
}

resource "examplecloud_thing" "test" {
  case_insensitive_string = "TEST" # apply 1, create
  # case_insensitive_string = "test" # apply 2, no-op
  # case_insensitive_string = "new value" # apply 3, update, computed_attr_two is marked as unknown
  # case_insensitive_string = "NEW value" # apply 4, no-op
}

output "test_out" {
  value = examplecloud_thing.test
}
