package cryptconverter

import (
	"context"
	"time"

	pb "github.com/temporalio/samples-go/encrypted-payloads/helloworld"
	"go.temporal.io/sdk/activity"
	"go.temporal.io/sdk/workflow"
)

// Workflow is a standard workflow definition.
// Note that the Workflow and Activity don't need to care that
// their inputs/results are being encrypted/decrypted.
func Workflow(ctx workflow.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	ao := workflow.ActivityOptions{
		ScheduleToStartTimeout: time.Minute,
		StartToCloseTimeout:    time.Minute,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	logger := workflow.GetLogger(ctx)
	logger.Info("Encrypted Payloads workflow started", "name", in.Name)

	var result *pb.HelloReply
	err := workflow.ExecuteActivity(ctx, Activity, in).Get(ctx, &result)
	if err != nil {
		logger.Error("Activity failed.", "Error", err)
		return nil, err
	}

	logger.Info("Encrypted Payloads workflow completed.", "result", result)

	return result, nil
}

func Activity(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	logger := activity.GetLogger(ctx)
	logger.Info("Activity", "name", in.Name)
	return &pb.HelloReply{
		Message: "Hello " + in.Name + "!",
	}, nil

}
