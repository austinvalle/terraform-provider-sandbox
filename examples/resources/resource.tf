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
  # ipv6_example = "1080:0:0:0:8:800:200C:417A"
  # ipv6_example = "0:0:0:0:0:0:0:0"
  # ipv6_example = "0:0:0:0:0:0:0:1"

  # IPv4-compatible addresses
  # ipv6_example = "::13.1.68.3"
  # ipv6_example = "0:0:0:0:0:0:13.1.68.3"
  # IPv4-mapped addresses
  ipv6_example = "0:0:0:0:0:FFFF:129.144.52.38"
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
