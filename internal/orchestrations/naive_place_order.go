package orchestrations

//func (o *Orchestrations) NaivePlaceOrder(ctx workflow.Context, params *messages.ProvisionApplicationRequest) error {
//	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
//		StartToCloseTimeout: time.Second * 3,
//	})
//	var orderInfo *messages.GetTeamInformationResponse
//	var fulfillment *messages.ProvisionFoundationResourcesResponse
//	if err := workflow.ExecuteActivity(ctx, teams.TypeHandlers.GetTeamInformation, &messages.GetTeamInformationRequest{
//		RequesterID:  params.RequesterID,
//		TeamID:       params.TeamID,
//		AuthorizerID: params.AuthorizerID,
//		ProductID:    params.AuthorizerID,
//	}).Get(ctx, &orderInfo); err != nil {
//		return fmt.Errorf("failed to get order details %w", err)
//	}
//
//	if err := workflow.ExecuteActivity(ctx, provider_aws.TypeHandlers.ProvisionFoundationResources, &messages.ProvisionFoundationResourcesRequest{
//		ApplicationID: orderInfo.TeamName,
//		Region:        params.ApplicationID,
//		Profile:       params.TeamID,
//		TeamID:        orderInfo.TeamEmail,
//		DelaySeconds:  0, // synchronous
//	}).Get(ctx, &fulfillment); err != nil {
//		return fmt.Errorf("failed to fulfill prescription %w", err)
//	}
//	workflow.GetLogger(ctx).Info("order completed!")
//	return nil
//}
