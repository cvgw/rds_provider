package instance

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/rds"
	log "github.com/sirupsen/logrus"
)

//NewDBInstanceInput contains the data to be used when creating the RDS instance
type NewDBInstanceInput struct {
	// Whether minor version upgrades should be automatically applied to the instance (optional)
	AutoMinorVersionUpgrade bool
	// Identifier of the RDS cluster to add the instance to
	ClusterIdentifier string
	// Whether instance tags should be copied to DB snapshots (optional)
	CopyTagsToSnapshot bool
	// Engine type
	Engine string
	// Engine version (optional)
	EngineVersion string
	// Instance class (optional)
	InstanceClass string
	// Identifier for the RDS instance
	InstanceIdentifier string
	// Whether enhanced monitoring should be enabled on this instance (optional)
	EnhancedMonitoring bool
	// Interval at which to collecting monitoring data from the instance (optional)
	MonitoringInterval int64
	// The IAM Role ARN to use for collecting monitoring (optional)
	MonitoringRoleArn string
	// Name of RDS Instance ParameterGroup to use (optional)
	ParameterGroupName string
}

// CreateDBClusterInstance create a new RDS instance from the supplied NewDBInstanceInput
func CreateDBClusterInstance(svc *rds.RDS, input NewDBInstanceInput) (*rds.DBInstance, error) {
	instanceInput := NewCreateDBInstanceInput(input)
	instanceOutput, err := svc.CreateDBInstance(instanceInput)
	if err != nil {
		log.Warn(err)
		return nil, err
	}

	return instanceOutput.DBInstance, nil
}

func NewCreateDBInstanceInput(input NewDBInstanceInput) *rds.CreateDBInstanceInput {
	instanceInput := &rds.CreateDBInstanceInput{
		AutoMinorVersionUpgrade: aws.Bool(input.AutoMinorVersionUpgrade),
		CopyTagsToSnapshot:      aws.Bool(input.CopyTagsToSnapshot),
		DBInstanceIdentifier:    aws.String(input.InstanceIdentifier),
		DBClusterIdentifier:     aws.String(input.ClusterIdentifier),
		Engine:                  aws.String(input.Engine),
	}

	if input.EngineVersion != "" {
		instanceInput.EngineVersion = aws.String(input.EngineVersion)
	}

	if input.InstanceClass != "" {
		instanceInput.DBInstanceClass = aws.String(input.InstanceClass)
	}

	if input.ParameterGroupName != "" {
		instanceInput.DBParameterGroupName = aws.String(input.ParameterGroupName)
	}

	if input.EnhancedMonitoring {
		if input.MonitoringInterval > 0 {
			instanceInput.MonitoringInterval = aws.Int64(input.MonitoringInterval)
		}

		instanceInput.MonitoringRoleArn = aws.String(input.MonitoringRoleArn)
	}

	return instanceInput
}
