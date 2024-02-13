terraform {
  required_providers {
    examplecloud = {
      source = "austinvalle/sandbox"
    }
  }
}

output "test_1" {
  value = provider :: examplecloud :: dynamic_test({ bool_attr = false }, tolist(["it's", "a", "list"]))
}

output "test_2" {
  value = provider :: examplecloud :: dynamic_test({ string_attr = "hello" }, toset(["it's", "a", "set"]))
}


output "test_3" {
  value = provider :: examplecloud :: dynamic_test("two", "strings")
}
