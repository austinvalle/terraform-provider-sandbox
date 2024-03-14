terraform {
  required_providers {
    examplecloud = {
      source = "austinvalle/sandbox"
    }
  }
}

# Managed resource with dynamic attributes
# resource "examplecloud_thing" "this" {
#   dynamic = [true, { "a" : "hello there" }]
#   static_obj = {
#     dynamic = ["check", "this", "out"]
#     # dynamic = tolist(["check", "this", "out"])
#     # dynamic = toset(["check", "this", "out"])
#   }

#   static_list = ["this", "is", "a", "list"]
# }

# output "dynamic_output" {
#   value = [for num in examplecloud_thing.this.dynamic_computed : num]
# }

# Function with a static string param and a dynamic variadic parameter
output "func_test" {
  value = provider::examplecloud::check_things(null, null, "hey")
}
