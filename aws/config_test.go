package aws

import (
	"net/http"
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go/aws/credentials"
)

var testCredentials = credentials.NewChainCredentials([]credentials.Provider{
	&credentials.EnvProvider{},
	&credentials.SharedCredentialsProvider{
		Filename: "TestFilename",
		Profile:  "TestProfile"},
	&credentials.EC2RoleProvider{ExpiryWindow: 5 * time.Minute},
})

var mergeTestZeroValueConfig = Config{}

var mergeTestConfig = Config{
	Credentials:             testCredentials,
	Endpoint:                NewString("MergeTestEndpoint"),
	Region:                  NewString("MERGE_TEST_AWS_REGION"),
	DisableSSL:              NewBool(true),
	ManualSend:              NewBool(true),
	HTTPClient:              http.DefaultClient,
	LogHTTPBody:             NewBool(true),
	LogLevel:                NewInt(2),
	Logger:                  os.Stdout,
	MaxRetries:              NewInt(10),
	DisableParamValidation:  NewBool(true),
	DisableComputeChecksums: NewBool(true),
	S3ForcePathStyle:        NewBool(true),
}

var mergeTests = []struct {
	cfg  *Config
	in   *Config
	want *Config
}{
	{&Config{}, nil, &Config{}},
	{&Config{}, &mergeTestZeroValueConfig, &Config{}},
	{&Config{}, &mergeTestConfig, &mergeTestConfig},
}

var mergeErrorFmt = `
Merge Failed [%d]
  Config  %+v
    Merge(%+v)
      got %+v
     want %+v
`

func TestMerge(t *testing.T) {
	for i, tt := range mergeTests {
		got := tt.cfg.Merge(tt.in)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf(mergeErrorFmt, i, tt.cfg, tt.in, got, tt.want)
		}
	}
}
