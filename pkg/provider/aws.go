package provider

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"github.com/aws/aws-sdk-go/aws/session"
	log "github.com/sirupsen/logrus"
)

const (
	profile       = "dev"
	defaultRegion = "us-west-2"
)

type AwsSessionEnv struct {
	Aki     string
	Sak     string
	Region  string
	RoleArn string
}

func (env AwsSessionEnv) PopulateEnv() AwsSessionEnv {
	env.Aki = os.Getenv("AWS_ACCESS_KEY_ID")
	env.Sak = os.Getenv("AWS_SECRET_ACCESS_KEY")
	env.Region = os.Getenv("AWS_REGION")
	env.RoleArn = os.Getenv("AWS_ROLE_ARN")
	return env
}

func NewSessionFromEnv(env AwsSessionEnv) *session.Session {
	return newKeySession(env.Aki, env.Sak, env.Region, env.RoleArn)
}

func NewSession() *session.Session {
	if os.Getenv("AWS_KEY_SESSION") == "true" {
		log.Debug("building session from keys")
		env := AwsSessionEnv{}.PopulateEnv()
		return NewSessionFromEnv(env)
	}
	log.Debug("building session from profile")
	return newSessionFromProfile()
}

func newSessionFromProfile() *session.Session {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		Config:  aws.Config{Region: aws.String(defaultRegion)},
		Profile: profile,
	}))

	return sess
}

func newKeySession(aki, sak, region, roleArn string) *session.Session {
	creds := credentials.NewStaticCredentials(aki, sak, "")
	cfg := aws.NewConfig().WithCredentials(creds).WithRegion(region)

	sess := session.Must(session.NewSession(cfg))
	creds = stscreds.NewCredentials(sess, roleArn)

	assumeCfg := cfg.Copy()
	assumeCfg.Credentials = creds

	assumeSess := session.Must(session.NewSession(assumeCfg))
	return assumeSess
}
