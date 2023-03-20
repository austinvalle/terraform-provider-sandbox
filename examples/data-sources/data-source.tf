terraform {
  required_providers {
    example = {
      source = "austinvalle/sandbox"
    }
  }
}


data "example_datasource" "example" {}

resource "example_resource" "example" {
  string_attribute = data.example_datasource.example.string_attribute
}

# data "example_datasource" "example" {
#   // Simple/primitive attributes.
#   bool_attribute = true

#   float64_attribute = 1234.5

#   int64_attribute = 9223372036854775807

#   list_attribute = ["list-element", "list-element"]

#   map_attribute = { "map-key-1" : "map-value-1" }

#   number_attribute = 1.79769313486231570814527423731704356798070e+1000

#   // Treated as an atomic unit.
#   // Attributes are NOT configurable (i.e., cannot define required, optional,
#   // computed, PlanModifiers or Validators on individual fields).
#   object_attribute = {
#     bool_attribute    = true
#     float64_attribute = 1234.5
#     int64_attribute   = 9223372036854775807
#     list_attribute    = ["obj-list-element", "obj-list-element"]
#     map_attribute     = { "obj-map-key-1" : "obj-map-value-1" }
#     number_attribute  = 1.79769313486231570814527423731704356798070e+1000
#     set_attribute     = ["obj-set-element-1", "obj-set-element-2"]
#     string_attribute  = "obj-string"
#   }

#   set_attribute = ["set-element-1", "set-element-2"]

#   string_attribute = "string"

#   // Nested attributes - can contain simple/primitive attributes or other
#   // nested attributes.
#   // See https://www.terraform.io/plugin/framework/schemas
#   list_nested_attribute = [
#     {
#       int64_attribute  = 9223372036854775807
#       list_attribute   = ["list-element", "list-element"]
#       map_attribute    = { "map-key-1" : "map-value-1" }
#       number_attribute = 1.79769313486231570814527423731704356798070e+1000
#     },
#     {
#       int64_attribute  = 9223372036854775807
#       list_attribute   = ["list-element", "list-element"]
#       map_attribute    = { "map-key-1" : "map-value-1" }
#       number_attribute = 1.79769313486231570814527423731704356798070e+1000
#     }
#   ]

#   map_nested_attribute = {
#     "one" = {
#       map_attribute    = { "map-key-1" : "map-value-1" }
#       number_attribute = 1.79769313486231570814527423731704356798070e+1000
#     },
#     "two" = {
#       map_attribute    = { "map-key-1" : "map-value-1" }
#       number_attribute = 1.79769313486231570814527423731704356798070e+1000
#     }
#   }

#   // Each element in the set must be unique hence true/false for bool_attribute.
#   set_nested_attribute = [
#     {
#       object_attribute = {
#         bool_attribute    = true
#         float64_attribute = 1234.5
#         int64_attribute   = 9223372036854775807
#         list_attribute    = ["obj-list-element", "obj-list-element"]
#         map_attribute     = { "obj-map-key-1" : "obj-map-value-1" }
#         number_attribute  = 1.79769313486231570814527423731704356798070e+1000
#         set_attribute     = ["obj-set-element-1", "obj-set-element-2"]
#         string_attribute  = "obj-string"
#       }
#       set_attribute = ["set-element-1", "set-element-2"]
#     },
#     {
#       object_attribute = {
#         bool_attribute    = false
#         float64_attribute = 1234.5
#         int64_attribute   = 9223372036854775807
#         list_attribute    = ["obj-list-element", "obj-list-element"]
#         map_attribute     = { "obj-map-key-1" : "obj-map-value-1" }
#         number_attribute  = 1.79769313486231570814527423731704356798070e+1000
#         set_attribute     = ["obj-set-element-1", "obj-set-element-2"]
#         string_attribute  = "obj-string"
#       }
#       set_attribute = ["set-element-1", "set-element-2"]
#     }
#   ]

#   single_nested_attribute = {
#     bool_attribute    = true
#     float64_attribute = 1234.5
#   }

#   // Blocks

#   // The syntax for list and set nested block within the configuration is confusing
#   // as it doesn't appear as a list or set, rather separate fields with the same
#   // identifier is used in the configuration (e.g., list_nested_block,
#   // set_nested_block)and within the state (e.g., the .tfstate file) they are stored
#   // within an array.
#   // https://github.com/hashicorp/terraform-plugin-docs/issues/47
#   list_nested_block {
#     bool_attribute    = true
#     float64_attribute = 1234.5
#     int64_attribute   = 9223372036854775807
#     list_attribute    = ["list-element", "list-element"]
#     list_nested_nested_block {
#       bool_attribute = true
#     }
#     list_nested_nested_block {
#       bool_attribute = false
#     }
#   }
#   list_nested_block {
#     bool_attribute    = true
#     float64_attribute = 1234.5
#     int64_attribute   = 9223372036854775807
#     list_attribute    = ["list-element", "list-element"]
#     list_nested_nested_block {
#       bool_attribute = true
#     }
#     list_nested_nested_block {
#       bool_attribute = false
#     }
#   }

#   // Each element in the set must be unique hence true/false for object bool_attribute.
#   set_nested_block {
#     map_attribute    = { "map-key-1" : "map-value-1" }
#     number_attribute = 1.79769313486231570814527423731704356798070e+1000
#     object_attribute = {
#       bool_attribute    = true
#       float64_attribute = 1234.5
#       int64_attribute   = 9223372036854775807
#       list_attribute    = ["obj-list-element", "obj-list-element"]
#       map_attribute     = { "obj-map-key-1" : "obj-map-value-1" }
#       number_attribute  = 1.79769313486231570814527423731704356798070e+1000
#       set_attribute     = ["obj-set-element-1", "obj-set-element-2"]
#       string_attribute  = "obj-string"
#     }
#     set_attribute = ["set-element-1", "set-element-2"]
#     set_nested_nested_block {
#       bool_attribute = true
#     }
#     set_nested_nested_block {
#       bool_attribute = false
#     }
#   }
#   set_nested_block {
#     map_attribute    = { "map-key-1" : "map-value-1" }
#     number_attribute = 1.79769313486231570814527423731704356798070e+1000
#     object_attribute = {
#       bool_attribute    = false
#       float64_attribute = 1234.5
#       int64_attribute   = 9223372036854775807
#       list_attribute    = ["obj-list-element", "obj-list-element"]
#       map_attribute     = { "obj-map-key-1" : "obj-map-value-1" }
#       number_attribute  = 1.79769313486231570814527423731704356798070e+1000
#       set_attribute     = ["obj-set-element-1", "obj-set-element-2"]
#       string_attribute  = "obj-string"
#     }
#     set_attribute = ["set-element-1", "set-element-2"]
#     set_nested_nested_block {
#       bool_attribute = true
#     }
#     set_nested_nested_block {
#       bool_attribute = false
#     }
#   }

#   single_nested_block {
#     bool_attribute    = true
#     float64_attribute = 1234.5
#   }
# }
