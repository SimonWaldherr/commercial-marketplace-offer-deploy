//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package sdk

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type marketplaceOfferDeploymentManagerClient struct {
	pl runtime.Pipeline
}

// newMarketplaceOfferDeploymentManagerClient creates a new instance of marketplaceOfferDeploymentManagerClient with the specified values.
//   - pl - the pipeline used for sending requests and handling responses.
func newMarketplaceOfferDeploymentManagerClient(pl runtime.Pipeline) *marketplaceOfferDeploymentManagerClient {
	client := &marketplaceOfferDeploymentManagerClient{
		pl: pl,
	}
	return client
}

// AddDeployment - Add a new deployment to the deployment
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
//   - body - Deployment object that needs to be added to the store
//   - options - MarketplaceOfferDeploymentManagerClientAddDeploymentOptions contains the optional parameters for the marketplaceOfferDeploymentManagerClient.AddDeployment
//     method.
func (client *marketplaceOfferDeploymentManagerClient) AddDeployment(ctx context.Context, body Deployment, options *MarketplaceOfferDeploymentManagerClientAddDeploymentOptions) (MarketplaceOfferDeploymentManagerClientAddDeploymentResponse, error) {
	req, err := client.addDeploymentCreateRequest(ctx, body, options)
	if err != nil {
		return MarketplaceOfferDeploymentManagerClientAddDeploymentResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return MarketplaceOfferDeploymentManagerClientAddDeploymentResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusMethodNotAllowed) {
		return MarketplaceOfferDeploymentManagerClientAddDeploymentResponse{}, runtime.NewResponseError(resp)
	}
	return MarketplaceOfferDeploymentManagerClientAddDeploymentResponse{}, nil
}

// addDeploymentCreateRequest creates the AddDeployment request.
func (client *marketplaceOfferDeploymentManagerClient) addDeploymentCreateRequest(ctx context.Context, body Deployment, options *MarketplaceOfferDeploymentManagerClientAddDeploymentOptions) (*policy.Request, error) {
	urlPath := "/deployments"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(	host, urlPath))
	if err != nil {
		return nil, err
	}
	return req, runtime.MarshalAsJSON(req, body)
}

// AddEventSubscription - Registers a subscription for a particular topic
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
//   - body - Deployment object that needs to be added to the store
//   - options - MarketplaceOfferDeploymentManagerClientAddEventSubscriptionOptions contains the optional parameters for the marketplaceOfferDeploymentManagerClient.AddEventSubscription
//     method.
func (client *marketplaceOfferDeploymentManagerClient) AddEventSubscription(ctx context.Context, body EventSubscription, options *MarketplaceOfferDeploymentManagerClientAddEventSubscriptionOptions) (MarketplaceOfferDeploymentManagerClientAddEventSubscriptionResponse, error) {
	req, err := client.addEventSubscriptionCreateRequest(ctx, body, options)
	if err != nil {
		return MarketplaceOfferDeploymentManagerClientAddEventSubscriptionResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return MarketplaceOfferDeploymentManagerClientAddEventSubscriptionResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusMethodNotAllowed) {
		return MarketplaceOfferDeploymentManagerClientAddEventSubscriptionResponse{}, runtime.NewResponseError(resp)
	}
	return MarketplaceOfferDeploymentManagerClientAddEventSubscriptionResponse{}, nil
}

// addEventSubscriptionCreateRequest creates the AddEventSubscription request.
func (client *marketplaceOfferDeploymentManagerClient) addEventSubscriptionCreateRequest(ctx context.Context, body EventSubscription, options *MarketplaceOfferDeploymentManagerClientAddEventSubscriptionOptions) (*policy.Request, error) {
	urlPath := "/events/subscriptions"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(	host, urlPath))
	if err != nil {
		return nil, err
	}
	return req, runtime.MarshalAsJSON(req, body)
}

// GetDeploymentByID - Returns a single deployment
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
//   - deploymentID - ID of deployment to return
//   - options - MarketplaceOfferDeploymentManagerClientGetDeploymentByIDOptions contains the optional parameters for the marketplaceOfferDeploymentManagerClient.GetDeploymentByID
//     method.
func (client *marketplaceOfferDeploymentManagerClient) GetDeploymentByID(ctx context.Context, deploymentID int64, options *MarketplaceOfferDeploymentManagerClientGetDeploymentByIDOptions) (MarketplaceOfferDeploymentManagerClientGetDeploymentByIDResponse, error) {
	req, err := client.getDeploymentByIDCreateRequest(ctx, deploymentID, options)
	if err != nil {
		return MarketplaceOfferDeploymentManagerClientGetDeploymentByIDResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return MarketplaceOfferDeploymentManagerClientGetDeploymentByIDResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusBadRequest, http.StatusNotFound) {
		return MarketplaceOfferDeploymentManagerClientGetDeploymentByIDResponse{}, runtime.NewResponseError(resp)
	}
	return client.getDeploymentByIDHandleResponse(resp)
}

// getDeploymentByIDCreateRequest creates the GetDeploymentByID request.
func (client *marketplaceOfferDeploymentManagerClient) getDeploymentByIDCreateRequest(ctx context.Context, deploymentID int64, options *MarketplaceOfferDeploymentManagerClientGetDeploymentByIDOptions) (*policy.Request, error) {
	urlPath := "/deployments/{deploymentId}"
	urlPath = strings.ReplaceAll(urlPath, "{deploymentId}", url.PathEscape(strconv.FormatInt(deploymentID, 10)))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(	host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getDeploymentByIDHandleResponse handles the GetDeploymentByID response.
func (client *marketplaceOfferDeploymentManagerClient) getDeploymentByIDHandleResponse(resp *http.Response) (MarketplaceOfferDeploymentManagerClientGetDeploymentByIDResponse, error) {
	result := MarketplaceOfferDeploymentManagerClientGetDeploymentByIDResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Deployment); err != nil {
		return MarketplaceOfferDeploymentManagerClientGetDeploymentByIDResponse{}, err
	}
	return result, nil
}

// GetDeployments - List all deployments
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
//   - options - MarketplaceOfferDeploymentManagerClientGetDeploymentsOptions contains the optional parameters for the marketplaceOfferDeploymentManagerClient.GetDeployments
//     method.
func (client *marketplaceOfferDeploymentManagerClient) GetDeployments(ctx context.Context, options *MarketplaceOfferDeploymentManagerClientGetDeploymentsOptions) (MarketplaceOfferDeploymentManagerClientGetDeploymentsResponse, error) {
	req, err := client.getDeploymentsCreateRequest(ctx, options)
	if err != nil {
		return MarketplaceOfferDeploymentManagerClientGetDeploymentsResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return MarketplaceOfferDeploymentManagerClientGetDeploymentsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return MarketplaceOfferDeploymentManagerClientGetDeploymentsResponse{}, runtime.NewResponseError(resp)
	}
	return client.getDeploymentsHandleResponse(resp)
}

// getDeploymentsCreateRequest creates the GetDeployments request.
func (client *marketplaceOfferDeploymentManagerClient) getDeploymentsCreateRequest(ctx context.Context, options *MarketplaceOfferDeploymentManagerClientGetDeploymentsOptions) (*policy.Request, error) {
	urlPath := "/deployments"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(	host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Status != nil {
			for _, qv := range options.Status {
		reqQP.Add("status", string(qv))
	}
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getDeploymentsHandleResponse handles the GetDeployments response.
func (client *marketplaceOfferDeploymentManagerClient) getDeploymentsHandleResponse(resp *http.Response) (MarketplaceOfferDeploymentManagerClientGetDeploymentsResponse, error) {
	result := MarketplaceOfferDeploymentManagerClientGetDeploymentsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DeploymentArray); err != nil {
		return MarketplaceOfferDeploymentManagerClientGetDeploymentsResponse{}, err
	}
	return result, nil
}

// GetEvents - Returns a list of available operations that can be performed on a deployment
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
//   - options - MarketplaceOfferDeploymentManagerClientGetEventsOptions contains the optional parameters for the marketplaceOfferDeploymentManagerClient.GetEvents
//     method.
func (client *marketplaceOfferDeploymentManagerClient) GetEvents(ctx context.Context, options *MarketplaceOfferDeploymentManagerClientGetEventsOptions) (MarketplaceOfferDeploymentManagerClientGetEventsResponse, error) {
	req, err := client.getEventsCreateRequest(ctx, options)
	if err != nil {
		return MarketplaceOfferDeploymentManagerClientGetEventsResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return MarketplaceOfferDeploymentManagerClientGetEventsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return MarketplaceOfferDeploymentManagerClientGetEventsResponse{}, runtime.NewResponseError(resp)
	}
	return client.getEventsHandleResponse(resp)
}

// getEventsCreateRequest creates the GetEvents request.
func (client *marketplaceOfferDeploymentManagerClient) getEventsCreateRequest(ctx context.Context, options *MarketplaceOfferDeploymentManagerClientGetEventsOptions) (*policy.Request, error) {
	urlPath := "/events"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(	host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getEventsHandleResponse handles the GetEvents response.
func (client *marketplaceOfferDeploymentManagerClient) getEventsHandleResponse(resp *http.Response) (MarketplaceOfferDeploymentManagerClientGetEventsResponse, error) {
	result := MarketplaceOfferDeploymentManagerClientGetEventsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AvailableDeploymentOperationArray); err != nil {
		return MarketplaceOfferDeploymentManagerClientGetEventsResponse{}, err
	}
	return result, nil
}

// GetOperations - Returns a list of available operations that can be performed on a deployment
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
//   - options - MarketplaceOfferDeploymentManagerClientGetOperationsOptions contains the optional parameters for the marketplaceOfferDeploymentManagerClient.GetOperations
//     method.
func (client *marketplaceOfferDeploymentManagerClient) GetOperations(ctx context.Context, options *MarketplaceOfferDeploymentManagerClientGetOperationsOptions) (MarketplaceOfferDeploymentManagerClientGetOperationsResponse, error) {
	req, err := client.getOperationsCreateRequest(ctx, options)
	if err != nil {
		return MarketplaceOfferDeploymentManagerClientGetOperationsResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return MarketplaceOfferDeploymentManagerClientGetOperationsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return MarketplaceOfferDeploymentManagerClientGetOperationsResponse{}, runtime.NewResponseError(resp)
	}
	return client.getOperationsHandleResponse(resp)
}

// getOperationsCreateRequest creates the GetOperations request.
func (client *marketplaceOfferDeploymentManagerClient) getOperationsCreateRequest(ctx context.Context, options *MarketplaceOfferDeploymentManagerClientGetOperationsOptions) (*policy.Request, error) {
	urlPath := "/deployments/operations"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(	host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getOperationsHandleResponse handles the GetOperations response.
func (client *marketplaceOfferDeploymentManagerClient) getOperationsHandleResponse(resp *http.Response) (MarketplaceOfferDeploymentManagerClientGetOperationsResponse, error) {
	result := MarketplaceOfferDeploymentManagerClientGetOperationsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AvailableDeploymentOperationArray); err != nil {
		return MarketplaceOfferDeploymentManagerClientGetOperationsResponse{}, err
	}
	return result, nil
}

// TriggerDeploymentOperation - Triggers a command operation with parameters for a deployment
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
//   - deploymentID - ID of deployment to return
//   - body - Deployment operation trigger
//   - options - MarketplaceOfferDeploymentManagerClientTriggerDeploymentOperationOptions contains the optional parameters for
//     the marketplaceOfferDeploymentManagerClient.TriggerDeploymentOperation method.
func (client *marketplaceOfferDeploymentManagerClient) TriggerDeploymentOperation(ctx context.Context, deploymentID int64, body DeploymentOperation, options *MarketplaceOfferDeploymentManagerClientTriggerDeploymentOperationOptions) (MarketplaceOfferDeploymentManagerClientTriggerDeploymentOperationResponse, error) {
	req, err := client.triggerDeploymentOperationCreateRequest(ctx, deploymentID, body, options)
	if err != nil {
		return MarketplaceOfferDeploymentManagerClientTriggerDeploymentOperationResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return MarketplaceOfferDeploymentManagerClientTriggerDeploymentOperationResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return MarketplaceOfferDeploymentManagerClientTriggerDeploymentOperationResponse{}, runtime.NewResponseError(resp)
	}
	return MarketplaceOfferDeploymentManagerClientTriggerDeploymentOperationResponse{}, nil
}

// triggerDeploymentOperationCreateRequest creates the TriggerDeploymentOperation request.
func (client *marketplaceOfferDeploymentManagerClient) triggerDeploymentOperationCreateRequest(ctx context.Context, deploymentID int64, body DeploymentOperation, options *MarketplaceOfferDeploymentManagerClientTriggerDeploymentOperationOptions) (*policy.Request, error) {
	urlPath := "/deployments/{deploymentId}/operations"
	urlPath = strings.ReplaceAll(urlPath, "{deploymentId}", url.PathEscape(strconv.FormatInt(deploymentID, 10)))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(	host, urlPath))
	if err != nil {
		return nil, err
	}
	return req, runtime.MarshalAsJSON(req, body)
}

// UpdateDeployment - Update an existing deployment
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
//   - body - Deployment object that needs to be added to the store
//   - options - MarketplaceOfferDeploymentManagerClientUpdateDeploymentOptions contains the optional parameters for the marketplaceOfferDeploymentManagerClient.UpdateDeployment
//     method.
func (client *marketplaceOfferDeploymentManagerClient) UpdateDeployment(ctx context.Context, body Deployment, options *MarketplaceOfferDeploymentManagerClientUpdateDeploymentOptions) (MarketplaceOfferDeploymentManagerClientUpdateDeploymentResponse, error) {
	req, err := client.updateDeploymentCreateRequest(ctx, body, options)
	if err != nil {
		return MarketplaceOfferDeploymentManagerClientUpdateDeploymentResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return MarketplaceOfferDeploymentManagerClientUpdateDeploymentResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusBadRequest, http.StatusNotFound, http.StatusMethodNotAllowed) {
		return MarketplaceOfferDeploymentManagerClientUpdateDeploymentResponse{}, runtime.NewResponseError(resp)
	}
	return MarketplaceOfferDeploymentManagerClientUpdateDeploymentResponse{}, nil
}

// updateDeploymentCreateRequest creates the UpdateDeployment request.
func (client *marketplaceOfferDeploymentManagerClient) updateDeploymentCreateRequest(ctx context.Context, body Deployment, options *MarketplaceOfferDeploymentManagerClientUpdateDeploymentOptions) (*policy.Request, error) {
	urlPath := "/deployments"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(	host, urlPath))
	if err != nil {
		return nil, err
	}
	return req, runtime.MarshalAsJSON(req, body)
}

