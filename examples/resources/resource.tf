terraform {
  required_providers {
    examplecloud = {
      source = "austinvalle/sandbox"
    }
  }
}

resource "examplecloud_thing" "test" {
  # Terraform will receive the type constraint as "list(dynamic)"
  # Resulting in the final type determination here as "list(string)"
  list_with_dynamics = ["hello", "world", true, 123, null]
}
