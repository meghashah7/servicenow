package servicenow

import (
	"encoding/json"
	"net/url"
)

const TableSysProperties = "sys_properties"

func (c Client) GetSysPropertiesItems(query url.Values) ([]SysPropertiesItem, error) {
	var res struct {
		Records []SysPropertiesItem
	}
	err := c.GetRecordsFor(TableSysProperties, query, &res)
	return res.Records, err
}

type SysPropertiesItem struct {
	Application      string      `json:"application"`
	Choices          string      `json:"choices"`
	Class            string      `json:"class"`
	Created          SNTime      `json:"created"`
	CreatedBy        string      `json:"created_by"`
	Description      string      `json:"description"`
	DisplayName      string      `json:"display_name"`
	Name             string      `json:"name"`
	Package          string      `json:"package"`
	Private          bool        `json:"private"`
	Protectionpolicy string      `json:"protection_policy"`
	Readroles        string      `json:"read_roles"`
	Suffix           string      `json:"suffix"`
	SysID            string      `json:"sys_id"`
	Type             string      `json:"type"`
	Updates          json.Number `json:"sys_tags"`
	Value            string      `json:"value"`
	Writeroles       string      `json:"write_roles"`
}
