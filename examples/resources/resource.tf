terraform {
  required_providers {
    examplecloud = {
      source = "austinvalle/sandbox"
    }
  }
}

locals {
  # json_example = "abc"
  json_example = jsonencode(
    {
      hello   = "world-update",
      num     = 3,
      decimal = 2.2,
      itdobe  = true,
      obj = {
        wewant = "some-more"
      },
      arr = [
        {
          ayo = "itsyaboi-1"
        },
        {
          ayo = "itsyaboi-2"
        }
      ]
    }
  )
}


resource "examplecloud_thing" "this" {
  json_before           = local.json_example
  json_exact_after      = local.json_example
  json_normalized_after = local.json_example
}
