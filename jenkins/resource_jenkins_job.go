package jenkins

import (
	"errors"

	"github.com/bndr/gojenkins"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceXmlJob() *schema.Resource {
	return &schema.Resource{
		Create: resourceXmlJobCreate,
		Read:   resourceXmlJobRead,
		Update: resourceXmlJobUpdate,
		Delete: resourceXmlJobDelete,

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

func resourceXmlJobCreate(d *schema.ResourceData, m interface{}) error {

	client := m.(*gojenkins.Jenkins)
	name := d.Get("name").(string)

	xml := d.Get("xml").(string)
	_, err := client.CreateJob(xml, name)
	if err != nil {
		return err
	}

	d.SetId(name)
	_ = resourceXmlJobRead(d, m)
	return nil
}

func resourceXmlJobRead(d *schema.ResourceData, m interface{}) error {
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

	_ = d.Set("xml", xml)
	return nil
}

func resourceXmlJobUpdate(d *schema.ResourceData, m interface{}) error {
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

	_ = d.Set("xml", xml)
	return nil
}

func resourceXmlJobDelete(d *schema.ResourceData, m interface{}) error {
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
