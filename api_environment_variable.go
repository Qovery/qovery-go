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

// Linger please
var (
	_ context.Context
)

// EnvironmentVariableApiService EnvironmentVariableApi service
type EnvironmentVariableApiService service

type ApiCreateEnvironmentEnvironmentVariableRequest struct {
	ctx                        context.Context
	ApiService                 *EnvironmentVariableApiService
	environmentId              string
	environmentVariableRequest *EnvironmentVariableRequest
}

func (r ApiCreateEnvironmentEnvironmentVariableRequest) EnvironmentVariableRequest(environmentVariableRequest EnvironmentVariableRequest) ApiCreateEnvironmentEnvironmentVariableRequest {
	r.environmentVariableRequest = &environmentVariableRequest
	return r
}

func (r ApiCreateEnvironmentEnvironmentVariableRequest) Execute() (*EnvironmentVariableResponse, *http.Response, error) {
	return r.ApiService.CreateEnvironmentEnvironmentVariableExecute(r)
}

/*
CreateEnvironmentEnvironmentVariable Add an environment variable to the environment

- Add an environment variable to the environment.
  - If the environment variable key already exists, then it will be replaced by the new one.
  - If the environment variable value points toward an existing environment variable key, it will be considered as an alias.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param environmentId Environment ID
 @return ApiCreateEnvironmentEnvironmentVariableRequest
*/
func (a *EnvironmentVariableApiService) CreateEnvironmentEnvironmentVariable(ctx context.Context, environmentId string) ApiCreateEnvironmentEnvironmentVariableRequest {
	return ApiCreateEnvironmentEnvironmentVariableRequest{
		ApiService:    a,
		ctx:           ctx,
		environmentId: environmentId,
	}
}

// Execute executes the request
//  @return EnvironmentVariableResponse
func (a *EnvironmentVariableApiService) CreateEnvironmentEnvironmentVariableExecute(r ApiCreateEnvironmentEnvironmentVariableRequest) (*EnvironmentVariableResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *EnvironmentVariableResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EnvironmentVariableApiService.CreateEnvironmentEnvironmentVariable")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/environment/{environmentId}/environmentVariable"
	localVarPath = strings.Replace(localVarPath, "{"+"environmentId"+"}", url.PathEscape(parameterToString(r.environmentId, "")), -1)

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
	localVarPostBody = r.environmentVariableRequest
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

type ApiCreateEnvironmentEnvironmentVariableAliasRequest struct {
	ctx                   context.Context
	ApiService            *EnvironmentVariableApiService
	environmentId         string
	environmentVariableId string
	key                   *Key
}

func (r ApiCreateEnvironmentEnvironmentVariableAliasRequest) Key(key Key) ApiCreateEnvironmentEnvironmentVariableAliasRequest {
	r.key = &key
	return r
}

func (r ApiCreateEnvironmentEnvironmentVariableAliasRequest) Execute() (*EnvironmentVariableResponse, *http.Response, error) {
	return r.ApiService.CreateEnvironmentEnvironmentVariableAliasExecute(r)
}

/*
CreateEnvironmentEnvironmentVariableAlias Create an environment variable alias at the environment level

- Allows you to add an alias at environment level on an existing environment variable having higher scope, in order to customize its key.
- You only have to specify a key in the request body
- The system will create a new environment variable at environment level with the same value as the one corresponding to the variable id in the path
- The response body will contain the newly created variable
- Information regarding the aliased_variable will be exposed in the "aliased_variable" field of the newly created variable
- Only 1 alias level is allowed. You can't create an alias on an alias


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param environmentId Environment ID
 @param environmentVariableId Environment Variable ID
 @return ApiCreateEnvironmentEnvironmentVariableAliasRequest
*/
func (a *EnvironmentVariableApiService) CreateEnvironmentEnvironmentVariableAlias(ctx context.Context, environmentId string, environmentVariableId string) ApiCreateEnvironmentEnvironmentVariableAliasRequest {
	return ApiCreateEnvironmentEnvironmentVariableAliasRequest{
		ApiService:            a,
		ctx:                   ctx,
		environmentId:         environmentId,
		environmentVariableId: environmentVariableId,
	}
}

// Execute executes the request
//  @return EnvironmentVariableResponse
func (a *EnvironmentVariableApiService) CreateEnvironmentEnvironmentVariableAliasExecute(r ApiCreateEnvironmentEnvironmentVariableAliasRequest) (*EnvironmentVariableResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *EnvironmentVariableResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EnvironmentVariableApiService.CreateEnvironmentEnvironmentVariableAlias")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/environment/{environmentId}/environmentVariable/{environmentVariableId}/alias"
	localVarPath = strings.Replace(localVarPath, "{"+"environmentId"+"}", url.PathEscape(parameterToString(r.environmentId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"environmentVariableId"+"}", url.PathEscape(parameterToString(r.environmentVariableId, "")), -1)

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
	localVarPostBody = r.key
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

type ApiCreateEnvironmentEnvironmentVariableOverrideRequest struct {
	ctx                   context.Context
	ApiService            *EnvironmentVariableApiService
	environmentId         string
	environmentVariableId string
	value                 *Value
}

func (r ApiCreateEnvironmentEnvironmentVariableOverrideRequest) Value(value Value) ApiCreateEnvironmentEnvironmentVariableOverrideRequest {
	r.value = &value
	return r
}

func (r ApiCreateEnvironmentEnvironmentVariableOverrideRequest) Execute() (*EnvironmentVariableResponse, *http.Response, error) {
	return r.ApiService.CreateEnvironmentEnvironmentVariableOverrideExecute(r)
}

/*
CreateEnvironmentEnvironmentVariableOverride Create an environment variable override at the environment level

- Allows you to override at environment level an environment variable that has a higher scope.
- You only have to specify a value in the request body
- The system will create a new environment variable at environment level with the same key as the one corresponding to the variable id in the path
- The response body will contain the newly created variable
- Information regarding the overridden_variable will be exposed in the "overridden_variable" field of the newly created variable


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param environmentId Environment ID
 @param environmentVariableId Environment Variable ID
 @return ApiCreateEnvironmentEnvironmentVariableOverrideRequest
*/
func (a *EnvironmentVariableApiService) CreateEnvironmentEnvironmentVariableOverride(ctx context.Context, environmentId string, environmentVariableId string) ApiCreateEnvironmentEnvironmentVariableOverrideRequest {
	return ApiCreateEnvironmentEnvironmentVariableOverrideRequest{
		ApiService:            a,
		ctx:                   ctx,
		environmentId:         environmentId,
		environmentVariableId: environmentVariableId,
	}
}

// Execute executes the request
//  @return EnvironmentVariableResponse
func (a *EnvironmentVariableApiService) CreateEnvironmentEnvironmentVariableOverrideExecute(r ApiCreateEnvironmentEnvironmentVariableOverrideRequest) (*EnvironmentVariableResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *EnvironmentVariableResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EnvironmentVariableApiService.CreateEnvironmentEnvironmentVariableOverride")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/environment/{environmentId}/environmentVariable/{environmentVariableId}/override"
	localVarPath = strings.Replace(localVarPath, "{"+"environmentId"+"}", url.PathEscape(parameterToString(r.environmentId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"environmentVariableId"+"}", url.PathEscape(parameterToString(r.environmentVariableId, "")), -1)

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
	localVarPostBody = r.value
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

type ApiDeleteEnvironmentEnvironmentVariableRequest struct {
	ctx                   context.Context
	ApiService            *EnvironmentVariableApiService
	environmentId         string
	environmentVariableId string
}

func (r ApiDeleteEnvironmentEnvironmentVariableRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteEnvironmentEnvironmentVariableExecute(r)
}

/*
DeleteEnvironmentEnvironmentVariable Delete an environment variable from an environment

- To delete an environment variable you must have the project user permission
- You can't delete a BUILT_IN variable
- If you delete a variable having override or alias, the associated override/alias will be deleted as well


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param environmentId Environment ID
 @param environmentVariableId Environment Variable ID
 @return ApiDeleteEnvironmentEnvironmentVariableRequest
*/
func (a *EnvironmentVariableApiService) DeleteEnvironmentEnvironmentVariable(ctx context.Context, environmentId string, environmentVariableId string) ApiDeleteEnvironmentEnvironmentVariableRequest {
	return ApiDeleteEnvironmentEnvironmentVariableRequest{
		ApiService:            a,
		ctx:                   ctx,
		environmentId:         environmentId,
		environmentVariableId: environmentVariableId,
	}
}

// Execute executes the request
func (a *EnvironmentVariableApiService) DeleteEnvironmentEnvironmentVariableExecute(r ApiDeleteEnvironmentEnvironmentVariableRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EnvironmentVariableApiService.DeleteEnvironmentEnvironmentVariable")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/environment/{environmentId}/environmentVariable/{environmentVariableId}"
	localVarPath = strings.Replace(localVarPath, "{"+"environmentId"+"}", url.PathEscape(parameterToString(r.environmentId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"environmentVariableId"+"}", url.PathEscape(parameterToString(r.environmentVariableId, "")), -1)

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

type ApiEditEnvironmentEnvironmentVariableRequest struct {
	ctx                            context.Context
	ApiService                     *EnvironmentVariableApiService
	environmentId                  string
	environmentVariableId          string
	environmentVariableEditRequest *EnvironmentVariableEditRequest
}

func (r ApiEditEnvironmentEnvironmentVariableRequest) EnvironmentVariableEditRequest(environmentVariableEditRequest EnvironmentVariableEditRequest) ApiEditEnvironmentEnvironmentVariableRequest {
	r.environmentVariableEditRequest = &environmentVariableEditRequest
	return r
}

func (r ApiEditEnvironmentEnvironmentVariableRequest) Execute() (*EnvironmentVariableResponse, *http.Response, error) {
	return r.ApiService.EditEnvironmentEnvironmentVariableExecute(r)
}

/*
EditEnvironmentEnvironmentVariable Edit an environment variable belonging to the environment

- You can't edit a BUILT_IN variable
- For an override, you can't edit the key
- For an alias, you can't edit the value
- An override can only have a scope lower to the variable it is overriding (hierarchy is BUILT_IN > PROJECT > ENVIRONMENT > APPLICATION)


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param environmentId Environment ID
 @param environmentVariableId Environment Variable ID
 @return ApiEditEnvironmentEnvironmentVariableRequest
*/
func (a *EnvironmentVariableApiService) EditEnvironmentEnvironmentVariable(ctx context.Context, environmentId string, environmentVariableId string) ApiEditEnvironmentEnvironmentVariableRequest {
	return ApiEditEnvironmentEnvironmentVariableRequest{
		ApiService:            a,
		ctx:                   ctx,
		environmentId:         environmentId,
		environmentVariableId: environmentVariableId,
	}
}

// Execute executes the request
//  @return EnvironmentVariableResponse
func (a *EnvironmentVariableApiService) EditEnvironmentEnvironmentVariableExecute(r ApiEditEnvironmentEnvironmentVariableRequest) (*EnvironmentVariableResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *EnvironmentVariableResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EnvironmentVariableApiService.EditEnvironmentEnvironmentVariable")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/environment/{environmentId}/environmentVariable/{environmentVariableId}"
	localVarPath = strings.Replace(localVarPath, "{"+"environmentId"+"}", url.PathEscape(parameterToString(r.environmentId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"environmentVariableId"+"}", url.PathEscape(parameterToString(r.environmentVariableId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.environmentVariableEditRequest == nil {
		return localVarReturnValue, nil, reportError("environmentVariableEditRequest is required and must be specified")
	}

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
	localVarPostBody = r.environmentVariableEditRequest
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

type ApiListEnvironmentEnvironmentVariableRequest struct {
	ctx           context.Context
	ApiService    *EnvironmentVariableApiService
	environmentId string
}

func (r ApiListEnvironmentEnvironmentVariableRequest) Execute() (*EnvironmentVariableResponseList, *http.Response, error) {
	return r.ApiService.ListEnvironmentEnvironmentVariableExecute(r)
}

/*
ListEnvironmentEnvironmentVariable List environment variables

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param environmentId Environment ID
 @return ApiListEnvironmentEnvironmentVariableRequest
*/
func (a *EnvironmentVariableApiService) ListEnvironmentEnvironmentVariable(ctx context.Context, environmentId string) ApiListEnvironmentEnvironmentVariableRequest {
	return ApiListEnvironmentEnvironmentVariableRequest{
		ApiService:    a,
		ctx:           ctx,
		environmentId: environmentId,
	}
}

// Execute executes the request
//  @return EnvironmentVariableResponseList
func (a *EnvironmentVariableApiService) ListEnvironmentEnvironmentVariableExecute(r ApiListEnvironmentEnvironmentVariableRequest) (*EnvironmentVariableResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *EnvironmentVariableResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EnvironmentVariableApiService.ListEnvironmentEnvironmentVariable")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/environment/{environmentId}/environmentVariable"
	localVarPath = strings.Replace(localVarPath, "{"+"environmentId"+"}", url.PathEscape(parameterToString(r.environmentId, "")), -1)

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
