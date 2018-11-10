package servicenow

import (
	"encoding/json"
	"net/url"
)

const TableEvent = "sysevent"

func (c Client) GetEventItems(query url.Values) ([]EventItem, error) {
	var res struct {
		Records []EventItem
	}
	err := c.GetRecordsFor(TableEvent, query, &res)
	return res.Records, err
}

type EventItem struct {
	ClaimedBy          string      `json:"claimed_by"`
	Created            SNTime      `json:"created"`
	CreatedBy          string      `json:"created_by"`
	DescriptiveName    string      `json:"descriptive_name"`
	Instance           string      `json:"instance"`
	Name               string      `json:"name"`
	Parm1              string      `json:"parm1"`
	Parm2              string      `json:"parm2"`
	ProcessOn          SNTime      `json:"process_on"`
	Processed          string      `json:"processed"`
	ProcessingDuration json.Number `json:"processing_duration"`
	Queue              string      `json:"queue"`
	State              string      `json:"state"`
	SysID              string      `json:"sys_id"`
	Table              string      `json:"table"`
	Updated            SNTime      `json:"updated"`
	UpdatedBy          string      `json:"updated_by"`
	Updates            json.Number `json:"updates"`
	URI                string      `json:"uri"`
	UserID             string      `json:"user_id"`
	UserName           string      `json:"user_name"`
}
