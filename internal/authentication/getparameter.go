package authentication

import (
	"context"
	"dyno/internal/logger"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
)

func getSSMParameterValue(parameterName string) (string, error) {

	cfg, err := config.LoadDefaultConfig(context.TODO())//, config.WithRegion("ap-southeast-2")
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	svc := ssm.NewFromConfig(cfg)

	resp, err := svc.GetParameter(context.TODO(), &ssm.GetParameterInput{
		Name: aws.String(parameterName),
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	return *resp.Parameter.Value, nil

}
