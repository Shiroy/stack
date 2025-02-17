# \ScopesApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddTransientScope**](ScopesApi.md#AddTransientScope) | **Put** /scopes/{scopeId}/transient/{transientScopeId} | Add a transient scope to a scope
[**CreateScope**](ScopesApi.md#CreateScope) | **Post** /scopes | Create scope
[**DeleteScope**](ScopesApi.md#DeleteScope) | **Delete** /scopes/{scopeId} | Delete scope
[**DeleteTransientScope**](ScopesApi.md#DeleteTransientScope) | **Delete** /scopes/{scopeId}/transient/{transientScopeId} | Delete a transient scope from a scope
[**ListScopes**](ScopesApi.md#ListScopes) | **Get** /scopes | List scopes
[**ReadScope**](ScopesApi.md#ReadScope) | **Get** /scopes/{scopeId} | Read scope
[**UpdateScope**](ScopesApi.md#UpdateScope) | **Put** /scopes/{scopeId} | Update scope



## AddTransientScope

> AddTransientScope(ctx, scopeId, transientScopeId).Execute()

Add a transient scope to a scope



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/formancehq/auth/authclient"
)

func main() {
    scopeId := TODO // interface{} | Scope ID
    transientScopeId := TODO // interface{} | Transient scope ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ScopesApi.AddTransientScope(context.Background(), scopeId, transientScopeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScopesApi.AddTransientScope``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scopeId** | [**interface{}**](.md) | Scope ID | 
**transientScopeId** | [**interface{}**](.md) | Transient scope ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddTransientScopeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateScope

> CreateScopeResponse CreateScope(ctx).Body(body).Execute()

Create scope



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/formancehq/auth/authclient"
)

func main() {
    body := ScopeOptions(987) // ScopeOptions |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ScopesApi.CreateScope(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScopesApi.CreateScope``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateScope`: CreateScopeResponse
    fmt.Fprintf(os.Stdout, "Response from `ScopesApi.CreateScope`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateScopeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **ScopeOptions** |  | 

### Return type

[**CreateScopeResponse**](CreateScopeResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteScope

> DeleteScope(ctx, scopeId).Execute()

Delete scope



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/formancehq/auth/authclient"
)

func main() {
    scopeId := TODO // interface{} | Scope ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ScopesApi.DeleteScope(context.Background(), scopeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScopesApi.DeleteScope``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scopeId** | [**interface{}**](.md) | Scope ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteScopeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTransientScope

> DeleteTransientScope(ctx, scopeId, transientScopeId).Execute()

Delete a transient scope from a scope



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/formancehq/auth/authclient"
)

func main() {
    scopeId := TODO // interface{} | Scope ID
    transientScopeId := TODO // interface{} | Transient scope ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ScopesApi.DeleteTransientScope(context.Background(), scopeId, transientScopeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScopesApi.DeleteTransientScope``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scopeId** | [**interface{}**](.md) | Scope ID | 
**transientScopeId** | [**interface{}**](.md) | Transient scope ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTransientScopeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListScopes

> ListScopesResponse ListScopes(ctx).Execute()

List scopes



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/formancehq/auth/authclient"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ScopesApi.ListScopes(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScopesApi.ListScopes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListScopes`: ListScopesResponse
    fmt.Fprintf(os.Stdout, "Response from `ScopesApi.ListScopes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListScopesRequest struct via the builder pattern


### Return type

[**ListScopesResponse**](ListScopesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadScope

> CreateScopeResponse ReadScope(ctx, scopeId).Execute()

Read scope



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/formancehq/auth/authclient"
)

func main() {
    scopeId := TODO // interface{} | Scope ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ScopesApi.ReadScope(context.Background(), scopeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScopesApi.ReadScope``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadScope`: CreateScopeResponse
    fmt.Fprintf(os.Stdout, "Response from `ScopesApi.ReadScope`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scopeId** | [**interface{}**](.md) | Scope ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadScopeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreateScopeResponse**](CreateScopeResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateScope

> CreateScopeResponse UpdateScope(ctx, scopeId).Body(body).Execute()

Update scope



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/formancehq/auth/authclient"
)

func main() {
    scopeId := TODO // interface{} | Scope ID
    body := ScopeOptions(987) // ScopeOptions |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ScopesApi.UpdateScope(context.Background(), scopeId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScopesApi.UpdateScope``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateScope`: CreateScopeResponse
    fmt.Fprintf(os.Stdout, "Response from `ScopesApi.UpdateScope`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scopeId** | [**interface{}**](.md) | Scope ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateScopeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **ScopeOptions** |  | 

### Return type

[**CreateScopeResponse**](CreateScopeResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

