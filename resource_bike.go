package main

import (
	"fmt"
	"strconv"

	bikeindex "github.com/bendrucker/terraform-provider-bike-index/pkg/bike-index"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"

	_ "github.com/motemen/go-loghttp/global"
)

var tireSizes = []string{"wide", "narrow"}

func resourceBike() *schema.Resource {
	return &schema.Resource{
		Create: resourceBikeCreate,
		Read:   resourceBikeRead,
		Update: resourceBikeUpdate,
		Delete: resourceBikeDelete,

		Schema: map[string]*schema.Schema{
			"owner_email": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"serial": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"year": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"manufacturer_id": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"additional_registration": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"frame": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"colors": {
							Type:     schema.TypeList,
							Elem:     &schema.Schema{Type: schema.TypeString},
							MinItems: 1,
							MaxItems: 3,
							Required: true,
						},
						"model": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"size": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"material": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"rear_tire": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice(tireSizes, false),
				Default:      "narrow",
			},
			"front_tire": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice(tireSizes, false),
				Default:      "narrow",
			},
			"rear_gears": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"front_gears": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"handlebar_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"components": {
				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"description": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"serial_number": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"type": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"manufacturer_name": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"model_name": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"year": {
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
		},
	}
}

func resourceBikeCreate(d *schema.ResourceData, m interface{}) error {
	config := m.(*Config)
	bi := config.client

	bike := &bikeindex.CreateUpdateBikeRequest{}

	bike.TestBike = config.test

	bike.Serial = d.Get("serial").(string)
	bike.Manufacturer = strconv.Itoa(d.Get("manufacturer_id").(int))
	bike.OwnerEmail = d.Get("owner_email").(string)

	bike.Description = d.Get("description").(string)
	bike.Name = d.Get("name").(string)

	bike.Color = d.Get("frame.0.colors.0").(string)
	bike.PrimaryFrameColor = d.Get("frame.0.colors.0").(string)
	if color, ok := d.GetOk("frame.0.colors.1"); ok {
		bike.SecondaryFrameColor = color.(string)
	}
	if color, ok := d.GetOk("frame.0.colors.2"); ok {
		bike.TertiaryFrameColor = color.(string)
	}

	bike.FrontTireNarrow = bikeTireNarrow(d.Get("front_tire").(string))
	bike.RearTireNarrow = bikeTireNarrow(d.Get("rear_tire").(string))

	result, err := bi.Bikes.Create(bike)
	if err != nil {
		return fmt.Errorf("failed to create bike: %v", err)
	}

	d.SetId(strconv.Itoa(result.ID))

	return resourceBikeRead(d, m)
}

func resourceBikeRead(d *schema.ResourceData, m interface{}) error {
	bi := m.(*Config).client

	bike, err := bi.Bikes.Get(d.Id())
	if err != nil {
		return fmt.Errorf("Failed to load bike: %v", err)
	}

	d.Set("description", bike.Description)
	d.Set("serial", bike.Serial)
	d.Set("year", bike.Year)
	d.Set("manufacturer_id", bike.ManufacturerID)
	d.Set("name", bike.Name)
	d.Set("additional_registration", bike.AdditionalRegistration)

	d.Set("frame.0.colors", bike.FrameColors)
	d.Set("frame.0.model", bike.FrameModel)
	d.Set("frame.0.size", bike.FrameSize)
	d.Set("frame.0.material", bike.FrameMaterialSlug)

	d.Set("rear_tire", bikeTireWidth(bike.RearTireNarrow))
	d.Set("front_tire", bikeTireWidth(bike.FrontTireNarrow))

	d.Set("rear_gears", bike.RearGearTypeSlug)
	d.Set("front_gears", bike.FrontGearTypeSlug)
	d.Set("handlebar_type", bike.HandlebarTypeSlug)

	d.Set("components", flattenComponents(bike.Components))

	return nil
}

func resourceBikeUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceBikeRead(d, m)
}

func resourceBikeDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}

func bikeTireWidth(narrow bool) string {
	if narrow {
		return "narrow"
	}

	return "wide"
}

func bikeTireNarrow(width string) bool {
	return width == "narrow"
}

func flattenComponents(components []*bikeindex.Component) []map[string]interface{} {
	out := make([]map[string]interface{}, len(components), len(components))

	for i, component := range components {
		out[i] = map[string]interface{}{
			"description":       component.Description,
			"serial_number":     component.SerialNumber,
			"manufacturer_name": component.ManufacturerName,
			"model_name":        component.ModelName,
			"year":              component.Year,
		}
	}

	return out
}
