package main

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccBikeIndexDataSourceManufacturer_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccBikeIndexDataSourceManufacturer_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("bike_index_manufacturer.cannondale", "name", "Cannondale"),
					resource.TestCheckResourceAttr("bike_index_manufacturer.cannondale", "slug", "cannondale"),
				),
			},
		},
	})
}

const testAccBikeIndexDataSourceManufacturer_basic = `
data "bike_index_manufacturer" "cannondale" {
  q = "cannondale"
}
`
