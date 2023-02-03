package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccExampleResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `resource "example_resource" "example" {
  bool_attribute = true

  float64_attribute = 1234.5

  int64_attribute = 9223372036854775807

  list_attribute = ["list-element", "list-element"]

  map_attribute = { "map-key-1" : "map-value-1" }

  number_attribute = 1.7976931348623157e+308

  object_attribute = {
    bool_attribute    = true
    float64_attribute = 1234.5
    int64_attribute   = 9223372036854775807
    list_attribute    = ["obj-list-element", "obj-list-element"]
    map_attribute     = { "obj-map-key-1" : "obj-map-value-1" }
    number_attribute  = 1.7976931348623157e+308
    set_attribute     = ["obj-set-element-1", "obj-set-element-2"]
    string_attribute  = "obj-string"
  }

  set_attribute = ["set-element-1", "set-element-2"]

  string_attribute = "string"

  list_nested_attribute = [
    {
      int64_attribute  = 9223372036854775807
      list_attribute   = ["list-element", "list-element"]
    },
    {
      int64_attribute  = 9223372036854775807
      list_attribute   = ["list-element", "list-element"]
    }
  ]

  map_nested_attribute = {
    "one" = {
      map_attribute    = { "map-key-1" : "map-value-1" }
      number_attribute = 1.7976931348623157e+308
    },
    "two" = {
      map_attribute    = { "map-key-1" : "map-value-1" }
      number_attribute = 1.7976931348623157e+308
    }
  }

  set_nested_attribute = [
    {
      object_attribute = {
        bool_attribute    = true
        float64_attribute = 1234.5
        int64_attribute   = 9223372036854775807
        list_attribute    = ["obj-list-element", "obj-list-element"]
        map_attribute     = { "obj-map-key-1" : "obj-map-value-1" }
        number_attribute  = 1.7976931348623157e+308
        set_attribute     = ["obj-set-element-1", "obj-set-element-2"]
        string_attribute  = "obj-string"
      }
      set_attribute    = ["set-element-1", "set-element-2"]
      string_attribute = "string"
    },
    {
      object_attribute = {
        bool_attribute    = false
        float64_attribute = 1234.5
        int64_attribute   = 9223372036854775807
        list_attribute    = ["obj-list-element", "obj-list-element"]
        map_attribute     = { "obj-map-key-1" : "obj-map-value-1" }
        number_attribute  = 1.7976931348623157e+308
        set_attribute     = ["obj-set-element-1", "obj-set-element-2"]
        string_attribute  = "obj-string"
      }
      set_attribute    = ["set-element-1", "set-element-2"]
      string_attribute = "string"
    }
  ]

  single_nested_attribute = {
    bool_attribute    = true
    float64_attribute = 1234.5
  }

  list_nested_block {
    bool_attribute = true
    float64_attribute = 1234.5
    int64_attribute = 9223372036854775807
    list_attribute = ["list-element", "list-element"]
    list_nested_nested_block {
      bool_attribute = true
    }
    list_nested_nested_block {
      bool_attribute = false
    }
  }
  list_nested_block {
    bool_attribute = true
    float64_attribute = 1234.5
    int64_attribute = 9223372036854775807
    list_attribute = ["list-element", "list-element"]
    list_nested_nested_block {
      bool_attribute = true
    }
    list_nested_nested_block {
      bool_attribute = false
    }
  }

  set_nested_block {
    map_attribute = { "map-key-1" : "map-value-1" }
    number_attribute = 1.7976931348623157e+308
    object_attribute = {
      bool_attribute    = true
      float64_attribute = 1234.5
      int64_attribute   = 9223372036854775807
      list_attribute    = ["obj-list-element", "obj-list-element"]
      map_attribute     = { "obj-map-key-1" : "obj-map-value-1" }
      number_attribute  = 1.7976931348623157e+308
      set_attribute     = ["obj-set-element-1", "obj-set-element-2"]
      string_attribute  = "obj-string"
    }
    set_attribute = ["set-element-1", "set-element-2"]
    set_nested_nested_block {
      bool_attribute = true
    }
    set_nested_nested_block {
      bool_attribute = false
    }
  }
  set_nested_block {
    map_attribute = { "map-key-1" : "map-value-1" }
    number_attribute = 1.7976931348623157e+308
    object_attribute = {
      bool_attribute    = false
      float64_attribute = 1234.5
      int64_attribute   = 9223372036854775807
      list_attribute    = ["obj-list-element", "obj-list-element"]
      map_attribute     = { "obj-map-key-1" : "obj-map-value-1" }
      number_attribute  = 1.7976931348623157e+308
      set_attribute     = ["obj-set-element-1", "obj-set-element-2"]
      string_attribute  = "obj-string"
    }
    set_attribute = ["set-element-1", "set-element-2"]
    set_nested_nested_block {
      bool_attribute = true
    }
    set_nested_nested_block {
      bool_attribute = false
    }
  }

  single_nested_block {
    bool_attribute    = true
    float64_attribute = 1234.5
  }
}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// Id
					resource.TestCheckResourceAttr("example_resource.example", "id", "example-id"),

					// Simple/primitive attributes
					resource.TestCheckResourceAttr("example_resource.example", "bool_attribute", "true"),
					resource.TestCheckResourceAttr("example_resource.example", "float64_attribute", "1234.5"),
					resource.TestCheckResourceAttr("example_resource.example", "int64_attribute", "9223372036854775807"),
					resource.TestCheckResourceAttr("example_resource.example", "list_attribute.0", "list-element"),
					resource.TestCheckResourceAttr("example_resource.example", "list_attribute.1", "list-element"),
					resource.TestCheckResourceAttr("example_resource.example", "map_attribute.map-key-1", "map-value-1"),
					resource.TestCheckResourceAttr("example_resource.example", "number_attribute", "179769313486231570000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"),
					resource.TestCheckResourceAttr("example_resource.example", "object_attribute.bool_attribute", "true"),
					resource.TestCheckResourceAttr("example_resource.example", "object_attribute.float64_attribute", "1234.5"),
					resource.TestCheckResourceAttr("example_resource.example", "object_attribute.int64_attribute", "9223372036854775807"),
					resource.TestCheckResourceAttr("example_resource.example", "object_attribute.list_attribute.0", "obj-list-element"),
					resource.TestCheckResourceAttr("example_resource.example", "object_attribute.list_attribute.1", "obj-list-element"),
					resource.TestCheckResourceAttr("example_resource.example", "object_attribute.map_attribute.obj-map-key-1", "obj-map-value-1"),
					resource.TestCheckResourceAttr("example_resource.example", "object_attribute.number_attribute", "179769313486231570000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"),
					resource.TestCheckResourceAttr("example_resource.example", "object_attribute.set_attribute.0", "obj-set-element-1"),
					resource.TestCheckResourceAttr("example_resource.example", "object_attribute.set_attribute.1", "obj-set-element-2"),
					resource.TestCheckResourceAttr("example_resource.example", "object_attribute.string_attribute", "obj-string"),
					resource.TestCheckResourceAttr("example_resource.example", "set_attribute.0", "set-element-1"),
					resource.TestCheckResourceAttr("example_resource.example", "set_attribute.1", "set-element-2"),
					resource.TestCheckResourceAttr("example_resource.example", "string_attribute", "string"),

					// Nested Attributes
					// List Nested Attribute
					resource.TestCheckResourceAttr("example_resource.example", "list_nested_attribute.0.int64_attribute", "9223372036854775807"),
					resource.TestCheckResourceAttr("example_resource.example", "list_nested_attribute.0.list_attribute.0", "list-element"),
					resource.TestCheckResourceAttr("example_resource.example", "list_nested_attribute.0.list_attribute.1", "list-element"),
					// Map Nested Attribute
					resource.TestCheckResourceAttr("example_resource.example", "map_nested_attribute.one.map_attribute.map-key-1", "map-value-1"),
					resource.TestCheckResourceAttr("example_resource.example", "map_nested_attribute.one.number_attribute", "179769313486231570000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"),
					resource.TestCheckResourceAttr("example_resource.example", "map_nested_attribute.two.map_attribute.map-key-1", "map-value-1"),
					resource.TestCheckResourceAttr("example_resource.example", "map_nested_attribute.two.number_attribute", "179769313486231570000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"),
					// Set Nested Attribute
					resource.TestCheckResourceAttr("example_resource.example", "set_nested_attribute.0.object_attribute.bool_attribute", "false"),
					resource.TestCheckResourceAttr("example_resource.example", "set_nested_attribute.0.object_attribute.float64_attribute", "1234.5"),
					resource.TestCheckResourceAttr("example_resource.example", "set_nested_attribute.0.object_attribute.int64_attribute", "9223372036854775807"),
					resource.TestCheckResourceAttr("example_resource.example", "set_nested_attribute.0.object_attribute.list_attribute.0", "obj-list-element"),
					resource.TestCheckResourceAttr("example_resource.example", "set_nested_attribute.0.object_attribute.list_attribute.1", "obj-list-element"),
					resource.TestCheckResourceAttr("example_resource.example", "set_nested_attribute.0.object_attribute.map_attribute.obj-map-key-1", "obj-map-value-1"),
					resource.TestCheckResourceAttr("example_resource.example", "set_nested_attribute.0.object_attribute.number_attribute", "179769313486231570000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"),
					resource.TestCheckResourceAttr("example_resource.example", "set_nested_attribute.0.object_attribute.set_attribute.0", "obj-set-element-1"),
					resource.TestCheckResourceAttr("example_resource.example", "set_nested_attribute.0.object_attribute.set_attribute.1", "obj-set-element-2"),
					resource.TestCheckResourceAttr("example_resource.example", "set_nested_attribute.0.object_attribute.string_attribute", "obj-string"),
					resource.TestCheckResourceAttr("example_resource.example", "set_nested_attribute.0.set_attribute.0", "set-element-1"),
					resource.TestCheckResourceAttr("example_resource.example", "set_nested_attribute.0.set_attribute.1", "set-element-2"),
					resource.TestCheckResourceAttr("example_resource.example", "set_nested_attribute.0.string_attribute", "string"),
					resource.TestCheckResourceAttr("example_resource.example", "set_nested_attribute.1.object_attribute.bool_attribute", "true"),
					resource.TestCheckResourceAttr("example_resource.example", "set_nested_attribute.1.object_attribute.float64_attribute", "1234.5"),
					resource.TestCheckResourceAttr("example_resource.example", "set_nested_attribute.1.object_attribute.int64_attribute", "9223372036854775807"),
					resource.TestCheckResourceAttr("example_resource.example", "set_nested_attribute.1.object_attribute.list_attribute.0", "obj-list-element"),
					resource.TestCheckResourceAttr("example_resource.example", "set_nested_attribute.1.object_attribute.list_attribute.1", "obj-list-element"),
					resource.TestCheckResourceAttr("example_resource.example", "set_nested_attribute.1.object_attribute.map_attribute.obj-map-key-1", "obj-map-value-1"),
					resource.TestCheckResourceAttr("example_resource.example", "set_nested_attribute.1.object_attribute.number_attribute", "179769313486231570000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"),
					resource.TestCheckResourceAttr("example_resource.example", "set_nested_attribute.1.object_attribute.set_attribute.0", "obj-set-element-1"),
					resource.TestCheckResourceAttr("example_resource.example", "set_nested_attribute.1.object_attribute.set_attribute.1", "obj-set-element-2"),
					resource.TestCheckResourceAttr("example_resource.example", "set_nested_attribute.1.object_attribute.string_attribute", "obj-string"),
					resource.TestCheckResourceAttr("example_resource.example", "set_nested_attribute.1.set_attribute.0", "set-element-1"),
					resource.TestCheckResourceAttr("example_resource.example", "set_nested_attribute.1.set_attribute.1", "set-element-2"),
					resource.TestCheckResourceAttr("example_resource.example", "set_nested_attribute.1.string_attribute", "string"),
					// Single Nested Attribute
					resource.TestCheckResourceAttr("example_resource.example", "single_nested_attribute.bool_attribute", "true"),
					resource.TestCheckResourceAttr("example_resource.example", "single_nested_attribute.float64_attribute", "1234.5"),

					// Blocks
					// List Nested Block
					resource.TestCheckResourceAttr("example_resource.example", "list_nested_block.0.bool_attribute", "true"),
					resource.TestCheckResourceAttr("example_resource.example", "list_nested_block.0.float64_attribute", "1234.5"),
					resource.TestCheckResourceAttr("example_resource.example", "list_nested_block.0.int64_attribute", "9223372036854775807"),
					resource.TestCheckResourceAttr("example_resource.example", "list_nested_block.0.list_attribute.0", "list-element"),
					resource.TestCheckResourceAttr("example_resource.example", "list_nested_block.0.list_attribute.1", "list-element"),
					resource.TestCheckResourceAttr("example_resource.example", "list_nested_block.0.list_nested_nested_block.0.bool_attribute", "true"),
					resource.TestCheckResourceAttr("example_resource.example", "list_nested_block.1.bool_attribute", "true"),
					resource.TestCheckResourceAttr("example_resource.example", "list_nested_block.1.float64_attribute", "1234.5"),
					resource.TestCheckResourceAttr("example_resource.example", "list_nested_block.1.int64_attribute", "9223372036854775807"),
					resource.TestCheckResourceAttr("example_resource.example", "list_nested_block.1.list_attribute.0", "list-element"),
					resource.TestCheckResourceAttr("example_resource.example", "list_nested_block.1.list_attribute.1", "list-element"),
					resource.TestCheckResourceAttr("example_resource.example", "list_nested_block.1.list_nested_nested_block.0.bool_attribute", "true"),
					// Set Nested Block
					resource.TestCheckResourceAttr("example_resource.example", "set_nested_block.0.map_attribute.map-key-1", "map-value-1"),
					resource.TestCheckResourceAttr("example_resource.example", "set_nested_block.0.number_attribute", "179769313486231570000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"),
					resource.TestCheckResourceAttr("example_resource.example", "set_nested_block.0.object_attribute.bool_attribute", "false"),
					resource.TestCheckResourceAttr("example_resource.example", "set_nested_block.0.object_attribute.float64_attribute", "1234.5"),
					resource.TestCheckResourceAttr("example_resource.example", "set_nested_block.0.object_attribute.int64_attribute", "9223372036854775807"),
					resource.TestCheckResourceAttr("example_resource.example", "set_nested_block.0.object_attribute.list_attribute.0", "obj-list-element"),
					resource.TestCheckResourceAttr("example_resource.example", "set_nested_block.0.object_attribute.list_attribute.1", "obj-list-element"),
					resource.TestCheckResourceAttr("example_resource.example", "set_nested_block.0.object_attribute.map_attribute.obj-map-key-1", "obj-map-value-1"),
					resource.TestCheckResourceAttr("example_resource.example", "set_nested_block.0.object_attribute.number_attribute", "179769313486231570000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"),
					resource.TestCheckResourceAttr("example_resource.example", "set_nested_block.0.object_attribute.set_attribute.0", "obj-set-element-1"),
					resource.TestCheckResourceAttr("example_resource.example", "set_nested_block.0.object_attribute.set_attribute.1", "obj-set-element-2"),
					resource.TestCheckResourceAttr("example_resource.example", "set_nested_block.0.object_attribute.string_attribute", "obj-string"),
					resource.TestCheckResourceAttr("example_resource.example", "set_nested_block.0.set_attribute.0", "set-element-1"),
					resource.TestCheckResourceAttr("example_resource.example", "set_nested_block.0.set_attribute.1", "set-element-2"),
					resource.TestCheckResourceAttr("example_resource.example", "set_nested_block.0.set_nested_nested_block.0.bool_attribute", "false"),
					resource.TestCheckResourceAttr("example_resource.example", "set_nested_block.0.set_nested_nested_block.1.bool_attribute", "true"),
					resource.TestCheckResourceAttr("example_resource.example", "set_nested_block.1.map_attribute.map-key-1", "map-value-1"),
					resource.TestCheckResourceAttr("example_resource.example", "set_nested_block.1.number_attribute", "179769313486231570000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"),
					resource.TestCheckResourceAttr("example_resource.example", "set_nested_block.1.object_attribute.bool_attribute", "true"),
					resource.TestCheckResourceAttr("example_resource.example", "set_nested_block.1.object_attribute.float64_attribute", "1234.5"),
					resource.TestCheckResourceAttr("example_resource.example", "set_nested_block.1.object_attribute.int64_attribute", "9223372036854775807"),
					resource.TestCheckResourceAttr("example_resource.example", "set_nested_block.1.object_attribute.list_attribute.0", "obj-list-element"),
					resource.TestCheckResourceAttr("example_resource.example", "set_nested_block.1.object_attribute.list_attribute.1", "obj-list-element"),
					resource.TestCheckResourceAttr("example_resource.example", "set_nested_block.1.object_attribute.map_attribute.obj-map-key-1", "obj-map-value-1"),
					resource.TestCheckResourceAttr("example_resource.example", "set_nested_block.1.object_attribute.number_attribute", "179769313486231570000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"),
					resource.TestCheckResourceAttr("example_resource.example", "set_nested_block.1.object_attribute.set_attribute.0", "obj-set-element-1"),
					resource.TestCheckResourceAttr("example_resource.example", "set_nested_block.1.object_attribute.set_attribute.1", "obj-set-element-2"),
					resource.TestCheckResourceAttr("example_resource.example", "set_nested_block.1.object_attribute.string_attribute", "obj-string"),
					resource.TestCheckResourceAttr("example_resource.example", "set_nested_block.1.set_attribute.0", "set-element-1"),
					resource.TestCheckResourceAttr("example_resource.example", "set_nested_block.1.set_attribute.1", "set-element-2"),
					resource.TestCheckResourceAttr("example_resource.example", "set_nested_block.1.set_nested_nested_block.0.bool_attribute", "false"),
					resource.TestCheckResourceAttr("example_resource.example", "set_nested_block.1.set_nested_nested_block.1.bool_attribute", "true"),
					// Single Nested Block
					resource.TestCheckResourceAttr("example_resource.example", "single_nested_block.bool_attribute", "true"),
					resource.TestCheckResourceAttr("example_resource.example", "single_nested_block.float64_attribute", "1234.5"),
				),
			},
			//// ImportState testing
			//{
			//	ResourceName:      "example_resource.test",
			//	ImportState:       true,
			//	ImportStateVerify: true,
			//	// This is not normally necessary, but is here because this
			//	// example code does not have an actual upstream service.
			//	// Once the Read method is able to refresh information from
			//	// the upstream service, this can be removed.
			//	ImportStateVerifyIgnore: []string{"configurable_attribute"},
			//},
			//// Update and Read testing
			//{
			//	Config: testAccExampleResourceConfig("two"),
			//	Check: resource.ComposeAggregateTestCheckFunc(
			//		resource.TestCheckResourceAttr("example_resource.test", "configurable_attribute", "two"),
			//	),
			//},
			//// Delete testing automatically occurs in TestCase
		},
	})
}
