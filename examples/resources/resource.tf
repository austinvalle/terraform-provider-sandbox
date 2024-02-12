terraform {
  required_providers {
    examplecloud = {
      source = "austinvalle/sandbox"
    }
  }
}

resource "examplecloud_thing" "this" {
  dynamic_attr = tolist([1, 2, 3])
}
