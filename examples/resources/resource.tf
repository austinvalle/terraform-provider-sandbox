terraform {
  required_providers {
    examplecloud = {
      source = "austinvalle/sandbox"
    }
  }
}

resource "examplecloud_thing" "test" {
  name = "john"
  cities = {
    season = "spring"
  }
}

