package plugins

import (
	"github.com/bndr/gojenkins"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func Plugin() *schema.Resource {
	return &schema.Resource{
		Create: resourceServerCreate,
		Read:   resourceServerRead,
		Update: resourceServerUpdate,
		Delete: resourceServerDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"version": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceServerCreate(d *schema.ResourceData, m interface{}) error {

	client := m.(*gojenkins.Jenkins)
	name := d.Get("name").(string)
	version := d.Get("version").(string)

	err := client.InstallPlugin(name, version)
	if err != nil {
		return err
	}

	d.SetId(name)
	d.Set("version", version)
	return nil
}

func resourceServerRead(d *schema.ResourceData, m interface{}) error {

	client := m.(*gojenkins.Jenkins)
	name := d.Get("name").(string)

	plugs, err := client.GetPlugins(1)
	if err != nil {
		return err
	}

	plug := plugs.Contains(name)
	if plug == nil {
		d.SetId("")
		return nil
	}

	d.Set("version", plug.Version)
	return nil
}

func resourceServerUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceServerCreate(d, m)
}

func resourceServerDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*gojenkins.Jenkins)
	name := d.Get("name").(string)
	return client.UninstallPlugin(name)
}
