package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

var testAccProviders map[string]terraform.ResourceProvider

func init() {
	testAccProviders = map[string]terraform.ResourceProvider{
		"bike_index": Provider(),
	}
}
