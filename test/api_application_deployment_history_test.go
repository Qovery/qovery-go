/*
Qovery API

Testing ApplicationDeploymentHistoryAPIService

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

func Test_qovery_ApplicationDeploymentHistoryAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ApplicationDeploymentHistoryAPIService ListApplicationDeploymentHistory", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var applicationId string

		resp, httpRes, err := apiClient.ApplicationDeploymentHistoryAPI.ListApplicationDeploymentHistory(context.Background(), applicationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
