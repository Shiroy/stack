/*
Auth API

Testing ScopesApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package authclient

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/formancehq/auth/authclient"
)

func Test_authclient_ScopesApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ScopesApiService AddTransientScope", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var scopeId interface{}
		var transientScopeId interface{}

		httpRes, err := apiClient.ScopesApi.AddTransientScope(context.Background(), scopeId, transientScopeId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ScopesApiService CreateScope", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ScopesApi.CreateScope(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ScopesApiService DeleteScope", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var scopeId interface{}

		httpRes, err := apiClient.ScopesApi.DeleteScope(context.Background(), scopeId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ScopesApiService DeleteTransientScope", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var scopeId interface{}
		var transientScopeId interface{}

		httpRes, err := apiClient.ScopesApi.DeleteTransientScope(context.Background(), scopeId, transientScopeId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ScopesApiService ListScopes", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ScopesApi.ListScopes(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ScopesApiService ReadScope", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var scopeId interface{}

		resp, httpRes, err := apiClient.ScopesApi.ReadScope(context.Background(), scopeId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ScopesApiService UpdateScope", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var scopeId interface{}

		resp, httpRes, err := apiClient.ScopesApi.UpdateScope(context.Background(), scopeId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
