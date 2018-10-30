package servicenow

import (
	"encoding/json"
	"net/url"
)

const TableSecurityIncident = "sn_si_incident"

func (c Client) GetSecurityIncident(query url.Values) ([]SecurityIncidentItem, error) {
	var res struct {
		Records []SecurityIncidentItem
	}
	err := c.GetRecordsFor(TableSecurityIncident, query, &res)
	return res.Records, err
}

type SecurityIncidentItem struct {
	Status                 string      `json:"__status"`
	Opened                 SNTime      `json:"opened"`
	ShortDescription       string      `json:"short_description"`
	Caller                 string      `json:"caller"`
	Priority               string      `json:"priority"`
	State                  string      `json:"state"`
	Category               string      `json:"cagtegory"`
	AssignmentGroup        string      `json:"assignment_group"`
	AssignedTo             string      `json:"assigned_to"`
	Updated                SNTime      `json:"updated"`
	UpdatedBy              string      `json:"updated_by"`
	Active                 bool        `json:"active,string"`
	ActivityDue            string      `json:"activity_due"`
	ActualEnd              SNTime      `json:"actual_end"`
	ActualStart            SNTime      `json:"actual_start"`
	AdditionalAssigneeList string      `json:"additional_assignee_list"`
	AdditionalComments     string      `json:"additional_comments"`
	Approval               string      `json:"approval"`
	ApprovalHistory        string      `json:"approval_history"`
	ApprovalSet            string      `json:"approval_set"`
	BusinessDuration       string      `json:"business_duration"`
	BusinessResolveTime    string      `json:"business_resolve_time"`
	BusinessService        string      `json:"business_service"`
	CausedByChange         string      `json:"caused_by_change"`
	ChangeRequest          string      `json:"change_request"`
	ChildIncidents         string      `json:"child_incidents"`
	Closed                 SNTime      `json:"closed"`
	ClosedBy               string      `json:"closed_by"`
	CommentsAndWorkNotes   string      `json:"comments_and_work_notes"`
	Company                string      `json:"company"`
	ConfigurationItem      string      `json:"configuratin_item"`
	ContactType            string      `json:"contact_type"`
	CorrelationDisplay     string      `json:"correlation_display"`
	CorrelationID          string      `json:"correlation_id"`
	Created                SNTime      `json:"created"`
	CreatedBy              string      `json:"created_by"`
	DeliveryPlan           string      `json:"delivery_plan"`
	DeliveryTask           string      `json:"delivery_task"`
	Description            string      `json:"description"`
	Domain                 string      `json:"domain"`
	DomainPath             string      `json:"domain_path"`
	DueDate                SNTime      `json:"due_date"`
	Duration               json.Number `json:"duration"`
	Escalation             json.Number `json:"escalation"`
	ExpectedStart          SNTime      `json:"expected_start"`
	ExpectedEnd            SNTime      `json:"expected_end"`
	FollowUp               string      `json:"follow_up"`
	GroupList              string      `json:"group_list"`
	Impact                 json.Number `json:"impact"`
	IncidentState          string      `json:"incident_state"`
	Knowledge              bool        `json:"knowledge,string"`
	Location               string      `json:"location"`
	MadeSLA                bool        `json:"made_sla,string"`
	Number                 string      `json:"number"`
	Notify                 string      `json:"notify"`
	OnHoldReason           string      `json:"on_hold_reason"`
	OpenedBy               string      `json:"opened_by"`
	Order                  string      `json:"order"`
	Parent                 string      `json:"parent"`
	ParentIncident         string      `json:"parent_incident"`
	ReassignmentCount      json.Number `json:"reassignment_count"`
	ReopenCount            json.Number `json:"reopen_count"`
	ResolutionCode         string      `json:"resolution_code"`
	ResolutionNotes        string      `json:"resolution_notes"`
	ResolveTime            string      `json:"resolve_time"`
	Resolved               string      `json:"resolved"`
	ResolvedBy             string      `json:"resolved_by"`
	SLADue                 string      `json:"sla_due"`
	Severity               string      `json:"severity"`
	Subcategory            string      `json:"subcategory"`
	Tags                   string      `json:"tags"`
	TaskType               string      `json:"task_type"`
	TimeWorked             string      `json:"time_worked"`
	Updates                string      `json:"updates"`
	UponApproval           string      `json:"upon_approval"`
	UponReject             string      `json:"upon_reject"`
	Urgency                string      `json:"urgency"`
	UserInput              string      `json:"user_input"`
	WatchList              string      `json:"watch_list"`
	WorkNotes              string      `json:"work_notes"`
	WorkNotesList          string      `json:"work_notes_list"`
}
