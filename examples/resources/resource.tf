terraform {
  required_providers {
    examplecloud = {
      source = "austinvalle/sandbox"
    }
  }
}

resource "terraform_data" "force_deferral" {
  input = "hello"
}

data "examplecloud_thing" "test" {
  nested_map = {
    "key1" : {
      required_attr = terraform_data.force_deferral.output
      # computed_attr
    }
  }
  nested_obj = {
    required_attr = terraform_data.force_deferral.output
    # computed_attr
  }
}

output "test" {
  value = data.examplecloud_thing.test
}
