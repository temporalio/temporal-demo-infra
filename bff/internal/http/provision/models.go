package provision

type ProvisionPOST struct {
	ApplicationName string `json:"applicationName"`
	TeamID          string `json:"teamId"`
	AuthorizerID    string `json:"authorizerId""`
}
type ProvisionPATCH struct {
	Region        string `json:"region"`
	WorkflowID    string `json:"workflowId"`
	AuthorizerID  string `json:"authorizerId"`
	Profile       string `json:"profile"`
	ApplicationID string `json:"applicationId"`
	TeamID        string `json:"teamId"`
}
type ProvisionDELETE struct {
	Region        string `json:"region"`
	WorkflowID    string `json:"workflowId"`
	AuthorizerID  string `json:"authorizerId"`
	Profile       string `json:"profile"`
	ApplicationID string `json:"applicationId"`
	TeamID        string `json:"teamId"`
}
type ProvisionPOSTResponse struct {
	ApplicationName string `json:"applicationName"`
	TeamID          string `json:"teamId"`
	AuthorizerID    string `json:"authorizerId"`
	WorkflowID      string `json:"workflowId"`
	ApplicationID   string `json:"applicationId"`
}
