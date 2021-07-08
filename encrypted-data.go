package cryptconverter

import (
	"context"
	"time"

	"go.temporal.io/sdk/activity"
	"go.temporal.io/sdk/workflow"
)

// Workflow is a standard workflow definition.
// Note that the Workflow and Activity don't need to care that
// their inputs/results are being encrypted/decrypted.
func Workflow(ctx workflow.Context, in *SensitivePayload) (*SensitivePayload, error) {
	ao := workflow.ActivityOptions{
		ScheduleToStartTimeout: time.Minute,
		StartToCloseTimeout:    time.Minute,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	logger := workflow.GetLogger(ctx)
	logger.Info("Encrypted Payloads workflow started", "name", in.Secret)

	var result *SensitivePayload
	err := workflow.ExecuteActivity(ctx, Activity, in).Get(ctx, &result)
	if err != nil {
		logger.Error("Activity failed.", "Error", err)
		return nil, err
	}

	logger.Info("Encrypted Payloads workflow completed.", "result", result)

	return result, nil
}

func Activity(ctx context.Context, in *SensitivePayload) (*SensitivePayload, error) {
	logger := activity.GetLogger(ctx)
	logger.Info("Activity", "name", in.Secret)
	return &SensitivePayload{
		Secret: "Hello " + in.Secret + "!",
	}, nil

}
