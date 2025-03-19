terraform {
  required_providers {
    examplecloud = {
      source = "austinvalle/sandbox"
    }
  }
}

resource "examplecloud_thing" "test" {
  attr_one = "hellos"
}

output "test" {
  value = "${examplecloud_thing.test.attr_one} ${examplecloud_thing.test.attr_two}"
}
