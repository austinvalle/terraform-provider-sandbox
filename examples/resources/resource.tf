terraform {
  required_providers {
    examplecloud = {
      source = "austinvalle/sandbox"
    }
  }
}

locals {
  json_example = jsonencode(
    {
      hello   = "world",
      num     = 2,
      decimal = 2.2,
      itdobe  = true,
      obj = {
        wewant = "some-more"
      },
      arr = [
        {
          ayo = "itsyaboi"
        }
      ]
    }
  )
}


resource "examplecloud_thing" "this" {
  json_string = local.json_example
}
