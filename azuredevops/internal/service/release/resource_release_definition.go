package release

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/microsoft/azure-devops-go-api/azuredevops/release"
	"github.com/microsoft/terraform-provider-azuredevops/azuredevops/internal/client"
	"github.com/microsoft/terraform-provider-azuredevops/azuredevops/internal/utils/validate"
)

func ResourceReleaseDefinition() *schema.Resource {

	id := &schema.Schema{
		Type: schema.TypeInt,
		Computed: true,
	}

	rank := &schema.Schema{
		Type: schema.TypeInt,
		Computed: true,
	}

	return &schema.Resource{
		Create: resourceReleaseDefinitionCreate,
		Read: resourceReleaseDefinitionRead,
		Update: resourceReleaseDefinitionUpdate,
		Delete: resourceReleaseDefinitionDelete,
		Importer: resourceReleaseDefinitionImporter,
		Schema: map[string]*schema.Schema{
			"project_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"revision": {
				Type: schema.TypeInt,
				Computed: true,
			},
			"name": {
				Type: schema.TypeString,
				Optional: true,
				Default: "",
			},
			"path": {
				Type: schema.TypeString,
				Optional: true,
				Default: "\\",
				ValidateFunc: validate.Path,
			},
			"variable_groups": variableGroups,
			"source": {
				Type: schema.TypeString,
				Computed: ture,
			},
			"description": {
				Type: schema.TypeString,
				Optional: true,
				Default: "",
			},
			"variable": {

			},
			"release_name_format": {

			},
			"release_name_format": {

			},
			"url": {

			},
			"is_deleted": {

			},
			"tags": {

			},
			"properties": {

			},
			"comment": {

			},
			"created_on": {

			},
			"modified_on": {

			},
			"stage": {

			},
			"build_artifact": {

			},
			"triggers":
		},
	}
}

func createReleaseDefinition(clients *client.AggregatedClient, releaseDefinition *release.ReleaseDefinition, project string) (*release.ReleaseDefinition, error) {
	createdBuild, err := clients.ReleaseClient.CreateReleaseDefinition(clients.Ctx, release.CreateReleaseDefinitionArgs{
		ReleaseDefinition: releaseDefinition,
		Project: &project,
	})

	return createdBuild, err
}