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

// HiveApiService HiveApi service
type HiveApiService service

type HiveApiGenerateHiveAccessTokenRequest struct {
	ctx                            context.Context
	ApiService                     *HiveApiService
	generateHiveAccessTokenRequest *GenerateHiveAccessTokenRequest
}

func (r HiveApiGenerateHiveAccessTokenRequest) GenerateHiveAccessTokenRequest(generateHiveAccessTokenRequest GenerateHiveAccessTokenRequest) HiveApiGenerateHiveAccessTokenRequest {
	r.generateHiveAccessTokenRequest = &generateHiveAccessTokenRequest
	return r
}

func (r HiveApiGenerateHiveAccessTokenRequest) Execute() (*HiveAccessToken, *http.Response, error) {
	return r.ApiService.GenerateHiveAccessTokenExecute(r)
}

/*
GenerateHiveAccessToken Returns a token to authenticate with Hive.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return HiveApiGenerateHiveAccessTokenRequest
*/
func (a *HiveApiService) GenerateHiveAccessToken(ctx context.Context) HiveApiGenerateHiveAccessTokenRequest {
	return HiveApiGenerateHiveAccessTokenRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return HiveAccessToken
func (a *HiveApiService) GenerateHiveAccessTokenExecute(r HiveApiGenerateHiveAccessTokenRequest) (*HiveAccessToken, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *HiveAccessToken
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "HiveApiService.GenerateHiveAccessToken")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/internal/hive/token"

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
	localVarPostBody = r.generateHiveAccessTokenRequest
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

type HiveApiGenerateSpaceHiveReadAccessTokenRequest struct {
	ctx        context.Context
	ApiService *HiveApiService
	spaceId    string
}

func (r HiveApiGenerateSpaceHiveReadAccessTokenRequest) Execute() (*HiveAccessToken, *http.Response, error) {
	return r.ApiService.GenerateSpaceHiveReadAccessTokenExecute(r)
}

/*
GenerateSpaceHiveReadAccessToken Returns a token to authenticate with Hive to read content from a given space.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param spaceId The unique id of the space
	@return HiveApiGenerateSpaceHiveReadAccessTokenRequest
*/
func (a *HiveApiService) GenerateSpaceHiveReadAccessToken(ctx context.Context, spaceId string) HiveApiGenerateSpaceHiveReadAccessTokenRequest {
	return HiveApiGenerateSpaceHiveReadAccessTokenRequest{
		ApiService: a,
		ctx:        ctx,
		spaceId:    spaceId,
	}
}

// Execute executes the request
//
//	@return HiveAccessToken
func (a *HiveApiService) GenerateSpaceHiveReadAccessTokenExecute(r HiveApiGenerateSpaceHiveReadAccessTokenRequest) (*HiveAccessToken, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *HiveAccessToken
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "HiveApiService.GenerateSpaceHiveReadAccessToken")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/spaces/{spaceId}/hive/token"
	localVarPath = strings.Replace(localVarPath, "{"+"spaceId"+"}", url.PathEscape(parameterValueToString(r.spaceId, "spaceId")), -1)

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
