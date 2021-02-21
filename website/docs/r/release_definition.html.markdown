

# azuredevops_release_definition

Manages a Release Definition within Azure DevOps.

## Example Usage

### Tfs

```hcl
resource "azuredevops_project" "project" {
  name       = "Sample Project"
  visibility         = "private"
  version_control    = "Git"
  work_item_template = "Agile"
}

resource "azuredevops_git_repository" "repository" {
  project_id = azuredevops_project.project.id
  name       = "Sample Repository"
  initialization {
    init_type = "Clean"
  }
}

resource "azuredevops_variable_group" "vars" {
  project_id   = azuredevops_project.project.id
  name         = "Infrastructure Pipeline Variables"
  description  = "Managed by Terraform"
  allow_access = true

  variable {
    name  = "FOO"
    value = "BAR"
  }
}

resource "azuredevops_variable_group" "vars" {
    project_id     = data.azuredevops_project.project.id
    name           = var.variable_group_name
    description    = "Managed by Terraform"
    allow_access   = true
}

resource "azuredevops_release_definition" "p" {
  project_id = data.azuredevops_project.project.id
  name       = var.release_pipeline_name
  path       = "\\"

  variable_groups = [
    azuredevops_variable_group.vars.id
  ]

  description = "Run L2 test case"

  variable {
    name  = "app_name"
    value = "XBC"
  }

  tags = {
      
  }
}
```

## Argument Reference

The following arguments are supported:

- `project_id` - (Required) The project ID or project name.
- `name` - (Optional) The name of the build definition.
- `path` - (Optional) The folder path of the build definition.


