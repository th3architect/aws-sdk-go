package sts_test

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sts"
	"github.com/stretchr/testify/assert"
)

var svc = sts.New(&aws.Config{
	Region: aws.NewString("mock-region"),
})

func TestUnsignedRequest_AssumeRoleWithSAML(t *testing.T) {
	req, _ := svc.AssumeRoleWithSAMLRequest(&sts.AssumeRoleWithSAMLInput{
		PrincipalARN:  aws.StringPtr("ARN"),
		RoleARN:       aws.StringPtr("ARN"),
		SAMLAssertion: aws.StringPtr("ASSERT"),
	})

	err := req.Sign()
	assert.NoError(t, err)
	assert.Equal(t, "", req.HTTPRequest.Header.Get("Authorization"))
}

func TestUnsignedRequest_AssumeRoleWithWebIdentity(t *testing.T) {
	req, _ := svc.AssumeRoleWithWebIdentityRequest(&sts.AssumeRoleWithWebIdentityInput{
		RoleARN:          aws.StringPtr("ARN"),
		RoleSessionName:  aws.StringPtr("SESSION"),
		WebIdentityToken: aws.StringPtr("TOKEN"),
	})

	err := req.Sign()
	assert.NoError(t, err)
	assert.Equal(t, "", req.HTTPRequest.Header.Get("Authorization"))
}
