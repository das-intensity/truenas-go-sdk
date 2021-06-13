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

// SystemNtpserverApiService SystemNtpserverApi service
type SystemNtpserverApiService service

type ApiSystemNtpserverGetRequest struct {
	ctx _context.Context
	ApiService *SystemNtpserverApiService
	limit *int32
	offset *int32
	count *bool
	sort *string
}

func (r ApiSystemNtpserverGetRequest) Limit(limit int32) ApiSystemNtpserverGetRequest {
	r.limit = &limit
	return r
}
func (r ApiSystemNtpserverGetRequest) Offset(offset int32) ApiSystemNtpserverGetRequest {
	r.offset = &offset
	return r
}
func (r ApiSystemNtpserverGetRequest) Count(count bool) ApiSystemNtpserverGetRequest {
	r.count = &count
	return r
}
func (r ApiSystemNtpserverGetRequest) Sort(sort string) ApiSystemNtpserverGetRequest {
	r.sort = &sort
	return r
}

func (r ApiSystemNtpserverGetRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.SystemNtpserverGetExecute(r)
}

/*
 * SystemNtpserverGet Method for SystemNtpserverGet
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiSystemNtpserverGetRequest
 */
func (a *SystemNtpserverApiService) SystemNtpserverGet(ctx _context.Context) ApiSystemNtpserverGetRequest {
	return ApiSystemNtpserverGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 */
func (a *SystemNtpserverApiService) SystemNtpserverGetExecute(r ApiSystemNtpserverGetRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SystemNtpserverApiService.SystemNtpserverGet")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/system/ntpserver"

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

type ApiSystemNtpserverIdIdDeleteRequest struct {
	ctx _context.Context
	ApiService *SystemNtpserverApiService
	id int32
}


func (r ApiSystemNtpserverIdIdDeleteRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.SystemNtpserverIdIdDeleteExecute(r)
}

/*
 * SystemNtpserverIdIdDelete Method for SystemNtpserverIdIdDelete
 * Delete NTP server of `id`.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id
 * @return ApiSystemNtpserverIdIdDeleteRequest
 */
func (a *SystemNtpserverApiService) SystemNtpserverIdIdDelete(ctx _context.Context, id int32) ApiSystemNtpserverIdIdDeleteRequest {
	return ApiSystemNtpserverIdIdDeleteRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

/*
 * Execute executes the request
 */
func (a *SystemNtpserverApiService) SystemNtpserverIdIdDeleteExecute(r ApiSystemNtpserverIdIdDeleteRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SystemNtpserverApiService.SystemNtpserverIdIdDelete")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/system/ntpserver/id/{id}"
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

type ApiSystemNtpserverIdIdGetRequest struct {
	ctx _context.Context
	ApiService *SystemNtpserverApiService
	id []interface{}
	limit *int32
	offset *int32
	count *bool
	sort *string
}

func (r ApiSystemNtpserverIdIdGetRequest) Limit(limit int32) ApiSystemNtpserverIdIdGetRequest {
	r.limit = &limit
	return r
}
func (r ApiSystemNtpserverIdIdGetRequest) Offset(offset int32) ApiSystemNtpserverIdIdGetRequest {
	r.offset = &offset
	return r
}
func (r ApiSystemNtpserverIdIdGetRequest) Count(count bool) ApiSystemNtpserverIdIdGetRequest {
	r.count = &count
	return r
}
func (r ApiSystemNtpserverIdIdGetRequest) Sort(sort string) ApiSystemNtpserverIdIdGetRequest {
	r.sort = &sort
	return r
}

func (r ApiSystemNtpserverIdIdGetRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.SystemNtpserverIdIdGetExecute(r)
}

/*
 * SystemNtpserverIdIdGet Method for SystemNtpserverIdIdGet
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id
 * @return ApiSystemNtpserverIdIdGetRequest
 */
func (a *SystemNtpserverApiService) SystemNtpserverIdIdGet(ctx _context.Context, id []interface{}) ApiSystemNtpserverIdIdGetRequest {
	return ApiSystemNtpserverIdIdGetRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

/*
 * Execute executes the request
 */
func (a *SystemNtpserverApiService) SystemNtpserverIdIdGetExecute(r ApiSystemNtpserverIdIdGetRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SystemNtpserverApiService.SystemNtpserverIdIdGet")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/system/ntpserver/id/{id}"
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

type ApiSystemNtpserverIdIdPutRequest struct {
	ctx _context.Context
	ApiService *SystemNtpserverApiService
	id int32
	systemNtpserverUpdate1 *SystemNtpserverUpdate1
}

func (r ApiSystemNtpserverIdIdPutRequest) SystemNtpserverUpdate1(systemNtpserverUpdate1 SystemNtpserverUpdate1) ApiSystemNtpserverIdIdPutRequest {
	r.systemNtpserverUpdate1 = &systemNtpserverUpdate1
	return r
}

func (r ApiSystemNtpserverIdIdPutRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.SystemNtpserverIdIdPutExecute(r)
}

/*
 * SystemNtpserverIdIdPut Method for SystemNtpserverIdIdPut
 * Update NTP server of `id`.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id
 * @return ApiSystemNtpserverIdIdPutRequest
 */
func (a *SystemNtpserverApiService) SystemNtpserverIdIdPut(ctx _context.Context, id int32) ApiSystemNtpserverIdIdPutRequest {
	return ApiSystemNtpserverIdIdPutRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

/*
 * Execute executes the request
 */
func (a *SystemNtpserverApiService) SystemNtpserverIdIdPutExecute(r ApiSystemNtpserverIdIdPutRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SystemNtpserverApiService.SystemNtpserverIdIdPut")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/system/ntpserver/id/{id}"
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
	localVarPostBody = r.systemNtpserverUpdate1
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

type ApiSystemNtpserverPostRequest struct {
	ctx _context.Context
	ApiService *SystemNtpserverApiService
	systemNtpserverCreate0 *SystemNtpserverCreate0
}

func (r ApiSystemNtpserverPostRequest) SystemNtpserverCreate0(systemNtpserverCreate0 SystemNtpserverCreate0) ApiSystemNtpserverPostRequest {
	r.systemNtpserverCreate0 = &systemNtpserverCreate0
	return r
}

func (r ApiSystemNtpserverPostRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.SystemNtpserverPostExecute(r)
}

/*
 * SystemNtpserverPost Method for SystemNtpserverPost
 * Add an NTP Server.

`address` specifies the hostname/IP address of the NTP server.

`burst` when enabled makes sure that if server is reachable, sends a burst of eight packets instead of one.
This is designed to improve timekeeping quality with the server command.

`iburst` when enabled speeds up the initial synchronization, taking seconds rather than minutes.

`prefer` marks the specified server as preferred. When all other things are equal, this host is chosen
for synchronization acquisition with the server command. It is recommended that they be used for servers with
time monitoring hardware.

`minpoll` is minimum polling time in seconds. It must be a power of 2 and less than `maxpoll`.

`maxpoll` is maximum polling time in seconds. It must be a power of 2 and greater than `minpoll`.

`force` when enabled forces the addition of NTP server even if it is currently unreachable.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiSystemNtpserverPostRequest
 */
func (a *SystemNtpserverApiService) SystemNtpserverPost(ctx _context.Context) ApiSystemNtpserverPostRequest {
	return ApiSystemNtpserverPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 */
func (a *SystemNtpserverApiService) SystemNtpserverPostExecute(r ApiSystemNtpserverPostRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SystemNtpserverApiService.SystemNtpserverPost")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/system/ntpserver"

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
	localVarPostBody = r.systemNtpserverCreate0
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

type ApiSystemNtpserverTestNtpServerGetRequest struct {
	ctx _context.Context
	ApiService *SystemNtpserverApiService
}


func (r ApiSystemNtpserverTestNtpServerGetRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.SystemNtpserverTestNtpServerGetExecute(r)
}

/*
 * SystemNtpserverTestNtpServerGet Method for SystemNtpserverTestNtpServerGet
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiSystemNtpserverTestNtpServerGetRequest
 */
func (a *SystemNtpserverApiService) SystemNtpserverTestNtpServerGet(ctx _context.Context) ApiSystemNtpserverTestNtpServerGetRequest {
	return ApiSystemNtpserverTestNtpServerGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 */
func (a *SystemNtpserverApiService) SystemNtpserverTestNtpServerGetExecute(r ApiSystemNtpserverTestNtpServerGetRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SystemNtpserverApiService.SystemNtpserverTestNtpServerGet")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/system/ntpserver/test_ntp_server"

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