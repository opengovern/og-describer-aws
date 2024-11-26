package configs

import "github.com/opengovern/opencomply/services/integration/integration-type/aws-account/configs"

type IntegrationCredentials struct {
	configs.IntegrationCredentials
	AccountID string `json:"account_id"`
}
