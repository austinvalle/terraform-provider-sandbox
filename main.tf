terraform {
  required_providers {
    examplecloud = {
      source = "austinvalle/sandbox"
    }
  }
}

resource "examplecloud_thing" "test" {
  # Causes infinite plans
  number = 242.08120431461208
  # Does not cause infinite plans - last digit changed
  # number =  242.08120431461209
}
