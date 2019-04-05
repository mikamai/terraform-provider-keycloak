package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/mrparkers/terraform-provider-keycloak/keycloak"
)

func resourceKeycloakRole() *schema.Resource {
	return &schema.Resource{
		Create: resourceKeycloakRoleCreate,
		Read:   resourceKeycloakRoleRead,
		Delete: resourceKeycloakRoleDelete,
		Update: resourceKeycloakRoleUpdate,
		// This resource can be imported using {{realm}}/{{group_id}}. The Group ID is displayed in the URL when editing it from the GUI
		/*Importer: &schema.ResourceImporter{
			State: resourceKeycloakRoleImport,
		},*/
		Schema: map[string]*schema.Schema{
			"realm_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func mapFromDataToRole(data *schema.ResourceData) *keycloak.Role {
	role := &keycloak.Role{
		Id:      data.Id(),
		RealmId: data.Get("realm_id").(string),
		Name:    data.Get("name").(string),
	}
	return role
}

func mapFromRoleToData(data *schema.ResourceData, role *keycloak.Role) {
	data.SetId(role.Id)

	data.Set("realm_id", role.RealmId)
	data.Set("name", role.Name)
}

func resourceKeycloakRoleCreate(data *schema.ResourceData, meta interface{}) error {
	keycloakClient := meta.(*keycloak.KeycloakClient)
	role := mapFromDataToRole(data)
	err := keycloakClient.NewRole(role)
	if err != nil {
		return err
	}
	mapFromRoleToData(data, role)
	return resourceKeycloakRoleRead(data, meta)
}

func resourceKeycloakRoleRead(data *schema.ResourceData, meta interface{}) error {
	keycloakClient := meta.(*keycloak.KeycloakClient)
	realmId := data.Get("realm_id").(string)
	id := data.Id()
	roleName := data.Get("name").(string)
	role, err := keycloakClient.GetRole(realmId, id, roleName)
	if err != nil {
		return handleNotFoundError(err, data)
	}
	mapFromRoleToData(data, role)
	return nil
}

func resourceKeycloakRoleUpdate(data *schema.ResourceData, meta interface{}) error {
	keycloakClient := meta.(*keycloak.KeycloakClient)

	role := mapFromDataToRole(data)

	err := keycloakClient.UpdateRole(role)
	if err != nil {
		return err
	}

	mapFromRoleToData(data, role)

	return nil
}

func resourceKeycloakRoleDelete(data *schema.ResourceData, meta interface{}) error {
	keycloakClient := meta.(*keycloak.KeycloakClient)
	realmId := data.Get("realm_id").(string)
	id := data.Id()
	roleName := data.Get("name").(string)

	return keycloakClient.DeleteRole(realmId, id, roleName)

}
