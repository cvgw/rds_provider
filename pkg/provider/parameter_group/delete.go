package parameter_group

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/rds"
	log "github.com/sirupsen/logrus"
)

func DeleteDBParameterGroup(svc *rds.RDS, groupName string) error {
	input := &rds.DeleteDBParameterGroupInput{
		DBParameterGroupName: aws.String(groupName),
	}

	result, err := svc.DeleteDBParameterGroup(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case rds.ErrCodeInvalidDBParameterGroupStateFault:
				log.Info(rds.ErrCodeInvalidDBParameterGroupStateFault, aerr.Error())
				return err
			case rds.ErrCodeDBParameterGroupNotFoundFault:
				log.Debug(rds.ErrCodeDBParameterGroupNotFoundFault, aerr.Error())
				return NotFoundErr
			default:
				log.Warn(aerr.Error())
				return err
			}
		} else {
			log.Warn(err.Error())
			return err
		}
	}
	log.Debug(result)

	return nil
}
