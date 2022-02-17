/*
Qovery API

- Qovery is the fastest way to deploy your full-stack apps on any Cloud provider. - ℹ️ The API is stable and still in development.

API version: 1.0.0
Contact: support+api+documentation@qovery.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package qovery

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// Linger please
var (
	_ context.Context
)

// ProjectDeploymentRuleApiService ProjectDeploymentRuleApi service
type ProjectDeploymentRuleApiService service

type ApiCreateDeploymentRuleRequest struct {
	ctx                          context.Context
	ApiService                   *ProjectDeploymentRuleApiService
	projectId                    string
	projectDeploymentRuleRequest *ProjectDeploymentRuleRequest
}

func (r ApiCreateDeploymentRuleRequest) ProjectDeploymentRuleRequest(projectDeploymentRuleRequest ProjectDeploymentRuleRequest) ApiCreateDeploymentRuleRequest {
	r.projectDeploymentRuleRequest = &projectDeploymentRuleRequest
	return r
}

func (r ApiCreateDeploymentRuleRequest) Execute() (*ProjectDeploymentRuleResponse, *http.Response, error) {
	return r.ApiService.CreateDeploymentRuleExecute(r)
}

/*
CreateDeploymentRule Create a deployment rule

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param projectId Project ID
 @return ApiCreateDeploymentRuleRequest
*/
func (a *ProjectDeploymentRuleApiService) CreateDeploymentRule(ctx context.Context, projectId string) ApiCreateDeploymentRuleRequest {
	return ApiCreateDeploymentRuleRequest{
		ApiService: a,
		ctx:        ctx,
		projectId:  projectId,
	}
}

// Execute executes the request
//  @return ProjectDeploymentRuleResponse
func (a *ProjectDeploymentRuleApiService) CreateDeploymentRuleExecute(r ApiCreateDeploymentRuleRequest) (*ProjectDeploymentRuleResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ProjectDeploymentRuleResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProjectDeploymentRuleApiService.CreateDeploymentRule")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/project/{projectId}/deploymentRule"
	localVarPath = strings.Replace(localVarPath, "{"+"projectId"+"}", url.PathEscape(parameterToString(r.projectId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.projectDeploymentRuleRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiDeleteProjectDeploymentRuleRequest struct {
	ctx              context.Context
	ApiService       *ProjectDeploymentRuleApiService
	projectId        string
	deploymentRuleId string
}

func (r ApiDeleteProjectDeploymentRuleRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteProjectDeploymentRuleExecute(r)
}

/*
DeleteProjectDeploymentRule Delete a project deployment rule

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param projectId Project ID
 @param deploymentRuleId Deployment Rule ID
 @return ApiDeleteProjectDeploymentRuleRequest
*/
func (a *ProjectDeploymentRuleApiService) DeleteProjectDeploymentRule(ctx context.Context, projectId string, deploymentRuleId string) ApiDeleteProjectDeploymentRuleRequest {
	return ApiDeleteProjectDeploymentRuleRequest{
		ApiService:       a,
		ctx:              ctx,
		projectId:        projectId,
		deploymentRuleId: deploymentRuleId,
	}
}

// Execute executes the request
func (a *ProjectDeploymentRuleApiService) DeleteProjectDeploymentRuleExecute(r ApiDeleteProjectDeploymentRuleRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProjectDeploymentRuleApiService.DeleteProjectDeploymentRule")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/project/{projectId}/deploymentRule/{deploymentRuleId}"
	localVarPath = strings.Replace(localVarPath, "{"+"projectId"+"}", url.PathEscape(parameterToString(r.projectId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"deploymentRuleId"+"}", url.PathEscape(parameterToString(r.deploymentRuleId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiEditProjectDeployemtnRuleRequest struct {
	ctx                          context.Context
	ApiService                   *ProjectDeploymentRuleApiService
	projectId                    string
	deploymentRuleId             string
	projectDeploymentRuleRequest *ProjectDeploymentRuleRequest
}

func (r ApiEditProjectDeployemtnRuleRequest) ProjectDeploymentRuleRequest(projectDeploymentRuleRequest ProjectDeploymentRuleRequest) ApiEditProjectDeployemtnRuleRequest {
	r.projectDeploymentRuleRequest = &projectDeploymentRuleRequest
	return r
}

func (r ApiEditProjectDeployemtnRuleRequest) Execute() (*ProjectDeploymentRuleResponse, *http.Response, error) {
	return r.ApiService.EditProjectDeployemtnRuleExecute(r)
}

/*
EditProjectDeployemtnRule Edit a project deployment rule

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param projectId Project ID
 @param deploymentRuleId Deployment Rule ID
 @return ApiEditProjectDeployemtnRuleRequest
*/
func (a *ProjectDeploymentRuleApiService) EditProjectDeployemtnRule(ctx context.Context, projectId string, deploymentRuleId string) ApiEditProjectDeployemtnRuleRequest {
	return ApiEditProjectDeployemtnRuleRequest{
		ApiService:       a,
		ctx:              ctx,
		projectId:        projectId,
		deploymentRuleId: deploymentRuleId,
	}
}

// Execute executes the request
//  @return ProjectDeploymentRuleResponse
func (a *ProjectDeploymentRuleApiService) EditProjectDeployemtnRuleExecute(r ApiEditProjectDeployemtnRuleRequest) (*ProjectDeploymentRuleResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ProjectDeploymentRuleResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProjectDeploymentRuleApiService.EditProjectDeployemtnRule")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/project/{projectId}/deploymentRule/{deploymentRuleId}"
	localVarPath = strings.Replace(localVarPath, "{"+"projectId"+"}", url.PathEscape(parameterToString(r.projectId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"deploymentRuleId"+"}", url.PathEscape(parameterToString(r.deploymentRuleId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.projectDeploymentRuleRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetProjectDeploymentRuleRequest struct {
	ctx              context.Context
	ApiService       *ProjectDeploymentRuleApiService
	projectId        string
	deploymentRuleId string
}

func (r ApiGetProjectDeploymentRuleRequest) Execute() (*ProjectDeploymentRuleResponse, *http.Response, error) {
	return r.ApiService.GetProjectDeploymentRuleExecute(r)
}

/*
GetProjectDeploymentRule Get project deployment rule

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param projectId Project ID
 @param deploymentRuleId Deployment Rule ID
 @return ApiGetProjectDeploymentRuleRequest
*/
func (a *ProjectDeploymentRuleApiService) GetProjectDeploymentRule(ctx context.Context, projectId string, deploymentRuleId string) ApiGetProjectDeploymentRuleRequest {
	return ApiGetProjectDeploymentRuleRequest{
		ApiService:       a,
		ctx:              ctx,
		projectId:        projectId,
		deploymentRuleId: deploymentRuleId,
	}
}

// Execute executes the request
//  @return ProjectDeploymentRuleResponse
func (a *ProjectDeploymentRuleApiService) GetProjectDeploymentRuleExecute(r ApiGetProjectDeploymentRuleRequest) (*ProjectDeploymentRuleResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ProjectDeploymentRuleResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProjectDeploymentRuleApiService.GetProjectDeploymentRule")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/project/{projectId}/deploymentRule/{deploymentRuleId}"
	localVarPath = strings.Replace(localVarPath, "{"+"projectId"+"}", url.PathEscape(parameterToString(r.projectId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"deploymentRuleId"+"}", url.PathEscape(parameterToString(r.deploymentRuleId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiListProjectDeploymentRuleRequest struct {
	ctx        context.Context
	ApiService *ProjectDeploymentRuleApiService
	projectId  string
}

func (r ApiListProjectDeploymentRuleRequest) Execute() (*ProjectDeploymentRuleResponseList, *http.Response, error) {
	return r.ApiService.ListProjectDeploymentRuleExecute(r)
}

/*
ListProjectDeploymentRule List project deployment rules

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param projectId Project ID
 @return ApiListProjectDeploymentRuleRequest
*/
func (a *ProjectDeploymentRuleApiService) ListProjectDeploymentRule(ctx context.Context, projectId string) ApiListProjectDeploymentRuleRequest {
	return ApiListProjectDeploymentRuleRequest{
		ApiService: a,
		ctx:        ctx,
		projectId:  projectId,
	}
}

// Execute executes the request
//  @return ProjectDeploymentRuleResponseList
func (a *ProjectDeploymentRuleApiService) ListProjectDeploymentRuleExecute(r ApiListProjectDeploymentRuleRequest) (*ProjectDeploymentRuleResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ProjectDeploymentRuleResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProjectDeploymentRuleApiService.ListProjectDeploymentRule")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/project/{projectId}/deploymentRule"
	localVarPath = strings.Replace(localVarPath, "{"+"projectId"+"}", url.PathEscape(parameterToString(r.projectId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
