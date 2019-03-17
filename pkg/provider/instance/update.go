package instance

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/rds"
	log "github.com/sirupsen/logrus"
)

type UpdateDBInstanceRequest struct {
	id                 string
	clusterId          string
	allocatedStorage   int
	engine             string
	class              string
	parameterGroupName string
	publiclyAccessible bool
}

func (req *UpdateDBInstanceRequest) SetId(v string) *UpdateDBInstanceRequest {
	req.id = v
	return req
}

func (req *UpdateDBInstanceRequest) SetClusterId(v string) *UpdateDBInstanceRequest {
	req.clusterId = v
	return req
}

func (req *UpdateDBInstanceRequest) SetAllocatedStorage(v int) *UpdateDBInstanceRequest {
	req.allocatedStorage = v
	return req
}

func (req *UpdateDBInstanceRequest) SetEngine(v string) *UpdateDBInstanceRequest {
	req.engine = v
	return req
}

func (req *UpdateDBInstanceRequest) SetClass(v string) *UpdateDBInstanceRequest {
	req.class = v
	return req
}

func (req *UpdateDBInstanceRequest) SetParameterGroupName(v string) *UpdateDBInstanceRequest {
	req.parameterGroupName = v
	return req
}

func (req *UpdateDBInstanceRequest) SetPubliclyAccessible(v bool) *UpdateDBInstanceRequest {
	req.publiclyAccessible = v
	return req
}

func UpdateDBClusterInstance(svc *rds.RDS, req UpdateDBInstanceRequest) error {
	input := &rds.ModifyDBInstanceInput{
		ApplyImmediately: aws.Bool(true),
		//BackupRetentionPeriod:      aws.Int64(1),
		DBInstanceClass:      aws.String(req.class),
		DBInstanceIdentifier: aws.String(req.id),
		DBParameterGroupName: aws.String(req.parameterGroupName),
		//MasterUserPassword:         aws.String("mynewpassword"),
		//PreferredBackupWindow:      aws.String("04:00-04:30"),
		//PreferredMaintenanceWindow: aws.String("Tue:05:00-Tue:05:30"),
	}

	if req.clusterId == "" {
		input.AllocatedStorage = aws.Int64(int64(req.allocatedStorage))
	}

	result, err := svc.ModifyDBInstance(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case rds.ErrCodeInvalidDBInstanceStateFault:
				log.Warn(rds.ErrCodeInvalidDBInstanceStateFault, aerr.Error())
				return aerr
			case rds.ErrCodeInvalidDBSecurityGroupStateFault:
				log.Warn(rds.ErrCodeInvalidDBSecurityGroupStateFault, aerr.Error())
				return aerr
			case rds.ErrCodeDBInstanceAlreadyExistsFault:
				log.Warn(rds.ErrCodeDBInstanceAlreadyExistsFault, aerr.Error())
				return aerr
			case rds.ErrCodeDBInstanceNotFoundFault:
				log.Warn(rds.ErrCodeDBInstanceNotFoundFault, aerr.Error())
				return aerr
			case rds.ErrCodeDBSecurityGroupNotFoundFault:
				log.Warn(rds.ErrCodeDBSecurityGroupNotFoundFault, aerr.Error())
				return aerr
			case rds.ErrCodeDBParameterGroupNotFoundFault:
				log.Warn(rds.ErrCodeDBParameterGroupNotFoundFault, aerr.Error())
				return aerr
			case rds.ErrCodeInsufficientDBInstanceCapacityFault:
				log.Warn(rds.ErrCodeInsufficientDBInstanceCapacityFault, aerr.Error())
				return aerr
			case rds.ErrCodeStorageQuotaExceededFault:
				log.Warn(rds.ErrCodeStorageQuotaExceededFault, aerr.Error())
				return aerr
			case rds.ErrCodeInvalidVPCNetworkStateFault:
				log.Warn(rds.ErrCodeInvalidVPCNetworkStateFault, aerr.Error())
				return aerr
			case rds.ErrCodeProvisionedIopsNotAvailableInAZFault:
				log.Warn(rds.ErrCodeProvisionedIopsNotAvailableInAZFault, aerr.Error())
				return aerr
			case rds.ErrCodeOptionGroupNotFoundFault:
				log.Warn(rds.ErrCodeOptionGroupNotFoundFault, aerr.Error())
				return aerr
			case rds.ErrCodeDBUpgradeDependencyFailureFault:
				log.Warn(rds.ErrCodeDBUpgradeDependencyFailureFault, aerr.Error())
				return aerr
			case rds.ErrCodeStorageTypeNotSupportedFault:
				log.Warn(rds.ErrCodeStorageTypeNotSupportedFault, aerr.Error())
				return aerr
			case rds.ErrCodeAuthorizationNotFoundFault:
				log.Warn(rds.ErrCodeAuthorizationNotFoundFault, aerr.Error())
				return aerr
			case rds.ErrCodeCertificateNotFoundFault:
				log.Warn(rds.ErrCodeCertificateNotFoundFault, aerr.Error())
				return aerr
			case rds.ErrCodeDomainNotFoundFault:
				log.Warn(rds.ErrCodeDomainNotFoundFault, aerr.Error())
				return aerr
			case rds.ErrCodeBackupPolicyNotFoundFault:
				log.Warn(rds.ErrCodeBackupPolicyNotFoundFault, aerr.Error())
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
