terraform {
  required_providers {
    examplecloud = {
      source = "austinvalle/sandbox"
    }
  }
}
resource "random_string" "a" {
  length = 10
}

resource "examplecloud_thing" "test" {
  str_value  = random_string.a.result
  max_length = 5
}
