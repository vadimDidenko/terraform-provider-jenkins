package credentials

import (
	"github.com/bndr/gojenkins"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/vadimDidenko/terraform-provider-jenkins/jenkins/resources/credentials/util"
)

type secretProvider struct{}

func ResourceSecretCredential() *schema.Resource {

	manager := util.CreateCredsManager(secretProvider{})

	return &schema.Resource{
		Create: manager.ResourceServerCreate,
		Read:   manager.ResourceServerRead,
		Update: manager.ResourceServerUpdate,
		Delete: manager.ResourceServerDelete,

		Schema: map[string]*schema.Schema{
			"secret": {
				Type:     schema.TypeString,
				Required: true,
			},
			"identifier": {
				Type:     schema.TypeString,
				Required: true,
			},
			"domain": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "_",
			},
			"jobpath": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
			"scope": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "global",
			},
		},
	}
}

func (secretProvider) Empty() interface{} {
	return gojenkins.UsernameCredentials{}
}

func (secretProvider) FromResourceData(d *schema.ResourceData) (interface{}, error) {

	return gojenkins.StringCredentials{
		ID:          d.Get("identifier").(string),
		Scope:       d.Get("scope").(string),
		Secret:      d.Get("secret").(string),
		Description: d.Get("description").(string),
	}, nil
}
