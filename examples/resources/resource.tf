terraform {
  required_providers {
    examplecloud = {
      source = "austinvalle/sandbox"
    }
  }
}

locals {
  # json_example = "abc"
  json_example = "{\"hello\": \"world1\", \"abc\": 3}"

  ipv4_example = "192.0.1.246"
  # ipv6_example = "0:0:0:0:0:0:0:0"
  # ipv6_example = "0:0:0:0:0:0:0:1"

  # IPv4-compatible addresses
  # ipv6_example = "::13.1.68.3"
  # ipv6_example = "0:0:0:0:0:0:13.1.68.3"
  # IPv4-mapped addresses
  ipv6_example = "0:0:0:0:0:FFFF:129.144.52.38"


  ipv4_cidr = "174.16.0.0/24"

  # ipv6_cidr = "2001:db8:0:0:0:0:0:0/117"
  ipv6_cidr = "2001:db8::/115"

  # rfc3339_example = timestamp()
  rfc3339_example = "2023-07-25T23:43:16Z"
}


resource "examplecloud_thing" "this" {
  json_before     = local.json_example
  json_exact      = local.json_example
  json_normalized = local.json_example

  ipv4_address_before = local.ipv4_example
  ipv6_address_before = local.ipv6_example
  ipv4_address        = local.ipv4_example
  ipv6_address        = local.ipv6_example

  ipv4_cidr_before = local.ipv4_cidr
  ipv6_cidr_before = local.ipv6_cidr

  ipv4_cidr = local.ipv4_cidr
  ipv6_cidr = local.ipv6_cidr

  rfc3339_before = local.rfc3339_example
  rfc3339        = local.rfc3339_example
}
