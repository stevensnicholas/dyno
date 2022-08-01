package authentication

import (
    "context"
    "fmt"
    "log"

    "github.com/aws/aws-sdk-go-v2/config"
    "github.com/aws/aws-sdk-go-v2/service/ssm"
)

func getParameter(name string) (string, error){

    cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("ap-southeast-2"))
    if err != nil {
        return "", err
    }

    svc := ssm.NewFromConfig(cfg)

    resp, err := svc.GetParameter(context.TODO(), &ssm.GetParameterInput{
        Name: &name,
    })
    if err != nil {
        return "", err
    }

    return resp.Parameter.Value, nil

}

