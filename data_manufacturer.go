package main

import (
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceManufacturer() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManufacturerRead,

		Schema: map[string]*schema.Schema{
			"q": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"company_url": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"frame_marker": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"image": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"short_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"slug": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceManufacturerRead(d *schema.ResourceData, m interface{}) error {
	bi := m.(*Config).client
	manufacturer, err := bi.Manufacturers.Get(d.Get("q").(string))

	if err != nil {
		return fmt.Errorf("Failed to load manufacturer: %v", err)
	}

	d.SetId(strconv.Itoa(manufacturer.ID))
	d.Set("name", manufacturer.Name)
	d.Set("company_url", manufacturer.CompanyURL)
	d.Set("frame_marker", manufacturer.FrameMarker)
	d.Set("image", manufacturer.Image)
	d.Set("description", manufacturer.Description)
	d.Set("short_name", manufacturer.ShortName)
	d.Set("slug", manufacturer.Slug)

	return nil
}
