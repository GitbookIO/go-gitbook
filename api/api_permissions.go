/*
Copyright 2023 GitBook, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

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

// PermissionsApiService PermissionsApi service
type PermissionsApiService service

type PermissionsApiListPermissionsAggregateInCollectionRequest struct {
	ctx          context.Context
	ApiService   *PermissionsApiService
	collectionId string
	page         *string
	limit        *float32
	role         *MemberRole
}

// Identifier of the page results to fetch.
func (r PermissionsApiListPermissionsAggregateInCollectionRequest) Page(page string) PermissionsApiListPermissionsAggregateInCollectionRequest {
	r.page = &page
	return r
}

// The number of results per page
func (r PermissionsApiListPermissionsAggregateInCollectionRequest) Limit(limit float32) PermissionsApiListPermissionsAggregateInCollectionRequest {
	r.limit = &limit
	return r
}

// If defined, only members with this role will be returned.
func (r PermissionsApiListPermissionsAggregateInCollectionRequest) Role(role MemberRole) PermissionsApiListPermissionsAggregateInCollectionRequest {
	r.role = &role
	return r
}

func (r PermissionsApiListPermissionsAggregateInCollectionRequest) Execute() (*ListPermissionsAggregateInSpace200Response, *http.Response, error) {
	return r.ApiService.ListPermissionsAggregateInCollectionExecute(r)
}

/*
ListPermissionsAggregateInCollection List permissions for all users in a collection.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param collectionId The unique id of the collection
	@return PermissionsApiListPermissionsAggregateInCollectionRequest
*/
func (a *PermissionsApiService) ListPermissionsAggregateInCollection(ctx context.Context, collectionId string) PermissionsApiListPermissionsAggregateInCollectionRequest {
	return PermissionsApiListPermissionsAggregateInCollectionRequest{
		ApiService:   a,
		ctx:          ctx,
		collectionId: collectionId,
	}
}

// Execute executes the request
//
//	@return ListPermissionsAggregateInSpace200Response
func (a *PermissionsApiService) ListPermissionsAggregateInCollectionExecute(r PermissionsApiListPermissionsAggregateInCollectionRequest) (*ListPermissionsAggregateInSpace200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ListPermissionsAggregateInSpace200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PermissionsApiService.ListPermissionsAggregateInCollection")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/collections/{collectionId}/permissions/aggregate"
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
	if r.role != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "role", r.role, "")
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

type PermissionsApiListPermissionsAggregateInSpaceRequest struct {
	ctx        context.Context
	ApiService *PermissionsApiService
	spaceId    string
	page       *string
	limit      *float32
	role       *MemberRole
}

// Identifier of the page results to fetch.
func (r PermissionsApiListPermissionsAggregateInSpaceRequest) Page(page string) PermissionsApiListPermissionsAggregateInSpaceRequest {
	r.page = &page
	return r
}

// The number of results per page
func (r PermissionsApiListPermissionsAggregateInSpaceRequest) Limit(limit float32) PermissionsApiListPermissionsAggregateInSpaceRequest {
	r.limit = &limit
	return r
}

// If defined, only members with this role will be returned.
func (r PermissionsApiListPermissionsAggregateInSpaceRequest) Role(role MemberRole) PermissionsApiListPermissionsAggregateInSpaceRequest {
	r.role = &role
	return r
}

func (r PermissionsApiListPermissionsAggregateInSpaceRequest) Execute() (*ListPermissionsAggregateInSpace200Response, *http.Response, error) {
	return r.ApiService.ListPermissionsAggregateInSpaceExecute(r)
}

/*
ListPermissionsAggregateInSpace List permissions for all users in a space.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param spaceId The unique id of the space
	@return PermissionsApiListPermissionsAggregateInSpaceRequest
*/
func (a *PermissionsApiService) ListPermissionsAggregateInSpace(ctx context.Context, spaceId string) PermissionsApiListPermissionsAggregateInSpaceRequest {
	return PermissionsApiListPermissionsAggregateInSpaceRequest{
		ApiService: a,
		ctx:        ctx,
		spaceId:    spaceId,
	}
}

// Execute executes the request
//
//	@return ListPermissionsAggregateInSpace200Response
func (a *PermissionsApiService) ListPermissionsAggregateInSpaceExecute(r PermissionsApiListPermissionsAggregateInSpaceRequest) (*ListPermissionsAggregateInSpace200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ListPermissionsAggregateInSpace200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PermissionsApiService.ListPermissionsAggregateInSpace")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/spaces/{spaceId}/permissions/aggregate"
	localVarPath = strings.Replace(localVarPath, "{"+"spaceId"+"}", url.PathEscape(parameterValueToString(r.spaceId, "spaceId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	}
	if r.role != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "role", r.role, "")
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

type PermissionsApiListSpacesForOrganizationMemberRequest struct {
	ctx            context.Context
	ApiService     *PermissionsApiService
	organizationId string
	userId         string
	page           *string
	limit          *float32
}

// Identifier of the page results to fetch.
func (r PermissionsApiListSpacesForOrganizationMemberRequest) Page(page string) PermissionsApiListSpacesForOrganizationMemberRequest {
	r.page = &page
	return r
}

// The number of results per page
func (r PermissionsApiListSpacesForOrganizationMemberRequest) Limit(limit float32) PermissionsApiListSpacesForOrganizationMemberRequest {
	r.limit = &limit
	return r
}

func (r PermissionsApiListSpacesForOrganizationMemberRequest) Execute() (*ListSpacesForOrganizationMember200Response, *http.Response, error) {
	return r.ApiService.ListSpacesForOrganizationMemberExecute(r)
}

/*
ListSpacesForOrganizationMember List permissions accross all spaces for a member of an organization

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param organizationId The unique id of the organization
	@param userId The unique ID of the User
	@return PermissionsApiListSpacesForOrganizationMemberRequest
*/
func (a *PermissionsApiService) ListSpacesForOrganizationMember(ctx context.Context, organizationId string, userId string) PermissionsApiListSpacesForOrganizationMemberRequest {
	return PermissionsApiListSpacesForOrganizationMemberRequest{
		ApiService:     a,
		ctx:            ctx,
		organizationId: organizationId,
		userId:         userId,
	}
}

// Execute executes the request
//
//	@return ListSpacesForOrganizationMember200Response
func (a *PermissionsApiService) ListSpacesForOrganizationMemberExecute(r PermissionsApiListSpacesForOrganizationMemberRequest) (*ListSpacesForOrganizationMember200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ListSpacesForOrganizationMember200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PermissionsApiService.ListSpacesForOrganizationMember")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/orgs/{organizationId}/members/{userId}/spaces"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", url.PathEscape(parameterValueToString(r.organizationId, "organizationId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"userId"+"}", url.PathEscape(parameterValueToString(r.userId, "userId")), -1)

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
