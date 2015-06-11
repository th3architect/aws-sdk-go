package aws

import (
	"io"
	"net/http"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws/credentials"
)

// DefaultChainCredentials is a Credentials which will find the first available
// credentials Value from the list of Providers.
//
// This should be used in the default case. Once the type of credentials are
// known switching to the specific Credentials will be more efficient.
var DefaultChainCredentials = credentials.NewChainCredentials(
	[]credentials.Provider{
		&credentials.EnvProvider{},
		&credentials.SharedCredentialsProvider{Filename: "", Profile: ""},
		&credentials.EC2RoleProvider{ExpiryWindow: 5 * time.Minute},
	})

// The default number of retries for a service. The value of -1 indicates that
// the service specific retry default will be used.
const DefaultRetries = -1

// DefaultConfig is the default all service configuration will be based off of.
var DefaultConfig = &Config{
	Credentials:             DefaultChainCredentials,
	Region:                  NewString(os.Getenv("AWS_REGION")),
	HTTPClient:              http.DefaultClient,
	Logger:                  os.Stdout,
	MaxRetries:              NewInt(DefaultRetries),
}

// A Config provides service configuration
type Config struct {
	Credentials             *credentials.Credentials
	Endpoint                String
	Region                  String
	DisableSSL              Bool
	ManualSend              Bool
	HTTPClient              *http.Client
	LogHTTPBody             Bool
	LogLevel                Int
	Logger                  io.Writer
	MaxRetries              Int
	DisableParamValidation  Bool
	DisableComputeChecksums Bool
	S3ForcePathStyle        Bool
}

// Merge merges the newcfg attribute values into this Config. Each attribute
// will be merged into this config if the newcfg attribute's value is non-zero.
// Due to this, newcfg attributes with zero values cannot be merged in. For
// example bool attributes cannot be cleared using Merge, and must be explicitly
// set on the Config structure.
func (c Config) Merge(newcfg *Config) *Config {
	if newcfg == nil {
		return &c
	}

	cfg := c

	if newcfg.Credentials != nil {
		cfg.Credentials = newcfg.Credentials
	}

	if newcfg.Endpoint.IsSet() {
		cfg.Endpoint = newcfg.Endpoint
	}

	if newcfg.Region.IsSet() {
		cfg.Region = newcfg.Region
	}

	if newcfg.DisableSSL.IsSet() {
		cfg.DisableSSL = newcfg.DisableSSL
	}

	if newcfg.ManualSend.IsSet() {
		cfg.ManualSend = newcfg.ManualSend
	}

	if newcfg.HTTPClient != nil {
		cfg.HTTPClient = newcfg.HTTPClient
	}

	if newcfg.LogHTTPBody.IsSet() {
		cfg.LogHTTPBody = newcfg.LogHTTPBody
	}

	if newcfg.LogLevel.IsSet() {
		cfg.LogLevel = newcfg.LogLevel
	}

	if newcfg.Logger != nil {
		cfg.Logger = newcfg.Logger
	}

	if newcfg.MaxRetries.IsSet() {
		cfg.MaxRetries = newcfg.MaxRetries
	}

	if newcfg.DisableParamValidation.IsSet() {
		cfg.DisableParamValidation = newcfg.DisableParamValidation
	}

	if newcfg.DisableComputeChecksums.IsSet() {
		cfg.DisableComputeChecksums = newcfg.DisableComputeChecksums
	}

	if newcfg.S3ForcePathStyle.IsSet() {
		cfg.S3ForcePathStyle = newcfg.S3ForcePathStyle
	}

	return &cfg
}
