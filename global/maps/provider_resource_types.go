package maps
import (
	"github.com/opengovern/og-describer-aws/discovery/describers"
	"github.com/opengovern/og-describer-aws/discovery/provider"
	"github.com/opengovern/og-describer-aws/platform/constants"
	"github.com/opengovern/og-util/pkg/integration/interfaces"
	model "github.com/opengovern/og-describer-aws/discovery/pkg/models"
)
var ResourceTypes = map[string]model.ResourceType{

	"AWS::Redshift::Snapshot": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::Redshift::Snapshot",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.RedshiftSnapshot),
		GetDescriber:         nil,
	},

	"AWS::IAM::AccountSummary": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::IAM::AccountSummary",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.SequentialDescribeGlobal(describers.IAMAccountSummary),
		GetDescriber:         nil,
	},

	"AWS::Glacier::Vault": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::Glacier::Vault",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.GlacierVault),
		GetDescriber:         nil,
	},

	"AWS::Organizations::Organization": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::Organizations::Organization",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.SequentialDescribeGlobal(describers.OrganizationsOrganization),
		GetDescriber:         nil,
	},

	"AWS::Organizations::Policy": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::Organizations::Policy",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.SequentialDescribeGlobal(describers.OrganizationsPolicy),
		GetDescriber:         nil,
	},

	"AWS::Organizations::PolicyTarget": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::Organizations::PolicyTarget",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.SequentialDescribeGlobal(describers.OrganizationsPolicyTarget),
		GetDescriber:         nil,
	},

	"AWS::Organizations::OrganizationalUnit": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::Organizations::OrganizationalUnit",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.SequentialDescribeGlobal(describers.OrganizationsOrganizationalUnit),
		GetDescriber:         nil,
	},

	"AWS::Organizations::Root": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::Organizations::Root",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.SequentialDescribeGlobal(describers.OrganizationsRoot),
		GetDescriber:         nil,
	},

	"AWS::CloudSearch::Domain": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::CloudSearch::Domain",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.CloudSearchDomain),
		GetDescriber:         nil,
	},

	"AWS::DynamoDb::GlobalSecondaryIndex": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::DynamoDb::GlobalSecondaryIndex",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.DynamoDbGlobalSecondaryIndex),
		GetDescriber:         nil,
	},

	"AWS::EC2::RouteTable": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::EC2::RouteTable",
		Tags:                 map[string][]string{
            "category": {"Networking"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance/awsicons/master/svg-export/icons/Route53RouteTable.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.EC2RouteTable),
		GetDescriber:         nil,
	},

	"AWS::SecurityHub::Hub": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::SecurityHub::Hub",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.SecurityHubHub),
		GetDescriber:         nil,
	},

	"AWS::StorageGateway::StorageGateway": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::StorageGateway::StorageGateway",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.StorageGatewayStorageGateway),
		GetDescriber:         nil,
	},

	"AWS::Inspector::AssessmentTemplate": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::Inspector::AssessmentTemplate",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.InspectorAssessmentTemplate),
		GetDescriber:         nil,
	},

	"AWS::ElasticLoadBalancingV2::ListenerRule": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::ElasticLoadBalancingV2::ListenerRule",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.ElasticLoadBalancingV2ListenerRule),
		GetDescriber:         nil,
	},

	"AWS::IAM::Role": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::IAM::Role",
		Tags:                 map[string][]string{
            "category": {"Management & Governance"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance/awsicons/master/svg-export/icons/IdentityAccessManagementRole.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.SequentialDescribeGlobal(describers.IAMRole),
		GetDescriber:         nil,
	},

	"AWS::Backup::ProtectedResource": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::Backup::ProtectedResource",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.BackupProtectedResource),
		GetDescriber:         nil,
	},

	"AWS::CodeCommit::Repository": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::CodeCommit::Repository",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.CodeCommitRepository),
		GetDescriber:         nil,
	},

	"AWS::EC2::VPCEndpoint": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::EC2::VPCEndpoint",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.EC2VPCEndpoint),
		GetDescriber:         nil,
	},

	"AWS::EventBridge::EventRule": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::EventBridge::EventRule",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.EventBridgeRule),
		GetDescriber:         nil,
	},

	"AWS::CloudFront::OriginAccessControl": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::CloudFront::OriginAccessControl",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.SequentialDescribeGlobal(describers.CloudFrontOriginAccessControl),
		GetDescriber:         nil,
	},

	"AWS::CodeBuild::Project": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::CodeBuild::Project",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.CodeBuildProject),
		GetDescriber:         nil,
	},

	"AWS::CodeBuild::Build": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::CodeBuild::Build",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.CodeBuildBuild),
		GetDescriber:         nil,
	},

	"AWS::ElastiCache::ParameterGroup": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::ElastiCache::ParameterGroup",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.ElastiCacheParameterGroup),
		GetDescriber:         nil,
	},

	"AWS::MemoryDb::Cluster": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::MemoryDb::Cluster",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.MemoryDbCluster),
		GetDescriber:         nil,
	},

	"AWS::Glue::Crawler": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::Glue::Crawler",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.GlueCrawler),
		GetDescriber:         nil,
	},

	"AWS::DirectConnect::Gateway": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::DirectConnect::Gateway",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.DirectConnectGateway),
		GetDescriber:         nil,
	},

	"AWS::DynamoDb::BackUp": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::DynamoDb::BackUp",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.DynamoDbBackUp),
		GetDescriber:         nil,
	},

	"AWS::EC2::EIP": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::EC2::EIP",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.EC2EIP),
		GetDescriber:         nil,
	},

	"AWS::EC2::InternetGateway": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::EC2::InternetGateway",
		Tags:                 map[string][]string{
            "category": {"Networking"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance/awsicons/master/svg-export/icons/InternetGateway.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.EC2InternetGateway),
		GetDescriber:         nil,
	},

	"AWS::GuardDuty::PublishingDestination": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::GuardDuty::PublishingDestination",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.GuardDutyPublishingDestination),
		GetDescriber:         nil,
	},

	"AWS::KinesisAnalyticsV2::Application": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::KinesisAnalyticsV2::Application",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.KinesisAnalyticsV2Application),
		GetDescriber:         nil,
	},

	"AWS::EMR::Instance": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::EMR::Instance",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.EMRInstance),
		GetDescriber:         nil,
	},

	"AWS::EMR::BlockPublicAccessConfiguration": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::EMR::BlockPublicAccessConfiguration",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.EMRBlockPublicAccessConfiguration),
		GetDescriber:         nil,
	},

	"AWS::ApiGateway::RestApi": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::ApiGateway::RestApi",
		Tags:                 map[string][]string{
            "category": {"Networking"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance/awsicons/master/svg-export/icons/ApiGateway.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.ApiGatewayRestAPI),
		GetDescriber:         nil,
	},

	"AWS::ApiGatewayV2::Integration": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::ApiGatewayV2::Integration",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.ApiGatewayV2Integration),
		GetDescriber:         nil,
	},

	"AWS::AutoScaling::AutoScalingGroup": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::AutoScaling::AutoScalingGroup",
		Tags:                 map[string][]string{
            "category": {"Compute"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance/awsicons/master/svg-export/icons/Ec2AutoScaling.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.AutoScalingAutoScalingGroup),
		GetDescriber:         nil,
	},

	"AWS::DynamoDb::TableExport": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::DynamoDb::TableExport",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.DynamoDbTableExport),
		GetDescriber:         nil,
	},

	"AWS::EC2::KeyPair": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::EC2::KeyPair",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.EC2KeyPair),
		GetDescriber:         nil,
	},

	"AWS::EFS::FileSystem": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::EFS::FileSystem",
		Tags:                 map[string][]string{
            "category": {"Storage"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance/awsicons/master/svg-export/icons/ElasticFileSystem.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.EFSFileSystem),
		GetDescriber:         nil,
	},

	"AWS::Kafka::Cluster": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::Kafka::Cluster",
		Tags:                 map[string][]string{
            "category": {"PaaS"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance/awsicons/master/svg-export/icons/ManagedStreamingForKafka.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.KafkaCluster),
		GetDescriber:         nil,
	},

	"AWS::SecretsManager::Secret": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::SecretsManager::Secret",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.SecretsManagerSecret),
		GetDescriber:         nil,
	},

	"AWS::Backup::LegalHold": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::Backup::LegalHold",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.BackupLegalHold),
		GetDescriber:         nil,
	},

	"AWS::CloudFront::Function": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::CloudFront::Function",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.SequentialDescribeGlobal(describers.CloudFrontFunction),
		GetDescriber:         nil,
	},

	"AWS::GlobalAccelerator::EndpointGroup": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::GlobalAccelerator::EndpointGroup",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.GlobalAcceleratorEndpointGroup),
		GetDescriber:         nil,
	},

	"AWS::DAX::ParameterGroup": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::DAX::ParameterGroup",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.DAXParameterGroup),
		GetDescriber:         nil,
	},

	"AWS::SQS::Queue": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::SQS::Queue",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.SQSQueue),
		GetDescriber:         nil,
	},

	"AWS::Config::Rule": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::Config::Rule",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.ConfigRule),
		GetDescriber:         nil,
	},

	"AWS::GuardDuty::Member": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::GuardDuty::Member",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.GuardDutyMember),
		GetDescriber:         nil,
	},

	"AWS::Inspector::Exclusion": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::Inspector::Exclusion",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.InspectorExclusion),
		GetDescriber:         nil,
	},

	"AWS::DirectoryService::Directory": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::DirectoryService::Directory",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.DirectoryServiceDirectory),
		GetDescriber:         nil,
	},

	"AWS::DirectoryService::Certificate": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::DirectoryService::Certificate",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.DirectoryServiceCertificate),
		GetDescriber:         nil,
	},

	"AWS::DirectoryService::LogSubscription": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::DirectoryService::LogSubscription",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.DirectoryServiceLogSubscription),
		GetDescriber:         nil,
	},

	"AWS::EFS::AccessPoint": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::EFS::AccessPoint",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.EFSAccessPoint),
		GetDescriber:         nil,
	},

	"AWS::IAM::PolicyAttachment": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::IAM::PolicyAttachment",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.SequentialDescribeGlobal(describers.IAMPolicyAttachment),
		GetDescriber:         nil,
	},

	"AWS::IAM::CredentialReport": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::IAM::CredentialReport",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.SequentialDescribeGlobal(describers.IAMCredentialReport),
		GetDescriber:         nil,
	},

	"AWS::RDS::GlobalCluster": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::RDS::GlobalCluster",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.RDSGlobalCluster),
		GetDescriber:         nil,
	},

	"AWS::CostExplorer::ForcastDaily": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::CostExplorer::ForcastDaily",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.SequentialDescribeGlobal(describers.CostForecastDaily),
		GetDescriber:         nil,
	},

	"AWS::GuardDuty::Detector": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::GuardDuty::Detector",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.GuardDutyDetector),
		GetDescriber:         nil,
	},

	"AWS::SNS::Topic": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::SNS::Topic",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.SNSTopic),
		GetDescriber:         nil,
	},

	"AWS::AppConfig::Application": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::AppConfig::Application",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.AppConfigApplication),
		GetDescriber:         nil,
	},

	"AWS::Batch::Job": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::Batch::Job",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.BatchJob),
		GetDescriber:         nil,
	},

	"AWS::Batch::JobQueue": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::Batch::JobQueue",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.BatchJobQueue),
		GetDescriber:         nil,
	},

	"AWS::ECS::Service": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::ECS::Service",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.ECSService),
		GetDescriber:         nil,
	},

	"AWS::FSX::Task": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::FSX::Task",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.FSXTask),
		GetDescriber:         nil,
	},

	"AWS::IAM::VirtualMFADevice": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::IAM::VirtualMFADevice",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.SequentialDescribeGlobal(describers.IAMVirtualMFADevice),
		GetDescriber:         nil,
	},

	"AWS::WAFv2::WebACL": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::WAFv2::WebACL",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.WAFv2WebACL),
		GetDescriber:         nil,
	},

	"AWS::ApplicationAutoScaling::Target": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::ApplicationAutoScaling::Target",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.ApplicationAutoScalingTarget),
		GetDescriber:         nil,
	},

	"AWS::ApplicationAutoScaling::Policy": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::ApplicationAutoScaling::Policy",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.ApplicationAutoScalingPolicy),
		GetDescriber:         nil,
	},

	"AWS::Backup::Vault": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::Backup::Vault",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.BackupVault),
		GetDescriber:         nil,
	},

	"AWS::ElastiCache::Cluster": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::ElastiCache::Cluster",
		Tags:                 map[string][]string{
            "category": {"Database"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance/awsicons/master/svg-export/icons/ElastiCache.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.ElastiCacheCluster),
		GetDescriber:         nil,
	},

	"AWS::Logs::LogGroup": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::Logs::LogGroup",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.CloudWatchLogsLogGroup),
		GetDescriber:         nil,
	},

	"AWS::S3::Bucket": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::S3::Bucket",
		Tags:                 map[string][]string{
            "category": {"Storage"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance/awsicons/master/svg-export/icons/SimpleStorageBucket.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.SequentialDescribeGlobal(describers.S3Bucket),
		GetDescriber:         nil,
	},

	"AWS::S3::Object": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::S3::Object",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.SequentialDescribeGlobal(describers.S3Object),
		GetDescriber:         nil,
	},

	"AWS::S3::BucketIntelligentTieringConfiguration": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::S3::BucketIntelligentTieringConfiguration",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.SequentialDescribeGlobal(describers.S3BucketIntelligentTieringConfiguration),
		GetDescriber:         nil,
	},

	"AWS::S3::MultiRegionAccessPoint": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::S3::MultiRegionAccessPoint",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.SequentialDescribeGlobal(describers.S3MultiRegionAccessPoint),
		GetDescriber:         nil,
	},

	"AWS::CertificateManager::Certificate": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::CertificateManager::Certificate",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.CertificateManagerCertificate),
		GetDescriber:         nil,
	},

	"AWS::EKS::AddonVersion": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::EKS::AddonVersion",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.EKSAddonVersion),
		GetDescriber:         nil,
	},

	"AWS::ApiGatewayV2::Api": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::ApiGatewayV2::Api",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.ApiGatewayV2API),
		GetDescriber:         nil,
	},

	"AWS::EC2::Volume": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::EC2::Volume",
		Tags:                 map[string][]string{
            "category": {"Storage"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance/awsicons/master/svg-export/icons/ElasticBlockStoreVolume.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.EC2Volume),
		GetDescriber:         nil,
	},

	"AWS::ApiGateway::ApiKey": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::ApiGateway::ApiKey",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.ApiGatewayApiKey),
		GetDescriber:         nil,
	},

	"AWS::Glue::Connection": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::Glue::Connection",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.GlueConnection),
		GetDescriber:         nil,
	},

	"AWS::ECS::Task": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::ECS::Task",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.ECSTask),
		GetDescriber:         nil,
	},

	"AWS::SSM::ManagedInstance": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::SSM::ManagedInstance",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.SSMManagedInstance),
		GetDescriber:         nil,
	},

	"AWS::SSM::Inventory": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::SSM::Inventory",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.SSMInventory),
		GetDescriber:         nil,
	},

	"AWS::SSM::InventoryEntry": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::SSM::InventoryEntry",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.SSMInventoryEntry),
		GetDescriber:         nil,
	},

	"AWS::SSM::MaintenanceWindow": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::SSM::MaintenanceWindow",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.SSMMaintenanceWindow),
		GetDescriber:         nil,
	},

	"AWS::SSM::PatchBaseline": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::SSM::PatchBaseline",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.SSMPatchBaseline),
		GetDescriber:         nil,
	},

	"AWS::SSM::Parameter": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::SSM::Parameter",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.SSMParameter),
		GetDescriber:         nil,
	},

	"AWS::Lambda::Function": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::Lambda::Function",
		Tags:                 map[string][]string{
            "category": {"Serverless"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance/awsicons/master/svg-export/icons/Lambda.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.LambdaFunction),
		GetDescriber:         nil,
	},

	"AWS::RDS::DBSnapshot": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::RDS::DBSnapshot",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.RDSDBSnapshot),
		GetDescriber:         nil,
	},

	"AWS::CodeDeploy::Application": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::CodeDeploy::Application",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.CodeDeployApplication),
		GetDescriber:         nil,
	},

	"AWS::CodeDeploy::DeploymentConfig": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::CodeDeploy::DeploymentConfig",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.CodeDeployDeploymentConfig),
		GetDescriber:         nil,
	},

	"AWS::EMR::Cluster": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::EMR::Cluster",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.EMRCluster),
		GetDescriber:         nil,
	},

	"AWS::IAM::AccessKey": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::IAM::AccessKey",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.SequentialDescribeGlobal(describers.IAMAccessKey),
		GetDescriber:         nil,
	},

	"AWS::IAM::SSHPublicKey": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::IAM::SSHPublicKey",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.SequentialDescribeGlobal(describers.IAMSSHPublicKey),
		GetDescriber:         nil,
	},

	"AWS::Glue::CatalogTable": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::Glue::CatalogTable",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.GlueCatalogTable),
		GetDescriber:         nil,
	},

	"AWS::CloudTrail::Channel": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::CloudTrail::Channel",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.CloudTrailChannel),
		GetDescriber:         nil,
	},

	"AWS::EC2::NetworkAcl": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::EC2::NetworkAcl",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.EC2NetworkAcl),
		GetDescriber:         nil,
	},

	"AWS::ECS::ContainerInstance": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::ECS::ContainerInstance",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.ECSContainerInstance),
		GetDescriber:         nil,
	},

	"AWS::RedshiftServerless::Snapshot": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::RedshiftServerless::Snapshot",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.RedshiftServerlessSnapshot),
		GetDescriber:         nil,
	},

	"AWS::Workspaces::Bundle": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::Workspaces::Bundle",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.WorkspacesBundle),
		GetDescriber:         nil,
	},

	"AWS::CloudTrail::Trail": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::CloudTrail::Trail",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.CloudTrailTrail),
		GetDescriber:         nil,
	},

	"AWS::DAX::Parameter": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::DAX::Parameter",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.DAXParameter),
		GetDescriber:         nil,
	},

	"AWS::ECR::Image": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::ECR::Image",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.ECRImage),
		GetDescriber:         nil,
	},

	"AWS::IAM::ServerCertificate": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::IAM::ServerCertificate",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.SequentialDescribeGlobal(describers.IAMServerCertificate),
		GetDescriber:         nil,
	},

	"AWS::Keyspaces::Keyspace": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::Keyspaces::Keyspace",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.KeyspacesKeyspace),
		GetDescriber:         nil,
	},

	"AWS::S3::AccessPoint": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::S3::AccessPoint",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.S3AccessPoint),
		GetDescriber:         nil,
	},

	"AWS::SageMaker::EndpointConfiguration": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::SageMaker::EndpointConfiguration",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.SageMakerEndpointConfiguration),
		GetDescriber:         nil,
	},

	"AWS::ElastiCache::ReservedCacheNode": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::ElastiCache::ReservedCacheNode",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.ElastiCacheReservedCacheNode),
		GetDescriber:         nil,
	},

	"AWS::EMR::InstanceFleet": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::EMR::InstanceFleet",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.EMRInstanceFleet),
		GetDescriber:         nil,
	},

	"AWS::Account::Account": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::Account::Account",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.SequentialDescribeGlobal(describers.IAMAccount),
		GetDescriber:         nil,
	},

	"AWS::EC2::VPCPeeringConnection": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::EC2::VPCPeeringConnection",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.EC2VPCPeeringConnection),
		GetDescriber:         nil,
	},

	"AWS::EKS::FargateProfile": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::EKS::FargateProfile",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.EKSFargateProfile),
		GetDescriber:         nil,
	},

	"AWS::IAM::AccountPasswordPolicy": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::IAM::AccountPasswordPolicy",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.SequentialDescribeGlobal(describers.IAMAccountPasswordPolicy),
		GetDescriber:         nil,
	},

	"AWS::CodePipeline::Pipeline": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::CodePipeline::Pipeline",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.CodePipelinePipeline),
		GetDescriber:         nil,
	},

	"AWS::DAX::Cluster": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::DAX::Cluster",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.DAXCluster),
		GetDescriber:         nil,
	},

	"AWS::DLM::LifecyclePolicy": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::DLM::LifecyclePolicy",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.DLMLifecyclePolicy),
		GetDescriber:         nil,
	},

	"AWS::OpsWorksCM::Server": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::OpsWorksCM::Server",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.OpsWorksCMServer),
		GetDescriber:         nil,
	},

	"AWS::AccessAnalyzer::Analyzer": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::AccessAnalyzer::Analyzer",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.AccessAnalyzerAnalyzer),
		GetDescriber:         nil,
	},

	"AWS::AccessAnalyzer::Finding": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::AccessAnalyzer::Finding",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.AccessAnalyzerAnalyzerFinding),
		GetDescriber:         nil,
	},

	"AWS::ElastiCache::SubnetGroup": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::ElastiCache::SubnetGroup",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.ElastiCacheSubnetGroup),
		GetDescriber:         nil,
	},

	"AWS::FSX::Volume": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::FSX::Volume",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.FSXVolume),
		GetDescriber:         nil,
	},

	"AWS::Amplify::App": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::Amplify::App",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.AmplifyApp),
		GetDescriber:         nil,
	},

	"AWS::CloudTrail::Query": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::CloudTrail::Query",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.CloudTrailQuery),
		GetDescriber:         nil,
	},

	"AWS::CostExplorer::ByAccountMonthly": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::CostExplorer::ByAccountMonthly",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.SequentialDescribeGlobal(describers.CostByAccountLastMonth),
		GetDescriber:         nil,
	},

	"AWS::ECR::PublicRegistry": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::ECR::PublicRegistry",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.ECRPublicRegistry),
		GetDescriber:         nil,
	},

	"AWS::EC2::NetworkInterface": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::EC2::NetworkInterface",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.EC2NetworkInterface),
		GetDescriber:         nil,
	},

	"AWS::EC2::VPNConnection": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::EC2::VPNConnection",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.EC2VPNConnection),
		GetDescriber:         nil,
	},

	"AWS::FSX::StorageVirtualMachine": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::FSX::StorageVirtualMachine",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.FSXStorageVirtualMachine),
		GetDescriber:         nil,
	},

	"AWS::ApiGateway::Authorizer": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::ApiGateway::Authorizer",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.ApiGatewayAuthorizer),
		GetDescriber:         nil,
	},

	"AWS::AppStream::Stack": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::AppStream::Stack",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.AppStreamStack),
		GetDescriber:         nil,
	},

	"AWS::Athena::WorkGroup": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::Athena::WorkGroup",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.AthenaWrokgroup),
		GetDescriber:         nil,
	},

	"AWS::Athena::QueryExecution": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::Athena::QueryExecution",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.AthenaQueryExecution),
		GetDescriber:         nil,
	},

	"AWS::AppStream::Image": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::AppStream::Image",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.AppStreamImage),
		GetDescriber:         nil,
	},

	"AWS::CloudWatch::Alarm": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::CloudWatch::Alarm",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.CloudWatchAlarm),
		GetDescriber:         nil,
	},

	"AWS::CloudWatch::LogSubscriptionFilter": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::CloudWatch::LogSubscriptionFilter",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.SequentialDescribeRegional(describers.CloudWatchLogsSubscriptionFilter),
		GetDescriber:         nil,
	},

	"AWS::CostExplorer::ByRecordTypeMonthly": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::CostExplorer::ByRecordTypeMonthly",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.SequentialDescribeGlobal(describers.CostByRecordTypeLastMonth),
		GetDescriber:         nil,
	},

	"AWS::RDS::DBCluster": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::RDS::DBCluster",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.RDSDBCluster),
		GetDescriber:         nil,
	},

	"AWS::RDS::DBClusterSnapshot": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::RDS::DBClusterSnapshot",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.RDSDBClusterSnapshot),
		GetDescriber:         nil,
	},

	"AWS::Backup::Framework": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::Backup::Framework",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.BackupFramework),
		GetDescriber:         nil,
	},

	"AWS::CodeBuild::SourceCredential": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::CodeBuild::SourceCredential",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.CodeBuildSourceCredential),
		GetDescriber:         nil,
	},

	"AWS::IAM::ServiceSpecificCredential": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::IAM::ServiceSpecificCredential",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.SequentialDescribeGlobal(describers.IAMServiceSpecificCredential),
		GetDescriber:         nil,
	},

	"AWS::EC2::CapacityReservationFleet": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::EC2::CapacityReservationFleet",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.EC2CapacityReservationFleet),
		GetDescriber:         nil,
	},

	"AWS::NetworkFirewall::Firewall": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::NetworkFirewall::Firewall",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.NetworkFirewallFirewall),
		GetDescriber:         nil,
	},

	"AWS::Workspaces::Workspace": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::Workspaces::Workspace",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.WorkspacesWorkspace),
		GetDescriber:         nil,
	},

	"AWS::ElasticSearch::Domain": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::ElasticSearch::Domain",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.ESDomain),
		GetDescriber:         nil,
	},

	"AWS::RDS::DBInstance": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::RDS::DBInstance",
		Tags:                 map[string][]string{
            "category": {"Database"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance/awsicons/master/svg-export/icons/Rds.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.RDSDBInstance),
		GetDescriber:         nil,
	},

	"AWS::RDS::DBInstanceAutomatedBackup": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::RDS::DBInstanceAutomatedBackup",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.RDSDBInstanceAutomatedBackup),
		GetDescriber:         nil,
	},

	"AWS::EFS::MountTarget": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::EFS::MountTarget",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.EFSMountTarget),
		GetDescriber:         nil,
	},

	"AWS::GlobalAccelerator::Listener": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::GlobalAccelerator::Listener",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.GlobalAcceleratorListener),
		GetDescriber:         nil,
	},

	"AWS::CostExplorer::ByUsageTypeDaily": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::CostExplorer::ByUsageTypeDaily",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.SequentialDescribeGlobal(describers.CostByServiceUsageLastDay),
		GetDescriber:         nil,
	},

	"AWS::EKS::Addon": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::EKS::Addon",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.EKSAddon),
		GetDescriber:         nil,
	},

	"AWS::CostExplorer::ByServiceMonthly": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::CostExplorer::ByServiceMonthly",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.SequentialDescribeGlobal(describers.CostByServiceLastMonth),
		GetDescriber:         nil,
	},

	"AWS::IAM::Policy": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::IAM::Policy",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.SequentialDescribeGlobal(describers.IAMPolicy),
		GetDescriber:         nil,
	},

	"AWS::Redshift::Cluster": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::Redshift::Cluster",
		Tags:                 map[string][]string{
            "category": {"Big Data"},
            "logo_uri": {""},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.RedshiftCluster),
		GetDescriber:         nil,
	},

	"AWS::WAFRegional::Rule": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::WAFRegional::Rule",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.WAFRegionalRule),
		GetDescriber:         nil,
	},

	"AWS::WAFRegional::RuleGroup": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::WAFRegional::RuleGroup",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.WAFRegionalRuleGroup),
		GetDescriber:         nil,
	},

	"AWS::Glue::DataCatalogEncryptionSettings": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::Glue::DataCatalogEncryptionSettings",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.GlueDataCatalogEncryptionSettings),
		GetDescriber:         nil,
	},

	"AWS::EC2::FlowLog": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::EC2::FlowLog",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.EC2FlowLog),
		GetDescriber:         nil,
	},

	"AWS::EC2::IpamPool": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::EC2::IpamPool",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.EC2IpamPool),
		GetDescriber:         nil,
	},

	"AWS::IAM::SamlProvider": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::IAM::SamlProvider",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.SequentialDescribeGlobal(describers.IAMSamlProvider),
		GetDescriber:         nil,
	},

	"AWS::Route53::HostedZone": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::Route53::HostedZone",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.SequentialDescribeGlobal(describers.Route53HostedZone),
		GetDescriber:         nil,
	},

	"AWS::Route53::QueryLog": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::Route53::QueryLog",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.SequentialDescribeGlobal(describers.Route53QueryLog),
		GetDescriber:         nil,
	},

	"AWS::EC2::PlacementGroup": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::EC2::PlacementGroup",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.EC2PlacementGroup),
		GetDescriber:         nil,
	},

	"AWS::FSX::Snapshot": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::FSX::Snapshot",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.FSXSnapshot),
		GetDescriber:         nil,
	},

	"AWS::KMS::Key": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::KMS::Key",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.KMSKey),
		GetDescriber:         nil,
	},

	"AWS::KMS::KeyRotation": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::KMS::KeyRotation",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.KMSKeyRotation),
		GetDescriber:         nil,
	},

	"AWS::EC2::Ipam": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::EC2::Ipam",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.EC2Ipam),
		GetDescriber:         nil,
	},

	"AWS::ElasticBeanstalk::Environment": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::ElasticBeanstalk::Environment",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.ElasticBeanstalkEnvironment),
		GetDescriber:         nil,
	},

	"AWS::ElasticBeanstalk::ApplicationVersion": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::ElasticBeanstalk::ApplicationVersion",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.ElasticBeanstalkApplicationVersion),
		GetDescriber:         nil,
	},

	"AWS::Lambda::FunctionVersion": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::Lambda::FunctionVersion",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.LambdaFunctionVersion),
		GetDescriber:         nil,
	},

	"AWS::Glue::DevEndpoint": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::Glue::DevEndpoint",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.GlueDevEndpoint),
		GetDescriber:         nil,
	},

	"AWS::Backup::RecoveryPoint": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::Backup::RecoveryPoint",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.BackupRecoveryPoint),
		GetDescriber:         nil,
	},

	"AWS::Backup::ReportPlan": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::Backup::ReportPlan",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.BackupReportPlan),
		GetDescriber:         nil,
	},

	"AWS::Backup::RegionSetting": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::Backup::RegionSetting",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.BackupRegionSetting),
		GetDescriber:         nil,
	},

	"AWS::DynamoDbStreams::Stream": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::DynamoDbStreams::Stream",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.DynamoDbStream),
		GetDescriber:         nil,
	},

	"AWS::EC2::EgressOnlyInternetGateway": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::EC2::EgressOnlyInternetGateway",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.EC2EgressOnlyInternetGateway),
		GetDescriber:         nil,
	},

	"AWS::CloudFront::Distribution": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::CloudFront::Distribution",
		Tags:                 map[string][]string{
            "category": {"Networking"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance/awsicons/master/svg-export/icons/CloudFront.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.SequentialDescribeGlobal(describers.CloudFrontDistribution),
		GetDescriber:         nil,
	},

	"AWS::CloudFront::StreamingDistribution": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::CloudFront::StreamingDistribution",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.SequentialDescribeGlobal(describers.CloudFrontStreamingDistribution),
		GetDescriber:         nil,
	},

	"AWS::Glue::Job": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::Glue::Job",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.GlueJob),
		GetDescriber:         nil,
	},

	"AWS::AppStream::Fleet": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::AppStream::Fleet",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.AppStreamFleet),
		GetDescriber:         nil,
	},

	"AWS::SES::ConfigurationSet": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::SES::ConfigurationSet",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.SESConfigurationSet),
		GetDescriber:         nil,
	},

	"AWS::IAM::User": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::IAM::User",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.SequentialDescribeGlobal(describers.IAMUser),
		GetDescriber:         nil,
	},

	"AWS::CloudFront::OriginRequestPolicy": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::CloudFront::OriginRequestPolicy",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.SequentialDescribeGlobal(describers.CloudFrontOriginRequestPolicy),
		GetDescriber:         nil,
	},

	"AWS::EC2::SecurityGroup": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::EC2::SecurityGroup",
		Tags:                 map[string][]string{
            "category": {"Networking"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.EC2SecurityGroup),
		GetDescriber:         nil,
	},

	"AWS::GuardDuty::IPSet": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::GuardDuty::IPSet",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.GuardDutyIPSet),
		GetDescriber:         nil,
	},

	"AWS::EKS::Cluster": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::EKS::Cluster",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.EKSCluster),
		GetDescriber:         nil,
	},

	"AWS::Grafana::Workspace": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::Grafana::Workspace",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.GrafanaWorkspace),
		GetDescriber:         nil,
	},

	"AWS::Glue::CatalogDatabase": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::Glue::CatalogDatabase",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.GlueCatalogDatabase),
		GetDescriber:         nil,
	},

	"AWS::Health::Event": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::Health::Event",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.HealthEvent),
		GetDescriber:         nil,
	},

	"AWS::Health::AffectedEntity": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::Health::AffectedEntity",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.HealthAffectedEntity),
		GetDescriber:         nil,
	},

	"AWS::CloudFormation::StackSet": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::CloudFormation::StackSet",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.CloudFormationStackSet),
		GetDescriber:         nil,
	},

	"AWS::EC2::AvailabilityZone": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::EC2::AvailabilityZone",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.EC2AvailabilityZone),
		GetDescriber:         nil,
	},

	"AWS::EC2::TransitGateway": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::EC2::TransitGateway",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.EC2TransitGateway),
		GetDescriber:         nil,
	},

	"AWS::ApiGateway::UsagePlan": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::ApiGateway::UsagePlan",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.ApiGatewayUsagePlan),
		GetDescriber:         nil,
	},

	"AWS::Inspector::Finding": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::Inspector::Finding",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.InspectorFinding),
		GetDescriber:         nil,
	},

	"AWS::EC2::Fleet": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::EC2::Fleet",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.EC2Fleet),
		GetDescriber:         nil,
	},

	"AWS::ElasticBeanstalk::Application": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::ElasticBeanstalk::Application",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.ElasticBeanstalkApplication),
		GetDescriber:         nil,
	},

	"AWS::ElasticLoadBalancingV2::LoadBalancer": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::ElasticLoadBalancingV2::LoadBalancer",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.ElasticLoadBalancingV2LoadBalancer),
		GetDescriber:         nil,
	},

	"AWS::OpenSearch::Domain": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::OpenSearch::Domain",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.OpenSearchDomain),
		GetDescriber:         nil,
	},

	"AWS::RDS::DBEventSubscription": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::RDS::DBEventSubscription",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.RDSDBEventSubscription),
		GetDescriber:         nil,
	},

	"AWS::RDS::DBEngineVersion": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::RDS::DBEngineVersion",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.RDSDBEngineVersion),
		GetDescriber:         nil,
	},

	"AWS::EC2::RegionalSettings": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::EC2::RegionalSettings",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.EC2RegionalSettings),
		GetDescriber:         nil,
	},

	"AWS::EC2::SecurityGroupRule": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::EC2::SecurityGroupRule",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.EC2SecurityGroupRule),
		GetDescriber:         nil,
	},

	"AWS::EC2::TransitGatewayAttachment": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::EC2::TransitGatewayAttachment",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.EC2TransitGatewayAttachment),
		GetDescriber:         nil,
	},

	"AWS::SES::Identity": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::SES::Identity",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.SESIdentity),
		GetDescriber:         nil,
	},

	"AWS::SESv2::EmailIdentities": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::SESv2::EmailIdentities",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.SESv2EmailIdentities),
		GetDescriber:         nil,
	},

	"AWS::WAF::Rule": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::WAF::Rule",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.WAFRule),
		GetDescriber:         nil,
	},

	"AWS::WAF::RuleGroup": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::WAF::RuleGroup",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.WAFRuleGroup),
		GetDescriber:         nil,
	},

	"AWS::WAF::RateBasedRule": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::WAF::RateBasedRule",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.WAFRateBasedRule),
		GetDescriber:         nil,
	},

	"AWS::WAF::WebACL": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::WAF::WebACL",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.WAFWebACL),
		GetDescriber:         nil,
	},

	"AWS::WAFRegional::WebACL": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::WAFRegional::WebACL",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.WAFRegionalWebACL),
		GetDescriber:         nil,
	},

	"AWS::WellArchitected::Workload": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::WellArchitected::Workload",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.WellArchitectedWorkload),
		GetDescriber:         nil,
	},

	"AWS::WellArchitected::Answer": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::WellArchitected::Answer",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.WellArchitectedAnswer),
		GetDescriber:         nil,
	},

	"AWS::WellArchitected::CheckDetail": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::WellArchitected::CheckDetail",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.WellArchitectedCheckDetail),
		GetDescriber:         nil,
	},

	"AWS::WellArchitected::CheckSummary": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::WellArchitected::CheckSummary",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.WellArchitectedCheckSummary),
		GetDescriber:         nil,
	},

	"AWS::WellArchitected::ConsolidatedReport": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::WellArchitected::ConsolidatedReport",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.WellArchitectedConsolidatedReport),
		GetDescriber:         nil,
	},

	"AWS::WellArchitected::Lens": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::WellArchitected::Lens",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.WellArchitectedLens),
		GetDescriber:         nil,
	},

	"AWS::WellArchitected::LensReview": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::WellArchitected::LensReview",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.WellArchitectedLensReview),
		GetDescriber:         nil,
	},

	"AWS::WellArchitected::LensReviewImprovement": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::WellArchitected::LensReviewImprovement",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.WellArchitectedLensReviewImprovement),
		GetDescriber:         nil,
	},

	"AWS::WellArchitected::LensReviewReport": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::WellArchitected::LensReviewReport",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.WellArchitectedLensReviewReport),
		GetDescriber:         nil,
	},

	"AWS::WellArchitected::LensShare": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::WellArchitected::LensShare",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.WellArchitectedLensShare),
		GetDescriber:         nil,
	},

	"AWS::WellArchitected::Milestone": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::WellArchitected::Milestone",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.WellArchitectedMilestone),
		GetDescriber:         nil,
	},

	"AWS::WellArchitected::Notification": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::WellArchitected::Notification",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.WellArchitectedNotification),
		GetDescriber:         nil,
	},

	"AWS::WellArchitected::ShareInvitation": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::WellArchitected::ShareInvitation",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.WellArchitectedShareInvitation),
		GetDescriber:         nil,
	},

	"AWS::WellArchitected::WorkloadShare": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::WellArchitected::WorkloadShare",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.WellArchitectedWorkloadShare),
		GetDescriber:         nil,
	},

	"AWS::AutoScaling::LaunchConfiguration": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::AutoScaling::LaunchConfiguration",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.AutoScalingLaunchConfiguration),
		GetDescriber:         nil,
	},

	"AWS::CloudTrail::EventDataStore": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::CloudTrail::EventDataStore",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.CloudTrailEventDataStore),
		GetDescriber:         nil,
	},

	"AWS::CodeDeploy::DeploymentGroup": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::CodeDeploy::DeploymentGroup",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.CodeDeployDeploymentGroup),
		GetDescriber:         nil,
	},

	"AWS::ImageBuilder::Image": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::ImageBuilder::Image",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.ImageBuilderImage),
		GetDescriber:         nil,
	},

	"AWS::Redshift::ClusterParameterGroup": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::Redshift::ClusterParameterGroup",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.RedshiftClusterParameterGroup),
		GetDescriber:         nil,
	},

	"AWS::Account::AlternateContact": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::Account::AlternateContact",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.SequentialDescribeGlobal(describers.AccountAlternateContact),
		GetDescriber:         nil,
	},

	"AWS::Inspector::AssessmentTarget": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::Inspector::AssessmentTarget",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.InspectorAssessmentTarget),
		GetDescriber:         nil,
	},

	"AWS::CloudFront::ResponseHeadersPolicy": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::CloudFront::ResponseHeadersPolicy",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.SequentialDescribeGlobal(describers.CloudFrontResponseHeadersPolicy),
		GetDescriber:         nil,
	},

	"AWS::EC2::Instance": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::EC2::Instance",
		Tags:                 map[string][]string{
            "category": {"Compute"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance/awsicons/master/svg-export/icons/Ec2Instance.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.EC2Instance),
		GetDescriber:         nil,
	},

	"AWS::EC2::InstanceAvailability": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::EC2::InstanceAvailability",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.EC2InstanceAvailability),
		GetDescriber:         nil,
	},

	"AWS::CostExplorer::ByRecordTypeDaily": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::CostExplorer::ByRecordTypeDaily",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.SequentialDescribeGlobal(describers.CostByRecordTypeLastDay),
		GetDescriber:         nil,
	},

	"AWS::EC2::ReservedInstances": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::EC2::ReservedInstances",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.EC2ReservedInstances),
		GetDescriber:         nil,
	},

	"AWS::ECR::Repository": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::ECR::Repository",
		Tags:                 map[string][]string{
            "category": {"Containers"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance/awsicons/master/svg-export/icons/ElasticContainerRegistry.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.ECRRepository),
		GetDescriber:         nil,
	},

	"AWS::ECR::Registry": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::ECR::Registry",
		Tags:                 map[string][]string{
            "category": {"Containers"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance/awsicons/master/svg-export/icons/ElasticContainerRegistry.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.ECRRegistry),
		GetDescriber:         nil,
	},

	"AWS::ECR::RegistryScanningConfiguration": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::ECR::RegistryScanningConfiguration",
		Tags:                 map[string][]string{
            "category": {"Containers"},
            "logo_uri": {},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.ECRRegistryScanningConfiguration),
		GetDescriber:         nil,
	},

	"AWS::ElasticLoadBalancingV2::Listener": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::ElasticLoadBalancingV2::Listener",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.ElasticLoadBalancingV2Listener),
		GetDescriber:         nil,
	},

	"AWS::IAM::Group": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::IAM::Group",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.SequentialDescribeGlobal(describers.IAMGroup),
		GetDescriber:         nil,
	},

	"AWS::IAM::OpenIdConnectProvider": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::IAM::OpenIdConnectProvider",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.SequentialDescribeGlobal(describers.IAMOpenIdConnectProvider),
		GetDescriber:         nil,
	},

	"AWS::Backup::Plan": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::Backup::Plan",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.BackupPlan),
		GetDescriber:         nil,
	},

	"AWS::Config::ConformancePack": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::Config::ConformancePack",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.ConfigConformancePack),
		GetDescriber:         nil,
	},

	"AWS::Config::RetentionConfiguration": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::Config::RetentionConfiguration",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.ConfigRetentionConfiguration),
		GetDescriber:         nil,
	},

	"AWS::CostExplorer::ByAccountDaily": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::CostExplorer::ByAccountDaily",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.SequentialDescribeGlobal(describers.CostByAccountLastDay),
		GetDescriber:         nil,
	},

	"AWS::Account::Contact": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::Account::Contact",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.SequentialDescribeGlobal(describers.AccountContact),
		GetDescriber:         nil,
	},

	"AWS::Glue::DataQualityRuleset": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::Glue::DataQualityRuleset",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.GlueDataQualityRuleset),
		GetDescriber:         nil,
	},

	"AWS::EventBridge::EventBus": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::EventBridge::EventBus",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.EventBridgeBus),
		GetDescriber:         nil,
	},

	"AWS::ApiGateway::Stage": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::ApiGateway::Stage",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.ApiGatewayStage),
		GetDescriber:         nil,
	},

	"AWS::ApiGatewayV2::Stage": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::ApiGatewayV2::Stage",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.ApiGatewayV2Stage),
		GetDescriber:         nil,
	},

	"AWS::DynamoDb::LocalSecondaryIndex": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::DynamoDb::LocalSecondaryIndex",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.DynamoDbLocalSecondaryIndex),
		GetDescriber:         nil,
	},

	"AWS::ResourceGroups::Groups": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::ResourceGroups::Groups",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.ResourceGroups),
		GetDescriber:         nil,
	},

	"AWS::Timestream::Database": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::Timestream::Database",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.TimestreamDatabase),
		GetDescriber:         nil,
	},

	"AWS::OpenSearchServerless::Collection": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::OpenSearchServerless::Collection",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.OpenSearchServerlessCollection),
		GetDescriber:         nil,
	},

	"AWS::EC2::ElasticIP": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::EC2::ElasticIP",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.EC2ElasticIP),
		GetDescriber:         nil,
	},

	"AWS::EC2::LocalGateway": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::EC2::LocalGateway",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.EC2LocalGateway),
		GetDescriber:         nil,
	},

	"AWS::EC2::Image": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::EC2::Image",
		Tags:                 map[string][]string{
            "category": {"Compute"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance/awsicons/master/svg-export/icons/Ec2AmiResource.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.EC2AMI),
		GetDescriber:         nil,
	},

	"AWS::EC2::Subnet": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::EC2::Subnet",
		Tags:                 map[string][]string{
            "category": {"Networking"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.EC2Subnet),
		GetDescriber:         nil,
	},

	"AWS::ECS::TaskSet": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::ECS::TaskSet",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.ECSTaskSet),
		GetDescriber:         nil,
	},

	"AWS::Kinesis::Stream": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::Kinesis::Stream",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.KinesisStream),
		GetDescriber:         nil,
	},

	"AWS::Kinesis::Consumer": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::Kinesis::Consumer",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.KinesisConsumer),
		GetDescriber:         nil,
	},

	"AWS::DocDB::Cluster": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::DocDB::Cluster",
		Tags:                 map[string][]string{
            "category": {"Database"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance/awsicons/master/svg-export/icons/DocumentDb.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.DocDBCluster),
		GetDescriber:         nil,
	},

	"AWS::DocDB::ClusterSnapshot": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::DocDB::ClusterSnapshot",
		Tags:                 map[string][]string{
            "category": {"Database"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance/awsicons/master/svg-export/icons/DocumentDb.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.DocDBClusterSnapshot),
		GetDescriber:         nil,
	},

	"AWS::DocDB::ClusterInstance": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::DocDB::ClusterInstance",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.DocDBClusterInstance),
		GetDescriber:         nil,
	},

	"AWS::ElastiCache::ReplicationGroup": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::ElastiCache::ReplicationGroup",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.ElastiCacheReplicationGroup),
		GetDescriber:         nil,
	},

	"AWS::GlobalAccelerator::Accelerator": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::GlobalAccelerator::Accelerator",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.GlobalAcceleratorAccelerator),
		GetDescriber:         nil,
	},

	"AWS::CloudWatch::Metric": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::CloudWatch::Metric",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.CloudWatchMetrics),
		GetDescriber:         nil,
	},

	"AWS::CostExplorer::ForcastMonthly": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::CostExplorer::ForcastMonthly",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.SequentialDescribeGlobal(describers.CostForecastMonthly),
		GetDescriber:         nil,
	},

	"AWS::EMR::InstanceGroup": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::EMR::InstanceGroup",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.EMRInstanceGroup),
		GetDescriber:         nil,
	},

	"AWS::EC2::ManagedPrefixList": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::EC2::ManagedPrefixList",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.EC2ManagedPrefixList),
		GetDescriber:         nil,
	},

	"AWS::EC2::ManagedPrefixListEntry": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::EC2::ManagedPrefixListEntry",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.EC2ManagedPrefixListEntry),
		GetDescriber:         nil,
	},

	"AWS::EC2::ClientVpnEndpoint": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::EC2::ClientVpnEndpoint",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.EC2ClientVpnEndpoint),
		GetDescriber:         nil,
	},

	"AWS::MWAA::Environment": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::MWAA::Environment",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.MWAAEnvironment),
		GetDescriber:         nil,
	},

	"AWS::CloudWatch::LogResourcePolicy": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::CloudWatch::LogResourcePolicy",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.CloudWatchLogsResourcePolicy),
		GetDescriber:         nil,
	},

	"AWS::CodeArtifact::Domain": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::CodeArtifact::Domain",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.CodeArtifactDomain),
		GetDescriber:         nil,
	},

	"AWS::CodeStar::Project": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::CodeStar::Project",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.CodeStarProject),
		GetDescriber:         nil,
	},

	"AWS::Neptune::Database": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::Neptune::Database",
		Tags:                 map[string][]string{
            "category": {"Database"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance/awsicons/master/svg-export/icons/Neptune.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.NeptuneDatabase),
		GetDescriber:         nil,
	},

	"AWS::Neptune::DBCluster": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::Neptune::DBCluster",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.NeptuneDatabaseCluster),
		GetDescriber:         nil,
	},

	"AWS::Neptune::DBClusterSnapshot": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::Neptune::DBClusterSnapshot",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.NeptuneDatabaseClusterSnapshot),
		GetDescriber:         nil,
	},

	"AWS::NetworkFirewall::FirewallPolicy": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::NetworkFirewall::FirewallPolicy",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.NetworkFirewallPolicy),
		GetDescriber:         nil,
	},

	"AWS::NetworkFirewall::RuleGroup": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::NetworkFirewall::RuleGroup",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.NetworkFirewallRuleGroup),
		GetDescriber:         nil,
	},

	"AWS::Oam::Link": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::Oam::Link",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.OAMLink),
		GetDescriber:         nil,
	},

	"AWS::Oam::Sink": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::Oam::Sink",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.OAMSink),
		GetDescriber:         nil,
	},

	"AWS::Organizations::Account": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::Organizations::Account",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.OrganizationsAccount),
		GetDescriber:         nil,
	},

	"AWS::Pinpoint::App": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::Pinpoint::App",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.PinpointApp),
		GetDescriber:         nil,
	},

	"AWS::Pipes::Pipe": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::Pipes::Pipe",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.PipesPipe),
		GetDescriber:         nil,
	},

	"AWS::RDS::DBClusterParameterGroup": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::RDS::DBClusterParameterGroup",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.RDSDBClusterParameterGroup),
		GetDescriber:         nil,
	},

	"AWS::RDS::OptionGroup": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::RDS::OptionGroup",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.RDSOptionGroup),
		GetDescriber:         nil,
	},

	"AWS::RDS::DBParameterGroup": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::RDS::DBParameterGroup",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.RDSDBParameterGroup),
		GetDescriber:         nil,
	},

	"AWS::RDS::DBProxy": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::RDS::DBProxy",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.RDSDBProxy),
		GetDescriber:         nil,
	},

	"AWS::RDS::DBSubnetGroup": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::RDS::DBSubnetGroup",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.RDSDBSubnetGroup),
		GetDescriber:         nil,
	},

	"AWS::RDS::DBRecommendation": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::RDS::DBRecommendation",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.RDSDBRecommendation),
		GetDescriber:         nil,
	},

	"AWS::Redshift::EventSubscription": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::Redshift::EventSubscription",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.RedshiftEventSubscription),
		GetDescriber:         nil,
	},

	"AWS::RedshiftServerless::Workgroup": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::RedshiftServerless::Workgroup",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.RedshiftServerlessWorkgroup),
		GetDescriber:         nil,
	},

	"AWS::ResourceExplorer2::Index": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::ResourceExplorer2::Index",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.ResourceExplorerIndex),
		GetDescriber:         nil,
	},

	"AWS::Route53::HealthCheck": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::Route53::HealthCheck",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.Route53HealthCheck),
		GetDescriber:         nil,
	},

	"AWS::Route53Resolver::ResolverRule": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::Route53Resolver::ResolverRule",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.Route53ResolverResolverRule),
		GetDescriber:         nil,
	},

	"AWS::Route53Resolver::QueryLogConfig": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::Route53Resolver::QueryLogConfig",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.Route53ResolverQueryLogConfig),
		GetDescriber:         nil,
	},

	"AWS::SageMaker::App": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::SageMaker::App",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.SageMakerApp),
		GetDescriber:         nil,
	},

	"AWS::SageMaker::Domain": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::SageMaker::Domain",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.SageMakerDomain),
		GetDescriber:         nil,
	},

	"AWS::StepFunctions::StateMachine": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::StepFunctions::StateMachine",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.StepFunctionsStateMachine),
		GetDescriber:         nil,
	},

	"AWS::StepFunctions::StateMachineExecution": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::StepFunctions::StateMachineExecution",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.StepFunctionsStateMachineExecution),
		GetDescriber:         nil,
	},

	"AWS::StepFunctions::StateMachineExecutionHistories": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::StepFunctions::StateMachineExecutionHistories",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.StepFunctionsStateMachineExecutionHistories),
		GetDescriber:         nil,
	},

	"AWS::SimSpaceWeaver::Simulation": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::SimSpaceWeaver::Simulation",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.SimSpaceWeaverSimulation),
		GetDescriber:         nil,
	},

	"AWS::SSM::Association": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::SSM::Association",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.SSMAssociation),
		GetDescriber:         nil,
	},

	"AWS::SSM::Document": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::SSM::Document",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.SSMDocument),
		GetDescriber:         nil,
	},

	"AWS::SSM::DocumentPermission": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::SSM::DocumentPermission",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.SSMDocumentPermission),
		GetDescriber:         nil,
	},

	"AWS::EC2::CustomerGateway": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::EC2::CustomerGateway",
		Tags:                 map[string][]string{
            "category": {"Networking"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance/awsicons/master/svg-export/icons/VpcCustomerGateway.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.EC2CustomerGateway),
		GetDescriber:         nil,
	},

	"AWS::EC2::VerifiedAccessInstance": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::EC2::VerifiedAccessInstance",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.EC2VerifiedAccessInstance),
		GetDescriber:         nil,
	},

	"AWS::EC2::VerifiedAccessEndpoint": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::EC2::VerifiedAccessEndpoint",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.EC2VerifiedAccessEndpoint),
		GetDescriber:         nil,
	},

	"AWS::EC2::VerifiedAccessGroup": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::EC2::VerifiedAccessGroup",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.EC2VerifiedAccessGroup),
		GetDescriber:         nil,
	},

	"AWS::EC2::VerifiedAccessTrustProvider": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::EC2::VerifiedAccessTrustProvider",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.EC2VerifiedAccessTrustProvider),
		GetDescriber:         nil,
	},

	"AWS::EC2::VPNGateway": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::EC2::VPNGateway",
		Tags:                 map[string][]string{
            "category": {"Networking"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance/awsicons/master/svg-export/icons/VpcVpnGateway.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.EC2VPNGateway),
		GetDescriber:         nil,
	},

	"AWS::WAFv2::IPSet": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::WAFv2::IPSet",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.WAFv2IPSet),
		GetDescriber:         nil,
	},

	"AWS::WAFv2::RegexPatternSet": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::WAFv2::RegexPatternSet",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.WAFv2RegexPatternSet),
		GetDescriber:         nil,
	},

	"AWS::WAFv2::RuleGroup": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::WAFv2::RuleGroup",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.WAFv2RuleGroup),
		GetDescriber:         nil,
	},

	"AWS::EC2::TransitGatewayRoute": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::EC2::TransitGatewayRoute",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.EC2TransitGatewayRoute),
		GetDescriber:         nil,
	},

	"AWS::GuardDuty::Filter": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::GuardDuty::Filter",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.GuardDutyFilter),
		GetDescriber:         nil,
	},

	"AWS::ECS::TaskDefinition": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::ECS::TaskDefinition",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.ECSTaskDefinition),
		GetDescriber:         nil,
	},

	"AWS::GuardDuty::ThreatIntelSet": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::GuardDuty::ThreatIntelSet",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.GuardDutyThreatIntelSet),
		GetDescriber:         nil,
	},

	"AWS::ApiGatewayV2::DomainName": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::ApiGatewayV2::DomainName",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.ApiGatewayV2DomainName),
		GetDescriber:         nil,
	},

	"AWS::ApiGateway::DomainName": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::ApiGateway::DomainName",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.ApiGatewayDomainName),
		GetDescriber:         nil,
	},

	"AWS::ApiGatewayV2::Route": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::ApiGatewayV2::Route",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.ApiGatewayV2Route),
		GetDescriber:         nil,
	},

	"AWS::MQ::Broker": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::MQ::Broker",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.MQBroker),
		GetDescriber:         nil,
	},

	"AWS::ACMPCA::CertificateAuthority": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::ACMPCA::CertificateAuthority",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.ACMPCACertificateAuthority),
		GetDescriber:         nil,
	},

	"AWS::CloudFormation::Stack": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::CloudFormation::Stack",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.CloudFormationStack),
		GetDescriber:         nil,
	},

	"AWS::CloudFormation::StackResource": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::CloudFormation::StackResource",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.CloudFormationStackResource),
		GetDescriber:         nil,
	},

	"AWS::DirectConnect::Connection": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::DirectConnect::Connection",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.DirectConnectConnection),
		GetDescriber:         nil,
	},

	"AWS::FSX::FileSystem": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::FSX::FileSystem",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.FSXFileSystem),
		GetDescriber:         nil,
	},

	"AWS::Glue::SecurityConfiguration": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::Glue::SecurityConfiguration",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.GlueSecurityConfiguration),
		GetDescriber:         nil,
	},

	"AWS::Inspector::AssessmentRun": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::Inspector::AssessmentRun",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.InspectorAssessmentRun),
		GetDescriber:         nil,
	},

	"AWS::Inspector2::Coverage": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::Inspector2::Coverage",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.Inspector2Coverage),
		GetDescriber:         nil,
	},

	"AWS::Inspector2::CoverageStatistics": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::Inspector2::CoverageStatistics",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.Inspector2CoverageStatistic),
		GetDescriber:         nil,
	},

	"AWS::Inspector2::Member": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::Inspector2::Member",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.Inspector2CoverageMember),
		GetDescriber:         nil,
	},

	"AWS::Inspector2::Finding": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::Inspector2::Finding",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.Inspector2Finding),
		GetDescriber:         nil,
	},

	"AWS::Config::ConfigurationRecorder": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::Config::ConfigurationRecorder",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.ConfigConfigurationRecorder),
		GetDescriber:         nil,
	},

	"AWS::EC2::NatGateway": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::EC2::NatGateway",
		Tags:                 map[string][]string{
            "category": {"Networking"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance/awsicons/master/svg-export/icons/VpcNatGateway.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.EC2NatGateway),
		GetDescriber:         nil,
	},

	"AWS::ECR::PublicRepository": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::ECR::PublicRepository",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.ECRPublicRepository),
		GetDescriber:         nil,
	},

	"AWS::ECS::Cluster": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::ECS::Cluster",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.ECSCluster),
		GetDescriber:         nil,
	},

	"AWS::ElasticLoadBalancingV2::TargetGroup": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::ElasticLoadBalancingV2::TargetGroup",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.ElasticLoadBalancingV2TargetGroup),
		GetDescriber:         nil,
	},

	"AWS::CloudFront::CachePolicy": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::CloudFront::CachePolicy",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.SequentialDescribeGlobal(describers.CloudFrontCachePolicy),
		GetDescriber:         nil,
	},

	"AWS::CodeArtifact::Repository": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::CodeArtifact::Repository",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.CodeArtifactRepository),
		GetDescriber:         nil,
	},

	"AWS::AMP::Workspace": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::AMP::Workspace",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.AMPWorkspace),
		GetDescriber:         nil,
	},

	"AWS::EC2::CapacityReservation": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::EC2::CapacityReservation",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.EC2CapacityReservation),
		GetDescriber:         nil,
	},

	"AWS::SageMaker::NotebookInstance": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::SageMaker::NotebookInstance",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.SageMakerNotebookInstance),
		GetDescriber:         nil,
	},

	"AWS::IAM::AccessAdvisor": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::IAM::AccessAdvisor",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.IAMAccessAdvisor),
		GetDescriber:         nil,
	},

	"AWS::EC2::VolumeSnapshot": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::EC2::VolumeSnapshot",
		Tags:                 map[string][]string{
            "category": {"Storage"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance/awsicons/master/svg-export/icons/ElasticBlockStoreSnapshot.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.EC2VolumeSnapshot),
		GetDescriber:         nil,
	},

	"AWS::EC2::Region": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::EC2::Region",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.EC2Region),
		GetDescriber:         nil,
	},

	"AWS::Keyspaces::Table": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::Keyspaces::Table",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.KeyspacesTable),
		GetDescriber:         nil,
	},

	"AWS::Config::AggregationAuthorization": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::Config::AggregationAuthorization",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.ConfigAggregateAuthorization),
		GetDescriber:         nil,
	},

	"AWS::DAX::SubnetGroup": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::DAX::SubnetGroup",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.DAXSubnetGroup),
		GetDescriber:         nil,
	},

	"AWS::DynamoDb::GlobalTable": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::DynamoDb::GlobalTable",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.DynamoDbGlobalTable),
		GetDescriber:         nil,
	},

	"AWS::ElasticLoadBalancing::LoadBalancer": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::ElasticLoadBalancing::LoadBalancer",
		Tags:                 map[string][]string{
            "category": {"Networking"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance/awsicons/master/svg-export/icons/ElasticLoadBalancingClassicLoadBalancer.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.ElasticLoadBalancingLoadBalancer),
		GetDescriber:         nil,
	},

	"AWS::AppStream::Application": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::AppStream::Application",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.AppStreamApplication),
		GetDescriber:         nil,
	},

	"AWS::RedshiftServerless::Namespace": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::RedshiftServerless::Namespace",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.RedshiftServerlessNamespace),
		GetDescriber:         nil,
	},

	"AWS::CloudFront::OriginAccessIdentity": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::CloudFront::OriginAccessIdentity",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.SequentialDescribeGlobal(describers.CloudFrontOriginAccessIdentity),
		GetDescriber:         nil,
	},

	"AWS::EC2::Host": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::EC2::Host",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.EC2Host),
		GetDescriber:         nil,
	},

	"AWS::EC2::VPC": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::EC2::VPC",
		Tags:                 map[string][]string{
            "category": {"Networking"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance/awsicons/master/svg-export/icons/Vpc.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.EC2VPC),
		GetDescriber:         nil,
	},

	"AWS::EC2::TransitGatewayRouteTable": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::EC2::TransitGatewayRouteTable",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.EC2TransitGatewayRouteTable),
		GetDescriber:         nil,
	},

	"AWS::EKS::Nodegroup": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::EKS::Nodegroup",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.EKSNodegroup),
		GetDescriber:         nil,
	},

	"AWS::Backup::Selection": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::Backup::Selection",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.BackupSelection),
		GetDescriber:         nil,
	},

	"AWS::CloudTrail::Import": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::CloudTrail::Import",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.CloudTrailImport),
		GetDescriber:         nil,
	},

	"AWS::CostExplorer::ByServiceDaily": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::CostExplorer::ByServiceDaily",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.SequentialDescribeGlobal(describers.CostByServiceLastDay),
		GetDescriber:         nil,
	},

	"AWS::ElasticLoadBalancingV2::SslPolicy": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::ElasticLoadBalancingV2::SslPolicy",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.ElasticLoadBalancingV2SslPolicy),
		GetDescriber:         nil,
	},

	"AWS::GuardDuty::Finding": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::GuardDuty::Finding",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.GuardDutyFinding),
		GetDescriber:         nil,
	},

	"AWS::EC2::DHCPOptions": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::EC2::DHCPOptions",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.EC2DHCPOptions),
		GetDescriber:         nil,
	},

	"AWS::EC2::InstanceType": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::EC2::InstanceType",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.EC2InstanceType),
		GetDescriber:         nil,
	},

	"AWS::Batch::ComputeEnvironment": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::Batch::ComputeEnvironment",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.BatchComputeEnvironment),
		GetDescriber:         nil,
	},

	"AWS::DMS::ReplicationInstance": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::DMS::ReplicationInstance",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.DMSReplicationInstance),
		GetDescriber:         nil,
	},

	"AWS::DMS::Endpoint": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::DMS::Endpoint",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.DMSEndpoint),
		GetDescriber:         nil,
	},

	"AWS::DMS::ReplicationTask": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::DMS::ReplicationTask",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.DMSReplicationTask),
		GetDescriber:         nil,
	},

	"AWS::DynamoDb::Table": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::DynamoDb::Table",
		Tags:                 map[string][]string{
            "category": {"Database"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance/awsicons/master/svg-export/icons/DynamoDbTable.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.DynamoDbTable),
		GetDescriber:         nil,
	},

	"AWS::Shield::ProtectionGroup": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::Shield::ProtectionGroup",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.SequentialDescribeGlobal(describers.ShieldProtectionGroup),
		GetDescriber:         nil,
	},

	"AWS::Firehose::DeliveryStream": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::Firehose::DeliveryStream",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.FirehoseDeliveryStream),
		GetDescriber:         nil,
	},

	"AWS::KinesisVideo::Stream": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::KinesisVideo::Stream",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.KinesisVideoStream),
		GetDescriber:         nil,
	},

	"AWS::KMS::Alias": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::KMS::Alias",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.KMSAlias),
		GetDescriber:         nil,
	},

	"AWS::Lambda::Alias": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::Lambda::Alias",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.LambdaAlias),
		GetDescriber:         nil,
	},

	"AWS::Lambda::LambdaLayer": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::Lambda::LambdaLayer",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.LambdaLayer),
		GetDescriber:         nil,
	},

	"AWS::Lambda::LayerVersion": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::Lambda::LayerVersion",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.LambdaLayerVersion),
		GetDescriber:         nil,
	},

	"AWS::Lightsail::Instance": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::Lightsail::Instance",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.LightsailInstance),
		GetDescriber:         nil,
	},

	"AWS::Macie2::ClassificationJob": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::Macie2::ClassificationJob",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.Macie2ClassificationJob),
		GetDescriber:         nil,
	},

	"AWS::MediaStore::Container": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::MediaStore::Container",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.MediaStoreContainer),
		GetDescriber:         nil,
	},

	"AWS::Mgn::Application": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::Mgn::Application",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.MGNApplication),
		GetDescriber:         nil,
	},

	"AWS::ResourceExplorer2::SupportedResourceType": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::ResourceExplorer2::SupportedResourceType",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.ResourceExplorer2SupportedResourceType),
		GetDescriber:         nil,
	},

	"AWS::Route53Resolver::ResolverEndpoint": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::Route53Resolver::ResolverEndpoint",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.Route53ResolverResolverEndpoint),
		GetDescriber:         nil,
	},

	"AWS::Route53Domains::Domain": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::Route53Domains::Domain",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.SequentialDescribeGlobal(describers.Route53Domain),
		GetDescriber:         nil,
	},

	"AWS::Route53::Record": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::Route53::Record",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.SequentialDescribeGlobal(describers.Route53Record),
		GetDescriber:         nil,
	},

	"AWS::Route53::TrafficPolicy": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::Route53::TrafficPolicy",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.SequentialDescribeGlobal(describers.Route53TrafficPolicy),
		GetDescriber:         nil,
	},

	"AWS::Route53::TrafficPolicyInstance": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::Route53::TrafficPolicyInstance",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.SequentialDescribeGlobal(describers.Route53TrafficPolicyInstance),
		GetDescriber:         nil,
	},

	"AWS::SageMaker::Model": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::SageMaker::Model",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.SageMakerModel),
		GetDescriber:         nil,
	},

	"AWS::SageMaker::TrainingJob": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::SageMaker::TrainingJob",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.SageMakerTrainingJob),
		GetDescriber:         nil,
	},

	"AWS::SecurityHub::ActionTarget": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::SecurityHub::ActionTarget",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.SecurityHubActionTarget),
		GetDescriber:         nil,
	},

	"AWS::SecurityHub::Finding": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::SecurityHub::Finding",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.SecurityHubFinding),
		GetDescriber:         nil,
	},

	"AWS::SecurityHub::FindingAggregator": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::SecurityHub::FindingAggregator",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.SecurityHubFindingAggregator),
		GetDescriber:         nil,
	},

	"AWS::SecurityHub::Insight": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::SecurityHub::Insight",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.SecurityHubInsight),
		GetDescriber:         nil,
	},

	"AWS::SecurityHub::Member": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::SecurityHub::Member",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.SecurityHubMember),
		GetDescriber:         nil,
	},

	"AWS::SecurityHub::Product": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::SecurityHub::Product",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.SecurityHubProduct),
		GetDescriber:         nil,
	},

	"AWS::SecurityHub::StandardsControl": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::SecurityHub::StandardsControl",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.SecurityHubStandardsControl),
		GetDescriber:         nil,
	},

	"AWS::SecurityHub::StandardsSubscription": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::SecurityHub::StandardsSubscription",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.SecurityHubStandardsSubscription),
		GetDescriber:         nil,
	},

	"AWS::SecurityLake::DataLake": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::SecurityLake::DataLake",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.SecurityLakeDataLake),
		GetDescriber:         nil,
	},

	"AWS::SecurityLake::Subscriber": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::SecurityLake::Subscriber",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.SecurityLakeSubscriber),
		GetDescriber:         nil,
	},

	"AWS::Ram::PrincipalAssociation": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::Ram::PrincipalAssociation",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.RamPrincipalAssociation),
		GetDescriber:         nil,
	},

	"AWS::Ram::ResourceAssociation": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::Ram::ResourceAssociation",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.RamResourceAssociation),
		GetDescriber:         nil,
	},

	"AWS::RDS::ReservedDBInstance": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::RDS::ReservedDBInstance",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.RDSReservedDBInstance),
		GetDescriber:         nil,
	},

	"AWS::Redshift::SubnetGroup": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::Redshift::SubnetGroup",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.RedshiftSubnetGroup),
		GetDescriber:         nil,
	},

	"AWS::SeverlessApplicationRepository::Application": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::SeverlessApplicationRepository::Application",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.ServerlessApplicationRepositoryApplication),
		GetDescriber:         nil,
	},

	"AWS::AuditManager::Framework": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::AuditManager::Framework",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.AuditManagerFramework),
		GetDescriber:         nil,
	},

	"AWS::AuditManager::EvidenceFolder": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::AuditManager::EvidenceFolder",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.AuditManagerEvidenceFolder),
		GetDescriber:         nil,
	},

	"AWS::AuditManager::Evidence": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::AuditManager::Evidence",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.AuditManagerEvidence),
		GetDescriber:         nil,
	},

	"AWS::AuditManager::Control": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::AuditManager::Control",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.AuditManagerControl),
		GetDescriber:         nil,
	},

	"AWS::AuditManager::Assessment": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::AuditManager::Assessment",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.AuditManagerAssessment),
		GetDescriber:         nil,
	},

	"AWS::CloudTrail::TrailEvent": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::CloudTrail::TrailEvent",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.CloudTrailTrailEvent),
		GetDescriber:         nil,
	},

	"AWS::CloudWatch::LogEvent": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::CloudWatch::LogEvent",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.CloudWatchLogEvent),
		GetDescriber:         nil,
	},

	"AWS::Logs::MetricFilter": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::Logs::MetricFilter",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.CloudWatchLogsMetricFilter),
		GetDescriber:         nil,
	},

	"AWS::CloudWatch::LogStream": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::CloudWatch::LogStream",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.CloudWatchLogStream),
		GetDescriber:         nil,
	},

	"AWS::CostExplorer::ByUsageTypeMonthly": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::CostExplorer::ByUsageTypeMonthly",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.SequentialDescribeGlobal(describers.CostByServiceUsageLastMonth),
		GetDescriber:         nil,
	},

	"AWS::ServiceQuotas::ServiceQuotaChangeRequest": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::ServiceQuotas::ServiceQuotaChangeRequest",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.ServiceQuotasServiceQuotaChangeRequest),
		GetDescriber:         nil,
	},

	"AWS::ServiceQuotas::Service": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::ServiceQuotas::Service",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.ServiceQuotasService),
		GetDescriber:         nil,
	},

	"AWS::EC2::VPCEndpointService": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::EC2::VPCEndpointService",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.EC2VPCEndpointService),
		GetDescriber:         nil,
	},

	"AWS::EC2::LaunchTemplate": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::EC2::LaunchTemplate",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.EC2LaunchTemplate),
		GetDescriber:         nil,
	},

	"AWS::EC2::LaunchTemplateVersion": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::EC2::LaunchTemplateVersion",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.EC2LaunchTemplateVersion),
		GetDescriber:         nil,
	},

	"AWS::SNS::Subscription": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::SNS::Subscription",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.SNSSubscription),
		GetDescriber:         nil,
	},

	"AWS::S3::AccountSetting": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::S3::AccountSetting",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.SequentialDescribeGlobal(describers.S3AccountSetting),
		GetDescriber:         nil,
	},

	"AWS::SSM::ManagedInstanceCompliance": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::SSM::ManagedInstanceCompliance",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.SSMManagedInstanceCompliance),
		GetDescriber:         nil,
	},

	"AWS::SSM::ManagedInstancePatchState": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::SSM::ManagedInstancePatchState",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.SSMManagedInstancePatchState),
		GetDescriber:         nil,
	},

	"AWS::SSOAdmin::AccountAssignment": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::SSOAdmin::AccountAssignment",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.SSOAdminAccountAssignment),
		GetDescriber:         nil,
	},

	"AWS::SSOAdmin::UserEffectiveAccess": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::SSOAdmin::UserEffectiveAccess",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.UserEffectiveAccess),
		GetDescriber:         nil,
	},

	"AWS::SSOAdmin::Instance": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::SSOAdmin::Instance",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.SSOAdminInstance),
		GetDescriber:         nil,
	},

	"AWS::SSOAdmin::PermissionSet": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::SSOAdmin::PermissionSet",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.SSOAdminPermissionSet),
		GetDescriber:         nil,
	},

	"AWS::SSOAdmin::AttachedManagedPolicy": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::SSOAdmin::AttachedManagedPolicy",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.SSOAdminManagedPolicyAttachment),
		GetDescriber:         nil,
	},

	"AWS::ServiceDiscovery::Service": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::ServiceDiscovery::Service",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.ServiceDiscoveryService),
		GetDescriber:         nil,
	},

	"AWS::ServiceDiscovery::Namespace": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::ServiceDiscovery::Namespace",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.ServiceDiscoveryNamespace),
		GetDescriber:         nil,
	},

	"AWS::ServiceDiscovery::Instance": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::ServiceDiscovery::Instance",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.ServiceDiscoveryInstance),
		GetDescriber:         nil,
	},

	"AWS::ServiceCatalog::Portfolio": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::ServiceCatalog::Portfolio",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.ServiceCatalogPortfolio),
		GetDescriber:         nil,
	},

	"AWS::ServiceCatalog::Product": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::ServiceCatalog::Product",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.ServiceCatalogProduct),
		GetDescriber:         nil,
	},

	"AWS::IdentityStore::User": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::IdentityStore::User",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.IdentityStoreUser),
		GetDescriber:         nil,
	},

	"AWS::IdentityStore::Group": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::IdentityStore::Group",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.IdentityStoreGroup),
		GetDescriber:         nil,
	},

	"AWS::IdentityStore::GroupMembership": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "AWS::IdentityStore::GroupMembership",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.ParallelDescribeRegional(describers.IdentityStoreGroupMembership),
		GetDescriber:         nil,
	},
}


var ResourceTypeConfigs = map[string]*interfaces.ResourceTypeConfiguration{

	"AWS::Redshift::Snapshot": {
		Name:         "AWS::Redshift::Snapshot",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::IAM::AccountSummary": {
		Name:         "AWS::IAM::AccountSummary",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::Glacier::Vault": {
		Name:         "AWS::Glacier::Vault",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::Organizations::Organization": {
		Name:         "AWS::Organizations::Organization",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::Organizations::Policy": {
		Name:         "AWS::Organizations::Policy",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::Organizations::PolicyTarget": {
		Name:         "AWS::Organizations::PolicyTarget",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::Organizations::OrganizationalUnit": {
		Name:         "AWS::Organizations::OrganizationalUnit",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::Organizations::Root": {
		Name:         "AWS::Organizations::Root",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::CloudSearch::Domain": {
		Name:         "AWS::CloudSearch::Domain",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::DynamoDb::GlobalSecondaryIndex": {
		Name:         "AWS::DynamoDb::GlobalSecondaryIndex",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::EC2::RouteTable": {
		Name:         "AWS::EC2::RouteTable",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::SecurityHub::Hub": {
		Name:         "AWS::SecurityHub::Hub",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::StorageGateway::StorageGateway": {
		Name:         "AWS::StorageGateway::StorageGateway",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::Inspector::AssessmentTemplate": {
		Name:         "AWS::Inspector::AssessmentTemplate",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::ElasticLoadBalancingV2::ListenerRule": {
		Name:         "AWS::ElasticLoadBalancingV2::ListenerRule",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::IAM::Role": {
		Name:         "AWS::IAM::Role",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::Backup::ProtectedResource": {
		Name:         "AWS::Backup::ProtectedResource",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::CodeCommit::Repository": {
		Name:         "AWS::CodeCommit::Repository",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::EC2::VPCEndpoint": {
		Name:         "AWS::EC2::VPCEndpoint",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::EventBridge::EventRule": {
		Name:         "AWS::EventBridge::EventRule",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::CloudFront::OriginAccessControl": {
		Name:         "AWS::CloudFront::OriginAccessControl",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::CodeBuild::Project": {
		Name:         "AWS::CodeBuild::Project",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::CodeBuild::Build": {
		Name:         "AWS::CodeBuild::Build",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::ElastiCache::ParameterGroup": {
		Name:         "AWS::ElastiCache::ParameterGroup",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::MemoryDb::Cluster": {
		Name:         "AWS::MemoryDb::Cluster",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::Glue::Crawler": {
		Name:         "AWS::Glue::Crawler",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::DirectConnect::Gateway": {
		Name:         "AWS::DirectConnect::Gateway",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::DynamoDb::BackUp": {
		Name:         "AWS::DynamoDb::BackUp",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::EC2::EIP": {
		Name:         "AWS::EC2::EIP",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::EC2::InternetGateway": {
		Name:         "AWS::EC2::InternetGateway",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::GuardDuty::PublishingDestination": {
		Name:         "AWS::GuardDuty::PublishingDestination",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::KinesisAnalyticsV2::Application": {
		Name:         "AWS::KinesisAnalyticsV2::Application",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::EMR::Instance": {
		Name:         "AWS::EMR::Instance",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::EMR::BlockPublicAccessConfiguration": {
		Name:         "AWS::EMR::BlockPublicAccessConfiguration",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::ApiGateway::RestApi": {
		Name:         "AWS::ApiGateway::RestApi",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::ApiGatewayV2::Integration": {
		Name:         "AWS::ApiGatewayV2::Integration",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::AutoScaling::AutoScalingGroup": {
		Name:         "AWS::AutoScaling::AutoScalingGroup",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::DynamoDb::TableExport": {
		Name:         "AWS::DynamoDb::TableExport",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::EC2::KeyPair": {
		Name:         "AWS::EC2::KeyPair",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::EFS::FileSystem": {
		Name:         "AWS::EFS::FileSystem",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::Kafka::Cluster": {
		Name:         "AWS::Kafka::Cluster",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::SecretsManager::Secret": {
		Name:         "AWS::SecretsManager::Secret",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::Backup::LegalHold": {
		Name:         "AWS::Backup::LegalHold",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::CloudFront::Function": {
		Name:         "AWS::CloudFront::Function",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::GlobalAccelerator::EndpointGroup": {
		Name:         "AWS::GlobalAccelerator::EndpointGroup",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::DAX::ParameterGroup": {
		Name:         "AWS::DAX::ParameterGroup",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::SQS::Queue": {
		Name:         "AWS::SQS::Queue",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::Config::Rule": {
		Name:         "AWS::Config::Rule",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::GuardDuty::Member": {
		Name:         "AWS::GuardDuty::Member",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::Inspector::Exclusion": {
		Name:         "AWS::Inspector::Exclusion",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::DirectoryService::Directory": {
		Name:         "AWS::DirectoryService::Directory",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::DirectoryService::Certificate": {
		Name:         "AWS::DirectoryService::Certificate",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::DirectoryService::LogSubscription": {
		Name:         "AWS::DirectoryService::LogSubscription",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::EFS::AccessPoint": {
		Name:         "AWS::EFS::AccessPoint",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::IAM::PolicyAttachment": {
		Name:         "AWS::IAM::PolicyAttachment",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::IAM::CredentialReport": {
		Name:         "AWS::IAM::CredentialReport",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::RDS::GlobalCluster": {
		Name:         "AWS::RDS::GlobalCluster",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::CostExplorer::ForcastDaily": {
		Name:         "AWS::CostExplorer::ForcastDaily",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::GuardDuty::Detector": {
		Name:         "AWS::GuardDuty::Detector",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::SNS::Topic": {
		Name:         "AWS::SNS::Topic",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::AppConfig::Application": {
		Name:         "AWS::AppConfig::Application",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::Batch::Job": {
		Name:         "AWS::Batch::Job",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::Batch::JobQueue": {
		Name:         "AWS::Batch::JobQueue",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::ECS::Service": {
		Name:         "AWS::ECS::Service",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::FSX::Task": {
		Name:         "AWS::FSX::Task",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::IAM::VirtualMFADevice": {
		Name:         "AWS::IAM::VirtualMFADevice",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::WAFv2::WebACL": {
		Name:         "AWS::WAFv2::WebACL",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::ApplicationAutoScaling::Target": {
		Name:         "AWS::ApplicationAutoScaling::Target",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::ApplicationAutoScaling::Policy": {
		Name:         "AWS::ApplicationAutoScaling::Policy",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::Backup::Vault": {
		Name:         "AWS::Backup::Vault",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::ElastiCache::Cluster": {
		Name:         "AWS::ElastiCache::Cluster",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::Logs::LogGroup": {
		Name:         "AWS::Logs::LogGroup",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::S3::Bucket": {
		Name:         "AWS::S3::Bucket",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::S3::Object": {
		Name:         "AWS::S3::Object",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::S3::BucketIntelligentTieringConfiguration": {
		Name:         "AWS::S3::BucketIntelligentTieringConfiguration",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::S3::MultiRegionAccessPoint": {
		Name:         "AWS::S3::MultiRegionAccessPoint",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::CertificateManager::Certificate": {
		Name:         "AWS::CertificateManager::Certificate",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::EKS::AddonVersion": {
		Name:         "AWS::EKS::AddonVersion",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::ApiGatewayV2::Api": {
		Name:         "AWS::ApiGatewayV2::Api",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::EC2::Volume": {
		Name:         "AWS::EC2::Volume",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::ApiGateway::ApiKey": {
		Name:         "AWS::ApiGateway::ApiKey",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::Glue::Connection": {
		Name:         "AWS::Glue::Connection",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::ECS::Task": {
		Name:         "AWS::ECS::Task",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::SSM::ManagedInstance": {
		Name:         "AWS::SSM::ManagedInstance",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::SSM::Inventory": {
		Name:         "AWS::SSM::Inventory",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::SSM::InventoryEntry": {
		Name:         "AWS::SSM::InventoryEntry",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::SSM::MaintenanceWindow": {
		Name:         "AWS::SSM::MaintenanceWindow",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::SSM::PatchBaseline": {
		Name:         "AWS::SSM::PatchBaseline",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::SSM::Parameter": {
		Name:         "AWS::SSM::Parameter",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::Lambda::Function": {
		Name:         "AWS::Lambda::Function",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::RDS::DBSnapshot": {
		Name:         "AWS::RDS::DBSnapshot",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::CodeDeploy::Application": {
		Name:         "AWS::CodeDeploy::Application",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::CodeDeploy::DeploymentConfig": {
		Name:         "AWS::CodeDeploy::DeploymentConfig",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::EMR::Cluster": {
		Name:         "AWS::EMR::Cluster",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::IAM::AccessKey": {
		Name:         "AWS::IAM::AccessKey",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::IAM::SSHPublicKey": {
		Name:         "AWS::IAM::SSHPublicKey",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::Glue::CatalogTable": {
		Name:         "AWS::Glue::CatalogTable",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::CloudTrail::Channel": {
		Name:         "AWS::CloudTrail::Channel",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::EC2::NetworkAcl": {
		Name:         "AWS::EC2::NetworkAcl",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::ECS::ContainerInstance": {
		Name:         "AWS::ECS::ContainerInstance",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::RedshiftServerless::Snapshot": {
		Name:         "AWS::RedshiftServerless::Snapshot",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::Workspaces::Bundle": {
		Name:         "AWS::Workspaces::Bundle",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::CloudTrail::Trail": {
		Name:         "AWS::CloudTrail::Trail",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::DAX::Parameter": {
		Name:         "AWS::DAX::Parameter",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::ECR::Image": {
		Name:         "AWS::ECR::Image",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::IAM::ServerCertificate": {
		Name:         "AWS::IAM::ServerCertificate",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::Keyspaces::Keyspace": {
		Name:         "AWS::Keyspaces::Keyspace",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::S3::AccessPoint": {
		Name:         "AWS::S3::AccessPoint",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::SageMaker::EndpointConfiguration": {
		Name:         "AWS::SageMaker::EndpointConfiguration",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::ElastiCache::ReservedCacheNode": {
		Name:         "AWS::ElastiCache::ReservedCacheNode",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::EMR::InstanceFleet": {
		Name:         "AWS::EMR::InstanceFleet",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::Account::Account": {
		Name:         "AWS::Account::Account",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::EC2::VPCPeeringConnection": {
		Name:         "AWS::EC2::VPCPeeringConnection",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::EKS::FargateProfile": {
		Name:         "AWS::EKS::FargateProfile",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::IAM::AccountPasswordPolicy": {
		Name:         "AWS::IAM::AccountPasswordPolicy",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::CodePipeline::Pipeline": {
		Name:         "AWS::CodePipeline::Pipeline",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::DAX::Cluster": {
		Name:         "AWS::DAX::Cluster",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::DLM::LifecyclePolicy": {
		Name:         "AWS::DLM::LifecyclePolicy",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::OpsWorksCM::Server": {
		Name:         "AWS::OpsWorksCM::Server",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::AccessAnalyzer::Analyzer": {
		Name:         "AWS::AccessAnalyzer::Analyzer",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::AccessAnalyzer::Finding": {
		Name:         "AWS::AccessAnalyzer::Finding",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::ElastiCache::SubnetGroup": {
		Name:         "AWS::ElastiCache::SubnetGroup",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::FSX::Volume": {
		Name:         "AWS::FSX::Volume",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::Amplify::App": {
		Name:         "AWS::Amplify::App",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::CloudTrail::Query": {
		Name:         "AWS::CloudTrail::Query",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::CostExplorer::ByAccountMonthly": {
		Name:         "AWS::CostExplorer::ByAccountMonthly",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::ECR::PublicRegistry": {
		Name:         "AWS::ECR::PublicRegistry",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::EC2::NetworkInterface": {
		Name:         "AWS::EC2::NetworkInterface",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::EC2::VPNConnection": {
		Name:         "AWS::EC2::VPNConnection",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::FSX::StorageVirtualMachine": {
		Name:         "AWS::FSX::StorageVirtualMachine",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::ApiGateway::Authorizer": {
		Name:         "AWS::ApiGateway::Authorizer",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::AppStream::Stack": {
		Name:         "AWS::AppStream::Stack",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::Athena::WorkGroup": {
		Name:         "AWS::Athena::WorkGroup",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::Athena::QueryExecution": {
		Name:         "AWS::Athena::QueryExecution",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::AppStream::Image": {
		Name:         "AWS::AppStream::Image",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::CloudWatch::Alarm": {
		Name:         "AWS::CloudWatch::Alarm",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::CloudWatch::LogSubscriptionFilter": {
		Name:         "AWS::CloudWatch::LogSubscriptionFilter",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::CostExplorer::ByRecordTypeMonthly": {
		Name:         "AWS::CostExplorer::ByRecordTypeMonthly",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::RDS::DBCluster": {
		Name:         "AWS::RDS::DBCluster",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::RDS::DBClusterSnapshot": {
		Name:         "AWS::RDS::DBClusterSnapshot",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::Backup::Framework": {
		Name:         "AWS::Backup::Framework",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::CodeBuild::SourceCredential": {
		Name:         "AWS::CodeBuild::SourceCredential",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::IAM::ServiceSpecificCredential": {
		Name:         "AWS::IAM::ServiceSpecificCredential",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::EC2::CapacityReservationFleet": {
		Name:         "AWS::EC2::CapacityReservationFleet",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::NetworkFirewall::Firewall": {
		Name:         "AWS::NetworkFirewall::Firewall",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::Workspaces::Workspace": {
		Name:         "AWS::Workspaces::Workspace",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::ElasticSearch::Domain": {
		Name:         "AWS::ElasticSearch::Domain",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::RDS::DBInstance": {
		Name:         "AWS::RDS::DBInstance",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::RDS::DBInstanceAutomatedBackup": {
		Name:         "AWS::RDS::DBInstanceAutomatedBackup",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::EFS::MountTarget": {
		Name:         "AWS::EFS::MountTarget",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::GlobalAccelerator::Listener": {
		Name:         "AWS::GlobalAccelerator::Listener",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::CostExplorer::ByUsageTypeDaily": {
		Name:         "AWS::CostExplorer::ByUsageTypeDaily",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::EKS::Addon": {
		Name:         "AWS::EKS::Addon",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::CostExplorer::ByServiceMonthly": {
		Name:         "AWS::CostExplorer::ByServiceMonthly",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::IAM::Policy": {
		Name:         "AWS::IAM::Policy",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::Redshift::Cluster": {
		Name:         "AWS::Redshift::Cluster",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::WAFRegional::Rule": {
		Name:         "AWS::WAFRegional::Rule",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::WAFRegional::RuleGroup": {
		Name:         "AWS::WAFRegional::RuleGroup",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::Glue::DataCatalogEncryptionSettings": {
		Name:         "AWS::Glue::DataCatalogEncryptionSettings",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::EC2::FlowLog": {
		Name:         "AWS::EC2::FlowLog",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::EC2::IpamPool": {
		Name:         "AWS::EC2::IpamPool",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::IAM::SamlProvider": {
		Name:         "AWS::IAM::SamlProvider",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::Route53::HostedZone": {
		Name:         "AWS::Route53::HostedZone",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::Route53::QueryLog": {
		Name:         "AWS::Route53::QueryLog",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::EC2::PlacementGroup": {
		Name:         "AWS::EC2::PlacementGroup",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::FSX::Snapshot": {
		Name:         "AWS::FSX::Snapshot",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::KMS::Key": {
		Name:         "AWS::KMS::Key",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::KMS::KeyRotation": {
		Name:         "AWS::KMS::KeyRotation",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::EC2::Ipam": {
		Name:         "AWS::EC2::Ipam",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::ElasticBeanstalk::Environment": {
		Name:         "AWS::ElasticBeanstalk::Environment",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::ElasticBeanstalk::ApplicationVersion": {
		Name:         "AWS::ElasticBeanstalk::ApplicationVersion",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::Lambda::FunctionVersion": {
		Name:         "AWS::Lambda::FunctionVersion",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::Glue::DevEndpoint": {
		Name:         "AWS::Glue::DevEndpoint",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::Backup::RecoveryPoint": {
		Name:         "AWS::Backup::RecoveryPoint",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::Backup::ReportPlan": {
		Name:         "AWS::Backup::ReportPlan",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::Backup::RegionSetting": {
		Name:         "AWS::Backup::RegionSetting",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::DynamoDbStreams::Stream": {
		Name:         "AWS::DynamoDbStreams::Stream",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::EC2::EgressOnlyInternetGateway": {
		Name:         "AWS::EC2::EgressOnlyInternetGateway",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::CloudFront::Distribution": {
		Name:         "AWS::CloudFront::Distribution",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::CloudFront::StreamingDistribution": {
		Name:         "AWS::CloudFront::StreamingDistribution",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::Glue::Job": {
		Name:         "AWS::Glue::Job",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::AppStream::Fleet": {
		Name:         "AWS::AppStream::Fleet",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::SES::ConfigurationSet": {
		Name:         "AWS::SES::ConfigurationSet",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::IAM::User": {
		Name:         "AWS::IAM::User",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::CloudFront::OriginRequestPolicy": {
		Name:         "AWS::CloudFront::OriginRequestPolicy",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::EC2::SecurityGroup": {
		Name:         "AWS::EC2::SecurityGroup",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::GuardDuty::IPSet": {
		Name:         "AWS::GuardDuty::IPSet",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::EKS::Cluster": {
		Name:         "AWS::EKS::Cluster",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::Grafana::Workspace": {
		Name:         "AWS::Grafana::Workspace",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::Glue::CatalogDatabase": {
		Name:         "AWS::Glue::CatalogDatabase",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::Health::Event": {
		Name:         "AWS::Health::Event",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::Health::AffectedEntity": {
		Name:         "AWS::Health::AffectedEntity",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::CloudFormation::StackSet": {
		Name:         "AWS::CloudFormation::StackSet",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::EC2::AvailabilityZone": {
		Name:         "AWS::EC2::AvailabilityZone",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::EC2::TransitGateway": {
		Name:         "AWS::EC2::TransitGateway",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::ApiGateway::UsagePlan": {
		Name:         "AWS::ApiGateway::UsagePlan",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::Inspector::Finding": {
		Name:         "AWS::Inspector::Finding",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::EC2::Fleet": {
		Name:         "AWS::EC2::Fleet",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::ElasticBeanstalk::Application": {
		Name:         "AWS::ElasticBeanstalk::Application",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::ElasticLoadBalancingV2::LoadBalancer": {
		Name:         "AWS::ElasticLoadBalancingV2::LoadBalancer",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::OpenSearch::Domain": {
		Name:         "AWS::OpenSearch::Domain",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::RDS::DBEventSubscription": {
		Name:         "AWS::RDS::DBEventSubscription",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::RDS::DBEngineVersion": {
		Name:         "AWS::RDS::DBEngineVersion",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::EC2::RegionalSettings": {
		Name:         "AWS::EC2::RegionalSettings",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::EC2::SecurityGroupRule": {
		Name:         "AWS::EC2::SecurityGroupRule",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::EC2::TransitGatewayAttachment": {
		Name:         "AWS::EC2::TransitGatewayAttachment",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::SES::Identity": {
		Name:         "AWS::SES::Identity",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::SESv2::EmailIdentities": {
		Name:         "AWS::SESv2::EmailIdentities",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::WAF::Rule": {
		Name:         "AWS::WAF::Rule",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::WAF::RuleGroup": {
		Name:         "AWS::WAF::RuleGroup",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::WAF::RateBasedRule": {
		Name:         "AWS::WAF::RateBasedRule",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::WAF::WebACL": {
		Name:         "AWS::WAF::WebACL",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::WAFRegional::WebACL": {
		Name:         "AWS::WAFRegional::WebACL",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::WellArchitected::Workload": {
		Name:         "AWS::WellArchitected::Workload",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::WellArchitected::Answer": {
		Name:         "AWS::WellArchitected::Answer",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::WellArchitected::CheckDetail": {
		Name:         "AWS::WellArchitected::CheckDetail",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::WellArchitected::CheckSummary": {
		Name:         "AWS::WellArchitected::CheckSummary",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::WellArchitected::ConsolidatedReport": {
		Name:         "AWS::WellArchitected::ConsolidatedReport",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::WellArchitected::Lens": {
		Name:         "AWS::WellArchitected::Lens",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::WellArchitected::LensReview": {
		Name:         "AWS::WellArchitected::LensReview",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::WellArchitected::LensReviewImprovement": {
		Name:         "AWS::WellArchitected::LensReviewImprovement",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::WellArchitected::LensReviewReport": {
		Name:         "AWS::WellArchitected::LensReviewReport",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::WellArchitected::LensShare": {
		Name:         "AWS::WellArchitected::LensShare",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::WellArchitected::Milestone": {
		Name:         "AWS::WellArchitected::Milestone",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::WellArchitected::Notification": {
		Name:         "AWS::WellArchitected::Notification",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::WellArchitected::ShareInvitation": {
		Name:         "AWS::WellArchitected::ShareInvitation",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::WellArchitected::WorkloadShare": {
		Name:         "AWS::WellArchitected::WorkloadShare",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::AutoScaling::LaunchConfiguration": {
		Name:         "AWS::AutoScaling::LaunchConfiguration",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::CloudTrail::EventDataStore": {
		Name:         "AWS::CloudTrail::EventDataStore",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::CodeDeploy::DeploymentGroup": {
		Name:         "AWS::CodeDeploy::DeploymentGroup",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::ImageBuilder::Image": {
		Name:         "AWS::ImageBuilder::Image",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::Redshift::ClusterParameterGroup": {
		Name:         "AWS::Redshift::ClusterParameterGroup",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::Account::AlternateContact": {
		Name:         "AWS::Account::AlternateContact",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::Inspector::AssessmentTarget": {
		Name:         "AWS::Inspector::AssessmentTarget",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::CloudFront::ResponseHeadersPolicy": {
		Name:         "AWS::CloudFront::ResponseHeadersPolicy",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::EC2::Instance": {
		Name:         "AWS::EC2::Instance",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::EC2::InstanceAvailability": {
		Name:         "AWS::EC2::InstanceAvailability",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::CostExplorer::ByRecordTypeDaily": {
		Name:         "AWS::CostExplorer::ByRecordTypeDaily",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::EC2::ReservedInstances": {
		Name:         "AWS::EC2::ReservedInstances",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::ECR::Repository": {
		Name:         "AWS::ECR::Repository",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::ECR::Registry": {
		Name:         "AWS::ECR::Registry",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::ECR::RegistryScanningConfiguration": {
		Name:         "AWS::ECR::RegistryScanningConfiguration",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::ElasticLoadBalancingV2::Listener": {
		Name:         "AWS::ElasticLoadBalancingV2::Listener",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::IAM::Group": {
		Name:         "AWS::IAM::Group",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::IAM::OpenIdConnectProvider": {
		Name:         "AWS::IAM::OpenIdConnectProvider",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::Backup::Plan": {
		Name:         "AWS::Backup::Plan",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::Config::ConformancePack": {
		Name:         "AWS::Config::ConformancePack",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::Config::RetentionConfiguration": {
		Name:         "AWS::Config::RetentionConfiguration",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::CostExplorer::ByAccountDaily": {
		Name:         "AWS::CostExplorer::ByAccountDaily",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::Account::Contact": {
		Name:         "AWS::Account::Contact",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::Glue::DataQualityRuleset": {
		Name:         "AWS::Glue::DataQualityRuleset",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::EventBridge::EventBus": {
		Name:         "AWS::EventBridge::EventBus",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::ApiGateway::Stage": {
		Name:         "AWS::ApiGateway::Stage",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::ApiGatewayV2::Stage": {
		Name:         "AWS::ApiGatewayV2::Stage",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::DynamoDb::LocalSecondaryIndex": {
		Name:         "AWS::DynamoDb::LocalSecondaryIndex",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::ResourceGroups::Groups": {
		Name:         "AWS::ResourceGroups::Groups",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::Timestream::Database": {
		Name:         "AWS::Timestream::Database",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::OpenSearchServerless::Collection": {
		Name:         "AWS::OpenSearchServerless::Collection",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::EC2::ElasticIP": {
		Name:         "AWS::EC2::ElasticIP",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::EC2::LocalGateway": {
		Name:         "AWS::EC2::LocalGateway",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::EC2::Image": {
		Name:         "AWS::EC2::Image",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::EC2::Subnet": {
		Name:         "AWS::EC2::Subnet",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::ECS::TaskSet": {
		Name:         "AWS::ECS::TaskSet",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::Kinesis::Stream": {
		Name:         "AWS::Kinesis::Stream",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::Kinesis::Consumer": {
		Name:         "AWS::Kinesis::Consumer",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::DocDB::Cluster": {
		Name:         "AWS::DocDB::Cluster",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::DocDB::ClusterSnapshot": {
		Name:         "AWS::DocDB::ClusterSnapshot",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::DocDB::ClusterInstance": {
		Name:         "AWS::DocDB::ClusterInstance",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::ElastiCache::ReplicationGroup": {
		Name:         "AWS::ElastiCache::ReplicationGroup",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::GlobalAccelerator::Accelerator": {
		Name:         "AWS::GlobalAccelerator::Accelerator",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::CloudWatch::Metric": {
		Name:         "AWS::CloudWatch::Metric",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::CostExplorer::ForcastMonthly": {
		Name:         "AWS::CostExplorer::ForcastMonthly",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::EMR::InstanceGroup": {
		Name:         "AWS::EMR::InstanceGroup",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::EC2::ManagedPrefixList": {
		Name:         "AWS::EC2::ManagedPrefixList",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::EC2::ManagedPrefixListEntry": {
		Name:         "AWS::EC2::ManagedPrefixListEntry",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::EC2::ClientVpnEndpoint": {
		Name:         "AWS::EC2::ClientVpnEndpoint",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::MWAA::Environment": {
		Name:         "AWS::MWAA::Environment",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::CloudWatch::LogResourcePolicy": {
		Name:         "AWS::CloudWatch::LogResourcePolicy",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::CodeArtifact::Domain": {
		Name:         "AWS::CodeArtifact::Domain",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::CodeStar::Project": {
		Name:         "AWS::CodeStar::Project",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::Neptune::Database": {
		Name:         "AWS::Neptune::Database",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::Neptune::DBCluster": {
		Name:         "AWS::Neptune::DBCluster",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::Neptune::DBClusterSnapshot": {
		Name:         "AWS::Neptune::DBClusterSnapshot",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::NetworkFirewall::FirewallPolicy": {
		Name:         "AWS::NetworkFirewall::FirewallPolicy",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::NetworkFirewall::RuleGroup": {
		Name:         "AWS::NetworkFirewall::RuleGroup",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::Oam::Link": {
		Name:         "AWS::Oam::Link",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::Oam::Sink": {
		Name:         "AWS::Oam::Sink",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::Organizations::Account": {
		Name:         "AWS::Organizations::Account",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::Pinpoint::App": {
		Name:         "AWS::Pinpoint::App",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::Pipes::Pipe": {
		Name:         "AWS::Pipes::Pipe",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::RDS::DBClusterParameterGroup": {
		Name:         "AWS::RDS::DBClusterParameterGroup",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::RDS::OptionGroup": {
		Name:         "AWS::RDS::OptionGroup",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::RDS::DBParameterGroup": {
		Name:         "AWS::RDS::DBParameterGroup",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::RDS::DBProxy": {
		Name:         "AWS::RDS::DBProxy",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::RDS::DBSubnetGroup": {
		Name:         "AWS::RDS::DBSubnetGroup",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::RDS::DBRecommendation": {
		Name:         "AWS::RDS::DBRecommendation",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::Redshift::EventSubscription": {
		Name:         "AWS::Redshift::EventSubscription",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::RedshiftServerless::Workgroup": {
		Name:         "AWS::RedshiftServerless::Workgroup",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::ResourceExplorer2::Index": {
		Name:         "AWS::ResourceExplorer2::Index",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::Route53::HealthCheck": {
		Name:         "AWS::Route53::HealthCheck",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::Route53Resolver::ResolverRule": {
		Name:         "AWS::Route53Resolver::ResolverRule",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::Route53Resolver::QueryLogConfig": {
		Name:         "AWS::Route53Resolver::QueryLogConfig",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::SageMaker::App": {
		Name:         "AWS::SageMaker::App",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::SageMaker::Domain": {
		Name:         "AWS::SageMaker::Domain",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::StepFunctions::StateMachine": {
		Name:         "AWS::StepFunctions::StateMachine",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::StepFunctions::StateMachineExecution": {
		Name:         "AWS::StepFunctions::StateMachineExecution",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::StepFunctions::StateMachineExecutionHistories": {
		Name:         "AWS::StepFunctions::StateMachineExecutionHistories",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::SimSpaceWeaver::Simulation": {
		Name:         "AWS::SimSpaceWeaver::Simulation",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::SSM::Association": {
		Name:         "AWS::SSM::Association",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::SSM::Document": {
		Name:         "AWS::SSM::Document",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::SSM::DocumentPermission": {
		Name:         "AWS::SSM::DocumentPermission",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::EC2::CustomerGateway": {
		Name:         "AWS::EC2::CustomerGateway",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::EC2::VerifiedAccessInstance": {
		Name:         "AWS::EC2::VerifiedAccessInstance",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::EC2::VerifiedAccessEndpoint": {
		Name:         "AWS::EC2::VerifiedAccessEndpoint",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::EC2::VerifiedAccessGroup": {
		Name:         "AWS::EC2::VerifiedAccessGroup",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::EC2::VerifiedAccessTrustProvider": {
		Name:         "AWS::EC2::VerifiedAccessTrustProvider",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::EC2::VPNGateway": {
		Name:         "AWS::EC2::VPNGateway",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::WAFv2::IPSet": {
		Name:         "AWS::WAFv2::IPSet",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::WAFv2::RegexPatternSet": {
		Name:         "AWS::WAFv2::RegexPatternSet",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::WAFv2::RuleGroup": {
		Name:         "AWS::WAFv2::RuleGroup",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::EC2::TransitGatewayRoute": {
		Name:         "AWS::EC2::TransitGatewayRoute",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::GuardDuty::Filter": {
		Name:         "AWS::GuardDuty::Filter",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::ECS::TaskDefinition": {
		Name:         "AWS::ECS::TaskDefinition",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::GuardDuty::ThreatIntelSet": {
		Name:         "AWS::GuardDuty::ThreatIntelSet",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::ApiGatewayV2::DomainName": {
		Name:         "AWS::ApiGatewayV2::DomainName",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::ApiGateway::DomainName": {
		Name:         "AWS::ApiGateway::DomainName",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::ApiGatewayV2::Route": {
		Name:         "AWS::ApiGatewayV2::Route",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::MQ::Broker": {
		Name:         "AWS::MQ::Broker",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::ACMPCA::CertificateAuthority": {
		Name:         "AWS::ACMPCA::CertificateAuthority",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::CloudFormation::Stack": {
		Name:         "AWS::CloudFormation::Stack",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::CloudFormation::StackResource": {
		Name:         "AWS::CloudFormation::StackResource",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::DirectConnect::Connection": {
		Name:         "AWS::DirectConnect::Connection",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::FSX::FileSystem": {
		Name:         "AWS::FSX::FileSystem",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::Glue::SecurityConfiguration": {
		Name:         "AWS::Glue::SecurityConfiguration",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::Inspector::AssessmentRun": {
		Name:         "AWS::Inspector::AssessmentRun",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::Inspector2::Coverage": {
		Name:         "AWS::Inspector2::Coverage",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::Inspector2::CoverageStatistics": {
		Name:         "AWS::Inspector2::CoverageStatistics",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::Inspector2::Member": {
		Name:         "AWS::Inspector2::Member",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::Inspector2::Finding": {
		Name:         "AWS::Inspector2::Finding",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::Config::ConfigurationRecorder": {
		Name:         "AWS::Config::ConfigurationRecorder",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::EC2::NatGateway": {
		Name:         "AWS::EC2::NatGateway",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::ECR::PublicRepository": {
		Name:         "AWS::ECR::PublicRepository",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::ECS::Cluster": {
		Name:         "AWS::ECS::Cluster",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::ElasticLoadBalancingV2::TargetGroup": {
		Name:         "AWS::ElasticLoadBalancingV2::TargetGroup",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::CloudFront::CachePolicy": {
		Name:         "AWS::CloudFront::CachePolicy",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::CodeArtifact::Repository": {
		Name:         "AWS::CodeArtifact::Repository",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::AMP::Workspace": {
		Name:         "AWS::AMP::Workspace",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::EC2::CapacityReservation": {
		Name:         "AWS::EC2::CapacityReservation",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::SageMaker::NotebookInstance": {
		Name:         "AWS::SageMaker::NotebookInstance",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::IAM::AccessAdvisor": {
		Name:         "AWS::IAM::AccessAdvisor",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::EC2::VolumeSnapshot": {
		Name:         "AWS::EC2::VolumeSnapshot",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::EC2::Region": {
		Name:         "AWS::EC2::Region",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::Keyspaces::Table": {
		Name:         "AWS::Keyspaces::Table",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::Config::AggregationAuthorization": {
		Name:         "AWS::Config::AggregationAuthorization",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::DAX::SubnetGroup": {
		Name:         "AWS::DAX::SubnetGroup",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::DynamoDb::GlobalTable": {
		Name:         "AWS::DynamoDb::GlobalTable",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::ElasticLoadBalancing::LoadBalancer": {
		Name:         "AWS::ElasticLoadBalancing::LoadBalancer",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::AppStream::Application": {
		Name:         "AWS::AppStream::Application",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::RedshiftServerless::Namespace": {
		Name:         "AWS::RedshiftServerless::Namespace",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::CloudFront::OriginAccessIdentity": {
		Name:         "AWS::CloudFront::OriginAccessIdentity",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::EC2::Host": {
		Name:         "AWS::EC2::Host",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::EC2::VPC": {
		Name:         "AWS::EC2::VPC",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::EC2::TransitGatewayRouteTable": {
		Name:         "AWS::EC2::TransitGatewayRouteTable",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::EKS::Nodegroup": {
		Name:         "AWS::EKS::Nodegroup",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::Backup::Selection": {
		Name:         "AWS::Backup::Selection",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::CloudTrail::Import": {
		Name:         "AWS::CloudTrail::Import",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::CostExplorer::ByServiceDaily": {
		Name:         "AWS::CostExplorer::ByServiceDaily",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::ElasticLoadBalancingV2::SslPolicy": {
		Name:         "AWS::ElasticLoadBalancingV2::SslPolicy",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::GuardDuty::Finding": {
		Name:         "AWS::GuardDuty::Finding",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::EC2::DHCPOptions": {
		Name:         "AWS::EC2::DHCPOptions",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::EC2::InstanceType": {
		Name:         "AWS::EC2::InstanceType",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::Batch::ComputeEnvironment": {
		Name:         "AWS::Batch::ComputeEnvironment",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::DMS::ReplicationInstance": {
		Name:         "AWS::DMS::ReplicationInstance",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::DMS::Endpoint": {
		Name:         "AWS::DMS::Endpoint",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::DMS::ReplicationTask": {
		Name:         "AWS::DMS::ReplicationTask",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::DynamoDb::Table": {
		Name:         "AWS::DynamoDb::Table",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::Shield::ProtectionGroup": {
		Name:         "AWS::Shield::ProtectionGroup",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::Firehose::DeliveryStream": {
		Name:         "AWS::Firehose::DeliveryStream",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::KinesisVideo::Stream": {
		Name:         "AWS::KinesisVideo::Stream",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::KMS::Alias": {
		Name:         "AWS::KMS::Alias",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::Lambda::Alias": {
		Name:         "AWS::Lambda::Alias",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::Lambda::LambdaLayer": {
		Name:         "AWS::Lambda::LambdaLayer",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::Lambda::LayerVersion": {
		Name:         "AWS::Lambda::LayerVersion",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::Lightsail::Instance": {
		Name:         "AWS::Lightsail::Instance",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::Macie2::ClassificationJob": {
		Name:         "AWS::Macie2::ClassificationJob",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::MediaStore::Container": {
		Name:         "AWS::MediaStore::Container",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::Mgn::Application": {
		Name:         "AWS::Mgn::Application",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::ResourceExplorer2::SupportedResourceType": {
		Name:         "AWS::ResourceExplorer2::SupportedResourceType",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::Route53Resolver::ResolverEndpoint": {
		Name:         "AWS::Route53Resolver::ResolverEndpoint",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::Route53Domains::Domain": {
		Name:         "AWS::Route53Domains::Domain",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::Route53::Record": {
		Name:         "AWS::Route53::Record",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::Route53::TrafficPolicy": {
		Name:         "AWS::Route53::TrafficPolicy",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::Route53::TrafficPolicyInstance": {
		Name:         "AWS::Route53::TrafficPolicyInstance",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::SageMaker::Model": {
		Name:         "AWS::SageMaker::Model",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::SageMaker::TrainingJob": {
		Name:         "AWS::SageMaker::TrainingJob",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::SecurityHub::ActionTarget": {
		Name:         "AWS::SecurityHub::ActionTarget",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::SecurityHub::Finding": {
		Name:         "AWS::SecurityHub::Finding",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::SecurityHub::FindingAggregator": {
		Name:         "AWS::SecurityHub::FindingAggregator",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::SecurityHub::Insight": {
		Name:         "AWS::SecurityHub::Insight",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::SecurityHub::Member": {
		Name:         "AWS::SecurityHub::Member",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::SecurityHub::Product": {
		Name:         "AWS::SecurityHub::Product",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::SecurityHub::StandardsControl": {
		Name:         "AWS::SecurityHub::StandardsControl",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::SecurityHub::StandardsSubscription": {
		Name:         "AWS::SecurityHub::StandardsSubscription",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::SecurityLake::DataLake": {
		Name:         "AWS::SecurityLake::DataLake",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::SecurityLake::Subscriber": {
		Name:         "AWS::SecurityLake::Subscriber",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::Ram::PrincipalAssociation": {
		Name:         "AWS::Ram::PrincipalAssociation",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::Ram::ResourceAssociation": {
		Name:         "AWS::Ram::ResourceAssociation",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::RDS::ReservedDBInstance": {
		Name:         "AWS::RDS::ReservedDBInstance",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::Redshift::SubnetGroup": {
		Name:         "AWS::Redshift::SubnetGroup",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::SeverlessApplicationRepository::Application": {
		Name:         "AWS::SeverlessApplicationRepository::Application",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::AuditManager::Framework": {
		Name:         "AWS::AuditManager::Framework",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::AuditManager::EvidenceFolder": {
		Name:         "AWS::AuditManager::EvidenceFolder",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::AuditManager::Evidence": {
		Name:         "AWS::AuditManager::Evidence",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::AuditManager::Control": {
		Name:         "AWS::AuditManager::Control",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::AuditManager::Assessment": {
		Name:         "AWS::AuditManager::Assessment",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::CloudTrail::TrailEvent": {
		Name:         "AWS::CloudTrail::TrailEvent",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::CloudWatch::LogEvent": {
		Name:         "AWS::CloudWatch::LogEvent",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::Logs::MetricFilter": {
		Name:         "AWS::Logs::MetricFilter",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::CloudWatch::LogStream": {
		Name:         "AWS::CloudWatch::LogStream",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::CostExplorer::ByUsageTypeMonthly": {
		Name:         "AWS::CostExplorer::ByUsageTypeMonthly",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::ServiceQuotas::ServiceQuotaChangeRequest": {
		Name:         "AWS::ServiceQuotas::ServiceQuotaChangeRequest",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::ServiceQuotas::Service": {
		Name:         "AWS::ServiceQuotas::Service",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::EC2::VPCEndpointService": {
		Name:         "AWS::EC2::VPCEndpointService",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::EC2::LaunchTemplate": {
		Name:         "AWS::EC2::LaunchTemplate",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::EC2::LaunchTemplateVersion": {
		Name:         "AWS::EC2::LaunchTemplateVersion",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::SNS::Subscription": {
		Name:         "AWS::SNS::Subscription",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::S3::AccountSetting": {
		Name:         "AWS::S3::AccountSetting",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::SSM::ManagedInstanceCompliance": {
		Name:         "AWS::SSM::ManagedInstanceCompliance",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::SSM::ManagedInstancePatchState": {
		Name:         "AWS::SSM::ManagedInstancePatchState",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::SSOAdmin::AccountAssignment": {
		Name:         "AWS::SSOAdmin::AccountAssignment",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::SSOAdmin::UserEffectiveAccess": {
		Name:         "AWS::SSOAdmin::UserEffectiveAccess",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::SSOAdmin::Instance": {
		Name:         "AWS::SSOAdmin::Instance",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::SSOAdmin::PermissionSet": {
		Name:         "AWS::SSOAdmin::PermissionSet",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::SSOAdmin::AttachedManagedPolicy": {
		Name:         "AWS::SSOAdmin::AttachedManagedPolicy",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::ServiceDiscovery::Service": {
		Name:         "AWS::ServiceDiscovery::Service",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::ServiceDiscovery::Namespace": {
		Name:         "AWS::ServiceDiscovery::Namespace",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::ServiceDiscovery::Instance": {
		Name:         "AWS::ServiceDiscovery::Instance",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::ServiceCatalog::Portfolio": {
		Name:         "AWS::ServiceCatalog::Portfolio",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::ServiceCatalog::Product": {
		Name:         "AWS::ServiceCatalog::Product",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::IdentityStore::User": {
		Name:         "AWS::IdentityStore::User",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::IdentityStore::Group": {
		Name:         "AWS::IdentityStore::Group",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"AWS::IdentityStore::GroupMembership": {
		Name:         "AWS::IdentityStore::GroupMembership",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},
}


var ResourceTypesList = []string{
  "AWS::Redshift::Snapshot",
  "AWS::IAM::AccountSummary",
  "AWS::Glacier::Vault",
  "AWS::Organizations::Organization",
  "AWS::Organizations::Policy",
  "AWS::Organizations::PolicyTarget",
  "AWS::Organizations::OrganizationalUnit",
  "AWS::Organizations::Root",
  "AWS::CloudSearch::Domain",
  "AWS::DynamoDb::GlobalSecondaryIndex",
  "AWS::EC2::RouteTable",
  "AWS::SecurityHub::Hub",
  "AWS::StorageGateway::StorageGateway",
  "AWS::Inspector::AssessmentTemplate",
  "AWS::ElasticLoadBalancingV2::ListenerRule",
  "AWS::IAM::Role",
  "AWS::Backup::ProtectedResource",
  "AWS::CodeCommit::Repository",
  "AWS::EC2::VPCEndpoint",
  "AWS::EventBridge::EventRule",
  "AWS::CloudFront::OriginAccessControl",
  "AWS::CodeBuild::Project",
  "AWS::CodeBuild::Build",
  "AWS::ElastiCache::ParameterGroup",
  "AWS::MemoryDb::Cluster",
  "AWS::Glue::Crawler",
  "AWS::DirectConnect::Gateway",
  "AWS::DynamoDb::BackUp",
  "AWS::EC2::EIP",
  "AWS::EC2::InternetGateway",
  "AWS::GuardDuty::PublishingDestination",
  "AWS::KinesisAnalyticsV2::Application",
  "AWS::EMR::Instance",
  "AWS::EMR::BlockPublicAccessConfiguration",
  "AWS::ApiGateway::RestApi",
  "AWS::ApiGatewayV2::Integration",
  "AWS::AutoScaling::AutoScalingGroup",
  "AWS::DynamoDb::TableExport",
  "AWS::EC2::KeyPair",
  "AWS::EFS::FileSystem",
  "AWS::Kafka::Cluster",
  "AWS::SecretsManager::Secret",
  "AWS::Backup::LegalHold",
  "AWS::CloudFront::Function",
  "AWS::GlobalAccelerator::EndpointGroup",
  "AWS::DAX::ParameterGroup",
  "AWS::SQS::Queue",
  "AWS::Config::Rule",
  "AWS::GuardDuty::Member",
  "AWS::Inspector::Exclusion",
  "AWS::DirectoryService::Directory",
  "AWS::DirectoryService::Certificate",
  "AWS::DirectoryService::LogSubscription",
  "AWS::EFS::AccessPoint",
  "AWS::IAM::PolicyAttachment",
  "AWS::IAM::CredentialReport",
  "AWS::RDS::GlobalCluster",
  "AWS::CostExplorer::ForcastDaily",
  "AWS::GuardDuty::Detector",
  "AWS::SNS::Topic",
  "AWS::AppConfig::Application",
  "AWS::Batch::Job",
  "AWS::Batch::JobQueue",
  "AWS::ECS::Service",
  "AWS::FSX::Task",
  "AWS::IAM::VirtualMFADevice",
  "AWS::WAFv2::WebACL",
  "AWS::ApplicationAutoScaling::Target",
  "AWS::ApplicationAutoScaling::Policy",
  "AWS::Backup::Vault",
  "AWS::ElastiCache::Cluster",
  "AWS::Logs::LogGroup",
  "AWS::S3::Bucket",
  "AWS::S3::Object",
  "AWS::S3::BucketIntelligentTieringConfiguration",
  "AWS::S3::MultiRegionAccessPoint",
  "AWS::CertificateManager::Certificate",
  "AWS::EKS::AddonVersion",
  "AWS::ApiGatewayV2::Api",
  "AWS::EC2::Volume",
  "AWS::ApiGateway::ApiKey",
  "AWS::Glue::Connection",
  "AWS::ECS::Task",
  "AWS::SSM::ManagedInstance",
  "AWS::SSM::Inventory",
  "AWS::SSM::InventoryEntry",
  "AWS::SSM::MaintenanceWindow",
  "AWS::SSM::PatchBaseline",
  "AWS::SSM::Parameter",
  "AWS::Lambda::Function",
  "AWS::RDS::DBSnapshot",
  "AWS::CodeDeploy::Application",
  "AWS::CodeDeploy::DeploymentConfig",
  "AWS::EMR::Cluster",
  "AWS::IAM::AccessKey",
  "AWS::IAM::SSHPublicKey",
  "AWS::Glue::CatalogTable",
  "AWS::CloudTrail::Channel",
  "AWS::EC2::NetworkAcl",
  "AWS::ECS::ContainerInstance",
  "AWS::RedshiftServerless::Snapshot",
  "AWS::Workspaces::Bundle",
  "AWS::CloudTrail::Trail",
  "AWS::DAX::Parameter",
  "AWS::ECR::Image",
  "AWS::IAM::ServerCertificate",
  "AWS::Keyspaces::Keyspace",
  "AWS::S3::AccessPoint",
  "AWS::SageMaker::EndpointConfiguration",
  "AWS::ElastiCache::ReservedCacheNode",
  "AWS::EMR::InstanceFleet",
  "AWS::Account::Account",
  "AWS::EC2::VPCPeeringConnection",
  "AWS::EKS::FargateProfile",
  "AWS::IAM::AccountPasswordPolicy",
  "AWS::CodePipeline::Pipeline",
  "AWS::DAX::Cluster",
  "AWS::DLM::LifecyclePolicy",
  "AWS::OpsWorksCM::Server",
  "AWS::AccessAnalyzer::Analyzer",
  "AWS::AccessAnalyzer::Finding",
  "AWS::ElastiCache::SubnetGroup",
  "AWS::FSX::Volume",
  "AWS::Amplify::App",
  "AWS::CloudTrail::Query",
  "AWS::CostExplorer::ByAccountMonthly",
  "AWS::ECR::PublicRegistry",
  "AWS::EC2::NetworkInterface",
  "AWS::EC2::VPNConnection",
  "AWS::FSX::StorageVirtualMachine",
  "AWS::ApiGateway::Authorizer",
  "AWS::AppStream::Stack",
  "AWS::Athena::WorkGroup",
  "AWS::Athena::QueryExecution",
  "AWS::AppStream::Image",
  "AWS::CloudWatch::Alarm",
  "AWS::CloudWatch::LogSubscriptionFilter",
  "AWS::CostExplorer::ByRecordTypeMonthly",
  "AWS::RDS::DBCluster",
  "AWS::RDS::DBClusterSnapshot",
  "AWS::Backup::Framework",
  "AWS::CodeBuild::SourceCredential",
  "AWS::IAM::ServiceSpecificCredential",
  "AWS::EC2::CapacityReservationFleet",
  "AWS::NetworkFirewall::Firewall",
  "AWS::Workspaces::Workspace",
  "AWS::ElasticSearch::Domain",
  "AWS::RDS::DBInstance",
  "AWS::RDS::DBInstanceAutomatedBackup",
  "AWS::EFS::MountTarget",
  "AWS::GlobalAccelerator::Listener",
  "AWS::CostExplorer::ByUsageTypeDaily",
  "AWS::EKS::Addon",
  "AWS::CostExplorer::ByServiceMonthly",
  "AWS::IAM::Policy",
  "AWS::Redshift::Cluster",
  "AWS::WAFRegional::Rule",
  "AWS::WAFRegional::RuleGroup",
  "AWS::Glue::DataCatalogEncryptionSettings",
  "AWS::EC2::FlowLog",
  "AWS::EC2::IpamPool",
  "AWS::IAM::SamlProvider",
  "AWS::Route53::HostedZone",
  "AWS::Route53::QueryLog",
  "AWS::EC2::PlacementGroup",
  "AWS::FSX::Snapshot",
  "AWS::KMS::Key",
  "AWS::KMS::KeyRotation",
  "AWS::EC2::Ipam",
  "AWS::ElasticBeanstalk::Environment",
  "AWS::ElasticBeanstalk::ApplicationVersion",
  "AWS::Lambda::FunctionVersion",
  "AWS::Glue::DevEndpoint",
  "AWS::Backup::RecoveryPoint",
  "AWS::Backup::ReportPlan",
  "AWS::Backup::RegionSetting",
  "AWS::DynamoDbStreams::Stream",
  "AWS::EC2::EgressOnlyInternetGateway",
  "AWS::CloudFront::Distribution",
  "AWS::CloudFront::StreamingDistribution",
  "AWS::Glue::Job",
  "AWS::AppStream::Fleet",
  "AWS::SES::ConfigurationSet",
  "AWS::IAM::User",
  "AWS::CloudFront::OriginRequestPolicy",
  "AWS::EC2::SecurityGroup",
  "AWS::GuardDuty::IPSet",
  "AWS::EKS::Cluster",
  "AWS::Grafana::Workspace",
  "AWS::Glue::CatalogDatabase",
  "AWS::Health::Event",
  "AWS::Health::AffectedEntity",
  "AWS::CloudFormation::StackSet",
  "AWS::EC2::AvailabilityZone",
  "AWS::EC2::TransitGateway",
  "AWS::ApiGateway::UsagePlan",
  "AWS::Inspector::Finding",
  "AWS::EC2::Fleet",
  "AWS::ElasticBeanstalk::Application",
  "AWS::ElasticLoadBalancingV2::LoadBalancer",
  "AWS::OpenSearch::Domain",
  "AWS::RDS::DBEventSubscription",
  "AWS::RDS::DBEngineVersion",
  "AWS::EC2::RegionalSettings",
  "AWS::EC2::SecurityGroupRule",
  "AWS::EC2::TransitGatewayAttachment",
  "AWS::SES::Identity",
  "AWS::SESv2::EmailIdentities",
  "AWS::WAF::Rule",
  "AWS::WAF::RuleGroup",
  "AWS::WAF::RateBasedRule",
  "AWS::WAF::WebACL",
  "AWS::WAFRegional::WebACL",
  "AWS::WellArchitected::Workload",
  "AWS::WellArchitected::Answer",
  "AWS::WellArchitected::CheckDetail",
  "AWS::WellArchitected::CheckSummary",
  "AWS::WellArchitected::ConsolidatedReport",
  "AWS::WellArchitected::Lens",
  "AWS::WellArchitected::LensReview",
  "AWS::WellArchitected::LensReviewImprovement",
  "AWS::WellArchitected::LensReviewReport",
  "AWS::WellArchitected::LensShare",
  "AWS::WellArchitected::Milestone",
  "AWS::WellArchitected::Notification",
  "AWS::WellArchitected::ShareInvitation",
  "AWS::WellArchitected::WorkloadShare",
  "AWS::AutoScaling::LaunchConfiguration",
  "AWS::CloudTrail::EventDataStore",
  "AWS::CodeDeploy::DeploymentGroup",
  "AWS::ImageBuilder::Image",
  "AWS::Redshift::ClusterParameterGroup",
  "AWS::Account::AlternateContact",
  "AWS::Inspector::AssessmentTarget",
  "AWS::CloudFront::ResponseHeadersPolicy",
  "AWS::EC2::Instance",
  "AWS::EC2::InstanceAvailability",
  "AWS::CostExplorer::ByRecordTypeDaily",
  "AWS::EC2::ReservedInstances",
  "AWS::ECR::Repository",
  "AWS::ECR::Registry",
  "AWS::ECR::RegistryScanningConfiguration",
  "AWS::ElasticLoadBalancingV2::Listener",
  "AWS::IAM::Group",
  "AWS::IAM::OpenIdConnectProvider",
  "AWS::Backup::Plan",
  "AWS::Config::ConformancePack",
  "AWS::Config::RetentionConfiguration",
  "AWS::CostExplorer::ByAccountDaily",
  "AWS::Account::Contact",
  "AWS::Glue::DataQualityRuleset",
  "AWS::EventBridge::EventBus",
  "AWS::ApiGateway::Stage",
  "AWS::ApiGatewayV2::Stage",
  "AWS::DynamoDb::LocalSecondaryIndex",
  "AWS::ResourceGroups::Groups",
  "AWS::Timestream::Database",
  "AWS::OpenSearchServerless::Collection",
  "AWS::EC2::ElasticIP",
  "AWS::EC2::LocalGateway",
  "AWS::EC2::Image",
  "AWS::EC2::Subnet",
  "AWS::ECS::TaskSet",
  "AWS::Kinesis::Stream",
  "AWS::Kinesis::Consumer",
  "AWS::DocDB::Cluster",
  "AWS::DocDB::ClusterSnapshot",
  "AWS::DocDB::ClusterInstance",
  "AWS::ElastiCache::ReplicationGroup",
  "AWS::GlobalAccelerator::Accelerator",
  "AWS::CloudWatch::Metric",
  "AWS::CostExplorer::ForcastMonthly",
  "AWS::EMR::InstanceGroup",
  "AWS::EC2::ManagedPrefixList",
  "AWS::EC2::ManagedPrefixListEntry",
  "AWS::EC2::ClientVpnEndpoint",
  "AWS::MWAA::Environment",
  "AWS::CloudWatch::LogResourcePolicy",
  "AWS::CodeArtifact::Domain",
  "AWS::CodeStar::Project",
  "AWS::Neptune::Database",
  "AWS::Neptune::DBCluster",
  "AWS::Neptune::DBClusterSnapshot",
  "AWS::NetworkFirewall::FirewallPolicy",
  "AWS::NetworkFirewall::RuleGroup",
  "AWS::Oam::Link",
  "AWS::Oam::Sink",
  "AWS::Organizations::Account",
  "AWS::Pinpoint::App",
  "AWS::Pipes::Pipe",
  "AWS::RDS::DBClusterParameterGroup",
  "AWS::RDS::OptionGroup",
  "AWS::RDS::DBParameterGroup",
  "AWS::RDS::DBProxy",
  "AWS::RDS::DBSubnetGroup",
  "AWS::RDS::DBRecommendation",
  "AWS::Redshift::EventSubscription",
  "AWS::RedshiftServerless::Workgroup",
  "AWS::ResourceExplorer2::Index",
  "AWS::Route53::HealthCheck",
  "AWS::Route53Resolver::ResolverRule",
  "AWS::Route53Resolver::QueryLogConfig",
  "AWS::SageMaker::App",
  "AWS::SageMaker::Domain",
  "AWS::StepFunctions::StateMachine",
  "AWS::StepFunctions::StateMachineExecution",
  "AWS::StepFunctions::StateMachineExecutionHistories",
  "AWS::SimSpaceWeaver::Simulation",
  "AWS::SSM::Association",
  "AWS::SSM::Document",
  "AWS::SSM::DocumentPermission",
  "AWS::EC2::CustomerGateway",
  "AWS::EC2::VerifiedAccessInstance",
  "AWS::EC2::VerifiedAccessEndpoint",
  "AWS::EC2::VerifiedAccessGroup",
  "AWS::EC2::VerifiedAccessTrustProvider",
  "AWS::EC2::VPNGateway",
  "AWS::WAFv2::IPSet",
  "AWS::WAFv2::RegexPatternSet",
  "AWS::WAFv2::RuleGroup",
  "AWS::EC2::TransitGatewayRoute",
  "AWS::GuardDuty::Filter",
  "AWS::ECS::TaskDefinition",
  "AWS::GuardDuty::ThreatIntelSet",
  "AWS::ApiGatewayV2::DomainName",
  "AWS::ApiGateway::DomainName",
  "AWS::ApiGatewayV2::Route",
  "AWS::MQ::Broker",
  "AWS::ACMPCA::CertificateAuthority",
  "AWS::CloudFormation::Stack",
  "AWS::CloudFormation::StackResource",
  "AWS::DirectConnect::Connection",
  "AWS::FSX::FileSystem",
  "AWS::Glue::SecurityConfiguration",
  "AWS::Inspector::AssessmentRun",
  "AWS::Inspector2::Coverage",
  "AWS::Inspector2::CoverageStatistics",
  "AWS::Inspector2::Member",
  "AWS::Inspector2::Finding",
  "AWS::Config::ConfigurationRecorder",
  "AWS::EC2::NatGateway",
  "AWS::ECR::PublicRepository",
  "AWS::ECS::Cluster",
  "AWS::ElasticLoadBalancingV2::TargetGroup",
  "AWS::CloudFront::CachePolicy",
  "AWS::CodeArtifact::Repository",
  "AWS::AMP::Workspace",
  "AWS::EC2::CapacityReservation",
  "AWS::SageMaker::NotebookInstance",
  "AWS::IAM::AccessAdvisor",
  "AWS::EC2::VolumeSnapshot",
  "AWS::EC2::Region",
  "AWS::Keyspaces::Table",
  "AWS::Config::AggregationAuthorization",
  "AWS::DAX::SubnetGroup",
  "AWS::DynamoDb::GlobalTable",
  "AWS::ElasticLoadBalancing::LoadBalancer",
  "AWS::AppStream::Application",
  "AWS::RedshiftServerless::Namespace",
  "AWS::CloudFront::OriginAccessIdentity",
  "AWS::EC2::Host",
  "AWS::EC2::VPC",
  "AWS::EC2::TransitGatewayRouteTable",
  "AWS::EKS::Nodegroup",
  "AWS::Backup::Selection",
  "AWS::CloudTrail::Import",
  "AWS::CostExplorer::ByServiceDaily",
  "AWS::ElasticLoadBalancingV2::SslPolicy",
  "AWS::GuardDuty::Finding",
  "AWS::EC2::DHCPOptions",
  "AWS::EC2::InstanceType",
  "AWS::Batch::ComputeEnvironment",
  "AWS::DMS::ReplicationInstance",
  "AWS::DMS::Endpoint",
  "AWS::DMS::ReplicationTask",
  "AWS::DynamoDb::Table",
  "AWS::Shield::ProtectionGroup",
  "AWS::Firehose::DeliveryStream",
  "AWS::KinesisVideo::Stream",
  "AWS::KMS::Alias",
  "AWS::Lambda::Alias",
  "AWS::Lambda::LambdaLayer",
  "AWS::Lambda::LayerVersion",
  "AWS::Lightsail::Instance",
  "AWS::Macie2::ClassificationJob",
  "AWS::MediaStore::Container",
  "AWS::Mgn::Application",
  "AWS::ResourceExplorer2::SupportedResourceType",
  "AWS::Route53Resolver::ResolverEndpoint",
  "AWS::Route53Domains::Domain",
  "AWS::Route53::Record",
  "AWS::Route53::TrafficPolicy",
  "AWS::Route53::TrafficPolicyInstance",
  "AWS::SageMaker::Model",
  "AWS::SageMaker::TrainingJob",
  "AWS::SecurityHub::ActionTarget",
  "AWS::SecurityHub::Finding",
  "AWS::SecurityHub::FindingAggregator",
  "AWS::SecurityHub::Insight",
  "AWS::SecurityHub::Member",
  "AWS::SecurityHub::Product",
  "AWS::SecurityHub::StandardsControl",
  "AWS::SecurityHub::StandardsSubscription",
  "AWS::SecurityLake::DataLake",
  "AWS::SecurityLake::Subscriber",
  "AWS::Ram::PrincipalAssociation",
  "AWS::Ram::ResourceAssociation",
  "AWS::RDS::ReservedDBInstance",
  "AWS::Redshift::SubnetGroup",
  "AWS::SeverlessApplicationRepository::Application",
  "AWS::AuditManager::Framework",
  "AWS::AuditManager::EvidenceFolder",
  "AWS::AuditManager::Evidence",
  "AWS::AuditManager::Control",
  "AWS::AuditManager::Assessment",
  "AWS::CloudTrail::TrailEvent",
  "AWS::CloudWatch::LogEvent",
  "AWS::Logs::MetricFilter",
  "AWS::CloudWatch::LogStream",
  "AWS::CostExplorer::ByUsageTypeMonthly",
  "AWS::ServiceQuotas::ServiceQuotaChangeRequest",
  "AWS::ServiceQuotas::Service",
  "AWS::EC2::VPCEndpointService",
  "AWS::EC2::LaunchTemplate",
  "AWS::EC2::LaunchTemplateVersion",
  "AWS::SNS::Subscription",
  "AWS::S3::AccountSetting",
  "AWS::SSM::ManagedInstanceCompliance",
  "AWS::SSM::ManagedInstancePatchState",
  "AWS::SSOAdmin::AccountAssignment",
  "AWS::SSOAdmin::UserEffectiveAccess",
  "AWS::SSOAdmin::Instance",
  "AWS::SSOAdmin::PermissionSet",
  "AWS::SSOAdmin::AttachedManagedPolicy",
  "AWS::ServiceDiscovery::Service",
  "AWS::ServiceDiscovery::Namespace",
  "AWS::ServiceDiscovery::Instance",
  "AWS::ServiceCatalog::Portfolio",
  "AWS::ServiceCatalog::Product",
  "AWS::IdentityStore::User",
  "AWS::IdentityStore::Group",
  "AWS::IdentityStore::GroupMembership",
}