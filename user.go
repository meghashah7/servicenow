package servicenow

import (
	"encoding/json"
	"net/url"
)

const TableUser = "sys_user"

// GetUsers returns all users that match a query.
func (c Client) GetUsers(query url.Values) ([]UserItem, error) {

	var res struct {
		Records []UserItem
	}
	err := c.GetRecordsFor(TableUser, query, &res)
	return res.Records, err
}

type UserItem struct {
	AccumulatedRoles    string      `json:"accumulated_roles"`
	City                string      `json:"city"`
	Company             string      `json:"company"`
	Created             SNTime      `json:"created"`
	CreatedBy           string      `json:"created_by"`
	Department          string      `json:"department"`
	DescriptiveName     string      `json:"descriptive_name"`
	EmployeeNumber      string      `json:"employee_number"`
	FailedLoginAttempts json.Number `json:"failed_login_attempts"`
	FirstName           string      `json:"first_name"`
	LastName            string      `json:"last_name"`
	Email               string      `json:"email"`
	Gender              string      `json:"gender"`
	LastLogin           SNTime      `json:"last_login"`
	LastLoginDevice     string      `json:"last_login_device"`
	Manager             string      `json:"manager"`
	Name                string      `json:"name"`
	Roles               string      `json:"roles"`
	SysID               string      `json:"sys_id"`
	Updated             SNTime      `json:"updated"`
	UpdatedBy           string      `json:"updated_by"`
	Updates             json.Number `json:"updates"`
	UserID              string      `json:"user_id"`
}
