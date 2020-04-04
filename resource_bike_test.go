package main

import (
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccBikeIndexBike_basic(t *testing.T) {
	serial := acctest.RandString(10)

	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: strings.Replace(testAccBikeIndexBike_basic, "{{.Serial}}", serial, 1),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("bikeindex_bike.supersix", "description", "Test SuperSix"),
					resource.TestCheckResourceAttrSet("bikeindex_bike.supersix", "manufacturer_id"),
					resource.TestCheckResourceAttr("bikeindex_bike.supersix", "owner_email", "test@example.com"),
					resource.TestCheckResourceAttr("bikeindex_bike.supersix", "serial", serial),
					resource.TestCheckResourceAttr("bikeindex_bike.supersix", "frame.0.colors.0", "black"),
				),
			},
		},
	})
}

const testAccBikeIndexBike_basic = `
data "bikeindex_manufacturer" "cannondale" {
  q = "cannondale"
}

resource "bikeindex_bike" "supersix" {
	description     = "Test SuperSix"
	manufacturer_id = data.bikeindex_manufacturer.cannondale.id

	owner_email = "test@example.com"
	serial      = "{{.Serial}}"

	frame {
		colors = ["black"]
	}
}
`
