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

// BackupsApiService BackupsApi service
type BackupsApiService service

type ApiAddBackupDatabaseRequest struct {
	ctx           context.Context
	ApiService    *BackupsApiService
	databaseId    string
	backupRequest *BackupRequest
}

func (r ApiAddBackupDatabaseRequest) BackupRequest(backupRequest BackupRequest) ApiAddBackupDatabaseRequest {
	r.backupRequest = &backupRequest
	return r
}

func (r ApiAddBackupDatabaseRequest) Execute() (*BackupResponse, *http.Response, error) {
	return r.ApiService.AddBackupDatabaseExecute(r)
}

/*
AddBackupDatabase Add a backup to the Database

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param databaseId Database ID
 @return ApiAddBackupDatabaseRequest
*/
func (a *BackupsApiService) AddBackupDatabase(ctx context.Context, databaseId string) ApiAddBackupDatabaseRequest {
	return ApiAddBackupDatabaseRequest{
		ApiService: a,
		ctx:        ctx,
		databaseId: databaseId,
	}
}

// Execute executes the request
//  @return BackupResponse
func (a *BackupsApiService) AddBackupDatabaseExecute(r ApiAddBackupDatabaseRequest) (*BackupResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *BackupResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BackupsApiService.AddBackupDatabase")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/database/{databaseId}/backup"
	localVarPath = strings.Replace(localVarPath, "{"+"databaseId"+"}", url.PathEscape(parameterToString(r.databaseId, "")), -1)

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
	localVarPostBody = r.backupRequest
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

type ApiListDatabaseBackupRequest struct {
	ctx        context.Context
	ApiService *BackupsApiService
	databaseId string
	startId    *string
}

// Starting point after which to return results
func (r ApiListDatabaseBackupRequest) StartId(startId string) ApiListDatabaseBackupRequest {
	r.startId = &startId
	return r
}

func (r ApiListDatabaseBackupRequest) Execute() (*BackupPaginatedResponseList, *http.Response, error) {
	return r.ApiService.ListDatabaseBackupExecute(r)
}

/*
ListDatabaseBackup List database  backups

By default it returns the 20 last results. The response is paginated. In order to request the next page, you can use the startId query parameter

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param databaseId Database ID
 @return ApiListDatabaseBackupRequest
*/
func (a *BackupsApiService) ListDatabaseBackup(ctx context.Context, databaseId string) ApiListDatabaseBackupRequest {
	return ApiListDatabaseBackupRequest{
		ApiService: a,
		ctx:        ctx,
		databaseId: databaseId,
	}
}

// Execute executes the request
//  @return BackupPaginatedResponseList
func (a *BackupsApiService) ListDatabaseBackupExecute(r ApiListDatabaseBackupRequest) (*BackupPaginatedResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *BackupPaginatedResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BackupsApiService.ListDatabaseBackup")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/database/{databaseId}/backup"
	localVarPath = strings.Replace(localVarPath, "{"+"databaseId"+"}", url.PathEscape(parameterToString(r.databaseId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.startId != nil {
		localVarQueryParams.Add("startId", parameterToString(*r.startId, ""))
	}
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

type ApiRemoveDatabaseBackupRequest struct {
	ctx        context.Context
	ApiService *BackupsApiService
	databaseId string
	backupId   string
}

func (r ApiRemoveDatabaseBackupRequest) Execute() (*http.Response, error) {
	return r.ApiService.RemoveDatabaseBackupExecute(r)
}

/*
RemoveDatabaseBackup Remove database  backup

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param databaseId Database ID
 @param backupId Database Backup ID
 @return ApiRemoveDatabaseBackupRequest
*/
func (a *BackupsApiService) RemoveDatabaseBackup(ctx context.Context, databaseId string, backupId string) ApiRemoveDatabaseBackupRequest {
	return ApiRemoveDatabaseBackupRequest{
		ApiService: a,
		ctx:        ctx,
		databaseId: databaseId,
		backupId:   backupId,
	}
}

// Execute executes the request
func (a *BackupsApiService) RemoveDatabaseBackupExecute(r ApiRemoveDatabaseBackupRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BackupsApiService.RemoveDatabaseBackup")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/database/{databaseId}/backup/{backupId}"
	localVarPath = strings.Replace(localVarPath, "{"+"databaseId"+"}", url.PathEscape(parameterToString(r.databaseId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"backupId"+"}", url.PathEscape(parameterToString(r.backupId, "")), -1)

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
