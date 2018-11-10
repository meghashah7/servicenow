package servicenow

import (
	"encoding/json"
	"net/url"
)

const TableUserGroup = "sys_user_group"

// GetUsers returns all users that match a query.
func (c Client) GetUserGroups(query url.Values) ([]UserGroupItem, error) {

	var res struct {
		Records []UserGroupItem
	}
	err := c.GetRecordsFor(TableUserGroup, query, &res)
	return res.Records, err
}

type UserGroupItem struct {
	Created     SNTime      `json:"created"`
	CreatedBy   string      `json:"created_by"`
	Description string      `json:"description"`
	GroupEmail  string      `json:"group_email"`
	Manager     string      `json:"manager"`
	Name        string      `json:"name"`
	Roles       string      `json:"roles"`
	SysID       string      `json:"sys_id"`
	Updated     SNTime      `json:"updated"`
	UpdatedBy   string      `json:"updated_by"`
	Updates     json.Number `json:"updates"`
}
