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
)

// Linger please
var (
	_ _context.Context
)

// FilesystemApiService FilesystemApi service
type FilesystemApiService service

type ApiFilesystemAclIsTrivialPostRequest struct {
	ctx _context.Context
	ApiService *FilesystemApiService
	body *string
}

func (r ApiFilesystemAclIsTrivialPostRequest) Body(body string) ApiFilesystemAclIsTrivialPostRequest {
	r.body = &body
	return r
}

func (r ApiFilesystemAclIsTrivialPostRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.FilesystemAclIsTrivialPostExecute(r)
}

/*
 * FilesystemAclIsTrivialPost Method for FilesystemAclIsTrivialPost
 * Returns True if the ACL can be fully expressed as a file mode without losing
any access rules, or if the path does not support NFSv4 ACLs (for example
a path on a tmpfs filesystem).
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiFilesystemAclIsTrivialPostRequest
 */
func (a *FilesystemApiService) FilesystemAclIsTrivialPost(ctx _context.Context) ApiFilesystemAclIsTrivialPostRequest {
	return ApiFilesystemAclIsTrivialPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 */
func (a *FilesystemApiService) FilesystemAclIsTrivialPostExecute(r ApiFilesystemAclIsTrivialPostRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FilesystemApiService.FilesystemAclIsTrivialPost")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/filesystem/acl_is_trivial"

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
	localVarPostBody = r.body
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

type ApiFilesystemChownPostRequest struct {
	ctx _context.Context
	ApiService *FilesystemApiService
	filesystemChown0 *FilesystemChown0
}

func (r ApiFilesystemChownPostRequest) FilesystemChown0(filesystemChown0 FilesystemChown0) ApiFilesystemChownPostRequest {
	r.filesystemChown0 = &filesystemChown0
	return r
}

func (r ApiFilesystemChownPostRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.FilesystemChownPostExecute(r)
}

/*
 * FilesystemChownPost Method for FilesystemChownPost
 * Change owner or group of file at `path`.

`uid` and `gid` specify new owner of the file. If either
key is absent or None, then existing value on the file is not
changed.

`recursive` performs action recursively, but does
not traverse filesystem mount points.

If `traverse` and `recursive` are specified, then the chown
operation will traverse filesystem mount points.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiFilesystemChownPostRequest
 */
func (a *FilesystemApiService) FilesystemChownPost(ctx _context.Context) ApiFilesystemChownPostRequest {
	return ApiFilesystemChownPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 */
func (a *FilesystemApiService) FilesystemChownPostExecute(r ApiFilesystemChownPostRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FilesystemApiService.FilesystemChownPost")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/filesystem/chown"

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
	localVarPostBody = r.filesystemChown0
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

type ApiFilesystemDefaultAclChoicesGetRequest struct {
	ctx _context.Context
	ApiService *FilesystemApiService
}


func (r ApiFilesystemDefaultAclChoicesGetRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.FilesystemDefaultAclChoicesGetExecute(r)
}

/*
 * FilesystemDefaultAclChoicesGet Method for FilesystemDefaultAclChoicesGet
 * Get list of default ACL types.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiFilesystemDefaultAclChoicesGetRequest
 */
func (a *FilesystemApiService) FilesystemDefaultAclChoicesGet(ctx _context.Context) ApiFilesystemDefaultAclChoicesGetRequest {
	return ApiFilesystemDefaultAclChoicesGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 */
func (a *FilesystemApiService) FilesystemDefaultAclChoicesGetExecute(r ApiFilesystemDefaultAclChoicesGetRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FilesystemApiService.FilesystemDefaultAclChoicesGet")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/filesystem/default_acl_choices"

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

type ApiFilesystemGetDefaultAclPostRequest struct {
	ctx _context.Context
	ApiService *FilesystemApiService
	filesystemGetDefaultAcl *FilesystemGetDefaultAcl
}

func (r ApiFilesystemGetDefaultAclPostRequest) FilesystemGetDefaultAcl(filesystemGetDefaultAcl FilesystemGetDefaultAcl) ApiFilesystemGetDefaultAclPostRequest {
	r.filesystemGetDefaultAcl = &filesystemGetDefaultAcl
	return r
}

func (r ApiFilesystemGetDefaultAclPostRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.FilesystemGetDefaultAclPostExecute(r)
}

/*
 * FilesystemGetDefaultAclPost Method for FilesystemGetDefaultAclPost
 * Returns a default ACL depending on the usage specified by `acl_type`.
If an admin group is defined, then an entry granting it full control will
be placed at the top of the ACL. Optionally may pass `share_type` to argument
to get share-specific template ACL.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiFilesystemGetDefaultAclPostRequest
 */
func (a *FilesystemApiService) FilesystemGetDefaultAclPost(ctx _context.Context) ApiFilesystemGetDefaultAclPostRequest {
	return ApiFilesystemGetDefaultAclPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 */
func (a *FilesystemApiService) FilesystemGetDefaultAclPostExecute(r ApiFilesystemGetDefaultAclPostRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FilesystemApiService.FilesystemGetDefaultAclPost")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/filesystem/get_default_acl"

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
	localVarPostBody = r.filesystemGetDefaultAcl
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

type ApiFilesystemGetaclPostRequest struct {
	ctx _context.Context
	ApiService *FilesystemApiService
	filesystemGetacl *FilesystemGetacl
}

func (r ApiFilesystemGetaclPostRequest) FilesystemGetacl(filesystemGetacl FilesystemGetacl) ApiFilesystemGetaclPostRequest {
	r.filesystemGetacl = &filesystemGetacl
	return r
}

func (r ApiFilesystemGetaclPostRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.FilesystemGetaclPostExecute(r)
}

/*
 * FilesystemGetaclPost Method for FilesystemGetaclPost
 * Return ACL of a given path. This may return a POSIX1e ACL or a NFSv4 ACL. The acl type is indicated
by the `ACLType` key.

Errata about ACLType NFSv4:

`simplified` returns a shortened form of the ACL permset and flags.

`TRAVERSE` sufficient rights to traverse a directory, but not read contents.

`READ` sufficient rights to traverse a directory, and read file contents.

`MODIFIY` sufficient rights to traverse, read, write, and modify a file. Equivalent to modify_set.

`FULL_CONTROL` all permissions.

If the permisssions do not fit within one of the pre-defined simplified permissions types, then
the full ACL entry will be returned.

In all cases we replace USER_OBJ, GROUP_OBJ, and EVERYONE with owner@, group@, everyone@ for
consistency with getfacl and setfacl. If one of aforementioned special tags is used, 'id' must
be set to None.

An inheriting empty everyone@ ACE is appended to non-trivial ACLs in order to enforce Windows
expectations regarding permissions inheritance. This entry is removed from NT ACL returned
to SMB clients when 'ixnas' samba VFS module is enabled. We also remove it here to avoid confusion.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiFilesystemGetaclPostRequest
 */
func (a *FilesystemApiService) FilesystemGetaclPost(ctx _context.Context) ApiFilesystemGetaclPostRequest {
	return ApiFilesystemGetaclPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 */
func (a *FilesystemApiService) FilesystemGetaclPostExecute(r ApiFilesystemGetaclPostRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FilesystemApiService.FilesystemGetaclPost")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/filesystem/getacl"

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
	localVarPostBody = r.filesystemGetacl
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

type ApiFilesystemListdirPostRequest struct {
	ctx _context.Context
	ApiService *FilesystemApiService
	filesystemListdir *FilesystemListdir
}

func (r ApiFilesystemListdirPostRequest) FilesystemListdir(filesystemListdir FilesystemListdir) ApiFilesystemListdirPostRequest {
	r.filesystemListdir = &filesystemListdir
	return r
}

func (r ApiFilesystemListdirPostRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.FilesystemListdirPostExecute(r)
}

/*
 * FilesystemListdirPost Method for FilesystemListdirPost
 * Get the contents of a directory.

Each entry of the list consists of:
  name(str): name of the file
  path(str): absolute path of the entry
  realpath(str): absolute real path of the entry (if SYMLINK)
  type(str): DIRECTORY | FILESYSTEM | SYMLINK | OTHER
  size(int): size of the entry
  mode(int): file mode/permission
  uid(int): user id of entry owner
  gid(int): group id of entry onwer
  acl(bool): extended ACL is present on file
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiFilesystemListdirPostRequest
 */
func (a *FilesystemApiService) FilesystemListdirPost(ctx _context.Context) ApiFilesystemListdirPostRequest {
	return ApiFilesystemListdirPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 */
func (a *FilesystemApiService) FilesystemListdirPostExecute(r ApiFilesystemListdirPostRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FilesystemApiService.FilesystemListdirPost")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/filesystem/listdir"

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
	localVarPostBody = r.filesystemListdir
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

type ApiFilesystemSetaclPostRequest struct {
	ctx _context.Context
	ApiService *FilesystemApiService
	filesystemSetacl0 *FilesystemSetacl0
}

func (r ApiFilesystemSetaclPostRequest) FilesystemSetacl0(filesystemSetacl0 FilesystemSetacl0) ApiFilesystemSetaclPostRequest {
	r.filesystemSetacl0 = &filesystemSetacl0
	return r
}

func (r ApiFilesystemSetaclPostRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.FilesystemSetaclPostExecute(r)
}

/*
 * FilesystemSetaclPost Method for FilesystemSetaclPost
 * Set ACL of a given path. Takes the following parameters:
`path` full path to directory or file.

`dacl` "simplified" ACL here or a full ACL.

`uid` the desired UID of the file user. If set to None (the default), then user is not changed.

`gid` the desired GID of the file group. If set to None (the default), then group is not changed.

`recursive` apply the ACL recursively

`traverse` traverse filestem boundaries (ZFS datasets)

`strip` convert ACL to trivial. ACL is trivial if it can be expressed as a file mode without
losing any access rules.

`canonicalize` reorder ACL entries so that they are in concanical form as described
in the Microsoft documentation MS-DTYP 2.4.5 (ACL)

In all cases we replace USER_OBJ, GROUP_OBJ, and EVERYONE with owner@, group@, everyone@ for
consistency with getfacl and setfacl. If one of aforementioned special tags is used, 'id' must
be set to None.

An inheriting empty everyone@ ACE is appended to non-trivial ACLs in order to enforce Windows
expectations regarding permissions inheritance. This entry is removed from NT ACL returned
to SMB clients when 'ixnas' samba VFS module is enabled.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiFilesystemSetaclPostRequest
 */
func (a *FilesystemApiService) FilesystemSetaclPost(ctx _context.Context) ApiFilesystemSetaclPostRequest {
	return ApiFilesystemSetaclPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 */
func (a *FilesystemApiService) FilesystemSetaclPostExecute(r ApiFilesystemSetaclPostRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FilesystemApiService.FilesystemSetaclPost")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/filesystem/setacl"

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
	localVarPostBody = r.filesystemSetacl0
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

type ApiFilesystemSetpermPostRequest struct {
	ctx _context.Context
	ApiService *FilesystemApiService
	filesystemSetperm0 *FilesystemSetperm0
}

func (r ApiFilesystemSetpermPostRequest) FilesystemSetperm0(filesystemSetperm0 FilesystemSetperm0) ApiFilesystemSetpermPostRequest {
	r.filesystemSetperm0 = &filesystemSetperm0
	return r
}

func (r ApiFilesystemSetpermPostRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.FilesystemSetpermPostExecute(r)
}

/*
 * FilesystemSetpermPost Method for FilesystemSetpermPost
 * Remove extended ACL from specified path.

If `mode` is specified then the mode will be applied to the
path and files and subdirectories depending on which `options` are
selected. Mode should be formatted as string representation of octal
permissions bits.

`uid` the desired UID of the file user. If set to None (the default), then user is not changed.

`gid` the desired GID of the file group. If set to None (the default), then group is not changed.

`stripacl` setperm will fail if an extended ACL is present on `path`,
unless `stripacl` is set to True.

`recursive` remove ACLs recursively, but do not traverse dataset
boundaries.

`traverse` remove ACLs from child datasets.

If no `mode` is set, and `stripacl` is True, then non-trivial ACLs
will be converted to trivial ACLs. An ACL is trivial if it can be
expressed as a file mode without losing any access rules.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiFilesystemSetpermPostRequest
 */
func (a *FilesystemApiService) FilesystemSetpermPost(ctx _context.Context) ApiFilesystemSetpermPostRequest {
	return ApiFilesystemSetpermPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 */
func (a *FilesystemApiService) FilesystemSetpermPostExecute(r ApiFilesystemSetpermPostRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FilesystemApiService.FilesystemSetpermPost")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/filesystem/setperm"

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
	localVarPostBody = r.filesystemSetperm0
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

type ApiFilesystemStatPostRequest struct {
	ctx _context.Context
	ApiService *FilesystemApiService
	body *string
}

func (r ApiFilesystemStatPostRequest) Body(body string) ApiFilesystemStatPostRequest {
	r.body = &body
	return r
}

func (r ApiFilesystemStatPostRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.FilesystemStatPostExecute(r)
}

/*
 * FilesystemStatPost Method for FilesystemStatPost
 * Return the filesystem stat(2) for a given `path`.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiFilesystemStatPostRequest
 */
func (a *FilesystemApiService) FilesystemStatPost(ctx _context.Context) ApiFilesystemStatPostRequest {
	return ApiFilesystemStatPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 */
func (a *FilesystemApiService) FilesystemStatPostExecute(r ApiFilesystemStatPostRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FilesystemApiService.FilesystemStatPost")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/filesystem/stat"

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
	localVarPostBody = r.body
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

type ApiFilesystemStatfsPostRequest struct {
	ctx _context.Context
	ApiService *FilesystemApiService
	body *string
}

func (r ApiFilesystemStatfsPostRequest) Body(body string) ApiFilesystemStatfsPostRequest {
	r.body = &body
	return r
}

func (r ApiFilesystemStatfsPostRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.FilesystemStatfsPostExecute(r)
}

/*
 * FilesystemStatfsPost Method for FilesystemStatfsPost
 * Return stats from the filesystem of a given path.

Raises:
    CallError(ENOENT) - Path not found
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiFilesystemStatfsPostRequest
 */
func (a *FilesystemApiService) FilesystemStatfsPost(ctx _context.Context) ApiFilesystemStatfsPostRequest {
	return ApiFilesystemStatfsPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 */
func (a *FilesystemApiService) FilesystemStatfsPostExecute(r ApiFilesystemStatfsPostRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FilesystemApiService.FilesystemStatfsPost")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/filesystem/statfs"

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
	localVarPostBody = r.body
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