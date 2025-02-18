/*
Qovery API

Testing DeploymentQueueActionsAPIService

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

func Test_qovery_DeploymentQueueActionsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test DeploymentQueueActionsAPIService CancelDeploymentRequest", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var deploymentRequestId string

		httpRes, err := apiClient.DeploymentQueueActionsAPI.CancelDeploymentRequest(context.Background(), deploymentRequestId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
