package jenkins

import (
	"github.com/bndr/gojenkins"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func Provider() terraform.ResourceProvider {

	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"url": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "http://localhost:8080",
				Description: "host address of jenkins instance",
			},
			"username": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("JENKINS_USERNAME", ""),
				Description: "username which should be used to loginto jenkins instance",
			},
			"password": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("JENKINS_PASSWORD", ""),
				Description: "password which should be used to loginto jenkins instance",
			},
			"ca_cert": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "server self-signed certificate",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"jenkins_job_xml":             resourceXmlJob(),
			"jenkins_plugin":              jenkinsPlugin(),
			"jenkins_username_credential": resourceUsernameCredential(),
			"jenkins_secret_credential":   resourceSecretCredential(),
			"jenkins_ssh_credential":      resourceSSHCredential(),
			"jenkins_docker_credential":   resourceDockerCredential(),
		},
		ConfigureFunc: configureFunc,
	}
}

func configureFunc(rd *schema.ResourceData) (interface{}, error) {

	config := JenkinsConfig{
		URL:      rd.Get("url").(string),
		Username: rd.Get("username").(string),
		Password: rd.Get("password").(string),
		CaCert:   rd.Get("ca_cert").(string),
	}

	jenkins := gojenkins.CreateJenkins(nil, config.URL, config.Username, config.Password)

	// Provide CA certificate if server is using self-signed certificate
	if len(config.CaCert) > 0 {
		jenkins.Requester.CACert = []byte(config.CaCert)
	}

	_, err := jenkins.Init()
	return jenkins, err
}

type JenkinsConfig struct {
	URL      string
	Username string
	Password string
	CaCert   string
}
