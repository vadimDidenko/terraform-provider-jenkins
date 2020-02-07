package job

import (
	"errors"

	"github.com/bndr/gojenkins"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func ResourceXmlJob() *schema.Resource {
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
			"xml": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceServerCreate(d *schema.ResourceData, m interface{}) error {

	client := m.(*gojenkins.Jenkins)
	name := d.Get("name").(string)

	xml := d.Get("xml").(string)
	_, err := client.CreateJob(xml, name)
	if err != nil {
		return err
	}

	d.SetId(name)
	resourceServerRead(d, m)
	return nil
}

func resourceServerRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*gojenkins.Jenkins)
	name := d.Get("name").(string)

	job, err := client.GetJob(name)
	if err != nil {
		return err
	}

	xml, err := job.GetConfig()
	if err != nil {
		return err
	}

	d.Set("xml", xml)
	return nil
}

func resourceServerUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*gojenkins.Jenkins)
	name := d.Get("name").(string)

	job, err := client.GetJob(name)
	if err != nil {
		return err
	}

	xml, err := job.GetConfig()
	if err != nil {
		return err
	}

	d.Set("xml", xml)
	return nil
}

func resourceServerDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*gojenkins.Jenkins)
	name := d.Get("name").(string)

	deleted, err := client.DeleteJob(name)
	if err != nil {
		return err
	}

	if !deleted {
		return errors.New("Could not delete job")
	}

	return nil
}
