/*
Qovery API

Testing DefaultAPIService

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

func Test_qovery_DefaultAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test DefaultAPIService GetContainerRegistryAssociatedServices", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var organizationId string
		var containerRegistryId string

		resp, httpRes, err := apiClient.DefaultAPI.GetContainerRegistryAssociatedServices(context.Background(), organizationId, containerRegistryId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultAPIService GetHelmRepositoryAssociatedServices", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var organizationId string
		var helmRepositoryId string

		resp, httpRes, err := apiClient.DefaultAPI.GetHelmRepositoryAssociatedServices(context.Background(), organizationId, helmRepositoryId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
