/*
Qovery API

Testing JobsAPIService

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

func Test_qovery_JobsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test JobsAPIService AutoDeployJobEnvironments", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var organizationId string

		resp, httpRes, err := apiClient.JobsAPI.AutoDeployJobEnvironments(context.Background(), organizationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test JobsAPIService CloneJob", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var jobId string

		resp, httpRes, err := apiClient.JobsAPI.CloneJob(context.Background(), jobId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test JobsAPIService CreateJob", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var environmentId string

		resp, httpRes, err := apiClient.JobsAPI.CreateJob(context.Background(), environmentId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test JobsAPIService GetDefaultJobAdvancedSettings", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.JobsAPI.GetDefaultJobAdvancedSettings(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test JobsAPIService GetEnvironmentJobStatus", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var environmentId string

		resp, httpRes, err := apiClient.JobsAPI.GetEnvironmentJobStatus(context.Background(), environmentId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test JobsAPIService ListJobs", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var environmentId string

		resp, httpRes, err := apiClient.JobsAPI.ListJobs(context.Background(), environmentId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
