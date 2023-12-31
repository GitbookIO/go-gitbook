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

// CollectionsApiService CollectionsApi service
type CollectionsApiService service

type CollectionsApiGetCollectionByIdRequest struct {
	ctx          context.Context
	ApiService   *CollectionsApiService
	collectionId string
}

func (r CollectionsApiGetCollectionByIdRequest) Execute() (*Collection, *http.Response, error) {
	return r.ApiService.GetCollectionByIdExecute(r)
}

/*
GetCollectionById Get the details about a collection using its ID

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param collectionId The unique id of the collection
	@return CollectionsApiGetCollectionByIdRequest
*/
func (a *CollectionsApiService) GetCollectionById(ctx context.Context, collectionId string) CollectionsApiGetCollectionByIdRequest {
	return CollectionsApiGetCollectionByIdRequest{
		ApiService:   a,
		ctx:          ctx,
		collectionId: collectionId,
	}
}

// Execute executes the request
//
//	@return Collection
func (a *CollectionsApiService) GetCollectionByIdExecute(r CollectionsApiGetCollectionByIdRequest) (*Collection, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Collection
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CollectionsApiService.GetCollectionById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/collections/{collectionId}"
	localVarPath = strings.Replace(localVarPath, "{"+"collectionId"+"}", url.PathEscape(parameterValueToString(r.collectionId, "collectionId")), -1)

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

type CollectionsApiGetCollectionPublishingAuthByIdRequest struct {
	ctx          context.Context
	ApiService   *CollectionsApiService
	collectionId string
}

func (r CollectionsApiGetCollectionPublishingAuthByIdRequest) Execute() (*ContentPublishingAuth, *http.Response, error) {
	return r.ApiService.GetCollectionPublishingAuthByIdExecute(r)
}

/*
GetCollectionPublishingAuthById Get the publishing authentication configuration for a collection.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param collectionId The unique id of the collection
	@return CollectionsApiGetCollectionPublishingAuthByIdRequest
*/
func (a *CollectionsApiService) GetCollectionPublishingAuthById(ctx context.Context, collectionId string) CollectionsApiGetCollectionPublishingAuthByIdRequest {
	return CollectionsApiGetCollectionPublishingAuthByIdRequest{
		ApiService:   a,
		ctx:          ctx,
		collectionId: collectionId,
	}
}

// Execute executes the request
//
//	@return ContentPublishingAuth
func (a *CollectionsApiService) GetCollectionPublishingAuthByIdExecute(r CollectionsApiGetCollectionPublishingAuthByIdRequest) (*ContentPublishingAuth, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ContentPublishingAuth
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CollectionsApiService.GetCollectionPublishingAuthById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/collections/{collectionId}/publishing/auth"
	localVarPath = strings.Replace(localVarPath, "{"+"collectionId"+"}", url.PathEscape(parameterValueToString(r.collectionId, "collectionId")), -1)

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
		if localVarHTTPResponse.StatusCode == 400 {
			var v GetSpacePublishingAuthById400Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
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

type CollectionsApiListSpacesInCollectionByIdRequest struct {
	ctx          context.Context
	ApiService   *CollectionsApiService
	collectionId string
	page         *string
	limit        *float32
}

// Identifier of the page results to fetch.
func (r CollectionsApiListSpacesInCollectionByIdRequest) Page(page string) CollectionsApiListSpacesInCollectionByIdRequest {
	r.page = &page
	return r
}

// The number of results per page
func (r CollectionsApiListSpacesInCollectionByIdRequest) Limit(limit float32) CollectionsApiListSpacesInCollectionByIdRequest {
	r.limit = &limit
	return r
}

func (r CollectionsApiListSpacesInCollectionByIdRequest) Execute() (*ListSpacesForAuthenticatedUser200Response, *http.Response, error) {
	return r.ApiService.ListSpacesInCollectionByIdExecute(r)
}

/*
ListSpacesInCollectionById List all the spaces in a collection by its ID.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param collectionId The unique id of the collection
	@return CollectionsApiListSpacesInCollectionByIdRequest
*/
func (a *CollectionsApiService) ListSpacesInCollectionById(ctx context.Context, collectionId string) CollectionsApiListSpacesInCollectionByIdRequest {
	return CollectionsApiListSpacesInCollectionByIdRequest{
		ApiService:   a,
		ctx:          ctx,
		collectionId: collectionId,
	}
}

// Execute executes the request
//
//	@return ListSpacesForAuthenticatedUser200Response
func (a *CollectionsApiService) ListSpacesInCollectionByIdExecute(r CollectionsApiListSpacesInCollectionByIdRequest) (*ListSpacesForAuthenticatedUser200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ListSpacesForAuthenticatedUser200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CollectionsApiService.ListSpacesInCollectionById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/collections/{collectionId}/spaces"
	localVarPath = strings.Replace(localVarPath, "{"+"collectionId"+"}", url.PathEscape(parameterValueToString(r.collectionId, "collectionId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
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

type CollectionsApiUpdateCollectionByIdRequest struct {
	ctx                         context.Context
	ApiService                  *CollectionsApiService
	collectionId                string
	updateCollectionByIdRequest *UpdateCollectionByIdRequest
}

func (r CollectionsApiUpdateCollectionByIdRequest) UpdateCollectionByIdRequest(updateCollectionByIdRequest UpdateCollectionByIdRequest) CollectionsApiUpdateCollectionByIdRequest {
	r.updateCollectionByIdRequest = &updateCollectionByIdRequest
	return r
}

func (r CollectionsApiUpdateCollectionByIdRequest) Execute() (*Collection, *http.Response, error) {
	return r.ApiService.UpdateCollectionByIdExecute(r)
}

/*
UpdateCollectionById Update the details of a collection

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param collectionId The unique id of the collection
	@return CollectionsApiUpdateCollectionByIdRequest
*/
func (a *CollectionsApiService) UpdateCollectionById(ctx context.Context, collectionId string) CollectionsApiUpdateCollectionByIdRequest {
	return CollectionsApiUpdateCollectionByIdRequest{
		ApiService:   a,
		ctx:          ctx,
		collectionId: collectionId,
	}
}

// Execute executes the request
//
//	@return Collection
func (a *CollectionsApiService) UpdateCollectionByIdExecute(r CollectionsApiUpdateCollectionByIdRequest) (*Collection, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Collection
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CollectionsApiService.UpdateCollectionById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/collections/{collectionId}"
	localVarPath = strings.Replace(localVarPath, "{"+"collectionId"+"}", url.PathEscape(parameterValueToString(r.collectionId, "collectionId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.updateCollectionByIdRequest == nil {
		return localVarReturnValue, nil, reportError("updateCollectionByIdRequest is required and must be specified")
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
	localVarPostBody = r.updateCollectionByIdRequest
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

type CollectionsApiUpdateCollectionPublishingAuthByIdRequest struct {
	ctx                                context.Context
	ApiService                         *CollectionsApiService
	collectionId                       string
	requestUpdateContentPublishingAuth *RequestUpdateContentPublishingAuth
}

func (r CollectionsApiUpdateCollectionPublishingAuthByIdRequest) RequestUpdateContentPublishingAuth(requestUpdateContentPublishingAuth RequestUpdateContentPublishingAuth) CollectionsApiUpdateCollectionPublishingAuthByIdRequest {
	r.requestUpdateContentPublishingAuth = &requestUpdateContentPublishingAuth
	return r
}

func (r CollectionsApiUpdateCollectionPublishingAuthByIdRequest) Execute() (*ContentPublishingAuth, *http.Response, error) {
	return r.ApiService.UpdateCollectionPublishingAuthByIdExecute(r)
}

/*
UpdateCollectionPublishingAuthById Update the publishing authentication configuration for a collection.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param collectionId The unique id of the collection
	@return CollectionsApiUpdateCollectionPublishingAuthByIdRequest
*/
func (a *CollectionsApiService) UpdateCollectionPublishingAuthById(ctx context.Context, collectionId string) CollectionsApiUpdateCollectionPublishingAuthByIdRequest {
	return CollectionsApiUpdateCollectionPublishingAuthByIdRequest{
		ApiService:   a,
		ctx:          ctx,
		collectionId: collectionId,
	}
}

// Execute executes the request
//
//	@return ContentPublishingAuth
func (a *CollectionsApiService) UpdateCollectionPublishingAuthByIdExecute(r CollectionsApiUpdateCollectionPublishingAuthByIdRequest) (*ContentPublishingAuth, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ContentPublishingAuth
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CollectionsApiService.UpdateCollectionPublishingAuthById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/collections/{collectionId}/publishing/auth"
	localVarPath = strings.Replace(localVarPath, "{"+"collectionId"+"}", url.PathEscape(parameterValueToString(r.collectionId, "collectionId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.requestUpdateContentPublishingAuth == nil {
		return localVarReturnValue, nil, reportError("requestUpdateContentPublishingAuth is required and must be specified")
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
	localVarPostBody = r.requestUpdateContentPublishingAuth
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v GetSpacePublishingAuthById400Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
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
