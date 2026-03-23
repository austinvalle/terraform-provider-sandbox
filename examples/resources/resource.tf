terraform {
  required_providers {
    examplecloud = {
      source = "austinvalle/sandbox"
    }
  }
}

resource "examplecloud_thing" "test" {
  config_attr = "hello"

  # break_data_consistency = true # ru roh
}

# resource "terraform_data" "test" {
#   input = examplecloud_thing.test.config_attr
# }

# import {
#   id = "hello"
#   to = examplecloud_thing.test
# }


## Data consistency rules ##
#
# 1. Any attribute that was non-null in the configuration
#    must either preserve the exact configuration value or
#    return the corresponding attribute value from the prior state.
#
# 2. Any attribute that is marked as computed in the schema and is null
#    in the configuration may be set by the provider to any arbitrary value
#    of the expected type.
#
# 3. If a computed attribute has any known value in the planned new state,
#    the provider will be required to ensure that it is unchanged in the new
#    state returned by ApplyResourceChange, or return an error explaining why
#    it changed. Set an attribute to an unknown value to indicate that its final
#    result will be determined during ApplyResourceChange.
