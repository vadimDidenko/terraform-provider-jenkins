package jenkins

import (
	"github.com/bndr/gojenkins"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func jenkinsPlugin() *schema.Resource {
	return &schema.Resource{
		Create: resourcePluginCreate,
		Read:   resourcePluginRead,
		Update: resourcePluginUpdate,
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

func resourcePluginCreate(d *schema.ResourceData, m interface{}) error {

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

func resourcePluginRead(d *schema.ResourceData, m interface{}) error {

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

func resourcePluginUpdate(d *schema.ResourceData, m interface{}) error {
	return resourcePluginCreate(d, m)
}

func resourceServerDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*gojenkins.Jenkins)
	name := d.Get("name").(string)
	return client.UninstallPlugin(name)
}
