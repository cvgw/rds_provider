package cluster

import (
	"errors"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/rds"
	log "github.com/sirupsen/logrus"
)

var (
	ClusterNotFoundErr error
)

func init() {
	ClusterNotFoundErr = errors.New("cluster not found")
}

func FindDBCluster(svc *rds.RDS, clusterId string) (*rds.DBCluster, error) {
	descClustersInput := &rds.DescribeDBClustersInput{
		DBClusterIdentifier: aws.String(clusterId),
	}

	descClusterOuput, err := svc.DescribeDBClusters(descClustersInput)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case rds.ErrCodeDBClusterNotFoundFault:
				log.Info(rds.ErrCodeDBClusterNotFoundFault, aerr.Error())
				return nil, ClusterNotFoundErr
			default:
				log.Warn(aerr.Error())
				return nil, aerr
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			log.Warn(err.Error())
			return nil, aerr
		}
	}

	return descClusterOuput.DBClusters[0], nil
}
