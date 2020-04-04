package main

import (
	bikeindex "github.com/bendrucker/terraform-provider-bike-index/pkg/bike-index"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"base_url": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("BIKE_INDEX_API_URL", bikeindex.BaseURL),
			},
			"token": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("BIKE_INDEX_API_TOKEN", nil),
			},
		},
		ResourcesMap:  map[string]*schema.Resource{},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	url := d.Get("base_url").(string)
	token := d.Get("token").(string)

	return bikeindex.New(url).SetToken(token), nil
}
