package parameter_group

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/rds"
	log "github.com/sirupsen/logrus"
)

const (
	Immediate applyMethod = "immediate"
	String    valType     = "string"
)

type applyMethod string
type valType string

type Param struct {
	Apply     applyMethod
	Name      string
	Value     interface{}
	ValueType valType
}

type UpdateRequest struct {
	name       string
	parameters []*rds.Parameter
}

func (r *UpdateRequest) SetName(v string) *UpdateRequest {
	r.name = v
	return r
}

func (r *UpdateRequest) SetParameters(params []Param) *UpdateRequest {
	awsParams := make([]*rds.Parameter, 0)
	for _, p := range params {
		switch p.ValueType {
		case String:
			awsParams = append(awsParams, &rds.Parameter{
				ApplyMethod:    aws.String(string(p.Apply)),
				ParameterName:  aws.String(p.Name),
				ParameterValue: aws.String(p.Value.(string)),
			})
		}
	}

	r.parameters = awsParams
	return r
}

func UpdateDBParameterGroup(svc *rds.RDS, req UpdateRequest) error {
	input := NewModifyDBParameterGroupInput(req)

	result, err := svc.ModifyDBParameterGroup(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case rds.ErrCodeDBParameterGroupNotFoundFault:
				log.Warn(rds.ErrCodeDBParameterGroupNotFoundFault, aerr.Error())
				return aerr
			case rds.ErrCodeInvalidDBParameterGroupStateFault:
				log.Warn(rds.ErrCodeInvalidDBParameterGroupStateFault, aerr.Error())
				return aerr
			default:
				log.Warn(aerr.Error())
				return aerr
			}
		} else {
			log.Warn(err.Error())
			return err
		}
	}
	log.Debug(result)

	return nil
}

func NewModifyDBParameterGroupInput(req UpdateRequest) *rds.ModifyDBParameterGroupInput {
	input := &rds.ModifyDBParameterGroupInput{
		DBParameterGroupName: aws.String(req.name),
		Parameters:           req.parameters,
	}

	return input
}
