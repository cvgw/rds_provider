package subnet_group

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/rds"
	log "github.com/sirupsen/logrus"
)

type CreateSubnetGroupRequest struct {
	Name        string   `json:"name,omitempty"`
	Description string   `json:"description,omitempty"`
	SubnetIds   []string `json:"subnet_ids,omitempty"`
}

func CreateSubnetGroup(svc *rds.RDS, req CreateSubnetGroupRequest) (*rds.DBSubnetGroup, error) {
	groupInput := NewCreateDBSubnetGroupInput(req)

	groupOutput, err := svc.CreateDBSubnetGroup(groupInput)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case rds.ErrCodeDBSubnetGroupAlreadyExistsFault:
				log.Warn(rds.ErrCodeDBSubnetGroupAlreadyExistsFault, aerr.Error())
				return nil, aerr
			case rds.ErrCodeDBSubnetGroupQuotaExceededFault:
				log.Warn(rds.ErrCodeDBSubnetGroupQuotaExceededFault, aerr.Error())
				return nil, aerr
			case rds.ErrCodeDBSubnetQuotaExceededFault:
				log.Warn(rds.ErrCodeDBSubnetQuotaExceededFault, aerr.Error())
				return nil, aerr
			case rds.ErrCodeDBSubnetGroupDoesNotCoverEnoughAZs:
				log.Warn(rds.ErrCodeDBSubnetGroupDoesNotCoverEnoughAZs, aerr.Error())
				return nil, aerr
			case rds.ErrCodeInvalidSubnet:
				log.Warn(rds.ErrCodeInvalidSubnet, aerr.Error())
				return nil, aerr
			default:
				log.Warn(aerr)
				return nil, aerr
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			log.Warn(err)
			return nil, aerr
		}
	}

	return groupOutput.DBSubnetGroup, nil
}

func NewCreateDBSubnetGroupInput(req CreateSubnetGroupRequest) *rds.CreateDBSubnetGroupInput {
	sIds := make([]*string, 0)
	for _, i := range req.SubnetIds {
		sIds = append(sIds, aws.String(i))
	}

	groupInput := &rds.CreateDBSubnetGroupInput{
		DBSubnetGroupName:        aws.String(req.Name),
		DBSubnetGroupDescription: aws.String(req.Description),
		SubnetIds:                sIds,
	}

	return groupInput
}
