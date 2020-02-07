package credentials

import (
	"github.com/bndr/gojenkins"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/vadimDidenko/terraform-provider-jenkins/jenkins/resources/credentials/util"
)

type dockerProvider struct{}

func ResourceDockerCredential() *schema.Resource {

	manager := util.CreateCredsManager(dockerProvider{})

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
			"server_ca_certificate": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
			"client_certificate": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
			"client_key": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
		},
	}
}

func (dockerProvider) Empty() interface{} {
	return gojenkins.DockerServerCredentials{}
}

func (dockerProvider) FromResourceData(d *schema.ResourceData) (interface{}, error) {

	return gojenkins.DockerServerCredentials{
		ID:                  d.Get("identifier").(string),
		Scope:               d.Get("scope").(string),
		Username:            d.Get("username").(string),
		Description:         d.Get("description").(string),
		ServerCaCertificate: d.Get("server_ca_certificate").(string),
		ClientCertificate:   d.Get("client_certificate").(string),
		ClientKey:           d.Get("client_key").(string),
	}, nil
}
