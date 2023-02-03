package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccExampleDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `data "example_datasource" "example" {
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
					resource.TestCheckResourceAttr("data.example_datasource.example", "id", "example-id"),

					// Simple/primitive attributes
					resource.TestCheckResourceAttr("data.example_datasource.example", "bool_attribute", "true"),
					resource.TestCheckResourceAttr("data.example_datasource.example", "float64_attribute", "1234.5"),
					resource.TestCheckResourceAttr("data.example_datasource.example", "int64_attribute", "9223372036854775807"),
					resource.TestCheckResourceAttr("data.example_datasource.example", "list_attribute.0", "list-element"),
					resource.TestCheckResourceAttr("data.example_datasource.example", "list_attribute.1", "list-element"),
					resource.TestCheckResourceAttr("data.example_datasource.example", "map_attribute.map-key-1", "map-value-1"),
					resource.TestCheckResourceAttr("data.example_datasource.example", "number_attribute", "179769313486231570000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"),
					resource.TestCheckResourceAttr("data.example_datasource.example", "object_attribute.bool_attribute", "true"),
					resource.TestCheckResourceAttr("data.example_datasource.example", "object_attribute.float64_attribute", "1234.5"),
					resource.TestCheckResourceAttr("data.example_datasource.example", "object_attribute.int64_attribute", "9223372036854775807"),
					resource.TestCheckResourceAttr("data.example_datasource.example", "object_attribute.list_attribute.0", "obj-list-element"),
					resource.TestCheckResourceAttr("data.example_datasource.example", "object_attribute.list_attribute.1", "obj-list-element"),
					resource.TestCheckResourceAttr("data.example_datasource.example", "object_attribute.map_attribute.obj-map-key-1", "obj-map-value-1"),
					resource.TestCheckResourceAttr("data.example_datasource.example", "object_attribute.number_attribute", "179769313486231570000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"),
					resource.TestCheckResourceAttr("data.example_datasource.example", "object_attribute.set_attribute.0", "obj-set-element-1"),
					resource.TestCheckResourceAttr("data.example_datasource.example", "object_attribute.set_attribute.1", "obj-set-element-2"),
					resource.TestCheckResourceAttr("data.example_datasource.example", "object_attribute.string_attribute", "obj-string"),
					resource.TestCheckResourceAttr("data.example_datasource.example", "set_attribute.0", "set-element-1"),
					resource.TestCheckResourceAttr("data.example_datasource.example", "set_attribute.1", "set-element-2"),
					resource.TestCheckResourceAttr("data.example_datasource.example", "string_attribute", "string"),

					// Nested Attributes
					// List Nested Attribute
					resource.TestCheckResourceAttr("data.example_datasource.example", "list_nested_attribute.0.int64_attribute", "9223372036854775807"),
					resource.TestCheckResourceAttr("data.example_datasource.example", "list_nested_attribute.0.list_attribute.0", "list-element"),
					resource.TestCheckResourceAttr("data.example_datasource.example", "list_nested_attribute.0.list_attribute.1", "list-element"),
					// Map Nested Attribute
					resource.TestCheckResourceAttr("data.example_datasource.example", "map_nested_attribute.one.map_attribute.map-key-1", "map-value-1"),
					resource.TestCheckResourceAttr("data.example_datasource.example", "map_nested_attribute.one.number_attribute", "179769313486231570000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"),
					resource.TestCheckResourceAttr("data.example_datasource.example", "map_nested_attribute.two.map_attribute.map-key-1", "map-value-1"),
					resource.TestCheckResourceAttr("data.example_datasource.example", "map_nested_attribute.two.number_attribute", "179769313486231570000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"),
					// Set Nested Attribute
					resource.TestCheckResourceAttr("data.example_datasource.example", "set_nested_attribute.0.object_attribute.bool_attribute", "false"),
					resource.TestCheckResourceAttr("data.example_datasource.example", "set_nested_attribute.0.object_attribute.float64_attribute", "1234.5"),
					resource.TestCheckResourceAttr("data.example_datasource.example", "set_nested_attribute.0.object_attribute.int64_attribute", "9223372036854775807"),
					resource.TestCheckResourceAttr("data.example_datasource.example", "set_nested_attribute.0.object_attribute.list_attribute.0", "obj-list-element"),
					resource.TestCheckResourceAttr("data.example_datasource.example", "set_nested_attribute.0.object_attribute.list_attribute.1", "obj-list-element"),
					resource.TestCheckResourceAttr("data.example_datasource.example", "set_nested_attribute.0.object_attribute.map_attribute.obj-map-key-1", "obj-map-value-1"),
					resource.TestCheckResourceAttr("data.example_datasource.example", "set_nested_attribute.0.object_attribute.number_attribute", "179769313486231570000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"),
					resource.TestCheckResourceAttr("data.example_datasource.example", "set_nested_attribute.0.object_attribute.set_attribute.0", "obj-set-element-1"),
					resource.TestCheckResourceAttr("data.example_datasource.example", "set_nested_attribute.0.object_attribute.set_attribute.1", "obj-set-element-2"),
					resource.TestCheckResourceAttr("data.example_datasource.example", "set_nested_attribute.0.object_attribute.string_attribute", "obj-string"),
					resource.TestCheckResourceAttr("data.example_datasource.example", "set_nested_attribute.0.set_attribute.0", "set-element-1"),
					resource.TestCheckResourceAttr("data.example_datasource.example", "set_nested_attribute.0.set_attribute.1", "set-element-2"),
					resource.TestCheckResourceAttr("data.example_datasource.example", "set_nested_attribute.0.string_attribute", "string"),
					resource.TestCheckResourceAttr("data.example_datasource.example", "set_nested_attribute.1.object_attribute.bool_attribute", "true"),
					resource.TestCheckResourceAttr("data.example_datasource.example", "set_nested_attribute.1.object_attribute.float64_attribute", "1234.5"),
					resource.TestCheckResourceAttr("data.example_datasource.example", "set_nested_attribute.1.object_attribute.int64_attribute", "9223372036854775807"),
					resource.TestCheckResourceAttr("data.example_datasource.example", "set_nested_attribute.1.object_attribute.list_attribute.0", "obj-list-element"),
					resource.TestCheckResourceAttr("data.example_datasource.example", "set_nested_attribute.1.object_attribute.list_attribute.1", "obj-list-element"),
					resource.TestCheckResourceAttr("data.example_datasource.example", "set_nested_attribute.1.object_attribute.map_attribute.obj-map-key-1", "obj-map-value-1"),
					resource.TestCheckResourceAttr("data.example_datasource.example", "set_nested_attribute.1.object_attribute.number_attribute", "179769313486231570000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"),
					resource.TestCheckResourceAttr("data.example_datasource.example", "set_nested_attribute.1.object_attribute.set_attribute.0", "obj-set-element-1"),
					resource.TestCheckResourceAttr("data.example_datasource.example", "set_nested_attribute.1.object_attribute.set_attribute.1", "obj-set-element-2"),
					resource.TestCheckResourceAttr("data.example_datasource.example", "set_nested_attribute.1.object_attribute.string_attribute", "obj-string"),
					resource.TestCheckResourceAttr("data.example_datasource.example", "set_nested_attribute.1.set_attribute.0", "set-element-1"),
					resource.TestCheckResourceAttr("data.example_datasource.example", "set_nested_attribute.1.set_attribute.1", "set-element-2"),
					resource.TestCheckResourceAttr("data.example_datasource.example", "set_nested_attribute.1.string_attribute", "string"),
					// Single Nested Attribute
					resource.TestCheckResourceAttr("data.example_datasource.example", "single_nested_attribute.bool_attribute", "true"),
					resource.TestCheckResourceAttr("data.example_datasource.example", "single_nested_attribute.float64_attribute", "1234.5"),

					// Blocks
					// List Nested Block
					resource.TestCheckResourceAttr("data.example_datasource.example", "list_nested_block.0.bool_attribute", "true"),
					resource.TestCheckResourceAttr("data.example_datasource.example", "list_nested_block.0.float64_attribute", "1234.5"),
					resource.TestCheckResourceAttr("data.example_datasource.example", "list_nested_block.0.int64_attribute", "9223372036854775807"),
					resource.TestCheckResourceAttr("data.example_datasource.example", "list_nested_block.0.list_attribute.0", "list-element"),
					resource.TestCheckResourceAttr("data.example_datasource.example", "list_nested_block.0.list_attribute.1", "list-element"),
					resource.TestCheckResourceAttr("data.example_datasource.example", "list_nested_block.0.list_nested_nested_block.0.bool_attribute", "true"),
					resource.TestCheckResourceAttr("data.example_datasource.example", "list_nested_block.1.bool_attribute", "true"),
					resource.TestCheckResourceAttr("data.example_datasource.example", "list_nested_block.1.float64_attribute", "1234.5"),
					resource.TestCheckResourceAttr("data.example_datasource.example", "list_nested_block.1.int64_attribute", "9223372036854775807"),
					resource.TestCheckResourceAttr("data.example_datasource.example", "list_nested_block.1.list_attribute.0", "list-element"),
					resource.TestCheckResourceAttr("data.example_datasource.example", "list_nested_block.1.list_attribute.1", "list-element"),
					resource.TestCheckResourceAttr("data.example_datasource.example", "list_nested_block.1.list_nested_nested_block.0.bool_attribute", "true"),
					// Set Nested Block
					resource.TestCheckResourceAttr("data.example_datasource.example", "set_nested_block.0.map_attribute.map-key-1", "map-value-1"),
					resource.TestCheckResourceAttr("data.example_datasource.example", "set_nested_block.0.number_attribute", "179769313486231570000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"),
					resource.TestCheckResourceAttr("data.example_datasource.example", "set_nested_block.0.object_attribute.bool_attribute", "false"),
					resource.TestCheckResourceAttr("data.example_datasource.example", "set_nested_block.0.object_attribute.float64_attribute", "1234.5"),
					resource.TestCheckResourceAttr("data.example_datasource.example", "set_nested_block.0.object_attribute.int64_attribute", "9223372036854775807"),
					resource.TestCheckResourceAttr("data.example_datasource.example", "set_nested_block.0.object_attribute.list_attribute.0", "obj-list-element"),
					resource.TestCheckResourceAttr("data.example_datasource.example", "set_nested_block.0.object_attribute.list_attribute.1", "obj-list-element"),
					resource.TestCheckResourceAttr("data.example_datasource.example", "set_nested_block.0.object_attribute.map_attribute.obj-map-key-1", "obj-map-value-1"),
					resource.TestCheckResourceAttr("data.example_datasource.example", "set_nested_block.0.object_attribute.number_attribute", "179769313486231570000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"),
					resource.TestCheckResourceAttr("data.example_datasource.example", "set_nested_block.0.object_attribute.set_attribute.0", "obj-set-element-1"),
					resource.TestCheckResourceAttr("data.example_datasource.example", "set_nested_block.0.object_attribute.set_attribute.1", "obj-set-element-2"),
					resource.TestCheckResourceAttr("data.example_datasource.example", "set_nested_block.0.object_attribute.string_attribute", "obj-string"),
					resource.TestCheckResourceAttr("data.example_datasource.example", "set_nested_block.0.set_attribute.0", "set-element-1"),
					resource.TestCheckResourceAttr("data.example_datasource.example", "set_nested_block.0.set_attribute.1", "set-element-2"),
					resource.TestCheckResourceAttr("data.example_datasource.example", "set_nested_block.0.set_nested_nested_block.0.bool_attribute", "false"),
					resource.TestCheckResourceAttr("data.example_datasource.example", "set_nested_block.0.set_nested_nested_block.1.bool_attribute", "true"),
					resource.TestCheckResourceAttr("data.example_datasource.example", "set_nested_block.1.map_attribute.map-key-1", "map-value-1"),
					resource.TestCheckResourceAttr("data.example_datasource.example", "set_nested_block.1.number_attribute", "179769313486231570000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"),
					resource.TestCheckResourceAttr("data.example_datasource.example", "set_nested_block.1.object_attribute.bool_attribute", "true"),
					resource.TestCheckResourceAttr("data.example_datasource.example", "set_nested_block.1.object_attribute.float64_attribute", "1234.5"),
					resource.TestCheckResourceAttr("data.example_datasource.example", "set_nested_block.1.object_attribute.int64_attribute", "9223372036854775807"),
					resource.TestCheckResourceAttr("data.example_datasource.example", "set_nested_block.1.object_attribute.list_attribute.0", "obj-list-element"),
					resource.TestCheckResourceAttr("data.example_datasource.example", "set_nested_block.1.object_attribute.list_attribute.1", "obj-list-element"),
					resource.TestCheckResourceAttr("data.example_datasource.example", "set_nested_block.1.object_attribute.map_attribute.obj-map-key-1", "obj-map-value-1"),
					resource.TestCheckResourceAttr("data.example_datasource.example", "set_nested_block.1.object_attribute.number_attribute", "179769313486231570000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"),
					resource.TestCheckResourceAttr("data.example_datasource.example", "set_nested_block.1.object_attribute.set_attribute.0", "obj-set-element-1"),
					resource.TestCheckResourceAttr("data.example_datasource.example", "set_nested_block.1.object_attribute.set_attribute.1", "obj-set-element-2"),
					resource.TestCheckResourceAttr("data.example_datasource.example", "set_nested_block.1.object_attribute.string_attribute", "obj-string"),
					resource.TestCheckResourceAttr("data.example_datasource.example", "set_nested_block.1.set_attribute.0", "set-element-1"),
					resource.TestCheckResourceAttr("data.example_datasource.example", "set_nested_block.1.set_attribute.1", "set-element-2"),
					resource.TestCheckResourceAttr("data.example_datasource.example", "set_nested_block.1.set_nested_nested_block.0.bool_attribute", "false"),
					resource.TestCheckResourceAttr("data.example_datasource.example", "set_nested_block.1.set_nested_nested_block.1.bool_attribute", "true"),
					// Single Nested Block
					resource.TestCheckResourceAttr("data.example_datasource.example", "single_nested_block.bool_attribute", "true"),
					resource.TestCheckResourceAttr("data.example_datasource.example", "single_nested_block.float64_attribute", "1234.5"),
				),
			},
		},
	})
}

const testAccExampleDataSourceConfig = `
data "example_datasource" "test" {
  configurable_attribute = "example"
}
`
