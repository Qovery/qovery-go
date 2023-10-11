/*
Qovery API

Testing DatabaseActionsAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package qovery

import (
	"context"
	openapiclient "github.com/qovery/qovery-client-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_qovery_DatabaseActionsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test DatabaseActionsAPIService DeployDatabase", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var databaseId string

		resp, httpRes, err := apiClient.DatabaseActionsAPI.DeployDatabase(context.Background(), databaseId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DatabaseActionsAPIService RebootDatabase", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var databaseId string

		resp, httpRes, err := apiClient.DatabaseActionsAPI.RebootDatabase(context.Background(), databaseId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DatabaseActionsAPIService RedeployDatabase", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var databaseId string

		resp, httpRes, err := apiClient.DatabaseActionsAPI.RedeployDatabase(context.Background(), databaseId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DatabaseActionsAPIService RestartDatabase", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var databaseId string

		resp, httpRes, err := apiClient.DatabaseActionsAPI.RestartDatabase(context.Background(), databaseId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DatabaseActionsAPIService StopDatabase", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var databaseId string

		resp, httpRes, err := apiClient.DatabaseActionsAPI.StopDatabase(context.Background(), databaseId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
