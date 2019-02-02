package parameter_group

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/rds"
	log "github.com/sirupsen/logrus"
)

type CreateRequest struct {
	Family      string
	Name        string
	Description string
}

func (r *CreateRequest) SetFamily(v string) *CreateRequest {
	r.Family = v
	return r
}

func (r *CreateRequest) SetName(v string) *CreateRequest {
	r.Name = v
	return r
}

func (r *CreateRequest) SetDescription(v string) *CreateRequest {
	r.Description = v
	return r
}

func CreateDBParameterGroup(svc *rds.RDS, req CreateRequest) (
	*rds.DBParameterGroup, error,
) {
	input := NewCreateDBParameterGroupInput(req)

	result, err := svc.CreateDBParameterGroup(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case rds.ErrCodeDBParameterGroupQuotaExceededFault:
				log.Warn(rds.ErrCodeDBParameterGroupQuotaExceededFault, aerr.Error())
				return nil, aerr
			case rds.ErrCodeDBParameterGroupAlreadyExistsFault:
				log.Warn(rds.ErrCodeDBParameterGroupAlreadyExistsFault, aerr.Error())
				return nil, aerr
			default:
				log.Warn(aerr.Error())
				return nil, aerr
			}
		} else {
			log.Warn(err.Error())
			return nil, err
		}
	}

	return result.DBParameterGroup, nil
}

func NewCreateDBParameterGroupInput(req CreateRequest) *rds.CreateDBParameterGroupInput {
	input := &rds.CreateDBParameterGroupInput{
		DBParameterGroupFamily: aws.String(req.Family),
		DBParameterGroupName:   aws.String(req.Name),
		Description:            aws.String(req.Description),
	}

	return input
}
