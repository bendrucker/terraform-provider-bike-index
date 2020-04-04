package main

import (
	"os"

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
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("BIKE_INDEX_API_TOKEN", nil),
			},
			"test": {
				Type:     schema.TypeBool,
				Optional: true,
				DefaultFunc: func() (interface{}, error) {
					return os.Getenv("TF_ACC") == "1", nil
				},
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"bikeindex_bike": resourceBike(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"bikeindex_manufacturer": dataSourceManufacturer(),
		},
		ConfigureFunc: providerConfigure,
	}
}

type Config struct {
	client *bikeindex.Client
	test   bool
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	url := d.Get("base_url").(string)
	token := d.Get("token").(string)
	test := d.Get("test").(bool)

	return &Config{
		client: bikeindex.New(bikeindex.Config{
			URL:   url,
			Token: token,
		}),
		test: test,
	}, nil
}
