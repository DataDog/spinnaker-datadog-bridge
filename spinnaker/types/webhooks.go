package types

// IncomingWebhook is a structure representing a Spinnaker echo rest Webhook
// You can view an example of the schema here:
// https://www.spinnaker.io/setup/features/notifications/#event-types
type IncomingWebhook struct {
	Details Details `json:"details"`
	Content Content `json:"content"`
}

// Details contains all of the details contained in the webhook
type Details struct {
	Source      string `json:"source"`
	Type        string `json:"type"`
	Application string `json:"application"`
	Created     string `json:"created"`
}

// Content is the main context of the given Webhook. It contains of the execution
// information and stage details as an example
type Content struct {
	ExecutionID string    `json:"executionId"`
	StartTime   Timestamp `json:"startTime"`
	EndTime     Timestamp `json:"endTime"`
	Execution   Execution `json:"execution,omitempty"`
}

// Execution represents an execution context for a spinnaker event
type Execution struct {
	Type             string         `json:"type"`
	ID               string         `json:"id"`
	Application      string         `json:"application"`
	StartTime        Timestamp      `json:"startTime"`
	EndTime          Timestamp      `json:"endTime"`
	Name             string         `json:"name"`
	Canceled         bool           `json:"cancelled"`
	CancelledBy      string         `json:"cancelledBy,omitempty"`
	PipelineConfigID string         `json:"pipelineConfigId"`
	Status           string         `json:"status"`
	Trigger          Trigger        `json:"trigger,omitempty"`
	Authentication   Authentication `json:"authentication,omitempty"`
	Stages           []interface{}  `json:"stages,omitempty"`
}

// Trigger represents a pipeline trigger
type Trigger struct {
	User       string                 `json:"user,omitempty"`
	Type       string                 `json:"type,omitempty"`
	Parameters map[string]interface{} `json:"parameters,omitempty"`
}

// Authentication holds potential authentication information
type Authentication struct {
	User            string   `json:"user,omitempty"`
	AllowedAccounts []string `json:"allowedAccounts,omitempty"`
}
