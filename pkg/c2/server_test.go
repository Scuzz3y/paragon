package c2_test

import (
	"context"
	"testing"

	"github.com/kcarretto/paragon/ent"
	"github.com/kcarretto/paragon/graphql/models"
	"github.com/kcarretto/paragon/pkg/c2"
	"github.com/kcarretto/paragon/pkg/c2/mocks"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

var sampleUUID = "ABCD"
var sampleHostname = "Box"
var sampleSessionID = "sesh"
var samplePrimaryIP = "127.257.12.34"
var samplePrimaryMAC = "mac-daddy"

// HandleAgentTestCase holds parameters used to test the server's HandleAgent method.
type HandleAgentTestCase struct {
	ReceivedMessage c2.AgentMessage
	ClaimedTasks    []*ent.Task
	ClaimTaskErr    error
	ExpectedParams  models.ClaimTasksRequest
	ExpectedErr     error
}

// Run the test case.
func (test HandleAgentTestCase) Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	teamserver := mocks.NewMockTeamserver(ctrl)
	teamserver.EXPECT().ClaimTasks(gomock.Any(), test.ExpectedParams).Return(test.ClaimedTasks, test.ClaimTaskErr)

	srv := c2.Server{
		teamserver,
	}

	resp, err := srv.HandleAgent(context.Background(), test.ReceivedMessage)
	require.Equal(t, test.ExpectedErr, err)
	require.Equal(t, len(test.ClaimedTasks), len(resp.Tasks))
	for i, task := range test.ClaimedTasks {
		require.Equal(t, int64(task.ID), resp.Tasks[i].Id)
		require.Equal(t, task.Content, resp.Tasks[i].Content)
	}
}

// TestHandleAgent method with various test cases.
func TestHandleAgent(t *testing.T) {
	metadata := &c2.AgentMetadata{
		MachineUUID: sampleUUID,
		SessionID:   sampleSessionID,
		Hostname:    sampleHostname,
		PrimaryIP:   samplePrimaryIP,
		PrimaryMAC:  samplePrimaryMAC,
	}
	testCases := []HandleAgentTestCase{
		{
			ReceivedMessage: c2.AgentMessage{
				Metadata: metadata,
			},
			ClaimedTasks: []*ent.Task{
				&ent.Task{
					ID:      123,
					Content: "print('testing')",
				},
			},
			ExpectedParams: models.ClaimTasksRequest{
				MachineUUID: &metadata.MachineUUID,
				SessionID:   &metadata.SessionID,
				Hostname:    &metadata.Hostname,
				PrimaryIP:   &metadata.PrimaryIP,
				PrimaryMac:  &metadata.PrimaryMAC,
			},
		},
	}

	for _, test := range testCases {
		test.Run(t)
	}
}
