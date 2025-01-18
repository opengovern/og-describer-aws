module github.com/opengovern/og-describer-aws

go 1.23.3

require (
	github.com/aws/aws-sdk-go-v2/service/accessanalyzer v1.36.8
	github.com/aws/aws-sdk-go-v2/service/account v1.22.4
	github.com/aws/aws-sdk-go-v2/service/acm v1.30.13
	github.com/aws/aws-sdk-go-v2/service/acmpca v1.37.13
	github.com/aws/aws-sdk-go-v2/service/amp v1.30.9
	github.com/aws/aws-sdk-go-v2/service/amplify v1.28.4
	github.com/aws/aws-sdk-go-v2/service/apigateway v1.28.7
	github.com/aws/aws-sdk-go-v2/service/apigatewayv2 v1.24.13
	github.com/aws/aws-sdk-go-v2/service/appconfig v1.36.7
	github.com/aws/aws-sdk-go-v2/service/applicationautoscaling v1.34.8
	github.com/aws/aws-sdk-go-v2/service/applicationinsights v1.29.9
	github.com/aws/aws-sdk-go-v2/service/appstream v1.42.4
	github.com/aws/aws-sdk-go-v2/service/athena v1.49.5
	github.com/aws/aws-sdk-go-v2/service/auditmanager v1.37.11
	github.com/aws/aws-sdk-go-v2/service/autoscaling v1.51.7
	github.com/aws/aws-sdk-go-v2/service/backup v1.40.4
	github.com/aws/aws-sdk-go-v2/service/batch v1.49.6
	github.com/aws/aws-sdk-go-v2/service/cloudcontrol v1.23.7
	github.com/aws/aws-sdk-go-v2/service/cloudformation v1.56.7
	github.com/aws/aws-sdk-go-v2/service/cloudfront v1.44.5
	github.com/aws/aws-sdk-go-v2/service/cloudsearch v1.26.10
	github.com/aws/aws-sdk-go-v2/service/cloudtrail v1.46.9
	github.com/aws/aws-sdk-go-v2/service/cloudwatch v1.43.9
	github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs v1.45.6
	github.com/aws/aws-sdk-go-v2/service/codeartifact v1.33.11
	github.com/aws/aws-sdk-go-v2/service/codebuild v1.50.2
	github.com/aws/aws-sdk-go-v2/service/codecommit v1.27.11
	github.com/aws/aws-sdk-go-v2/service/codedeploy v1.29.13
	github.com/aws/aws-sdk-go-v2/service/codepipeline v1.38.4
	github.com/aws/aws-sdk-go-v2/service/codestar v1.23.4
	github.com/aws/aws-sdk-go-v2/service/configservice v1.51.7
	github.com/aws/aws-sdk-go-v2/service/costexplorer v1.46.3
	github.com/aws/aws-sdk-go-v2/service/databasemigrationservice v1.45.6
	github.com/aws/aws-sdk-go-v2/service/dax v1.23.11
	github.com/aws/aws-sdk-go-v2/service/directconnect v1.30.7
	github.com/aws/aws-sdk-go-v2/service/directoryservice v1.30.12
	github.com/aws/aws-sdk-go-v2/service/dlm v1.29.6
	github.com/aws/aws-sdk-go-v2/service/docdb v1.40.5
	github.com/aws/aws-sdk-go-v2/service/drs v1.30.11
	github.com/aws/aws-sdk-go-v2/service/dynamodb v1.39.5
	github.com/aws/aws-sdk-go-v2/service/dynamodbstreams v1.24.15
	github.com/aws/aws-sdk-go-v2/service/ec2 v1.200.0
	github.com/aws/aws-sdk-go-v2/service/ecr v1.38.6
	github.com/aws/aws-sdk-go-v2/service/ecrpublic v1.29.4
	github.com/aws/aws-sdk-go-v2/service/ecs v1.53.8
	github.com/aws/aws-sdk-go-v2/service/efs v1.34.5
	github.com/aws/aws-sdk-go-v2/service/eks v1.56.5
	github.com/aws/aws-sdk-go-v2/service/elasticache v1.44.7
	github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk v1.28.12
	github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing v1.28.12
	github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2 v1.43.7
	github.com/aws/aws-sdk-go-v2/service/elasticsearchservice v1.32.13
	github.com/aws/aws-sdk-go-v2/service/emr v1.47.7
	github.com/aws/aws-sdk-go-v2/service/eventbridge v1.36.6
	github.com/aws/aws-sdk-go-v2/service/firehose v1.35.6
	github.com/aws/aws-sdk-go-v2/service/fms v1.39.2
	github.com/aws/aws-sdk-go-v2/service/fsx v1.51.5
	github.com/aws/aws-sdk-go-v2/service/glacier v1.26.12
	github.com/aws/aws-sdk-go-v2/service/globalaccelerator v1.29.11
	github.com/aws/aws-sdk-go-v2/service/glue v1.105.3
	github.com/aws/aws-sdk-go-v2/service/grafana v1.26.11
	github.com/aws/aws-sdk-go-v2/service/guardduty v1.52.5
	github.com/aws/aws-sdk-go-v2/service/health v1.29.5
	github.com/aws/aws-sdk-go-v2/service/iam v1.38.7
	github.com/aws/aws-sdk-go-v2/service/identitystore v1.27.12
	github.com/aws/aws-sdk-go-v2/service/imagebuilder v1.40.3
	github.com/aws/aws-sdk-go-v2/service/inspector v1.25.11
	github.com/aws/aws-sdk-go-v2/service/inspector2 v1.34.5
	github.com/aws/aws-sdk-go-v2/service/kafka v1.38.12
	github.com/aws/aws-sdk-go-v2/service/keyspaces v1.16.6
	github.com/aws/aws-sdk-go-v2/service/kinesis v1.32.13
	github.com/aws/aws-sdk-go-v2/service/kinesisanalyticsv2 v1.31.12
	github.com/aws/aws-sdk-go-v2/service/kinesisvideo v1.27.11
	github.com/aws/aws-sdk-go-v2/service/lakeformation v1.39.5
	github.com/aws/aws-sdk-go-v2/service/lambda v1.69.7
	github.com/aws/aws-sdk-go-v2/service/lightsail v1.42.11
	github.com/aws/aws-sdk-go-v2/service/macie2 v1.44.5
	github.com/aws/aws-sdk-go-v2/service/mediastore v1.24.11
	github.com/aws/aws-sdk-go-v2/service/memorydb v1.25.4
	github.com/aws/aws-sdk-go-v2/service/mgn v1.32.11
	github.com/aws/aws-sdk-go-v2/service/mq v1.27.12
	github.com/aws/aws-sdk-go-v2/service/mwaa v1.33.6
	github.com/aws/aws-sdk-go-v2/service/neptune v1.35.12
	github.com/aws/aws-sdk-go-v2/service/networkfirewall v1.44.10
	github.com/aws/aws-sdk-go-v2/service/oam v1.15.13
	github.com/aws/aws-sdk-go-v2/service/opensearch v1.45.6
	github.com/aws/aws-sdk-go-v2/service/opensearchserverless v1.17.9
	github.com/aws/aws-sdk-go-v2/service/opsworkscm v1.27.13
	github.com/aws/aws-sdk-go-v2/service/organizations v1.37.3
	github.com/aws/aws-sdk-go-v2/service/pinpoint v1.34.11
	github.com/aws/aws-sdk-go-v2/service/pipes v1.18.9
	github.com/aws/aws-sdk-go-v2/service/pricing v1.32.11
	github.com/aws/aws-sdk-go-v2/service/ram v1.29.13
	github.com/aws/aws-sdk-go-v2/service/rds v1.93.7
	github.com/aws/aws-sdk-go-v2/service/redshift v1.53.7
	github.com/aws/aws-sdk-go-v2/service/redshiftserverless v1.25.4
	github.com/aws/aws-sdk-go-v2/service/resourceexplorer2 v1.16.6
	github.com/aws/aws-sdk-go-v2/service/resourcegroups v1.27.13
	github.com/aws/aws-sdk-go-v2/service/resourcegroupstaggingapi v1.25.13
	github.com/aws/aws-sdk-go-v2/service/route53 v1.48.2
	github.com/aws/aws-sdk-go-v2/service/route53domains v1.28.5
	github.com/aws/aws-sdk-go-v2/service/route53resolver v1.34.8
	github.com/aws/aws-sdk-go-v2/service/s3 v1.48.1
	github.com/aws/aws-sdk-go-v2/service/s3control v1.52.6
	github.com/aws/aws-sdk-go-v2/service/sagemaker v1.173.1
	github.com/aws/aws-sdk-go-v2/service/secretsmanager v1.34.13
	github.com/aws/aws-sdk-go-v2/service/securityhub v1.55.5
	github.com/aws/aws-sdk-go-v2/service/securitylake v1.19.10
	github.com/aws/aws-sdk-go-v2/service/serverlessapplicationrepository v1.24.11
	github.com/aws/aws-sdk-go-v2/service/servicecatalog v1.32.11
	github.com/aws/aws-sdk-go-v2/service/servicediscovery v1.34.6
	github.com/aws/aws-sdk-go-v2/service/servicequotas v1.25.13
	github.com/aws/aws-sdk-go-v2/service/ses v1.29.6
	github.com/aws/aws-sdk-go-v2/service/sesv2 v1.41.1
	github.com/aws/aws-sdk-go-v2/service/sfn v1.34.7
	github.com/aws/aws-sdk-go-v2/service/shield v1.29.11
	github.com/aws/aws-sdk-go-v2/service/simspaceweaver v1.14.11
	github.com/aws/aws-sdk-go-v2/service/sns v1.33.13
	github.com/aws/aws-sdk-go-v2/service/sqs v1.37.9
	github.com/aws/aws-sdk-go-v2/service/ssm v1.56.7
	github.com/aws/aws-sdk-go-v2/service/ssoadmin v1.29.11
	github.com/aws/aws-sdk-go-v2/service/storagegateway v1.34.11
	github.com/aws/aws-sdk-go-v2/service/support v1.26.11
	github.com/aws/aws-sdk-go-v2/service/timestreamwrite v1.29.12
	github.com/aws/aws-sdk-go-v2/service/waf v1.25.11
	github.com/aws/aws-sdk-go-v2/service/wafregional v1.25.11
	github.com/aws/aws-sdk-go-v2/service/wafv2 v1.55.10
	github.com/aws/aws-sdk-go-v2/service/wellarchitected v1.34.11
	github.com/aws/aws-sdk-go-v2/service/workspaces v1.52.1
	github.com/go-errors/errors v1.5.1
	github.com/gocarina/gocsv v0.0.0-20240520201108-78e41c74b4b1
	github.com/golang-jwt/jwt/v5 v5.2.1
	github.com/google/go-github/v55 v55.0.0
	github.com/google/uuid v1.6.0
	github.com/hashicorp/go-hclog v1.6.3
	github.com/hashicorp/go-plugin v1.6.0
	github.com/jackc/pgtype v1.14.4
	github.com/nats-io/nats.go v1.36.0
	github.com/opengovern/og-util v1.11.0
	github.com/shurcooL/githubv4 v0.0.0-20240727222349-48295856cce7
	github.com/spf13/cobra v1.8.1
	github.com/turbot/steampipe-plugin-github v1.0.0
	github.com/turbot/steampipe-plugin-sdk/v5 v5.10.4
	go.uber.org/zap v1.27.0
	golang.org/x/net v0.33.0
	golang.org/x/oauth2 v0.23.0
	google.golang.org/grpc v1.67.1
	google.golang.org/protobuf v1.35.1
)

require (
	github.com/Azure/azure-sdk-for-go/sdk/azcore v1.16.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/azidentity v1.8.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/internal v1.10.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/security/keyvault/azsecrets v1.2.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/security/keyvault/internal v1.1.0 // indirect
	github.com/AzureAD/microsoft-authentication-library-for-go v1.2.3 // indirect
	github.com/acarl005/stripansi v0.0.0-20180116102854-5a71ef0e047d // indirect
	github.com/agext/levenshtein v1.2.3 // indirect
	github.com/allegro/bigcache/v3 v3.1.0 // indirect
	github.com/apparentlymart/go-textseg/v15 v15.0.0 // indirect
	github.com/aws/aws-sdk-go v1.55.5
	github.com/aws/aws-sdk-go-v2 v1.33.0
	github.com/aws/aws-sdk-go-v2/config v1.28.1
	github.com/aws/aws-sdk-go-v2/credentials v1.17.42
	github.com/aws/aws-sdk-go-v2/feature/ec2/imds v1.16.18 // indirect
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.3.28 // indirect
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.6.28 // indirect
	github.com/aws/aws-sdk-go-v2/internal/ini v1.8.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding v1.12.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.12.9 // indirect
	github.com/aws/aws-sdk-go-v2/service/kms v1.35.3
	github.com/aws/aws-sdk-go-v2/service/sso v1.24.3 // indirect
	github.com/aws/aws-sdk-go-v2/service/ssooidc v1.28.3 // indirect
	github.com/aws/aws-sdk-go-v2/service/sts v1.32.3
	github.com/aws/smithy-go v1.22.1
	github.com/ghodss/yaml v1.0.0
	github.com/golang/protobuf v1.5.4
	github.com/labstack/echo/v4 v4.12.0 // indirect
	github.com/turbot/go-kit v0.10.0-rc.0
	golang.org/x/time v0.8.0
	golang.org/x/tools v0.26.0
	google.golang.org/genproto v0.0.0-20241021214115-324edc3d5d38
)

require (
	cloud.google.com/go v0.116.0 // indirect
	cloud.google.com/go/auth v0.10.0 // indirect
	cloud.google.com/go/auth/oauth2adapt v0.2.5 // indirect
	cloud.google.com/go/compute/metadata v0.5.2 // indirect
	cloud.google.com/go/iam v1.2.1 // indirect
	cloud.google.com/go/longrunning v0.6.1 // indirect
	cloud.google.com/go/storage v1.43.0 // indirect
	github.com/ProtonMail/go-crypto v1.1.3 // indirect
	github.com/aws/aws-sdk-go-v2/aws/protocol/eventstream v1.6.7 // indirect
	github.com/aws/aws-sdk-go-v2/internal/v4a v1.3.28 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/checksum v1.2.10 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/endpoint-discovery v1.10.9 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/s3shared v1.18.9 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/bgentry/go-netrc v0.0.0-20140422174119-9fd32a8b3d3d // indirect
	github.com/bradleyfalzon/ghinstallation v1.1.1 // indirect
	github.com/btubbs/datetime v0.1.1 // indirect
	github.com/buildkite/go-pipeline v0.3.1 // indirect
	github.com/buildkite/interpolate v0.0.0-20200526001904-07f35b4ae251 // indirect
	github.com/cenkalti/backoff/v4 v4.3.0 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/cloudflare/circl v1.3.8 // indirect
	github.com/danwakefield/fnmatch v0.0.0-20160403171240-cbb64ac3d964 // indirect
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/dgraph-io/ristretto v0.1.1 // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible // indirect
	github.com/dustin/go-humanize v1.0.1 // indirect
	github.com/eko/gocache/lib/v4 v4.1.5 // indirect
	github.com/eko/gocache/store/bigcache/v4 v4.2.1 // indirect
	github.com/eko/gocache/store/ristretto/v4 v4.2.1 // indirect
	github.com/elastic/go-elasticsearch/v7 v7.17.10 // indirect
	github.com/fatih/color v1.18.0 // indirect
	github.com/felixge/httpsnoop v1.0.4 // indirect
	github.com/fsnotify/fsnotify v1.7.0 // indirect
	github.com/fxamacker/cbor/v2 v2.7.0 // indirect
	github.com/gertd/go-pluralize v0.2.1 // indirect
	github.com/globocom/echo-prometheus v0.1.2 // indirect
	github.com/go-jose/go-jose/v4 v4.0.1 // indirect
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-openapi/jsonpointer v0.21.0 // indirect
	github.com/go-openapi/swag v0.23.0 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang-jwt/jwt v3.2.2+incompatible // indirect
	github.com/golang/glog v1.2.2 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/golang/mock v1.6.0 // indirect
	github.com/google/go-cmp v0.6.0 // indirect
	github.com/google/go-github/v29 v29.0.2 // indirect
	github.com/google/go-querystring v1.1.0 // indirect
	github.com/google/gofuzz v1.2.0 // indirect
	github.com/google/s2a-go v0.1.8 // indirect
	github.com/googleapis/enterprise-certificate-proxy v0.3.4 // indirect
	github.com/googleapis/gax-go/v2 v2.13.0 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.22.0 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-cleanhttp v0.5.2 // indirect
	github.com/hashicorp/go-getter v1.7.5 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/hashicorp/go-retryablehttp v0.7.7 // indirect
	github.com/hashicorp/go-rootcerts v1.0.2 // indirect
	github.com/hashicorp/go-safetemp v1.0.0 // indirect
	github.com/hashicorp/go-secure-stdlib/parseutil v0.1.8 // indirect
	github.com/hashicorp/go-secure-stdlib/strutil v0.1.2 // indirect
	github.com/hashicorp/go-sockaddr v1.0.6 // indirect
	github.com/hashicorp/go-version v1.6.0 // indirect
	github.com/hashicorp/hcl v1.0.1-vault // indirect
	github.com/hashicorp/hcl/v2 v2.20.1 // indirect
	github.com/hashicorp/vault/api v1.15.0 // indirect
	github.com/hashicorp/vault/api/auth/kubernetes v0.7.0 // indirect
	github.com/hashicorp/yamux v0.1.1 // indirect
	github.com/iancoleman/strcase v0.3.0 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/jackc/pgio v1.0.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20240606120523-5a60cdf6a761 // indirect
	github.com/jackc/pgx/v5 v5.7.1 // indirect
	github.com/jackc/puddle/v2 v2.2.2 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/compress v1.17.11 // indirect
	github.com/kylelemons/godebug v1.1.0 // indirect
	github.com/labstack/gommon v0.4.2 // indirect
	github.com/leodido/go-urn v1.4.0 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mattn/go-runewidth v0.0.15 // indirect
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	github.com/mitchellh/go-testing-interface v1.14.1 // indirect
	github.com/mitchellh/go-wordwrap v1.0.1 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/munnerz/goautoneg v0.0.0-20191010083416-a7dc8b61c822 // indirect
	github.com/nats-io/nkeys v0.4.7 // indirect
	github.com/nats-io/nuid v1.0.1 // indirect
	github.com/oklog/run v1.1.0 // indirect
	github.com/oleiade/reflections v1.0.1 // indirect
	github.com/olekukonko/tablewriter v0.0.5 // indirect
	github.com/opensearch-project/opensearch-go/v2 v2.3.0 // indirect
	github.com/pganalyze/pg_query_go/v4 v4.2.3 // indirect
	github.com/pkg/browser v0.0.0-20240102092130-5ac0b6a4141c // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/prometheus/client_golang v1.20.5 // indirect
	github.com/prometheus/client_model v0.6.1 // indirect
	github.com/prometheus/common v0.60.1 // indirect
	github.com/prometheus/procfs v0.15.1 // indirect
	github.com/rivo/uniseg v0.4.7 // indirect
	github.com/ryanuber/go-glob v1.0.0 // indirect
	github.com/sethvargo/go-retry v0.2.4 // indirect
	github.com/shurcooL/graphql v0.0.0-20220606043923-3cf50f8a0a29 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/stevenle/topsort v0.2.0 // indirect
	github.com/tkrajina/go-reflector v0.5.6 // indirect
	github.com/ulikunitz/xz v0.5.11 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasttemplate v1.2.2 // indirect
	github.com/x448/float16 v0.8.4 // indirect
	github.com/zclconf/go-cty v1.14.4 // indirect
	go.opencensus.io v0.24.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/github.com/labstack/echo/otelecho v0.53.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.54.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.54.0 // indirect
	go.opentelemetry.io/otel v1.31.0 // indirect
	go.opentelemetry.io/otel/exporters/jaeger v1.17.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc v1.31.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace v1.28.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc v1.28.0 // indirect
	go.opentelemetry.io/otel/metric v1.31.0 // indirect
	go.opentelemetry.io/otel/sdk v1.31.0 // indirect
	go.opentelemetry.io/otel/sdk/metric v1.31.0 // indirect
	go.opentelemetry.io/otel/trace v1.31.0 // indirect
	go.opentelemetry.io/proto/otlp v1.3.1 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	golang.org/x/crypto v0.31.0 // indirect
	golang.org/x/exp v0.0.0-20240909161429-701f63a606c0 // indirect
	golang.org/x/mod v0.21.0 // indirect
	golang.org/x/sync v0.10.0 // indirect
	golang.org/x/sys v0.28.0 // indirect
	golang.org/x/term v0.27.0 // indirect
	golang.org/x/text v0.21.0 // indirect
	google.golang.org/api v0.204.0 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20241015192408-796eee8c2d53 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20241021214115-324edc3d5d38 // indirect
	gopkg.in/go-playground/validator.v9 v9.31.0 // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	k8s.io/apimachinery v0.31.2 // indirect
	k8s.io/client-go v0.31.2 // indirect
	k8s.io/klog/v2 v2.130.1 // indirect
	k8s.io/utils v0.0.0-20240921022957-49e7df575cb6 // indirect
	sigs.k8s.io/json v0.0.0-20221116044647-bc3834ca7abd // indirect
	sigs.k8s.io/structured-merge-diff/v4 v4.4.1 // indirect
	sigs.k8s.io/yaml v1.4.0 // indirect
)
