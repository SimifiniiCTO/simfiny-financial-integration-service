package grpc

import (
	"context"

	"github.com/pborman/uuid"
	"go.temporal.io/sdk/client"
)

// ExecuteWorkflow executes a workflow and returns the workflow run object via temporal
func (s *Server) ExecuteWorkflow(ctx context.Context, workflow interface{}, args ...interface{}) (client.WorkflowRun, error) {
	var (
		temporalClient = s.TransactionManager.Client
		options        = client.StartWorkflowOptions{
			ID:                       uuid.New(),
			TaskQueue:                string(s.TransactionManager.ServiceTaskQueue),
			WorkflowExecutionTimeout: s.config.WorkflowExecutionTimeout,
			WorkflowRunTimeout:       s.config.WorkflowRunTimeout,
			WorkflowTaskTimeout:      s.config.WorkflowTaskTimeout,
		}
	)

	wf, err := temporalClient.ExecuteWorkflow(ctx, options, workflow, args...)
	if err != nil {
		return nil, err
	}

	return wf, nil
}
