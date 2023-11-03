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
	RequesterID                       string
	ApplicationID                     string
	TeamID                            string
	ApplicationName                   string // eg, medicationID - TODO support `line items`
	AuthorizerID                      string
	AuthorizationTimeoutSeconds       int64
	DemoAuthorizationDelaySeconds     int64
	DemoFulfillmentDelaySeconds       int64
	DemoValidateInsuranceDelaySeconds int64
	ChangeOrder                       *ChangeOrderRequest
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
	ApplicationID   string
	Region          string
	RoleAdminArn    string
	TeamID          string
	ApplicationName string
	DelaySeconds    int64
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
	ApplicationID string
	Region        string
	RoleAdminArn  string
	TeamID        string
	IsApproved    bool
}

type OrderChangedResponse struct {
	OrderID        string
	Reason         string
	PrescriptionID string
}
type ProvisionFoundationResourcesResponse struct {
	ApplicationID      string
	CompletionDateTime string
}
type ValidateInsuranceResponse struct {
	OrderID    string
	CustomerID string
}
