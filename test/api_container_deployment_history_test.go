/*
Qovery API

Testing ContainerDeploymentHistoryAPIService

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

func Test_qovery_ContainerDeploymentHistoryAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ContainerDeploymentHistoryAPIService ListContainerDeploymentHistory", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var containerId string

		resp, httpRes, err := apiClient.ContainerDeploymentHistoryAPI.ListContainerDeploymentHistory(context.Background(), containerId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
