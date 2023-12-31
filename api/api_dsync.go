// Copyright 2023 GitBook, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gitbook

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// DsyncApiService DsyncApi service
type DsyncApiService service

type DsyncApiListDirectorySyncGroupsRequest struct {
	ctx            context.Context
	ApiService     *DsyncApiService
	organizationId string
}

func (r DsyncApiListDirectorySyncGroupsRequest) Execute() (*ListDirectorySyncGroups200Response, *http.Response, error) {
	return r.ApiService.ListDirectorySyncGroupsExecute(r)
}

/*
ListDirectorySyncGroups Lists the groups exposed to the synced Directory on an organization.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param organizationId The unique id of the organization
	@return DsyncApiListDirectorySyncGroupsRequest
*/
func (a *DsyncApiService) ListDirectorySyncGroups(ctx context.Context, organizationId string) DsyncApiListDirectorySyncGroupsRequest {
	return DsyncApiListDirectorySyncGroupsRequest{
		ApiService:     a,
		ctx:            ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//
//	@return ListDirectorySyncGroups200Response
func (a *DsyncApiService) ListDirectorySyncGroupsExecute(r DsyncApiListDirectorySyncGroupsRequest) (*ListDirectorySyncGroups200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ListDirectorySyncGroups200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DsyncApiService.ListDirectorySyncGroups")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/orgs/{organizationId}/dsync/groups"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", url.PathEscape(parameterValueToString(r.organizationId, "organizationId")), -1)

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

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v Error
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
		newErr.model = v
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

type DsyncApiSetupDirectorySyncRequest struct {
	ctx            context.Context
	ApiService     *DsyncApiService
	organizationId string
}

func (r DsyncApiSetupDirectorySyncRequest) Execute() (*SetupDirectorySync200Response, *http.Response, error) {
	return r.ApiService.SetupDirectorySyncExecute(r)
}

/*
SetupDirectorySync Set up Directory Sync in an organization.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param organizationId The unique id of the organization
	@return DsyncApiSetupDirectorySyncRequest
*/
func (a *DsyncApiService) SetupDirectorySync(ctx context.Context, organizationId string) DsyncApiSetupDirectorySyncRequest {
	return DsyncApiSetupDirectorySyncRequest{
		ApiService:     a,
		ctx:            ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//
//	@return SetupDirectorySync200Response
func (a *DsyncApiService) SetupDirectorySyncExecute(r DsyncApiSetupDirectorySyncRequest) (*SetupDirectorySync200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *SetupDirectorySync200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DsyncApiService.SetupDirectorySync")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/orgs/{organizationId}/dsync"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", url.PathEscape(parameterValueToString(r.organizationId, "organizationId")), -1)

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

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v Error
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
		newErr.model = v
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

type DsyncApiSyncDirectorySyncGroupsToTeamsRequest struct {
	ctx                                   context.Context
	ApiService                            *DsyncApiService
	organizationId                        string
	syncDirectorySyncGroupsToTeamsRequest *SyncDirectorySyncGroupsToTeamsRequest
}

func (r DsyncApiSyncDirectorySyncGroupsToTeamsRequest) SyncDirectorySyncGroupsToTeamsRequest(syncDirectorySyncGroupsToTeamsRequest SyncDirectorySyncGroupsToTeamsRequest) DsyncApiSyncDirectorySyncGroupsToTeamsRequest {
	r.syncDirectorySyncGroupsToTeamsRequest = &syncDirectorySyncGroupsToTeamsRequest
	return r
}

func (r DsyncApiSyncDirectorySyncGroupsToTeamsRequest) Execute() (*SyncDirectorySyncGroupsToTeams200Response, *http.Response, error) {
	return r.ApiService.SyncDirectorySyncGroupsToTeamsExecute(r)
}

/*
SyncDirectorySyncGroupsToTeams Syncs a list of group/team unique identifiers pairs together.

Each pair will lead the team being synced to the group (unless it is already synced to another) and the members of the group will then be synced to the team.
Each sync is independent, the endpoint will still succeed if some of the syncs fail.
Check the response to know more about the success or failure of the entire sync.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param organizationId The unique id of the organization
	@return DsyncApiSyncDirectorySyncGroupsToTeamsRequest
*/
func (a *DsyncApiService) SyncDirectorySyncGroupsToTeams(ctx context.Context, organizationId string) DsyncApiSyncDirectorySyncGroupsToTeamsRequest {
	return DsyncApiSyncDirectorySyncGroupsToTeamsRequest{
		ApiService:     a,
		ctx:            ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//
//	@return SyncDirectorySyncGroupsToTeams200Response
func (a *DsyncApiService) SyncDirectorySyncGroupsToTeamsExecute(r DsyncApiSyncDirectorySyncGroupsToTeamsRequest) (*SyncDirectorySyncGroupsToTeams200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *SyncDirectorySyncGroupsToTeams200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DsyncApiService.SyncDirectorySyncGroupsToTeams")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/orgs/{organizationId}/dsync/teams"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", url.PathEscape(parameterValueToString(r.organizationId, "organizationId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.syncDirectorySyncGroupsToTeamsRequest == nil {
		return localVarReturnValue, nil, reportError("syncDirectorySyncGroupsToTeamsRequest is required and must be specified")
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
	localVarPostBody = r.syncDirectorySyncGroupsToTeamsRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v Error
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
		newErr.model = v
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
