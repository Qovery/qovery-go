/*
Qovery API

- Qovery is the fastest way to deploy your full-stack apps on any Cloud provider. - ℹ️ The API is stable and still in development.

API version: 1.0.3
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

// OrganizationWebhookApiService OrganizationWebhookApi service
type OrganizationWebhookApiService service

type ApiCreateOrganizationWebhookRequest struct {
	ctx                              context.Context
	ApiService                       *OrganizationWebhookApiService
	organizationId                   string
	organizationWebhookCreateRequest *OrganizationWebhookCreateRequest
}

func (r ApiCreateOrganizationWebhookRequest) OrganizationWebhookCreateRequest(organizationWebhookCreateRequest OrganizationWebhookCreateRequest) ApiCreateOrganizationWebhookRequest {
	r.organizationWebhookCreateRequest = &organizationWebhookCreateRequest
	return r
}

func (r ApiCreateOrganizationWebhookRequest) Execute() (*OrganizationWebhookCreateResponse, *http.Response, error) {
	return r.ApiService.CreateOrganizationWebhookExecute(r)
}

/*
CreateOrganizationWebhook Create an organization webhook

Create an organization webhook.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId Organization ID
 @return ApiCreateOrganizationWebhookRequest
*/
func (a *OrganizationWebhookApiService) CreateOrganizationWebhook(ctx context.Context, organizationId string) ApiCreateOrganizationWebhookRequest {
	return ApiCreateOrganizationWebhookRequest{
		ApiService:     a,
		ctx:            ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return OrganizationWebhookCreateResponse
func (a *OrganizationWebhookApiService) CreateOrganizationWebhookExecute(r ApiCreateOrganizationWebhookRequest) (*OrganizationWebhookCreateResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *OrganizationWebhookCreateResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OrganizationWebhookApiService.CreateOrganizationWebhook")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organization/{organizationId}/webhook"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", url.PathEscape(parameterToString(r.organizationId, "")), -1)

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
	localVarPostBody = r.organizationWebhookCreateRequest
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Token"] = key
			}
		}
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

type ApiDeleteOrganizationWebhookRequest struct {
	ctx            context.Context
	ApiService     *OrganizationWebhookApiService
	organizationId string
	webhookId      string
}

func (r ApiDeleteOrganizationWebhookRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteOrganizationWebhookExecute(r)
}

/*
DeleteOrganizationWebhook Delete organization webhook

Delete organization webhook

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId Organization ID
 @param webhookId Webhook ID
 @return ApiDeleteOrganizationWebhookRequest
*/
func (a *OrganizationWebhookApiService) DeleteOrganizationWebhook(ctx context.Context, organizationId string, webhookId string) ApiDeleteOrganizationWebhookRequest {
	return ApiDeleteOrganizationWebhookRequest{
		ApiService:     a,
		ctx:            ctx,
		organizationId: organizationId,
		webhookId:      webhookId,
	}
}

// Execute executes the request
func (a *OrganizationWebhookApiService) DeleteOrganizationWebhookExecute(r ApiDeleteOrganizationWebhookRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OrganizationWebhookApiService.DeleteOrganizationWebhook")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organization/{organizationId}/webhook/{webhookId}"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", url.PathEscape(parameterToString(r.organizationId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"webhookId"+"}", url.PathEscape(parameterToString(r.webhookId, "")), -1)

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
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Token"] = key
			}
		}
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

type ApiEditOrganizationWebhookRequest struct {
	ctx                              context.Context
	ApiService                       *OrganizationWebhookApiService
	organizationId                   string
	webhookId                        string
	organizationWebhookCreateRequest *OrganizationWebhookCreateRequest
}

func (r ApiEditOrganizationWebhookRequest) OrganizationWebhookCreateRequest(organizationWebhookCreateRequest OrganizationWebhookCreateRequest) ApiEditOrganizationWebhookRequest {
	r.organizationWebhookCreateRequest = &organizationWebhookCreateRequest
	return r
}

func (r ApiEditOrganizationWebhookRequest) Execute() (*OrganizationWebhookCreateResponse, *http.Response, error) {
	return r.ApiService.EditOrganizationWebhookExecute(r)
}

/*
EditOrganizationWebhook Edit an organization webhook

Edit an organization webhook

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId Organization ID
 @param webhookId Webhook ID
 @return ApiEditOrganizationWebhookRequest
*/
func (a *OrganizationWebhookApiService) EditOrganizationWebhook(ctx context.Context, organizationId string, webhookId string) ApiEditOrganizationWebhookRequest {
	return ApiEditOrganizationWebhookRequest{
		ApiService:     a,
		ctx:            ctx,
		organizationId: organizationId,
		webhookId:      webhookId,
	}
}

// Execute executes the request
//  @return OrganizationWebhookCreateResponse
func (a *OrganizationWebhookApiService) EditOrganizationWebhookExecute(r ApiEditOrganizationWebhookRequest) (*OrganizationWebhookCreateResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *OrganizationWebhookCreateResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OrganizationWebhookApiService.EditOrganizationWebhook")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organization/{organizationId}/webhook/{webhookId}"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", url.PathEscape(parameterToString(r.organizationId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"webhookId"+"}", url.PathEscape(parameterToString(r.webhookId, "")), -1)

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
	localVarPostBody = r.organizationWebhookCreateRequest
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Token"] = key
			}
		}
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

type ApiGetOrganizationWebhookRequest struct {
	ctx            context.Context
	ApiService     *OrganizationWebhookApiService
	organizationId string
	webhookId      string
}

func (r ApiGetOrganizationWebhookRequest) Execute() (*OrganizationWebhookResponse, *http.Response, error) {
	return r.ApiService.GetOrganizationWebhookExecute(r)
}

/*
GetOrganizationWebhook Get an Organization webhook

Get an Organization webhook

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId Organization ID
 @param webhookId Webhook ID
 @return ApiGetOrganizationWebhookRequest
*/
func (a *OrganizationWebhookApiService) GetOrganizationWebhook(ctx context.Context, organizationId string, webhookId string) ApiGetOrganizationWebhookRequest {
	return ApiGetOrganizationWebhookRequest{
		ApiService:     a,
		ctx:            ctx,
		organizationId: organizationId,
		webhookId:      webhookId,
	}
}

// Execute executes the request
//  @return OrganizationWebhookResponse
func (a *OrganizationWebhookApiService) GetOrganizationWebhookExecute(r ApiGetOrganizationWebhookRequest) (*OrganizationWebhookResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *OrganizationWebhookResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OrganizationWebhookApiService.GetOrganizationWebhook")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organization/{organizationId}/webhook/{webhookId}"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", url.PathEscape(parameterToString(r.organizationId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"webhookId"+"}", url.PathEscape(parameterToString(r.webhookId, "")), -1)

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
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Token"] = key
			}
		}
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

type ApiListOrganizationWebHooksRequest struct {
	ctx            context.Context
	ApiService     *OrganizationWebhookApiService
	organizationId string
}

func (r ApiListOrganizationWebHooksRequest) Execute() (*OrganizationWebhookResponseList, *http.Response, error) {
	return r.ApiService.ListOrganizationWebHooksExecute(r)
}

/*
ListOrganizationWebHooks List organization webhooks

List organization webhooks

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId Organization ID
 @return ApiListOrganizationWebHooksRequest
*/
func (a *OrganizationWebhookApiService) ListOrganizationWebHooks(ctx context.Context, organizationId string) ApiListOrganizationWebHooksRequest {
	return ApiListOrganizationWebHooksRequest{
		ApiService:     a,
		ctx:            ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return OrganizationWebhookResponseList
func (a *OrganizationWebhookApiService) ListOrganizationWebHooksExecute(r ApiListOrganizationWebHooksRequest) (*OrganizationWebhookResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *OrganizationWebhookResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OrganizationWebhookApiService.ListOrganizationWebHooks")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organization/{organizationId}/webhook"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", url.PathEscape(parameterToString(r.organizationId, "")), -1)

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
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Token"] = key
			}
		}
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
