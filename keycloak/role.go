package keycloak

import (
	"fmt"
)

type Role struct {
	Id      string `json:"id,omitempty"`
	RealmId string `json:"-"`
	Name    string `json:"name"`
}

func (keycloakClient *KeycloakClient) NewRole(role *Role) error {
	var createRoleUrl string
	createRoleUrl = fmt.Sprintf("/%s/clients/%s/roles", role.Name, role.Id)
	location, err := keycloakClient.post(createRoleUrl, role)
	if err != nil {
		return err
	}
	role.Id = getIdFromLocationHeader(location)
	return nil
}

func (keycloakClient *KeycloakClient) GetRole(realmId, id, roleName string) (*Role, error) {
	var role Role

	err := keycloakClient.get(fmt.Sprintf("/%s/clients/%s/roles/%s", realmId, id, roleName), &role)
	if err != nil {
		return nil, err
	}

	return &role, nil
}

func (keycloakClient *KeycloakClient) UpdateRole(role *Role) error {
	return keycloakClient.put(fmt.Sprintf("/%s/clients/%s/roles/%s", role.RealmId, role.Id, role.Name), role)
}

func (keycloakClient *KeycloakClient) DeleteRole(realmId, id, roleName string) error {
	return keycloakClient.delete(fmt.Sprintf("/%s/clients/%s/roles/%s", realmId, id, roleName))
}
