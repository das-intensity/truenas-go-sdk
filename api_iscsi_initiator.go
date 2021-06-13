/*
 * TrueNAS RESTful API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v2.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package truenas

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
)

// Linger please
var (
	_ _context.Context
)

// IscsiInitiatorApiService IscsiInitiatorApi service
type IscsiInitiatorApiService service

type ApiIscsiInitiatorGetRequest struct {
	ctx _context.Context
	ApiService *IscsiInitiatorApiService
	limit *int32
	offset *int32
	count *bool
	sort *string
}

func (r ApiIscsiInitiatorGetRequest) Limit(limit int32) ApiIscsiInitiatorGetRequest {
	r.limit = &limit
	return r
}
func (r ApiIscsiInitiatorGetRequest) Offset(offset int32) ApiIscsiInitiatorGetRequest {
	r.offset = &offset
	return r
}
func (r ApiIscsiInitiatorGetRequest) Count(count bool) ApiIscsiInitiatorGetRequest {
	r.count = &count
	return r
}
func (r ApiIscsiInitiatorGetRequest) Sort(sort string) ApiIscsiInitiatorGetRequest {
	r.sort = &sort
	return r
}

func (r ApiIscsiInitiatorGetRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.IscsiInitiatorGetExecute(r)
}

/*
 * IscsiInitiatorGet Method for IscsiInitiatorGet
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiIscsiInitiatorGetRequest
 */
func (a *IscsiInitiatorApiService) IscsiInitiatorGet(ctx _context.Context) ApiIscsiInitiatorGetRequest {
	return ApiIscsiInitiatorGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 */
func (a *IscsiInitiatorApiService) IscsiInitiatorGetExecute(r ApiIscsiInitiatorGetRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IscsiInitiatorApiService.IscsiInitiatorGet")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/iscsi/initiator"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	if r.offset != nil {
		localVarQueryParams.Add("offset", parameterToString(*r.offset, ""))
	}
	if r.count != nil {
		localVarQueryParams.Add("count", parameterToString(*r.count, ""))
	}
	if r.sort != nil {
		localVarQueryParams.Add("sort", parameterToString(*r.sort, ""))
	}
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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiIscsiInitiatorIdIdDeleteRequest struct {
	ctx _context.Context
	ApiService *IscsiInitiatorApiService
	id int32
}


func (r ApiIscsiInitiatorIdIdDeleteRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.IscsiInitiatorIdIdDeleteExecute(r)
}

/*
 * IscsiInitiatorIdIdDelete Method for IscsiInitiatorIdIdDelete
 * Delete iSCSI initiator of `id`.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id
 * @return ApiIscsiInitiatorIdIdDeleteRequest
 */
func (a *IscsiInitiatorApiService) IscsiInitiatorIdIdDelete(ctx _context.Context, id int32) ApiIscsiInitiatorIdIdDeleteRequest {
	return ApiIscsiInitiatorIdIdDeleteRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

/*
 * Execute executes the request
 */
func (a *IscsiInitiatorApiService) IscsiInitiatorIdIdDeleteExecute(r ApiIscsiInitiatorIdIdDeleteRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IscsiInitiatorApiService.IscsiInitiatorIdIdDelete")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/iscsi/initiator/id/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", _neturl.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiIscsiInitiatorIdIdGetRequest struct {
	ctx _context.Context
	ApiService *IscsiInitiatorApiService
	id []interface{}
	limit *int32
	offset *int32
	count *bool
	sort *string
}

func (r ApiIscsiInitiatorIdIdGetRequest) Limit(limit int32) ApiIscsiInitiatorIdIdGetRequest {
	r.limit = &limit
	return r
}
func (r ApiIscsiInitiatorIdIdGetRequest) Offset(offset int32) ApiIscsiInitiatorIdIdGetRequest {
	r.offset = &offset
	return r
}
func (r ApiIscsiInitiatorIdIdGetRequest) Count(count bool) ApiIscsiInitiatorIdIdGetRequest {
	r.count = &count
	return r
}
func (r ApiIscsiInitiatorIdIdGetRequest) Sort(sort string) ApiIscsiInitiatorIdIdGetRequest {
	r.sort = &sort
	return r
}

func (r ApiIscsiInitiatorIdIdGetRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.IscsiInitiatorIdIdGetExecute(r)
}

/*
 * IscsiInitiatorIdIdGet Method for IscsiInitiatorIdIdGet
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id
 * @return ApiIscsiInitiatorIdIdGetRequest
 */
func (a *IscsiInitiatorApiService) IscsiInitiatorIdIdGet(ctx _context.Context, id []interface{}) ApiIscsiInitiatorIdIdGetRequest {
	return ApiIscsiInitiatorIdIdGetRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

/*
 * Execute executes the request
 */
func (a *IscsiInitiatorApiService) IscsiInitiatorIdIdGetExecute(r ApiIscsiInitiatorIdIdGetRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IscsiInitiatorApiService.IscsiInitiatorIdIdGet")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/iscsi/initiator/id/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", _neturl.PathEscape(parameterToString(r.id, "csv")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	if r.offset != nil {
		localVarQueryParams.Add("offset", parameterToString(*r.offset, ""))
	}
	if r.count != nil {
		localVarQueryParams.Add("count", parameterToString(*r.count, ""))
	}
	if r.sort != nil {
		localVarQueryParams.Add("sort", parameterToString(*r.sort, ""))
	}
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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiIscsiInitiatorIdIdPutRequest struct {
	ctx _context.Context
	ApiService *IscsiInitiatorApiService
	id int32
	iscsiInitiatorUpdate1 *IscsiInitiatorUpdate1
}

func (r ApiIscsiInitiatorIdIdPutRequest) IscsiInitiatorUpdate1(iscsiInitiatorUpdate1 IscsiInitiatorUpdate1) ApiIscsiInitiatorIdIdPutRequest {
	r.iscsiInitiatorUpdate1 = &iscsiInitiatorUpdate1
	return r
}

func (r ApiIscsiInitiatorIdIdPutRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.IscsiInitiatorIdIdPutExecute(r)
}

/*
 * IscsiInitiatorIdIdPut Method for IscsiInitiatorIdIdPut
 * Update iSCSI initiator of `id`.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id
 * @return ApiIscsiInitiatorIdIdPutRequest
 */
func (a *IscsiInitiatorApiService) IscsiInitiatorIdIdPut(ctx _context.Context, id int32) ApiIscsiInitiatorIdIdPutRequest {
	return ApiIscsiInitiatorIdIdPutRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

/*
 * Execute executes the request
 */
func (a *IscsiInitiatorApiService) IscsiInitiatorIdIdPutExecute(r ApiIscsiInitiatorIdIdPutRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IscsiInitiatorApiService.IscsiInitiatorIdIdPut")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/iscsi/initiator/id/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", _neturl.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.iscsiInitiatorUpdate1
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiIscsiInitiatorPostRequest struct {
	ctx _context.Context
	ApiService *IscsiInitiatorApiService
	iscsiInitiatorCreate0 *IscsiInitiatorCreate0
}

func (r ApiIscsiInitiatorPostRequest) IscsiInitiatorCreate0(iscsiInitiatorCreate0 IscsiInitiatorCreate0) ApiIscsiInitiatorPostRequest {
	r.iscsiInitiatorCreate0 = &iscsiInitiatorCreate0
	return r
}

func (r ApiIscsiInitiatorPostRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.IscsiInitiatorPostExecute(r)
}

/*
 * IscsiInitiatorPost Method for IscsiInitiatorPost
 * Create an iSCSI Initiator.

`initiators` is a list of initiator hostnames which are authorized to access an iSCSI Target. To allow all
possible initiators, `initiators` can be left empty.

`auth_network` is a list of IP/CIDR addresses which are allowed to use this initiator. If all networks are
to be allowed, this field should be left empty.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiIscsiInitiatorPostRequest
 */
func (a *IscsiInitiatorApiService) IscsiInitiatorPost(ctx _context.Context) ApiIscsiInitiatorPostRequest {
	return ApiIscsiInitiatorPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 */
func (a *IscsiInitiatorApiService) IscsiInitiatorPostExecute(r ApiIscsiInitiatorPostRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IscsiInitiatorApiService.IscsiInitiatorPost")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/iscsi/initiator"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.iscsiInitiatorCreate0
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}