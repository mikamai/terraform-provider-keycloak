package keycloak

import (
	"fmt"
)

type Role struct {
	Id      string `json:"id,omitempty"`
	RealmId string `json:"-"`
	Name    string `json:"name"`
}

/*
func (keycloakClient *KeycloakClient) NewRole(group *Group) error {
	var createGroupUrl string

	if group.ParentId == "" {
		createGroupUrl = fmt.Sprintf("/realms/%s/groups", group.RealmId)
	} else {
		createGroupUrl = fmt.Sprintf("/realms/%s/groups/%s/children", group.RealmId, group.ParentId)
	}

	location, err := keycloakClient.post(createGroupUrl, group)
	if err != nil {
		return err
	}

	group.Id = getIdFromLocationHeader(location)

	return nil
}*/

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
