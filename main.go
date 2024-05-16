package main

import (
	"fmt"
	"os"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

var (
	DomainName string
	Username   string
	Password   string
	Customerid string
)

func main() {

	// Perform authentication later
	//err := authenticate()
	//if err != nil {
	//	panic(fmt.Sprintf("failed to authenticate: %v", err))
	//} do not auth here - we need to get credentials first

	// Create a new plugin with a specific provider
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() *schema.Provider {
			return &schema.Provider{
				Schema: map[string]*schema.Schema{
					"domain_name": {
						Type:        schema.TypeString,
						Optional:    true,
						Default:     "radar.wandera.com",
						Description: "The JSC domain.",
					},
					"username": {
						Type:        schema.TypeString,
						Required:    true,
						Description: "The JSC username used for authentication.",
					},
					"password": {
						Type:        schema.TypeString,
						Required:    true,
						Sensitive:   true,
						Description: "The JSC password used for authentication.",
					},
					"customerid": {
						Type:        schema.TypeString,
						Required:    true,
						Description: "The required customerID.",
					},
				},
				// Define the resources that this provider manages
				ResourcesMap: map[string]*schema.Resource{
					"jsc_oktaidp":   resourceOktaIdp(),
					"jsc_uemc":      resourceUEMC(),
					"jsc_blockpage": resourceBlockPage(),
				},
				ConfigureFunc: providerConfigure,
			}

		},
	})

}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	// Read the domain_name field from the configuration and assign it to domainName
	DomainName = d.Get("domain_name").(string)

	// Assign username and password from configuration
	Username = d.Get("username").(string)
	Password = d.Get("password").(string)
	Customerid = d.Get("customerid").(string)

	return nil, nil
}

// GetClientPassword retrieves the 'password' value from the Terraform configuration.
// If it's not present in the configuration, it attempts to fetch it from the JAMFPRO_PASSWORD environment variable.
func GetClientPassword(d *schema.ResourceData) (string, error) {
	password := d.Get("password").(string)
	if password == "" {
		password = os.Getenv("JSC_PASSWORD")
		if password == "" {
			return "", fmt.Errorf("password must be provided either as an environment variable (JAMFPRO_PASSWORD) or in the Terraform configuration")
		}
	}
	return password, nil
}
