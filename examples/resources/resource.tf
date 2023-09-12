terraform {
  required_providers {
    examplecloud = {
      source = "austinvalle/sandbox"
    }
  }
}
resource "examplecloud_thing" "this" {

}
