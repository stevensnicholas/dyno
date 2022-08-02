package authentication

import (
    "context"
    "dyno/internal/logger"
	"github.com/aws/aws-sdk-go-v2/aws"
    "github.com/aws/aws-sdk-go-v2/config"
    "github.com/aws/aws-sdk-go-v2/service/ssm"
)

func getSSMParameterValue(parameterName string) (string){

    cfg, err := config.LoadDefaultConfig(context.TODO())
    if err != nil {
		logger.Error(err.Error())
    }

    svc := ssm.NewFromConfig(cfg)

    resp, err := svc.GetParameter(context.TODO(), &ssm.GetParameterInput{
        Name: aws.String(parameterName),
    })
    if err != nil {
		logger.Error(err.Error())
    }

    return *resp.Parameter.Value

}

