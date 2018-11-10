package servicenow

import (
	"net/url"
)

const TableUserRole = "sys_user_role"

// GetUsers returns all users that match a query.
func (c Client) GetUserRoles(query url.Values) ([]UserRoleItem, error) {

	var res struct {
		Records []UserRoleItem
	}
	err := c.GetRecordsFor(TableUserRole, query, &res)
	return res.Records, err
}

type UserRoleItem struct {
	Application   string `json:"application"`
	Created       SNTime `json:"created"`
	CreatedBy     string `json:"created_by"`
	Department    string `json:"department"`
	Description   string `json:"description"`
	DisplayName   string `json:"display_name"`
	IncludesRoles string `json:"includes_roles"`
	Name          string `json:"name"`
	SysID         string `json:"sys_id"`
	Updated       SNTime `json:"updated"`
	UpdateName    string `json:"update_name"`
}
