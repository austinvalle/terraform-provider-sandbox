terraform {
  required_providers {
    examplecloud = {
      source = "austinvalle/sandbox"
    }
  }
}

# resource "random_string" "old" {
#   length = 12
# }

resource "examplecloud_thing" "this" {
  email = "austin.valle@hashicorp.com"
  name  = "Austin Valle"
  age   = 32
}

output "ex" {
  value = provider::examplecloud::do_thing("AB", ["hey", "oh", "no"])
}

# moved {
#   from = random_string.old
#   to   = examplecloud_thing.new
# }
