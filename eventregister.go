package servicenow

import (
	"encoding/json"
	"net/url"
)

const TableEventRegister = "sysevent_register"

func (c Client) GetEventRegisterItems(query url.Values) ([]EventRegisterItem, error) {
	var res struct {
		Records []EventRegisterItem
	}
	err := c.GetRecordsFor(TableEventRegister, query, &res)
	return res.Records, err
}

type EventRegisterItem struct {
	Application      string      `json:"application"`
	Class            string      `json:"class"`
	Created          SNTime      `json:"created"`
	CreatedBy        string      `json:"created_by"`
	Description      string      `json:"description"`
	DisplayName      string      `json:"display_name"`
	EventName        string      `json:"event_name"`
	FiredBy          string      `json:"fired_by"`
	Package          string      `json:"package"`
	ProtectionPolicy string      `json:"protection_policy"`
	SysID            string      `json:"sys_id"`
	Table            string      `json:"table"`
	UpdateName       string      `json:"update_name"`
	Updated          SNTime      `json:"updated"`
	UpdatedBy        string      `json:"updated_by"`
	Updates          json.Number `json:"updates"`
}
