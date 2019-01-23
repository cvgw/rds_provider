package cluster

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/rds"
	log "github.com/sirupsen/logrus"
)

// NewDBClusterInput contains the data to be used when creating the RDS cluster
type NewDBClusterInput struct {
	// Identifier for the RDS cluster
	ClusterId string
	// Engine type
	Engine string
	// Engine version (optional)
	EngineVersion string
	// User name for root DB user
	MasterUsername string
	// Password for root DB user
	MasterUserPass string
	// List of Security groups to attach to the RDS cluster
	SecurityGroupIds []string
	// Name of the RDS Subnet Group to use
	SubnetGroupName string
	// Name of RDS Cluster Parameter Group to use (optional)
	ParameterGroupName string
	// List of Availability Zones in which to deploy the RDS cluster (optional)
	AvailabilityZones []string
	// Length of time (in days) to retain backups (optional)
	BackupRetentionPeriod int64
	// Whether the DB data should be encrypted at rest (optional)
	StorageEncrypted bool
}

func CreateDBCluster(svc *rds.RDS, input NewDBClusterInput) (*rds.DBCluster, error) {
	clusterInput := NewCreateClusterInput(input)
	clusterOutput, err := svc.CreateDBCluster(clusterInput)
	if err != nil {
		log.Warn(err)
		return nil, err
	}

	return clusterOutput.DBCluster, nil
}

func NewCreateClusterInput(input NewDBClusterInput) *rds.CreateDBClusterInput {
	sIds := make([]*string, 0)
	for _, i := range input.SecurityGroupIds {
		sIds = append(sIds, aws.String(i))
	}

	azs := make([]*string, 0)
	for _, i := range input.AvailabilityZones {
		azs = append(azs, aws.String(i))
	}

	clusterInput := &rds.CreateDBClusterInput{
		DBClusterIdentifier: aws.String(input.ClusterId),
		Engine:              aws.String(input.Engine),
		MasterUsername:      aws.String(input.MasterUsername),
		MasterUserPassword:  aws.String(input.MasterUserPass),
		DBSubnetGroupName:   aws.String(input.SubnetGroupName),
		VpcSecurityGroupIds: sIds,
		StorageEncrypted:    aws.Bool(input.StorageEncrypted),
	}

	if input.EngineVersion != "" {
		clusterInput.EngineVersion = aws.String(input.EngineVersion)
	}

	if input.ParameterGroupName != "" {
		clusterInput.DBClusterParameterGroupName = aws.String(input.ParameterGroupName)
	}

	if len(azs) > 0 {
		clusterInput.AvailabilityZones = azs
	}

	if input.BackupRetentionPeriod > 0 {
		clusterInput.BackupRetentionPeriod = aws.Int64(input.BackupRetentionPeriod)
	}

	return clusterInput
}
