package parameter_group

import (
	"errors"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/rds"
	log "github.com/sirupsen/logrus"
)

var (
	NotFoundErr error
)

func init() {
	NotFoundErr = errors.New("db parameter group not found")
}

func FindDBParameterGroup(svc *rds.RDS, paramGroupName string) (*rds.DBParameterGroup, error) {
	input := &rds.DescribeDBParameterGroupsInput{
		DBParameterGroupName: aws.String(paramGroupName),
	}

	result, err := svc.DescribeDBParameterGroups(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case rds.ErrCodeDBParameterGroupNotFoundFault:
				log.Info(rds.ErrCodeDBParameterGroupNotFoundFault, aerr.Error())
				return nil, NotFoundErr
			default:
				log.Warn(aerr.Error())
				return nil, aerr
			}
		} else {
			log.Warn(err.Error())
			return nil, err
		}
	}
	return result.DBParameterGroups[0], nil
}
