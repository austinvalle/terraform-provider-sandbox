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

  ipv4_example = "192.0.1.246"
  ipv6_example = "2002:db8::8a3f:362:7897"
}


resource "examplecloud_thing" "this" {
  json_before     = local.json_example
  json_exact      = local.json_example
  json_normalized = local.json_example

  ipv4_address_before = local.ipv4_example
  ipv6_address_before = local.ipv6_example

  ipv4_address = local.ipv4_example
  ipv6_address = local.ipv6_example
}
