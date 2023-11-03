package messages

import "reflect"

// helper
func MessageName(any interface{}) string {
	t := reflect.TypeOf(any)
	if t.Kind() == reflect.Ptr {
		return t.Elem().Name()
	}
	return t.Name()
}

// application service
type ProvisionApplicationRequest struct {
	RequesterID                   string `json:"requesterId"`
	ApplicationID                 string `json:"applicationId"`
	TeamID                        string `json:"teamId"`
	ApplicationName               string `json:"applicationName"` // eg, medicationID - TODO support `line items`
	AuthorizerID                  string `json:"authorizerId"`
	AuthorizationTimeoutSeconds   int64  `json:"authorizationTimeoutSeconds"`
	DemoAuthorizationDelaySeconds int64  `json:"demoAuthorizationDelaySeconds"`
}
type PlaceOrderResponse struct {
}

// query
type GetTeamInformationRequest struct {
	RequesterID string
	TeamID      string
}
type GetTeamInformationResponse struct {
	TeamName  string
	TeamEmail string
	Subdomain string
	TeamID    string
}

// commands
type RequestApplicationAuthorizationRequest struct {
	AuthorizerID  string // eg doctor
	TeamID        string
	CustomerID    string
	DelaySeconds  int64
	ApplicationID string
}
type DestroyResources struct {
	OrderID string
}
type ChangeOrderRequest struct {
	OrderID        string
	Reason         string
	PrescriptionID string
}
type ProvisionFoundationResourcesRequest struct {
	ApplicationID   string `json:"applicationId"`
	Region          string `json:"region"`
	Profile         string `json:"profile"`
	TeamID          string `json:"teamId"`
	ApplicationName string `json:"applicationName"`
	BucketName      string `json:"bucketName"`
}
type PauseOrderRequest struct {
	OrderID         string
	MaxPauseSeconds int64
}
type ResumeOrderRequest struct {
	OrderID string
}
type ValidateInsuranceRequest struct {
	OrderID        string
	CustomerID     string
	PrescriptionID string
	DelaySeconds   int64
}

/*events*/

type AuthorizationReceivedResponse struct {
	ApplicationID string `json:"applicationId"`
	Region        string `json:"region"`
	Profile       string `json:"profile"`
	TeamID        string `json:"teamId"`
	IsApproved    bool   `json:"isApproved"`
}

type OrderChangedResponse struct {
	OrderID        string
	Reason         string
	PrescriptionID string
}
type ProvisionFoundationResourcesResponse struct {
	ApplicationID      string   `json:"applicationId"`
	Summary            []string `json:"summary"`
	CompletionDateTime string   `json:"completionDateTime"`
	Region             string   `json:"region""`
}
type ValidateInsuranceResponse struct {
	OrderID    string
	CustomerID string
}
