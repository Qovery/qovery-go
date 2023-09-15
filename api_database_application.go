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

// DatabaseApplicationApiService DatabaseApplicationApi service
type DatabaseApplicationApiService service

type ApiListDatabaseApplicationRequest struct {
	ctx        context.Context
	ApiService *DatabaseApplicationApiService
	databaseId string
}

func (r ApiListDatabaseApplicationRequest) Execute() (*ApplicationResponseList, *http.Response, error) {
	return r.ApiService.ListDatabaseApplicationExecute(r)
}

/*
ListDatabaseApplication List applications using the database

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param databaseId Database ID
 @return ApiListDatabaseApplicationRequest
*/
func (a *DatabaseApplicationApiService) ListDatabaseApplication(ctx context.Context, databaseId string) ApiListDatabaseApplicationRequest {
	return ApiListDatabaseApplicationRequest{
		ApiService: a,
		ctx:        ctx,
		databaseId: databaseId,
	}
}

// Execute executes the request
//  @return ApplicationResponseList
func (a *DatabaseApplicationApiService) ListDatabaseApplicationExecute(r ApiListDatabaseApplicationRequest) (*ApplicationResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ApplicationResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DatabaseApplicationApiService.ListDatabaseApplication")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/database/{databaseId}/application"
	localVarPath = strings.Replace(localVarPath, "{"+"databaseId"+"}", url.PathEscape(parameterToString(r.databaseId, "")), -1)

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
				localVarHeaderParams["Authorization"] = key
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

type ApiRemoveApplicationFromDatabaseRequest struct {
	ctx                 context.Context
	ApiService          *DatabaseApplicationApiService
	databaseId          string
	targetApplicationId string
}

func (r ApiRemoveApplicationFromDatabaseRequest) Execute() (*http.Response, error) {
	return r.ApiService.RemoveApplicationFromDatabaseExecute(r)
}

/*
RemoveApplicationFromDatabase Remove an application from this database

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param databaseId Database ID
 @param targetApplicationId Target application ID
 @return ApiRemoveApplicationFromDatabaseRequest
*/
func (a *DatabaseApplicationApiService) RemoveApplicationFromDatabase(ctx context.Context, databaseId string, targetApplicationId string) ApiRemoveApplicationFromDatabaseRequest {
	return ApiRemoveApplicationFromDatabaseRequest{
		ApiService:          a,
		ctx:                 ctx,
		databaseId:          databaseId,
		targetApplicationId: targetApplicationId,
	}
}

// Execute executes the request
func (a *DatabaseApplicationApiService) RemoveApplicationFromDatabaseExecute(r ApiRemoveApplicationFromDatabaseRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DatabaseApplicationApiService.RemoveApplicationFromDatabase")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/database/{databaseId}/application/{targetApplicationId}"
	localVarPath = strings.Replace(localVarPath, "{"+"databaseId"+"}", url.PathEscape(parameterToString(r.databaseId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"targetApplicationId"+"}", url.PathEscape(parameterToString(r.targetApplicationId, "")), -1)

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
				localVarHeaderParams["Authorization"] = key
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
