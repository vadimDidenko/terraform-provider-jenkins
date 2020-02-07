package credentials

import (
	"github.com/bndr/gojenkins"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/vadimDidenko/terraform-provider-jenkins/jenkins/resources/credentials/util"
)

type usernameProvider struct{}

func ResourceUsernameCredential() *schema.Resource {

	manager := util.CreateCredsManager(usernameProvider{})

	return &schema.Resource{
		Create: manager.ResourceServerCreate,
		Read:   manager.ResourceServerRead,
		Update: manager.ResourceServerUpdate,
		Delete: manager.ResourceServerDelete,

		Schema: map[string]*schema.Schema{
			"username": {
				Type:     schema.TypeString,
				Required: true,
			},
			"identifier": {
				Type:     schema.TypeString,
				Required: true,
			},
			"password": {
				Type:      schema.TypeString,
				Required:  true,
				Sensitive: true,
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

func (usernameProvider) Empty() interface{} {
	return gojenkins.UsernameCredentials{}
}

func (usernameProvider) FromResourceData(d *schema.ResourceData) (interface{}, error) {

	return gojenkins.UsernameCredentials{
		ID:          d.Get("identifier").(string),
		Scope:       d.Get("scope").(string),
		Username:    d.Get("username").(string),
		Password:    d.Get("password").(string),
		Description: d.Get("description").(string),
	}, nil
}
