package jenkins

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

var testJenkinsProvider *schema.Provider
var testJenkinsProviders map[string]terraform.ResourceProvider
var testJenkinsProviderFactories func(providers *[]*schema.Provider) map[string]terraform.ResourceProviderFactory
var testJenkinsProviderFunc func() *schema.Provider

func init() {
	testJenkinsProvider = Provider().(*schema.Provider)
	testJenkinsProviders = map[string]terraform.ResourceProvider{
		"jenkins": testJenkinsProvider,
	}
	testJenkinsProviderFactories = func(providers *[]*schema.Provider) map[string]terraform.ResourceProviderFactory {
		return map[string]terraform.ResourceProviderFactory{
			"jenkins": func() (terraform.ResourceProvider, error) {
				p := Provider()
				*providers = append(*providers, p.(*schema.Provider))
				return p, nil
			},
		}
	}
	testJenkinsProviderFunc = func() *schema.Provider { return testJenkinsProvider }
}

func TestProvider(t *testing.T) {
	if err := Provider().(*schema.Provider).InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_impl(t *testing.T) {
	var _ terraform.ResourceProvider = Provider()
}
