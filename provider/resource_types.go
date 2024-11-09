package provider

import (
	model "github.com/opengovern/og-describer-aws/pkg/sdk/models"
	"github.com/opengovern/og-describer-aws/provider/configs"
	"github.com/opengovern/og-describer-aws/provider/describer"
)

var ResourceTypes = map[string]model.ResourceType{

	"AWS::Redshift::Snapshot": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::Redshift::Snapshot",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.RedshiftSnapshot),
		GetDescriber:    ParallelDescribeRegionalSingleResource(describer.GetRedshiftSnapshot),
	},

	"AWS::IAM::AccountSummary": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::IAM::AccountSummary",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   SequentialDescribeGlobal(describer.IAMAccountSummary),
		GetDescriber:    nil,
	},

	"AWS::Glacier::Vault": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::Glacier::Vault",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.GlacierVault),
		GetDescriber:    ParallelDescribeRegionalSingleResource(describer.GetGlacierVault),
	},

	"AWS::Organizations::Organization": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::Organizations::Organization",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   SequentialDescribeGlobal(describer.OrganizationsOrganization),
		GetDescriber:    nil,
	},

	"AWS::Organizations::Policy": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::Organizations::Policy",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   SequentialDescribeGlobal(describer.OrganizationsPolicy),
		GetDescriber:    nil,
	},

	"AWS::Organizations::PolicyTarget": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::Organizations::PolicyTarget",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   SequentialDescribeGlobal(describer.OrganizationsPolicyTarget),
		GetDescriber:    nil,
	},

	"AWS::Organizations::OrganizationalUnit": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::Organizations::OrganizationalUnit",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   SequentialDescribeGlobal(describer.OrganizationsOrganizationalUnit),
		GetDescriber:    nil,
	},

	"AWS::Organizations::Root": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::Organizations::Root",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   SequentialDescribeGlobal(describer.OrganizationsRoot),
		GetDescriber:    nil,
	},

	"AWS::CloudSearch::Domain": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::CloudSearch::Domain",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.CloudSearchDomain),
		GetDescriber:    nil,
	},

	"AWS::DynamoDb::GlobalSecondaryIndex": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::DynamoDb::GlobalSecondaryIndex",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.DynamoDbGlobalSecondaryIndex),
		GetDescriber:    ParallelDescribeRegionalSingleResource(describer.GetDynamoDbGlobalSecondaryIndex),
	},

	"AWS::EC2::RouteTable": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::EC2::RouteTable",
		Tags: map[string][]string{
			"category": {"Networking"},
			"logo_uri": {"https://raw.githubusercontent.com/kaytu-io/awsicons/master/svg-export/icons/Route53RouteTable.svg"},
		},
		Labels:        map[string]string{},
		Annotations:   map[string]string{},
		ListDescriber: ParallelDescribeRegional(describer.EC2RouteTable),
		GetDescriber:  ParallelDescribeRegionalSingleResource(describer.GetEC2RouteTable),
	},

	"AWS::SecurityHub::Hub": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::SecurityHub::Hub",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.SecurityHubHub),
		GetDescriber:    nil,
	},

	"AWS::StorageGateway::StorageGateway": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::StorageGateway::StorageGateway",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.StorageGatewayStorageGateway),
		GetDescriber:    nil,
	},

	"AWS::Inspector::AssessmentTemplate": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::Inspector::AssessmentTemplate",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.InspectorAssessmentTemplate),
		GetDescriber:    ParallelDescribeRegionalSingleResource(describer.GetInspectorAssessmentTemplate),
	},

	"AWS::ElasticLoadBalancingV2::ListenerRule": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::ElasticLoadBalancingV2::ListenerRule",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.ElasticLoadBalancingV2ListenerRule),
		GetDescriber:    nil,
	},

	"AWS::IAM::Role": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::IAM::Role",
		Tags: map[string][]string{
			"category": {"Management & Governance"},
			"logo_uri": {"https://raw.githubusercontent.com/kaytu-io/awsicons/master/svg-export/icons/IdentityAccessManagementRole.svg"},
		},
		Labels:        map[string]string{},
		Annotations:   map[string]string{},
		ListDescriber: SequentialDescribeGlobal(describer.IAMRole),
		GetDescriber:  nil,
	},

	"AWS::Backup::ProtectedResource": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::Backup::ProtectedResource",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.BackupProtectedResource),
		GetDescriber:    nil,
	},

	"AWS::CodeCommit::Repository": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::CodeCommit::Repository",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.CodeCommitRepository),
		GetDescriber:    nil,
	},

	"AWS::EC2::VPCEndpoint": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::EC2::VPCEndpoint",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.EC2VPCEndpoint),
		GetDescriber:    ParallelDescribeRegionalSingleResource(describer.GetEC2VPCEndpoint),
	},

	"AWS::EventBridge::EventRule": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::EventBridge::EventRule",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.EventBridgeRule),
		GetDescriber:    nil,
	},

	"AWS::CloudFront::OriginAccessControl": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::CloudFront::OriginAccessControl",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   SequentialDescribeGlobal(describer.CloudFrontOriginAccessControl),
		GetDescriber:    nil,
	},

	"AWS::CodeBuild::Project": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::CodeBuild::Project",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.CodeBuildProject),
		GetDescriber:    ParallelDescribeRegionalSingleResource(describer.GetCodeBuildProject),
	},

	"AWS::CodeBuild::Build": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::CodeBuild::Build",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.CodeBuildBuild),
		GetDescriber:    ParallelDescribeRegionalSingleResource(describer.GetCodeBuildBuild),
	},

	"AWS::ElastiCache::ParameterGroup": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::ElastiCache::ParameterGroup",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.ElastiCacheParameterGroup),
		GetDescriber:    nil,
	},

	"AWS::MemoryDb::Cluster": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::MemoryDb::Cluster",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.MemoryDbCluster),
		GetDescriber:    nil,
	},

	"AWS::Glue::Crawler": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::Glue::Crawler",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.GlueCrawler),
		GetDescriber:    ParallelDescribeRegionalSingleResource(describer.GetGlueCrawler),
	},

	"AWS::DirectConnect::Gateway": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::DirectConnect::Gateway",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.DirectConnectGateway),
		GetDescriber:    nil,
	},

	"AWS::DynamoDb::BackUp": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::DynamoDb::BackUp",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.DynamoDbBackUp),
		GetDescriber:    nil,
	},

	"AWS::EC2::EIP": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::EC2::EIP",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.EC2EIP),
		GetDescriber:    ParallelDescribeRegionalSingleResource(describer.GetEC2EIP),
	},

	"AWS::EC2::InternetGateway": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::EC2::InternetGateway",
		Tags: map[string][]string{
			"category": {"Networking"},
			"logo_uri": {"https://raw.githubusercontent.com/kaytu-io/awsicons/master/svg-export/icons/InternetGateway.svg"},
		},
		Labels:        map[string]string{},
		Annotations:   map[string]string{},
		ListDescriber: ParallelDescribeRegional(describer.EC2InternetGateway),
		GetDescriber:  ParallelDescribeRegionalSingleResource(describer.GetEC2InternetGateway),
	},

	"AWS::GuardDuty::PublishingDestination": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::GuardDuty::PublishingDestination",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.GuardDutyPublishingDestination),
		GetDescriber:    nil,
	},

	"AWS::KinesisAnalyticsV2::Application": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::KinesisAnalyticsV2::Application",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.KinesisAnalyticsV2Application),
		GetDescriber:    nil,
	},

	"AWS::EMR::Instance": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::EMR::Instance",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.EMRInstance),
		GetDescriber:    nil,
	},

	"AWS::EMR::BlockPublicAccessConfiguration": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::EMR::BlockPublicAccessConfiguration",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.EMRBlockPublicAccessConfiguration),
		GetDescriber:    nil,
	},

	"AWS::ApiGateway::RestApi": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::ApiGateway::RestApi",
		Tags: map[string][]string{
			"category": {"Networking"},
			"logo_uri": {"https://raw.githubusercontent.com/kaytu-io/awsicons/master/svg-export/icons/ApiGateway.svg"},
		},
		Labels:        map[string]string{},
		Annotations:   map[string]string{},
		ListDescriber: ParallelDescribeRegional(describer.ApiGatewayRestAPI),
		GetDescriber:  ParallelDescribeRegionalSingleResource(describer.GetApiGatewayRestAPI),
	},

	"AWS::ApiGatewayV2::Integration": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::ApiGatewayV2::Integration",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.ApiGatewayV2Integration),
		GetDescriber:    ParallelDescribeRegionalSingleResource(describer.GetApiGatewayV2Integration),
	},

	"AWS::AutoScaling::AutoScalingGroup": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::AutoScaling::AutoScalingGroup",
		Tags: map[string][]string{
			"category": {"Compute"},
			"logo_uri": {"https://raw.githubusercontent.com/kaytu-io/awsicons/master/svg-export/icons/Ec2AutoScaling.svg"},
		},
		Labels:        map[string]string{},
		Annotations:   map[string]string{},
		ListDescriber: ParallelDescribeRegional(describer.AutoScalingAutoScalingGroup),
		GetDescriber:  ParallelDescribeRegionalSingleResource(describer.GetAutoScalingAutoScalingGroup),
	},

	"AWS::DynamoDb::TableExport": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::DynamoDb::TableExport",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.DynamoDbTableExport),
		GetDescriber:    nil,
	},

	"AWS::EC2::KeyPair": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::EC2::KeyPair",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.EC2KeyPair),
		GetDescriber:    ParallelDescribeRegionalSingleResource(describer.GetEC2KeyPair),
	},

	"AWS::EFS::FileSystem": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::EFS::FileSystem",
		Tags: map[string][]string{
			"category": {"Storage"},
			"logo_uri": {"https://raw.githubusercontent.com/kaytu-io/awsicons/master/svg-export/icons/ElasticFileSystem.svg"},
		},
		Labels:        map[string]string{},
		Annotations:   map[string]string{},
		ListDescriber: ParallelDescribeRegional(describer.EFSFileSystem),
		GetDescriber:  nil,
	},

	"AWS::Kafka::Cluster": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::Kafka::Cluster",
		Tags: map[string][]string{
			"category": {"PaaS"},
			"logo_uri": {"https://raw.githubusercontent.com/kaytu-io/awsicons/master/svg-export/icons/ManagedStreamingForKafka.svg"},
		},
		Labels:        map[string]string{},
		Annotations:   map[string]string{},
		ListDescriber: ParallelDescribeRegional(describer.KafkaCluster),
		GetDescriber:  nil,
	},

	"AWS::SecretsManager::Secret": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::SecretsManager::Secret",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.SecretsManagerSecret),
		GetDescriber:    nil,
	},

	"AWS::Backup::LegalHold": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::Backup::LegalHold",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.BackupLegalHold),
		GetDescriber:    nil,
	},

	"AWS::CloudFront::Function": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::CloudFront::Function",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   SequentialDescribeGlobal(describer.CloudFrontFunction),
		GetDescriber:    nil,
	},

	"AWS::GlobalAccelerator::EndpointGroup": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::GlobalAccelerator::EndpointGroup",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.GlobalAcceleratorEndpointGroup),
		GetDescriber:    nil,
	},

	"AWS::DAX::ParameterGroup": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::DAX::ParameterGroup",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.DAXParameterGroup),
		GetDescriber:    nil,
	},

	"AWS::SQS::Queue": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::SQS::Queue",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.SQSQueue),
		GetDescriber:    nil,
	},

	"AWS::Config::Rule": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::Config::Rule",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.ConfigRule),
		GetDescriber:    nil,
	},

	"AWS::GuardDuty::Member": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::GuardDuty::Member",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.GuardDutyMember),
		GetDescriber:    nil,
	},

	"AWS::Inspector::Exclusion": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::Inspector::Exclusion",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.InspectorExclusion),
		GetDescriber:    nil,
	},

	"AWS::DirectoryService::Directory": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::DirectoryService::Directory",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.DirectoryServiceDirectory),
		GetDescriber:    nil,
	},

	"AWS::DirectoryService::Certificate": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::DirectoryService::Certificate",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.DirectoryServiceCertificate),
		GetDescriber:    nil,
	},

	"AWS::DirectoryService::LogSubscription": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::DirectoryService::LogSubscription",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.DirectoryServiceLogSubscription),
		GetDescriber:    nil,
	},

	"AWS::EFS::AccessPoint": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::EFS::AccessPoint",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.EFSAccessPoint),
		GetDescriber:    nil,
	},

	"AWS::IAM::PolicyAttachment": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::IAM::PolicyAttachment",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   SequentialDescribeGlobal(describer.IAMPolicyAttachment),
		GetDescriber:    nil,
	},

	"AWS::IAM::CredentialReport": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::IAM::CredentialReport",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   SequentialDescribeGlobal(describer.IAMCredentialReport),
		GetDescriber:    nil,
	},

	"AWS::RDS::GlobalCluster": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::RDS::GlobalCluster",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.RDSGlobalCluster),
		GetDescriber:    nil,
	},

	"AWS::CostExplorer::ForcastDaily": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::CostExplorer::ForcastDaily",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   SequentialDescribeGlobal(describer.CostForecastDaily),
		GetDescriber:    nil,
	},

	"AWS::GuardDuty::Detector": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::GuardDuty::Detector",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.GuardDutyDetector),
		GetDescriber:    nil,
	},

	"AWS::SNS::Topic": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::SNS::Topic",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.SNSTopic),
		GetDescriber:    nil,
	},

	"AWS::AppConfig::Application": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::AppConfig::Application",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.AppConfigApplication),
		GetDescriber:    nil,
	},

	"AWS::Batch::Job": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::Batch::Job",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.BatchJob),
		GetDescriber:    nil,
	},

	"AWS::Batch::JobQueue": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::Batch::JobQueue",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.BatchJobQueue),
		GetDescriber:    nil,
	},

	"AWS::ECS::Service": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::ECS::Service",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.ECSService),
		GetDescriber:    ParallelDescribeRegionalSingleResource(describer.GetECSService),
	},

	"AWS::FSX::Task": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::FSX::Task",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.FSXTask),
		GetDescriber:    nil,
	},

	"AWS::IAM::VirtualMFADevice": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::IAM::VirtualMFADevice",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   SequentialDescribeGlobal(describer.IAMVirtualMFADevice),
		GetDescriber:    nil,
	},

	"AWS::WAFv2::WebACL": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::WAFv2::WebACL",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.WAFv2WebACL),
		GetDescriber:    nil,
	},

	"AWS::ApplicationAutoScaling::Target": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::ApplicationAutoScaling::Target",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.ApplicationAutoScalingTarget),
		GetDescriber:    nil,
	},

	"AWS::ApplicationAutoScaling::Policy": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::ApplicationAutoScaling::Policy",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.ApplicationAutoScalingPolicy),
		GetDescriber:    nil,
	},

	"AWS::Backup::Vault": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::Backup::Vault",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.BackupVault),
		GetDescriber:    nil,
	},

	"AWS::ElastiCache::Cluster": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::ElastiCache::Cluster",
		Tags: map[string][]string{
			"category": {"Database"},
			"logo_uri": {"https://raw.githubusercontent.com/kaytu-io/awsicons/master/svg-export/icons/ElastiCache.svg"},
		},
		Labels:        map[string]string{},
		Annotations:   map[string]string{},
		ListDescriber: ParallelDescribeRegional(describer.ElastiCacheCluster),
		GetDescriber:  ParallelDescribeRegionalSingleResource(describer.GetElastiCacheCluster),
	},

	"AWS::Logs::LogGroup": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::Logs::LogGroup",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.CloudWatchLogsLogGroup),
		GetDescriber:    nil,
	},

	"AWS::S3::Bucket": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::S3::Bucket",
		Tags: map[string][]string{
			"category": {"Storage"},
			"logo_uri": {"https://raw.githubusercontent.com/kaytu-io/awsicons/master/svg-export/icons/SimpleStorageBucket.svg"},
		},
		Labels:        map[string]string{},
		Annotations:   map[string]string{},
		ListDescriber: SequentialDescribeGlobal(describer.S3Bucket),
		GetDescriber:  nil,
	},

	"AWS::S3::Object": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::S3::Object",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   SequentialDescribeGlobal(describer.S3Object),
		GetDescriber:    nil,
	},

	"AWS::S3::BucketIntelligentTieringConfiguration": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::S3::BucketIntelligentTieringConfiguration",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   SequentialDescribeGlobal(describer.S3BucketIntelligentTieringConfiguration),
		GetDescriber:    nil,
	},

	"AWS::S3::MultiRegionAccessPoint": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::S3::MultiRegionAccessPoint",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   SequentialDescribeGlobal(describer.S3MultiRegionAccessPoint),
		GetDescriber:    nil,
	},

	"AWS::CertificateManager::Certificate": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::CertificateManager::Certificate",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.CertificateManagerCertificate),
		GetDescriber:    nil,
	},

	"AWS::EKS::AddonVersion": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::EKS::AddonVersion",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.EKSAddonVersion),
		GetDescriber:    nil,
	},

	"AWS::ApiGatewayV2::Api": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::ApiGatewayV2::Api",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.ApiGatewayV2API),
		GetDescriber:    ParallelDescribeRegionalSingleResource(describer.GetApiGatewayV2API),
	},

	"AWS::EC2::Volume": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::EC2::Volume",
		Tags: map[string][]string{
			"category": {"Storage"},
			"logo_uri": {"https://raw.githubusercontent.com/kaytu-io/awsicons/master/svg-export/icons/ElasticBlockStoreVolume.svg"},
		},
		Labels:        map[string]string{},
		Annotations:   map[string]string{},
		ListDescriber: ParallelDescribeRegional(describer.EC2Volume),
		GetDescriber:  ParallelDescribeRegionalSingleResource(describer.GetEC2Volume),
	},

	"AWS::ApiGateway::ApiKey": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::ApiGateway::ApiKey",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.ApiGatewayApiKey),
		GetDescriber:    nil,
	},

	"AWS::Glue::Connection": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::Glue::Connection",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.GlueConnection),
		GetDescriber:    nil,
	},

	"AWS::ECS::Task": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::ECS::Task",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.ECSTask),
		GetDescriber:    nil,
	},

	"AWS::SSM::ManagedInstance": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::SSM::ManagedInstance",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.SSMManagedInstance),
		GetDescriber:    nil,
	},

	"AWS::SSM::Inventory": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::SSM::Inventory",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.SSMInventory),
		GetDescriber:    nil,
	},

	"AWS::SSM::InventoryEntry": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::SSM::InventoryEntry",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.SSMInventoryEntry),
		GetDescriber:    nil,
	},

	"AWS::SSM::MaintenanceWindow": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::SSM::MaintenanceWindow",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.SSMMaintenanceWindow),
		GetDescriber:    nil,
	},

	"AWS::SSM::PatchBaseline": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::SSM::PatchBaseline",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.SSMPatchBaseline),
		GetDescriber:    nil,
	},

	"AWS::SSM::Parameter": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::SSM::Parameter",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.SSMParameter),
		GetDescriber:    nil,
	},

	"AWS::Lambda::Function": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::Lambda::Function",
		Tags: map[string][]string{
			"category": {"Serverless"},
			"logo_uri": {"https://raw.githubusercontent.com/kaytu-io/awsicons/master/svg-export/icons/Lambda.svg"},
		},
		Labels:        map[string]string{},
		Annotations:   map[string]string{},
		ListDescriber: ParallelDescribeRegional(describer.LambdaFunction),
		GetDescriber:  ParallelDescribeRegionalSingleResource(describer.GetLambdaFunction),
	},

	"AWS::RDS::DBSnapshot": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::RDS::DBSnapshot",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.RDSDBSnapshot),
		GetDescriber:    nil,
	},

	"AWS::CodeDeploy::Application": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::CodeDeploy::Application",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.CodeDeployApplication),
		GetDescriber:    nil,
	},

	"AWS::CodeDeploy::DeploymentConfig": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::CodeDeploy::DeploymentConfig",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.CodeDeployDeploymentConfig),
		GetDescriber:    nil,
	},

	"AWS::EMR::Cluster": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::EMR::Cluster",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.EMRCluster),
		GetDescriber:    nil,
	},

	"AWS::IAM::AccessKey": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::IAM::AccessKey",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   SequentialDescribeGlobal(describer.IAMAccessKey),
		GetDescriber:    nil,
	},

	"AWS::IAM::SSHPublicKey": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::IAM::SSHPublicKey",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   SequentialDescribeGlobal(describer.IAMSSHPublicKey),
		GetDescriber:    nil,
	},

	"AWS::Glue::CatalogTable": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::Glue::CatalogTable",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.GlueCatalogTable),
		GetDescriber:    nil,
	},

	"AWS::CloudTrail::Channel": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::CloudTrail::Channel",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.CloudTrailChannel),
		GetDescriber:    nil,
	},

	"AWS::EC2::NetworkAcl": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::EC2::NetworkAcl",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.EC2NetworkAcl),
		GetDescriber:    ParallelDescribeRegionalSingleResource(describer.GetEC2NetworkAcl),
	},

	"AWS::ECS::ContainerInstance": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::ECS::ContainerInstance",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.ECSContainerInstance),
		GetDescriber:    nil,
	},

	"AWS::RedshiftServerless::Snapshot": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::RedshiftServerless::Snapshot",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.RedshiftServerlessSnapshot),
		GetDescriber:    nil,
	},

	"AWS::Workspaces::Bundle": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::Workspaces::Bundle",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.WorkspacesBundle),
		GetDescriber:    nil,
	},

	"AWS::CloudTrail::Trail": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::CloudTrail::Trail",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.CloudTrailTrail),
		GetDescriber:    nil,
	},

	"AWS::DAX::Parameter": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::DAX::Parameter",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.DAXParameter),
		GetDescriber:    nil,
	},

	"AWS::ECR::Image": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::ECR::Image",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.ECRImage),
		GetDescriber:    nil,
	},

	"AWS::IAM::ServerCertificate": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::IAM::ServerCertificate",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   SequentialDescribeGlobal(describer.IAMServerCertificate),
		GetDescriber:    nil,
	},

	"AWS::Keyspaces::Keyspace": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::Keyspaces::Keyspace",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.KeyspacesKeyspace),
		GetDescriber:    nil,
	},

	"AWS::S3::AccessPoint": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::S3::AccessPoint",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.S3AccessPoint),
		GetDescriber:    nil,
	},

	"AWS::SageMaker::EndpointConfiguration": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::SageMaker::EndpointConfiguration",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.SageMakerEndpointConfiguration),
		GetDescriber:    nil,
	},

	"AWS::ElastiCache::ReservedCacheNode": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::ElastiCache::ReservedCacheNode",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.ElastiCacheReservedCacheNode),
		GetDescriber:    nil,
	},

	"AWS::EMR::InstanceFleet": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::EMR::InstanceFleet",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.EMRInstanceFleet),
		GetDescriber:    nil,
	},

	"AWS::Account::Account": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::Account::Account",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   SequentialDescribeGlobal(describer.IAMAccount),
		GetDescriber:    nil,
	},

	"AWS::EC2::VPCPeeringConnection": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::EC2::VPCPeeringConnection",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.EC2VPCPeeringConnection),
		GetDescriber:    ParallelDescribeRegionalSingleResource(describer.GetEC2VPCPeeringConnection),
	},

	"AWS::EKS::FargateProfile": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::EKS::FargateProfile",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.EKSFargateProfile),
		GetDescriber:    nil,
	},

	"AWS::IAM::AccountPasswordPolicy": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::IAM::AccountPasswordPolicy",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   SequentialDescribeGlobal(describer.IAMAccountPasswordPolicy),
		GetDescriber:    nil,
	},

	"AWS::CodePipeline::Pipeline": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::CodePipeline::Pipeline",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.CodePipelinePipeline),
		GetDescriber:    nil,
	},

	"AWS::DAX::Cluster": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::DAX::Cluster",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.DAXCluster),
		GetDescriber:    nil,
	},

	"AWS::DLM::LifecyclePolicy": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::DLM::LifecyclePolicy",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.DLMLifecyclePolicy),
		GetDescriber:    nil,
	},

	"AWS::OpsWorksCM::Server": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::OpsWorksCM::Server",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.OpsWorksCMServer),
		GetDescriber:    nil,
	},

	"AWS::AccessAnalyzer::Analyzer": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::AccessAnalyzer::Analyzer",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.AccessAnalyzerAnalyzer),
		GetDescriber:    ParallelDescribeRegionalSingleResource(describer.GetAccessAnalyzerAnalyzer),
	},

	"AWS::AccessAnalyzer::Finding": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::AccessAnalyzer::Finding",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.AccessAnalyzerAnalyzerFinding),
		GetDescriber:    nil,
	},

	"AWS::ElastiCache::SubnetGroup": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::ElastiCache::SubnetGroup",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.ElastiCacheSubnetGroup),
		GetDescriber:    nil,
	},

	"AWS::FSX::Volume": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::FSX::Volume",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.FSXVolume),
		GetDescriber:    nil,
	},

	"AWS::Amplify::App": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::Amplify::App",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.AmplifyApp),
		GetDescriber:    nil,
	},

	"AWS::CloudTrail::Query": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::CloudTrail::Query",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.CloudTrailQuery),
		GetDescriber:    nil,
	},

	"AWS::CostExplorer::ByAccountMonthly": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::CostExplorer::ByAccountMonthly",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   SequentialDescribeGlobal(describer.CostByAccountLastMonth),
		GetDescriber:    nil,
	},

	"AWS::ECR::PublicRegistry": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::ECR::PublicRegistry",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.ECRPublicRegistry),
		GetDescriber:    nil,
	},

	"AWS::EC2::NetworkInterface": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::EC2::NetworkInterface",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.EC2NetworkInterface),
		GetDescriber:    ParallelDescribeRegionalSingleResource(describer.GetEC2NetworkInterface),
	},

	"AWS::EC2::VPNConnection": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::EC2::VPNConnection",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.EC2VPNConnection),
		GetDescriber:    ParallelDescribeRegionalSingleResource(describer.GetEC2VPNConnection),
	},

	"AWS::FSX::StorageVirtualMachine": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::FSX::StorageVirtualMachine",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.FSXStorageVirtualMachine),
		GetDescriber:    nil,
	},

	"AWS::ApiGateway::Authorizer": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::ApiGateway::Authorizer",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.ApiGatewayAuthorizer),
		GetDescriber:    nil,
	},

	"AWS::AppStream::Stack": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::AppStream::Stack",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.AppStreamStack),
		GetDescriber:    nil,
	},

	"AWS::Athena::WorkGroup": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::Athena::WorkGroup",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.AthenaWrokgroup),
		GetDescriber:    nil,
	},

	"AWS::Athena::QueryExecution": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::Athena::QueryExecution",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.AthenaQueryExecution),
		GetDescriber:    nil,
	},

	"AWS::AppStream::Image": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::AppStream::Image",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.AppStreamImage),
		GetDescriber:    ParallelDescribeRegionalSingleResource(describer.GetAppStreamImage),
	},

	"AWS::CloudWatch::Alarm": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::CloudWatch::Alarm",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.CloudWatchAlarm),
		GetDescriber:    ParallelDescribeRegionalSingleResource(describer.GetCloudWatchAlarm),
	},

	"AWS::CloudWatch::LogSubscriptionFilter": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::CloudWatch::LogSubscriptionFilter",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   SequentialDescribeRegional(describer.CloudWatchLogsSubscriptionFilter),
		GetDescriber:    nil,
	},

	"AWS::CostExplorer::ByRecordTypeMonthly": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::CostExplorer::ByRecordTypeMonthly",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   SequentialDescribeGlobal(describer.CostByRecordTypeLastMonth),
		GetDescriber:    nil,
	},

	"AWS::RDS::DBCluster": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::RDS::DBCluster",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.RDSDBCluster),
		GetDescriber:    ParallelDescribeRegionalSingleResource(describer.GetRDSDBCluster),
	},

	"AWS::RDS::DBClusterSnapshot": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::RDS::DBClusterSnapshot",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.RDSDBClusterSnapshot),
		GetDescriber:    nil,
	},

	"AWS::Backup::Framework": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::Backup::Framework",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.BackupFramework),
		GetDescriber:    nil,
	},

	"AWS::CodeBuild::SourceCredential": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::CodeBuild::SourceCredential",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.CodeBuildSourceCredential),
		GetDescriber:    nil,
	},

	"AWS::IAM::ServiceSpecificCredential": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::IAM::ServiceSpecificCredential",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   SequentialDescribeGlobal(describer.IAMServiceSpecificCredential),
		GetDescriber:    nil,
	},

	"AWS::EC2::CapacityReservationFleet": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::EC2::CapacityReservationFleet",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.EC2CapacityReservationFleet),
		GetDescriber:    ParallelDescribeRegionalSingleResource(describer.GetEC2CapacityReservationFleet),
	},

	"AWS::NetworkFirewall::Firewall": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::NetworkFirewall::Firewall",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.NetworkFirewallFirewall),
		GetDescriber:    nil,
	},

	"AWS::Workspaces::Workspace": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::Workspaces::Workspace",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.WorkspacesWorkspace),
		GetDescriber:    nil,
	},

	"AWS::ElasticSearch::Domain": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::ElasticSearch::Domain",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.ESDomain),
		GetDescriber:    nil,
	},

	"AWS::RDS::DBInstance": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::RDS::DBInstance",
		Tags: map[string][]string{
			"category": {"Database"},
			"logo_uri": {"https://raw.githubusercontent.com/kaytu-io/awsicons/master/svg-export/icons/Rds.svg"},
		},
		Labels:        map[string]string{},
		Annotations:   map[string]string{},
		ListDescriber: ParallelDescribeRegional(describer.RDSDBInstance),
		GetDescriber:  ParallelDescribeRegionalSingleResource(describer.GetRDSDBInstance),
	},

	"AWS::RDS::DBInstanceAutomatedBackup": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::RDS::DBInstanceAutomatedBackup",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.RDSDBInstanceAutomatedBackup),
		GetDescriber:    ParallelDescribeRegionalSingleResource(describer.GetRDSDBInstanceAutomatedBackup),
	},

	"AWS::EFS::MountTarget": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::EFS::MountTarget",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.EFSMountTarget),
		GetDescriber:    nil,
	},

	"AWS::GlobalAccelerator::Listener": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::GlobalAccelerator::Listener",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.GlobalAcceleratorListener),
		GetDescriber:    nil,
	},

	"AWS::CostExplorer::ByUsageTypeDaily": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::CostExplorer::ByUsageTypeDaily",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   SequentialDescribeGlobal(describer.CostByServiceUsageLastDay),
		GetDescriber:    nil,
	},

	"AWS::EKS::Addon": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::EKS::Addon",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.EKSAddon),
		GetDescriber:    nil,
	},

	"AWS::CostExplorer::ByServiceMonthly": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::CostExplorer::ByServiceMonthly",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   SequentialDescribeGlobal(describer.CostByServiceLastMonth),
		GetDescriber:    nil,
	},

	"AWS::IAM::Policy": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::IAM::Policy",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   SequentialDescribeGlobal(describer.IAMPolicy),
		GetDescriber:    nil,
	},

	"AWS::Redshift::Cluster": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::Redshift::Cluster",
		Tags: map[string][]string{
			"category": {"Big Data"},
			"logo_uri": {""},
		},
		Labels:        map[string]string{},
		Annotations:   map[string]string{},
		ListDescriber: ParallelDescribeRegional(describer.RedshiftCluster),
		GetDescriber:  nil,
	},

	"AWS::WAFRegional::Rule": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::WAFRegional::Rule",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.WAFRegionalRule),
		GetDescriber:    nil,
	},

	"AWS::WAFRegional::RuleGroup": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::WAFRegional::RuleGroup",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.WAFRegionalRuleGroup),
		GetDescriber:    nil,
	},

	"AWS::Glue::DataCatalogEncryptionSettings": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::Glue::DataCatalogEncryptionSettings",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.GlueDataCatalogEncryptionSettings),
		GetDescriber:    nil,
	},

	"AWS::EC2::FlowLog": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::EC2::FlowLog",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.EC2FlowLog),
		GetDescriber:    ParallelDescribeRegionalSingleResource(describer.GetEC2FlowLog),
	},

	"AWS::EC2::IpamPool": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::EC2::IpamPool",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.EC2IpamPool),
		GetDescriber:    ParallelDescribeRegionalSingleResource(describer.GetEC2IpamPool),
	},

	"AWS::IAM::SamlProvider": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::IAM::SamlProvider",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   SequentialDescribeGlobal(describer.IAMSamlProvider),
		GetDescriber:    nil,
	},

	"AWS::Route53::HostedZone": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::Route53::HostedZone",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   SequentialDescribeGlobal(describer.Route53HostedZone),
		GetDescriber:    nil,
	},

	"AWS::Route53::QueryLog": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::Route53::QueryLog",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   SequentialDescribeGlobal(describer.Route53QueryLog),
		GetDescriber:    nil,
	},

	"AWS::EC2::PlacementGroup": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::EC2::PlacementGroup",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.EC2PlacementGroup),
		GetDescriber:    ParallelDescribeRegionalSingleResource(describer.GetEC2PlacementGroup),
	},

	"AWS::FSX::Snapshot": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::FSX::Snapshot",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.FSXSnapshot),
		GetDescriber:    nil,
	},

	"AWS::KMS::Key": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::KMS::Key",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.KMSKey),
		GetDescriber:    ParallelDescribeRegionalSingleResource(describer.GetKMSKey),
	},

	"AWS::KMS::KeyRotation": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::KMS::KeyRotation",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.KMSKeyRotation),
		GetDescriber:    nil,
	},

	"AWS::EC2::Ipam": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::EC2::Ipam",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.EC2Ipam),
		GetDescriber:    ParallelDescribeRegionalSingleResource(describer.GetEC2Ipam),
	},

	"AWS::ElasticBeanstalk::Environment": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::ElasticBeanstalk::Environment",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.ElasticBeanstalkEnvironment),
		GetDescriber:    nil,
	},

	"AWS::ElasticBeanstalk::ApplicationVersion": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::ElasticBeanstalk::ApplicationVersion",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.ElasticBeanstalkApplicationVersion),
		GetDescriber:    nil,
	},

	"AWS::Lambda::FunctionVersion": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::Lambda::FunctionVersion",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.LambdaFunctionVersion),
		GetDescriber:    nil,
	},

	"AWS::Glue::DevEndpoint": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::Glue::DevEndpoint",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.GlueDevEndpoint),
		GetDescriber:    nil,
	},

	"AWS::Backup::RecoveryPoint": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::Backup::RecoveryPoint",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.BackupRecoveryPoint),
		GetDescriber:    nil,
	},

	"AWS::Backup::ReportPlan": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::Backup::ReportPlan",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.BackupReportPlan),
		GetDescriber:    ParallelDescribeRegionalSingleResource(describer.GetBackupReportPlan),
	},

	"AWS::Backup::RegionSetting": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::Backup::RegionSetting",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.BackupRegionSetting),
		GetDescriber:    nil,
	},

	"AWS::DynamoDbStreams::Stream": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::DynamoDbStreams::Stream",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.DynamoDbStream),
		GetDescriber:    nil,
	},

	"AWS::EC2::EgressOnlyInternetGateway": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::EC2::EgressOnlyInternetGateway",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.EC2EgressOnlyInternetGateway),
		GetDescriber:    ParallelDescribeRegionalSingleResource(describer.GetEC2EgressOnlyInternetGateway),
	},

	"AWS::CloudFront::Distribution": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::CloudFront::Distribution",
		Tags: map[string][]string{
			"category": {"Networking"},
			"logo_uri": {"https://raw.githubusercontent.com/kaytu-io/awsicons/master/svg-export/icons/CloudFront.svg"},
		},
		Labels:        map[string]string{},
		Annotations:   map[string]string{},
		ListDescriber: SequentialDescribeGlobal(describer.CloudFrontDistribution),
		GetDescriber:  nil,
	},

	"AWS::CloudFront::StreamingDistribution": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::CloudFront::StreamingDistribution",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   SequentialDescribeGlobal(describer.CloudFrontStreamingDistribution),
		GetDescriber:    nil,
	},

	"AWS::Glue::Job": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::Glue::Job",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.GlueJob),
		GetDescriber:    ParallelDescribeRegionalSingleResource(describer.GetGlueJob),
	},

	"AWS::AppStream::Fleet": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::AppStream::Fleet",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.AppStreamFleet),
		GetDescriber:    nil,
	},

	"AWS::SES::ConfigurationSet": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::SES::ConfigurationSet",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.SESConfigurationSet),
		GetDescriber:    nil,
	},

	"AWS::IAM::User": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::IAM::User",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   SequentialDescribeGlobal(describer.IAMUser),
		GetDescriber:    nil,
	},

	"AWS::CloudFront::OriginRequestPolicy": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::CloudFront::OriginRequestPolicy",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   SequentialDescribeGlobal(describer.CloudFrontOriginRequestPolicy),
		GetDescriber:    nil,
	},

	"AWS::EC2::SecurityGroup": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::EC2::SecurityGroup",
		Tags: map[string][]string{
			"category": {"Networking"},
		},
		Labels:        map[string]string{},
		Annotations:   map[string]string{},
		ListDescriber: ParallelDescribeRegional(describer.EC2SecurityGroup),
		GetDescriber:  ParallelDescribeRegionalSingleResource(describer.GetEC2SecurityGroup),
	},

	"AWS::GuardDuty::IPSet": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::GuardDuty::IPSet",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.GuardDutyIPSet),
		GetDescriber:    nil,
	},

	"AWS::EKS::Cluster": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::EKS::Cluster",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.EKSCluster),
		GetDescriber:    nil,
	},

	"AWS::Grafana::Workspace": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::Grafana::Workspace",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.GrafanaWorkspace),
		GetDescriber:    nil,
	},

	"AWS::Glue::CatalogDatabase": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::Glue::CatalogDatabase",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.GlueCatalogDatabase),
		GetDescriber:    nil,
	},

	"AWS::Health::Event": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::Health::Event",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.HealthEvent),
		GetDescriber:    nil,
	},

	"AWS::Health::AffectedEntity": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::Health::AffectedEntity",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.HealthAffectedEntity),
		GetDescriber:    nil,
	},

	"AWS::CloudFormation::StackSet": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::CloudFormation::StackSet",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.CloudFormationStackSet),
		GetDescriber:    nil,
	},

	"AWS::EC2::AvailabilityZone": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::EC2::AvailabilityZone",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.EC2AvailabilityZone),
		GetDescriber:    ParallelDescribeRegionalSingleResource(describer.GetEC2AvailabilityZone),
	},

	"AWS::EC2::TransitGateway": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::EC2::TransitGateway",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.EC2TransitGateway),
		GetDescriber:    ParallelDescribeRegionalSingleResource(describer.GetEC2TransitGateway),
	},

	"AWS::ApiGateway::UsagePlan": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::ApiGateway::UsagePlan",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.ApiGatewayUsagePlan),
		GetDescriber:    nil,
	},

	"AWS::Inspector::Finding": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::Inspector::Finding",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.InspectorFinding),
		GetDescriber:    nil,
	},

	"AWS::EC2::Fleet": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::EC2::Fleet",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.EC2Fleet),
		GetDescriber:    ParallelDescribeRegionalSingleResource(describer.GetEC2Fleet),
	},

	"AWS::ElasticBeanstalk::Application": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::ElasticBeanstalk::Application",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.ElasticBeanstalkApplication),
		GetDescriber:    nil,
	},

	"AWS::ElasticLoadBalancingV2::LoadBalancer": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::ElasticLoadBalancingV2::LoadBalancer",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.ElasticLoadBalancingV2LoadBalancer),
		GetDescriber:    ParallelDescribeRegionalSingleResource(describer.GetElasticLoadBalancingV2LoadBalancer),
	},

	"AWS::OpenSearch::Domain": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::OpenSearch::Domain",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.OpenSearchDomain),
		GetDescriber:    nil,
	},

	"AWS::RDS::DBEventSubscription": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::RDS::DBEventSubscription",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.RDSDBEventSubscription),
		GetDescriber:    nil,
	},

	"AWS::RDS::DBEngineVersion": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::RDS::DBEngineVersion",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.RDSDBEngineVersion),
		GetDescriber:    nil,
	},

	"AWS::EC2::RegionalSettings": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::EC2::RegionalSettings",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.EC2RegionalSettings),
		GetDescriber:    nil,
	},

	"AWS::EC2::SecurityGroupRule": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::EC2::SecurityGroupRule",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.EC2SecurityGroupRule),
		GetDescriber:    nil,
	},

	"AWS::EC2::TransitGatewayAttachment": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::EC2::TransitGatewayAttachment",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.EC2TransitGatewayAttachment),
		GetDescriber:    ParallelDescribeRegionalSingleResource(describer.GetEC2TransitGatewayAttachment),
	},

	"AWS::SES::Identity": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::SES::Identity",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.SESIdentity),
		GetDescriber:    nil,
	},

	"AWS::SESv2::EmailIdentities": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::SESv2::EmailIdentities",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.SESv2EmailIdentities),
		GetDescriber:    nil,
	},

	"AWS::WAF::Rule": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::WAF::Rule",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.WAFRule),
		GetDescriber:    nil,
	},

	"AWS::WAF::RuleGroup": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::WAF::RuleGroup",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.WAFRuleGroup),
		GetDescriber:    nil,
	},

	"AWS::WAF::RateBasedRule": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::WAF::RateBasedRule",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.WAFRateBasedRule),
		GetDescriber:    nil,
	},

	"AWS::WAF::WebACL": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::WAF::WebACL",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.WAFWebACL),
		GetDescriber:    nil,
	},

	"AWS::WAFRegional::WebACL": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::WAFRegional::WebACL",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.WAFRegionalWebACL),
		GetDescriber:    nil,
	},

	"AWS::WellArchitected::Workload": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::WellArchitected::Workload",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.WellArchitectedWorkload),
		GetDescriber:    nil,
	},

	"AWS::WellArchitected::Answer": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::WellArchitected::Answer",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.WellArchitectedAnswer),
		GetDescriber:    nil,
	},

	"AWS::WellArchitected::CheckDetail": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::WellArchitected::CheckDetail",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.WellArchitectedCheckDetail),
		GetDescriber:    nil,
	},

	"AWS::WellArchitected::CheckSummary": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::WellArchitected::CheckSummary",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.WellArchitectedCheckSummary),
		GetDescriber:    nil,
	},

	"AWS::WellArchitected::ConsolidatedReport": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::WellArchitected::ConsolidatedReport",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.WellArchitectedConsolidatedReport),
		GetDescriber:    nil,
	},

	"AWS::WellArchitected::Lens": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::WellArchitected::Lens",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.WellArchitectedLens),
		GetDescriber:    nil,
	},

	"AWS::WellArchitected::LensReview": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::WellArchitected::LensReview",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.WellArchitectedLensReview),
		GetDescriber:    nil,
	},

	"AWS::WellArchitected::LensReviewImprovement": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::WellArchitected::LensReviewImprovement",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.WellArchitectedLensReviewImprovement),
		GetDescriber:    nil,
	},

	"AWS::WellArchitected::LensReviewReport": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::WellArchitected::LensReviewReport",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.WellArchitectedLensReviewReport),
		GetDescriber:    nil,
	},

	"AWS::WellArchitected::LensShare": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::WellArchitected::LensShare",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.WellArchitectedLensShare),
		GetDescriber:    nil,
	},

	"AWS::WellArchitected::Milestone": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::WellArchitected::Milestone",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.WellArchitectedMilestone),
		GetDescriber:    nil,
	},

	"AWS::WellArchitected::Notification": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::WellArchitected::Notification",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.WellArchitectedNotification),
		GetDescriber:    nil,
	},

	"AWS::WellArchitected::ShareInvitation": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::WellArchitected::ShareInvitation",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.WellArchitectedShareInvitation),
		GetDescriber:    nil,
	},

	"AWS::WellArchitected::WorkloadShare": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::WellArchitected::WorkloadShare",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.WellArchitectedWorkloadShare),
		GetDescriber:    nil,
	},

	"AWS::AutoScaling::LaunchConfiguration": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::AutoScaling::LaunchConfiguration",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.AutoScalingLaunchConfiguration),
		GetDescriber:    ParallelDescribeRegionalSingleResource(describer.GetAutoScalingLaunchConfiguration),
	},

	"AWS::CloudTrail::EventDataStore": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::CloudTrail::EventDataStore",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.CloudTrailEventDataStore),
		GetDescriber:    nil,
	},

	"AWS::CodeDeploy::DeploymentGroup": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::CodeDeploy::DeploymentGroup",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.CodeDeployDeploymentGroup),
		GetDescriber:    nil,
	},

	"AWS::ImageBuilder::Image": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::ImageBuilder::Image",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.ImageBuilderImage),
		GetDescriber:    nil,
	},

	"AWS::Redshift::ClusterParameterGroup": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::Redshift::ClusterParameterGroup",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.RedshiftClusterParameterGroup),
		GetDescriber:    nil,
	},

	"AWS::Account::AlternateContact": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::Account::AlternateContact",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   SequentialDescribeGlobal(describer.AccountAlternateContact),
		GetDescriber:    nil,
	},

	"AWS::Inspector::AssessmentTarget": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::Inspector::AssessmentTarget",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.InspectorAssessmentTarget),
		GetDescriber:    nil,
	},

	"AWS::CloudFront::ResponseHeadersPolicy": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::CloudFront::ResponseHeadersPolicy",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   SequentialDescribeGlobal(describer.CloudFrontResponseHeadersPolicy),
		GetDescriber:    nil,
	},

	"AWS::EC2::Instance": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::EC2::Instance",
		Tags: map[string][]string{
			"category": {"Compute"},
			"logo_uri": {"https://raw.githubusercontent.com/kaytu-io/awsicons/master/svg-export/icons/Ec2Instance.svg"},
		},
		Labels:        map[string]string{},
		Annotations:   map[string]string{},
		ListDescriber: ParallelDescribeRegional(describer.EC2Instance),
		GetDescriber:  ParallelDescribeRegionalSingleResource(describer.GetEC2Instance),
	},

	"AWS::EC2::InstanceMetricCpuUtilizationHourly": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::EC2::InstanceMetricCpuUtilizationHourly",
		Tags: map[string][]string{
			"category": {"Compute"},
			"logo_uri": {"https://raw.githubusercontent.com/kaytu-io/awsicons/master/svg-export/icons/Ec2Instance.svg"},
		},
		Labels:        map[string]string{},
		Annotations:   map[string]string{},
		ListDescriber: ParallelDescribeRegional(describer.Ec2InstanceMetricCpuUtilizationHourly),
		GetDescriber:  nil,
	},

	"AWS::EC2::InstanceAvailability": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::EC2::InstanceAvailability",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.EC2InstanceAvailability),
		GetDescriber:    nil,
	},

	"AWS::CostExplorer::ByRecordTypeDaily": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::CostExplorer::ByRecordTypeDaily",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   SequentialDescribeGlobal(describer.CostByRecordTypeLastDay),
		GetDescriber:    nil,
	},

	"AWS::EC2::ReservedInstances": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::EC2::ReservedInstances",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.EC2ReservedInstances),
		GetDescriber:    ParallelDescribeRegionalSingleResource(describer.GetEC2ReservedInstances),
	},

	"AWS::ECR::Repository": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::ECR::Repository",
		Tags: map[string][]string{
			"category": {"Containers"},
			"logo_uri": {"https://raw.githubusercontent.com/kaytu-io/awsicons/master/svg-export/icons/ElasticContainerRegistry.svg"},
		},
		Labels:        map[string]string{},
		Annotations:   map[string]string{},
		ListDescriber: ParallelDescribeRegional(describer.ECRRepository),
		GetDescriber:  nil,
	},

	"AWS::ECR::Registry": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::ECR::Registry",
		Tags: map[string][]string{
			"category": {"Containers"},
			"logo_uri": {"https://raw.githubusercontent.com/kaytu-io/awsicons/master/svg-export/icons/ElasticContainerRegistry.svg"},
		},
		Labels:        map[string]string{},
		Annotations:   map[string]string{},
		ListDescriber: ParallelDescribeRegional(describer.ECRRegistry),
		GetDescriber:  nil,
	},

	"AWS::ECR::RegistryScanningConfiguration": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::ECR::RegistryScanningConfiguration",
		Tags: map[string][]string{
			"category": {"Containers"},
			"logo_uri": {},
		},
		Labels:        map[string]string{},
		Annotations:   map[string]string{},
		ListDescriber: ParallelDescribeRegional(describer.ECRRegistryScanningConfiguration),
		GetDescriber:  nil,
	},

	"AWS::ElasticLoadBalancingV2::Listener": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::ElasticLoadBalancingV2::Listener",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.ElasticLoadBalancingV2Listener),
		GetDescriber:    ParallelDescribeRegionalSingleResource(describer.GetElasticLoadBalancingV2Listener),
	},

	"AWS::IAM::Group": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::IAM::Group",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   SequentialDescribeGlobal(describer.IAMGroup),
		GetDescriber:    nil,
	},

	"AWS::IAM::OpenIdConnectProvider": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::IAM::OpenIdConnectProvider",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   SequentialDescribeGlobal(describer.IAMOpenIdConnectProvider),
		GetDescriber:    ParallelDescribeRegionalSingleResource(describer.GetIAMOpenIdConnectProvider),
	},

	"AWS::Backup::Plan": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::Backup::Plan",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.BackupPlan),
		GetDescriber:    nil,
	},

	"AWS::Config::ConformancePack": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::Config::ConformancePack",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.ConfigConformancePack),
		GetDescriber:    nil,
	},

	"AWS::Config::RetentionConfiguration": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::Config::RetentionConfiguration",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.ConfigRetentionConfiguration),
		GetDescriber:    nil,
	},

	"AWS::CostExplorer::ByAccountDaily": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::CostExplorer::ByAccountDaily",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   SequentialDescribeGlobal(describer.CostByAccountLastDay),
		GetDescriber:    nil,
	},

	"AWS::Account::Contact": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::Account::Contact",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   SequentialDescribeGlobal(describer.AccountContact),
		GetDescriber:    nil,
	},

	"AWS::Glue::DataQualityRuleset": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::Glue::DataQualityRuleset",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.GlueDataQualityRuleset),
		GetDescriber:    nil,
	},

	"AWS::EventBridge::EventBus": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::EventBridge::EventBus",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.EventBridgeBus),
		GetDescriber:    nil,
	},

	"AWS::ApiGateway::Stage": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::ApiGateway::Stage",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.ApiGatewayStage),
		GetDescriber:    ParallelDescribeRegionalSingleResource(describer.GetApiGatewayStage),
	},

	"AWS::ApiGatewayV2::Stage": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::ApiGatewayV2::Stage",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.ApiGatewayV2Stage),
		GetDescriber:    nil,
	},

	"AWS::DynamoDb::LocalSecondaryIndex": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::DynamoDb::LocalSecondaryIndex",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.DynamoDbLocalSecondaryIndex),
		GetDescriber:    nil,
	},

	"AWS::ResourceGroups::Groups": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::ResourceGroups::Groups",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.ResourceGroups),
		GetDescriber:    nil,
	},

	"AWS::Timestream::Database": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::Timestream::Database",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.TimestreamDatabase),
		GetDescriber:    nil,
	},

	"AWS::OpenSearchServerless::Collection": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::OpenSearchServerless::Collection",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.OpenSearchServerlessCollection),
		GetDescriber:    nil,
	},

	"AWS::EC2::ElasticIP": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::EC2::ElasticIP",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.EC2ElasticIP),
		GetDescriber:    nil,
	},

	"AWS::EC2::LocalGateway": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::EC2::LocalGateway",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.EC2LocalGateway),
		GetDescriber:    nil,
	},

	"AWS::EC2::Image": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::EC2::Image",
		Tags: map[string][]string{
			"category": {"Compute"},
			"logo_uri": {"https://raw.githubusercontent.com/kaytu-io/awsicons/master/svg-export/icons/Ec2AmiResource.svg"},
		},
		Labels:        map[string]string{},
		Annotations:   map[string]string{},
		ListDescriber: ParallelDescribeRegional(describer.EC2AMI),
		GetDescriber:  ParallelDescribeRegionalSingleResource(describer.GetEC2AMI),
	},

	"AWS::EC2::Subnet": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::EC2::Subnet",
		Tags: map[string][]string{
			"category": {"Networking"},
		},
		Labels:        map[string]string{},
		Annotations:   map[string]string{},
		ListDescriber: ParallelDescribeRegional(describer.EC2Subnet),
		GetDescriber:  ParallelDescribeRegionalSingleResource(describer.GetEC2Subnet),
	},

	"AWS::ECS::TaskSet": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::ECS::TaskSet",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.ECSTaskSet),
		GetDescriber:    nil,
	},

	"AWS::Kinesis::Stream": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::Kinesis::Stream",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.KinesisStream),
		GetDescriber:    nil,
	},

	"AWS::Kinesis::Consumer": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::Kinesis::Consumer",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.KinesisConsumer),
		GetDescriber:    nil,
	},

	"AWS::DocDB::Cluster": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::DocDB::Cluster",
		Tags: map[string][]string{
			"category": {"Database"},
			"logo_uri": {"https://raw.githubusercontent.com/kaytu-io/awsicons/master/svg-export/icons/DocumentDb.svg"},
		},
		Labels:        map[string]string{},
		Annotations:   map[string]string{},
		ListDescriber: ParallelDescribeRegional(describer.DocDBCluster),
		GetDescriber:  ParallelDescribeRegionalSingleResource(describer.GetDocDBCluster),
	},

	"AWS::DocDB::ClusterSnapshot": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::DocDB::ClusterSnapshot",
		Tags: map[string][]string{
			"category": {"Database"},
			"logo_uri": {"https://raw.githubusercontent.com/kaytu-io/awsicons/master/svg-export/icons/DocumentDb.svg"},
		},
		Labels:        map[string]string{},
		Annotations:   map[string]string{},
		ListDescriber: ParallelDescribeRegional(describer.DocDBClusterSnapshot),
		GetDescriber:  nil,
	},

	"AWS::DocDB::ClusterInstance": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::DocDB::ClusterInstance",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.DocDBClusterInstance),
		GetDescriber:    ParallelDescribeRegionalSingleResource(describer.GetDocDBClusterInstance),
	},

	"AWS::ElastiCache::ReplicationGroup": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::ElastiCache::ReplicationGroup",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.ElastiCacheReplicationGroup),
		GetDescriber:    nil,
	},

	"AWS::GlobalAccelerator::Accelerator": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::GlobalAccelerator::Accelerator",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.GlobalAcceleratorAccelerator),
		GetDescriber:    nil,
	},

	"AWS::CloudWatch::Metric": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::CloudWatch::Metric",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.CloudWatchMetrics),
		GetDescriber:    nil,
	},

	"AWS::CostExplorer::ForcastMonthly": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::CostExplorer::ForcastMonthly",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   SequentialDescribeGlobal(describer.CostForecastMonthly),
		GetDescriber:    nil,
	},

	"AWS::EMR::InstanceGroup": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::EMR::InstanceGroup",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.EMRInstanceGroup),
		GetDescriber:    nil,
	},

	"AWS::EC2::ManagedPrefixList": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::EC2::ManagedPrefixList",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.EC2ManagedPrefixList),
		GetDescriber:    ParallelDescribeRegionalSingleResource(describer.GetEC2ManagedPrefixList),
	},

	"AWS::EC2::ManagedPrefixListEntry": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::EC2::ManagedPrefixListEntry",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.EC2ManagedPrefixListEntry),
		GetDescriber:    nil,
	},

	"AWS::EC2::ClientVpnEndpoint": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::EC2::ClientVpnEndpoint",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.EC2ClientVpnEndpoint),
		GetDescriber:    nil,
	},

	"AWS::MWAA::Environment": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::MWAA::Environment",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.MWAAEnvironment),
		GetDescriber:    nil,
	},

	"AWS::CloudWatch::LogResourcePolicy": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::CloudWatch::LogResourcePolicy",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.CloudWatchLogsResourcePolicy),
		GetDescriber:    nil,
	},

	"AWS::CodeArtifact::Domain": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::CodeArtifact::Domain",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.CodeArtifactDomain),
		GetDescriber:    nil,
	},

	"AWS::CodeStar::Project": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::CodeStar::Project",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.CodeStarProject),
		GetDescriber:    nil,
	},

	"AWS::Neptune::Database": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::Neptune::Database",
		Tags: map[string][]string{
			"category": {"Database"},
			"logo_uri": {"https://raw.githubusercontent.com/kaytu-io/awsicons/master/svg-export/icons/Neptune.svg"},
		},
		Labels:        map[string]string{},
		Annotations:   map[string]string{},
		ListDescriber: ParallelDescribeRegional(describer.NeptuneDatabase),
		GetDescriber:  nil,
	},

	"AWS::Neptune::DBCluster": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::Neptune::DBCluster",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.NeptuneDatabaseCluster),
		GetDescriber:    nil,
	},

	"AWS::Neptune::DBClusterSnapshot": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::Neptune::DBClusterSnapshot",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.NeptuneDatabaseClusterSnapshot),
		GetDescriber:    nil,
	},

	"AWS::NetworkFirewall::FirewallPolicy": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::NetworkFirewall::FirewallPolicy",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.NetworkFirewallPolicy),
		GetDescriber:    nil,
	},

	"AWS::NetworkFirewall::RuleGroup": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::NetworkFirewall::RuleGroup",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.NetworkFirewallRuleGroup),
		GetDescriber:    nil,
	},

	"AWS::Oam::Link": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::Oam::Link",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.OAMLink),
		GetDescriber:    nil,
	},

	"AWS::Oam::Sink": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::Oam::Sink",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.OAMSink),
		GetDescriber:    nil,
	},

	"AWS::Organizations::Account": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::Organizations::Account",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.OrganizationsAccount),
		GetDescriber:    nil,
	},

	"AWS::Pinpoint::App": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::Pinpoint::App",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.PinpointApp),
		GetDescriber:    nil,
	},

	"AWS::Pipes::Pipe": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::Pipes::Pipe",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.PipesPipe),
		GetDescriber:    nil,
	},

	"AWS::RDS::DBClusterParameterGroup": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::RDS::DBClusterParameterGroup",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.RDSDBClusterParameterGroup),
		GetDescriber:    nil,
	},

	"AWS::RDS::OptionGroup": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::RDS::OptionGroup",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.RDSOptionGroup),
		GetDescriber:    nil,
	},

	"AWS::RDS::DBParameterGroup": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::RDS::DBParameterGroup",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.RDSDBParameterGroup),
		GetDescriber:    nil,
	},

	"AWS::RDS::DBProxy": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::RDS::DBProxy",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.RDSDBProxy),
		GetDescriber:    nil,
	},

	"AWS::RDS::DBSubnetGroup": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::RDS::DBSubnetGroup",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.RDSDBSubnetGroup),
		GetDescriber:    nil,
	},

	"AWS::RDS::DBRecommendation": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::RDS::DBRecommendation",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.RDSDBRecommendation),
		GetDescriber:    nil,
	},

	"AWS::Redshift::EventSubscription": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::Redshift::EventSubscription",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.RedshiftEventSubscription),
		GetDescriber:    nil,
	},

	"AWS::RedshiftServerless::Workgroup": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::RedshiftServerless::Workgroup",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.RedshiftServerlessWorkgroup),
		GetDescriber:    nil,
	},

	"AWS::ResourceExplorer2::Index": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::ResourceExplorer2::Index",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.ResourceExplorerIndex),
		GetDescriber:    nil,
	},

	"AWS::Route53::HealthCheck": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::Route53::HealthCheck",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.Route53HealthCheck),
		GetDescriber:    nil,
	},

	"AWS::Route53Resolver::ResolverRule": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::Route53Resolver::ResolverRule",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.Route53ResolverResolverRule),
		GetDescriber:    nil,
	},

	"AWS::Route53Resolver::QueryLogConfig": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::Route53Resolver::QueryLogConfig",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.Route53ResolverQueryLogConfig),
		GetDescriber:    ParallelDescribeRegionalSingleResource(describer.GetRoute53ResolverQueryLogConfig),
	},

	"AWS::SageMaker::App": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::SageMaker::App",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.SageMakerApp),
		GetDescriber:    nil,
	},

	"AWS::SageMaker::Domain": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::SageMaker::Domain",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.SageMakerDomain),
		GetDescriber:    nil,
	},

	"AWS::StepFunctions::StateMachine": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::StepFunctions::StateMachine",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.StepFunctionsStateMachine),
		GetDescriber:    nil,
	},

	"AWS::StepFunctions::StateMachineExecution": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::StepFunctions::StateMachineExecution",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.StepFunctionsStateMachineExecution),
		GetDescriber:    nil,
	},

	"AWS::StepFunctions::StateMachineExecutionHistories": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::StepFunctions::StateMachineExecutionHistories",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.StepFunctionsStateMachineExecutionHistories),
		GetDescriber:    nil,
	},

	"AWS::SimSpaceWeaver::Simulation": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::SimSpaceWeaver::Simulation",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.SimSpaceWeaverSimulation),
		GetDescriber:    nil,
	},

	"AWS::SSM::Association": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::SSM::Association",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.SSMAssociation),
		GetDescriber:    nil,
	},

	"AWS::SSM::Document": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::SSM::Document",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.SSMDocument),
		GetDescriber:    nil,
	},

	"AWS::SSM::DocumentPermission": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::SSM::DocumentPermission",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.SSMDocumentPermission),
		GetDescriber:    nil,
	},

	"AWS::EC2::CustomerGateway": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::EC2::CustomerGateway",
		Tags: map[string][]string{
			"category": {"Networking"},
			"logo_uri": {"https://raw.githubusercontent.com/kaytu-io/awsicons/master/svg-export/icons/VpcCustomerGateway.svg"},
		},
		Labels:        map[string]string{},
		Annotations:   map[string]string{},
		ListDescriber: ParallelDescribeRegional(describer.EC2CustomerGateway),
		GetDescriber:  ParallelDescribeRegionalSingleResource(describer.GetEC2CustomerGateway),
	},

	"AWS::EC2::VerifiedAccessInstance": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::EC2::VerifiedAccessInstance",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.EC2VerifiedAccessInstance),
		GetDescriber:    ParallelDescribeRegionalSingleResource(describer.GetEC2VerifiedAccessInstance),
	},

	"AWS::EC2::VerifiedAccessEndpoint": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::EC2::VerifiedAccessEndpoint",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.EC2VerifiedAccessEndpoint),
		GetDescriber:    ParallelDescribeRegionalSingleResource(describer.GetEC2VerifiedAccessEndpoint),
	},

	"AWS::EC2::VerifiedAccessGroup": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::EC2::VerifiedAccessGroup",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.EC2VerifiedAccessGroup),
		GetDescriber:    ParallelDescribeRegionalSingleResource(describer.GetEC2VerifiedAccessGroup),
	},

	"AWS::EC2::VerifiedAccessTrustProvider": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::EC2::VerifiedAccessTrustProvider",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.EC2VerifiedAccessTrustProvider),
		GetDescriber:    ParallelDescribeRegionalSingleResource(describer.GetEC2VerifiedAccessTrustProvider),
	},

	"AWS::EC2::VPNGateway": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::EC2::VPNGateway",
		Tags: map[string][]string{
			"category": {"Networking"},
			"logo_uri": {"https://raw.githubusercontent.com/kaytu-io/awsicons/master/svg-export/icons/VpcVpnGateway.svg"},
		},
		Labels:        map[string]string{},
		Annotations:   map[string]string{},
		ListDescriber: ParallelDescribeRegional(describer.EC2VPNGateway),
		GetDescriber:  ParallelDescribeRegionalSingleResource(describer.GetEC2VPNGateway),
	},

	"AWS::WAFv2::IPSet": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::WAFv2::IPSet",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.WAFv2IPSet),
		GetDescriber:    nil,
	},

	"AWS::WAFv2::RegexPatternSet": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::WAFv2::RegexPatternSet",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.WAFv2RegexPatternSet),
		GetDescriber:    nil,
	},

	"AWS::WAFv2::RuleGroup": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::WAFv2::RuleGroup",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.WAFv2RuleGroup),
		GetDescriber:    nil,
	},

	"AWS::EC2::TransitGatewayRoute": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::EC2::TransitGatewayRoute",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.EC2TransitGatewayRoute),
		GetDescriber:    ParallelDescribeRegionalSingleResource(describer.GetEC2TransitGatewayRoute),
	},

	"AWS::GuardDuty::Filter": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::GuardDuty::Filter",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.GuardDutyFilter),
		GetDescriber:    nil,
	},

	"AWS::ECS::TaskDefinition": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::ECS::TaskDefinition",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.ECSTaskDefinition),
		GetDescriber:    ParallelDescribeRegionalSingleResource(describer.GetECSTaskDefinition),
	},

	"AWS::GuardDuty::ThreatIntelSet": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::GuardDuty::ThreatIntelSet",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.GuardDutyThreatIntelSet),
		GetDescriber:    nil,
	},

	"AWS::ApiGatewayV2::DomainName": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::ApiGatewayV2::DomainName",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.ApiGatewayV2DomainName),
		GetDescriber:    ParallelDescribeRegionalSingleResource(describer.GetApiGatewayV2DomainName),
	},

	"AWS::ApiGateway::DomainName": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::ApiGateway::DomainName",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.ApiGatewayDomainName),
		GetDescriber:    ParallelDescribeRegionalSingleResource(describer.GetApiGatewayDomainName),
	},

	"AWS::ApiGatewayV2::Route": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::ApiGatewayV2::Route",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.ApiGatewayV2Route),
		GetDescriber:    nil,
	},

	"AWS::MQ::Broker": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::MQ::Broker",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.MQBroker),
		GetDescriber:    nil,
	},

	"AWS::ACMPCA::CertificateAuthority": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::ACMPCA::CertificateAuthority",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.ACMPCACertificateAuthority),
		GetDescriber:    nil,
	},

	"AWS::CloudFormation::Stack": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::CloudFormation::Stack",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.CloudFormationStack),
		GetDescriber:    ParallelDescribeRegionalSingleResource(describer.GetCloudFormationStack),
	},

	"AWS::CloudFormation::StackResource": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::CloudFormation::StackResource",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.CloudFormationStackResource),
		GetDescriber:    ParallelDescribeRegionalSingleResource(describer.GetCloudFormationStackResource),
	},

	"AWS::DirectConnect::Connection": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::DirectConnect::Connection",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.DirectConnectConnection),
		GetDescriber:    nil,
	},

	"AWS::FSX::FileSystem": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::FSX::FileSystem",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.FSXFileSystem),
		GetDescriber:    nil,
	},

	"AWS::Glue::SecurityConfiguration": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::Glue::SecurityConfiguration",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.GlueSecurityConfiguration),
		GetDescriber:    nil,
	},

	"AWS::Inspector::AssessmentRun": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::Inspector::AssessmentRun",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.InspectorAssessmentRun),
		GetDescriber:    nil,
	},

	"AWS::Inspector2::Coverage": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::Inspector2::Coverage",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.Inspector2Coverage),
		GetDescriber:    nil,
	},

	"AWS::Inspector2::CoverageStatistics": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::Inspector2::CoverageStatistics",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.Inspector2CoverageStatistic),
		GetDescriber:    nil,
	},

	"AWS::Inspector2::Member": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::Inspector2::Member",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.Inspector2CoverageMember),
		GetDescriber:    nil,
	},

	"AWS::Inspector2::Finding": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::Inspector2::Finding",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.Inspector2Finding),
		GetDescriber:    nil,
	},

	"AWS::Config::ConfigurationRecorder": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::Config::ConfigurationRecorder",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.ConfigConfigurationRecorder),
		GetDescriber:    nil,
	},

	"AWS::EC2::NatGateway": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::EC2::NatGateway",
		Tags: map[string][]string{
			"category": {"Networking"},
			"logo_uri": {"https://raw.githubusercontent.com/kaytu-io/awsicons/master/svg-export/icons/VpcNatGateway.svg"},
		},
		Labels:        map[string]string{},
		Annotations:   map[string]string{},
		ListDescriber: ParallelDescribeRegional(describer.EC2NatGateway),
		GetDescriber:  ParallelDescribeRegionalSingleResource(describer.GetEC2NatGateway),
	},

	"AWS::ECR::PublicRepository": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::ECR::PublicRepository",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.ECRPublicRepository),
		GetDescriber:    nil,
	},

	"AWS::ECS::Cluster": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::ECS::Cluster",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.ECSCluster),
		GetDescriber:    ParallelDescribeRegionalSingleResource(describer.GetECSCluster),
	},

	"AWS::ElasticLoadBalancingV2::TargetGroup": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::ElasticLoadBalancingV2::TargetGroup",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.ElasticLoadBalancingV2TargetGroup),
		GetDescriber:    nil,
	},

	"AWS::CloudFront::CachePolicy": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::CloudFront::CachePolicy",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   SequentialDescribeGlobal(describer.CloudFrontCachePolicy),
		GetDescriber:    nil,
	},

	"AWS::CodeArtifact::Repository": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::CodeArtifact::Repository",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.CodeArtifactRepository),
		GetDescriber:    nil,
	},

	"AWS::AMP::Workspace": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::AMP::Workspace",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.AMPWorkspace),
		GetDescriber:    nil,
	},

	"AWS::EC2::CapacityReservation": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::EC2::CapacityReservation",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.EC2CapacityReservation),
		GetDescriber:    ParallelDescribeRegionalSingleResource(describer.GetEC2CapacityReservation),
	},

	"AWS::SageMaker::NotebookInstance": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::SageMaker::NotebookInstance",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.SageMakerNotebookInstance),
		GetDescriber:    nil,
	},

	"AWS::IAM::AccessAdvisor": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::IAM::AccessAdvisor",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.IAMAccessAdvisor),
		GetDescriber:    nil,
	},

	"AWS::EC2::VolumeSnapshot": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::EC2::VolumeSnapshot",
		Tags: map[string][]string{
			"category": {"Storage"},
			"logo_uri": {"https://raw.githubusercontent.com/kaytu-io/awsicons/master/svg-export/icons/ElasticBlockStoreSnapshot.svg"},
		},
		Labels:        map[string]string{},
		Annotations:   map[string]string{},
		ListDescriber: ParallelDescribeRegional(describer.EC2VolumeSnapshot),
		GetDescriber:  ParallelDescribeRegionalSingleResource(describer.GetEC2VolumeSnapshot),
	},

	"AWS::EC2::Region": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::EC2::Region",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.EC2Region),
		GetDescriber:    nil,
	},

	"AWS::Keyspaces::Table": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::Keyspaces::Table",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.KeyspacesTable),
		GetDescriber:    nil,
	},

	"AWS::Config::AggregationAuthorization": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::Config::AggregationAuthorization",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.ConfigAggregateAuthorization),
		GetDescriber:    nil,
	},

	"AWS::DAX::SubnetGroup": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::DAX::SubnetGroup",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.DAXSubnetGroup),
		GetDescriber:    nil,
	},

	"AWS::DynamoDb::GlobalTable": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::DynamoDb::GlobalTable",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.DynamoDbGlobalTable),
		GetDescriber:    nil,
	},

	"AWS::ElasticLoadBalancing::LoadBalancer": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::ElasticLoadBalancing::LoadBalancer",
		Tags: map[string][]string{
			"category": {"Networking"},
			"logo_uri": {"https://raw.githubusercontent.com/kaytu-io/awsicons/master/svg-export/icons/ElasticLoadBalancingClassicLoadBalancer.svg"},
		},
		Labels:        map[string]string{},
		Annotations:   map[string]string{},
		ListDescriber: ParallelDescribeRegional(describer.ElasticLoadBalancingLoadBalancer),
		GetDescriber:  nil,
	},

	"AWS::AppStream::Application": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::AppStream::Application",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.AppStreamApplication),
		GetDescriber:    nil,
	},

	"AWS::RedshiftServerless::Namespace": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::RedshiftServerless::Namespace",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.RedshiftServerlessNamespace),
		GetDescriber:    nil,
	},

	"AWS::CloudFront::OriginAccessIdentity": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::CloudFront::OriginAccessIdentity",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   SequentialDescribeGlobal(describer.CloudFrontOriginAccessIdentity),
		GetDescriber:    nil,
	},

	"AWS::EC2::Host": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::EC2::Host",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.EC2Host),
		GetDescriber:    ParallelDescribeRegionalSingleResource(describer.GetEC2Host),
	},

	"AWS::EC2::VPC": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::EC2::VPC",
		Tags: map[string][]string{
			"category": {"Networking"},
			"logo_uri": {"https://raw.githubusercontent.com/kaytu-io/awsicons/master/svg-export/icons/Vpc.svg"},
		},
		Labels:        map[string]string{},
		Annotations:   map[string]string{},
		ListDescriber: ParallelDescribeRegional(describer.EC2VPC),
		GetDescriber:  ParallelDescribeRegionalSingleResource(describer.GetEC2VPC),
	},

	"AWS::EC2::TransitGatewayRouteTable": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::EC2::TransitGatewayRouteTable",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.EC2TransitGatewayRouteTable),
		GetDescriber:    ParallelDescribeRegionalSingleResource(describer.GetEC2TransitGatewayRouteTable),
	},

	"AWS::EKS::Nodegroup": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::EKS::Nodegroup",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.EKSNodegroup),
		GetDescriber:    nil,
	},

	"AWS::Backup::Selection": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::Backup::Selection",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.BackupSelection),
		GetDescriber:    nil,
	},

	"AWS::CloudTrail::Import": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::CloudTrail::Import",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.CloudTrailImport),
		GetDescriber:    nil,
	},

	"AWS::CostExplorer::ByServiceDaily": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::CostExplorer::ByServiceDaily",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   SequentialDescribeGlobal(describer.CostByServiceLastDay),
		GetDescriber:    nil,
	},

	"AWS::ElasticLoadBalancingV2::SslPolicy": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::ElasticLoadBalancingV2::SslPolicy",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.ElasticLoadBalancingV2SslPolicy),
		GetDescriber:    nil,
	},

	"AWS::GuardDuty::Finding": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::GuardDuty::Finding",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.GuardDutyFinding),
		GetDescriber:    nil,
	},

	"AWS::EC2::DHCPOptions": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::EC2::DHCPOptions",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.EC2DHCPOptions),
		GetDescriber:    ParallelDescribeRegionalSingleResource(describer.GetEC2DHCPOptions),
	},

	"AWS::EC2::InstanceType": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::EC2::InstanceType",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.EC2InstanceType),
		GetDescriber:    nil,
	},

	"AWS::Batch::ComputeEnvironment": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::Batch::ComputeEnvironment",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.BatchComputeEnvironment),
		GetDescriber:    ParallelDescribeRegionalSingleResource(describer.GetBatchComputeEnvironment),
	},

	"AWS::DMS::ReplicationInstance": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::DMS::ReplicationInstance",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.DMSReplicationInstance),
		GetDescriber:    nil,
	},

	"AWS::DMS::Endpoint": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::DMS::Endpoint",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.DMSEndpoint),
		GetDescriber:    nil,
	},

	"AWS::DMS::ReplicationTask": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::DMS::ReplicationTask",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.DMSReplicationTask),
		GetDescriber:    nil,
	},

	"AWS::DynamoDb::Table": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::DynamoDb::Table",
		Tags: map[string][]string{
			"category": {"Database"},
			"logo_uri": {"https://raw.githubusercontent.com/kaytu-io/awsicons/master/svg-export/icons/DynamoDbTable.svg"},
		},
		Labels:        map[string]string{},
		Annotations:   map[string]string{},
		ListDescriber: ParallelDescribeRegional(describer.DynamoDbTable),
		GetDescriber:  nil,
	},

	"AWS::Shield::ProtectionGroup": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::Shield::ProtectionGroup",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   SequentialDescribeGlobal(describer.ShieldProtectionGroup),
		GetDescriber:    nil,
	},

	"AWS::Firehose::DeliveryStream": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::Firehose::DeliveryStream",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.FirehoseDeliveryStream),
		GetDescriber:    nil,
	},

	"AWS::KinesisVideo::Stream": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::KinesisVideo::Stream",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.KinesisVideoStream),
		GetDescriber:    nil,
	},

	"AWS::KMS::Alias": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::KMS::Alias",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.KMSAlias),
		GetDescriber:    nil,
	},

	"AWS::Lambda::Alias": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::Lambda::Alias",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.LambdaAlias),
		GetDescriber:    nil,
	},

	"AWS::Lambda::LambdaLayer": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::Lambda::LambdaLayer",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.LambdaLayer),
		GetDescriber:    nil,
	},

	"AWS::Lambda::LayerVersion": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::Lambda::LayerVersion",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.LambdaLayerVersion),
		GetDescriber:    nil,
	},

	"AWS::Lightsail::Instance": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::Lightsail::Instance",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.LightsailInstance),
		GetDescriber:    nil,
	},

	"AWS::Macie2::ClassificationJob": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::Macie2::ClassificationJob",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.Macie2ClassificationJob),
		GetDescriber:    nil,
	},

	"AWS::MediaStore::Container": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::MediaStore::Container",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.MediaStoreContainer),
		GetDescriber:    nil,
	},

	"AWS::Mgn::Application": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::Mgn::Application",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.MGNApplication),
		GetDescriber:    nil,
	},

	"AWS::ResourceExplorer2::SupportedResourceType": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::ResourceExplorer2::SupportedResourceType",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.ResourceExplorer2SupportedResourceType),
		GetDescriber:    nil,
	},

	"AWS::Route53Resolver::ResolverEndpoint": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::Route53Resolver::ResolverEndpoint",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.Route53ResolverResolverEndpoint),
		GetDescriber:    nil,
	},

	"AWS::Route53Domains::Domain": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::Route53Domains::Domain",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   SequentialDescribeGlobal(describer.Route53Domain),
		GetDescriber:    nil,
	},

	"AWS::Route53::Record": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::Route53::Record",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   SequentialDescribeGlobal(describer.Route53Record),
		GetDescriber:    nil,
	},

	"AWS::Route53::TrafficPolicy": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::Route53::TrafficPolicy",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   SequentialDescribeGlobal(describer.Route53TrafficPolicy),
		GetDescriber:    nil,
	},

	"AWS::Route53::TrafficPolicyInstance": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::Route53::TrafficPolicyInstance",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   SequentialDescribeGlobal(describer.Route53TrafficPolicyInstance),
		GetDescriber:    nil,
	},

	"AWS::SageMaker::Model": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::SageMaker::Model",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.SageMakerModel),
		GetDescriber:    nil,
	},

	"AWS::SageMaker::TrainingJob": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::SageMaker::TrainingJob",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.SageMakerTrainingJob),
		GetDescriber:    nil,
	},

	"AWS::SecurityHub::ActionTarget": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::SecurityHub::ActionTarget",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.SecurityHubActionTarget),
		GetDescriber:    nil,
	},

	"AWS::SecurityHub::Finding": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::SecurityHub::Finding",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.SecurityHubFinding),
		GetDescriber:    nil,
	},

	"AWS::SecurityHub::FindingAggregator": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::SecurityHub::FindingAggregator",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.SecurityHubFindingAggregator),
		GetDescriber:    nil,
	},

	"AWS::SecurityHub::Insight": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::SecurityHub::Insight",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.SecurityHubInsight),
		GetDescriber:    nil,
	},

	"AWS::SecurityHub::Member": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::SecurityHub::Member",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.SecurityHubMember),
		GetDescriber:    nil,
	},

	"AWS::SecurityHub::Product": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::SecurityHub::Product",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.SecurityHubProduct),
		GetDescriber:    nil,
	},

	"AWS::SecurityHub::StandardsControl": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::SecurityHub::StandardsControl",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.SecurityHubStandardsControl),
		GetDescriber:    nil,
	},

	"AWS::SecurityHub::StandardsSubscription": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::SecurityHub::StandardsSubscription",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.SecurityHubStandardsSubscription),
		GetDescriber:    nil,
	},

	"AWS::SecurityLake::DataLake": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::SecurityLake::DataLake",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.SecurityLakeDataLake),
		GetDescriber:    nil,
	},

	"AWS::SecurityLake::Subscriber": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::SecurityLake::Subscriber",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.SecurityLakeSubscriber),
		GetDescriber:    nil,
	},

	"AWS::Ram::PrincipalAssociation": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::Ram::PrincipalAssociation",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.RamPrincipalAssociation),
		GetDescriber:    nil,
	},

	"AWS::Ram::ResourceAssociation": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::Ram::ResourceAssociation",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.RamResourceAssociation),
		GetDescriber:    nil,
	},

	"AWS::RDS::ReservedDBInstance": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::RDS::ReservedDBInstance",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.RDSReservedDBInstance),
		GetDescriber:    nil,
	},

	"AWS::Redshift::SubnetGroup": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::Redshift::SubnetGroup",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.RedshiftSubnetGroup),
		GetDescriber:    nil,
	},

	"AWS::SeverlessApplicationRepository::Application": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::SeverlessApplicationRepository::Application",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.ServerlessApplicationRepositoryApplication),
		GetDescriber:    nil,
	},

	"AWS::AuditManager::Framework": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::AuditManager::Framework",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.AuditManagerFramework),
		GetDescriber:    nil,
	},

	"AWS::AuditManager::EvidenceFolder": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::AuditManager::EvidenceFolder",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.AuditManagerEvidenceFolder),
		GetDescriber:    nil,
	},

	"AWS::AuditManager::Evidence": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::AuditManager::Evidence",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.AuditManagerEvidence),
		GetDescriber:    nil,
	},

	"AWS::AuditManager::Control": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::AuditManager::Control",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.AuditManagerControl),
		GetDescriber:    nil,
	},

	"AWS::AuditManager::Assessment": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::AuditManager::Assessment",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.AuditManagerAssessment),
		GetDescriber:    nil,
	},

	"AWS::CloudTrail::TrailEvent": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::CloudTrail::TrailEvent",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.CloudTrailTrailEvent),
		GetDescriber:    nil,
	},

	"AWS::CloudWatch::LogEvent": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::CloudWatch::LogEvent",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.CloudWatchLogEvent),
		GetDescriber:    nil,
	},

	"AWS::Logs::MetricFilter": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::Logs::MetricFilter",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.CloudWatchLogsMetricFilter),
		GetDescriber:    nil,
	},

	"AWS::CloudWatch::LogStream": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::CloudWatch::LogStream",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.CloudWatchLogStream),
		GetDescriber:    nil,
	},

	"AWS::CostExplorer::ByUsageTypeMonthly": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::CostExplorer::ByUsageTypeMonthly",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   SequentialDescribeGlobal(describer.CostByServiceUsageLastMonth),
		GetDescriber:    nil,
	},

	"AWS::ServiceQuotas::ServiceQuotaChangeRequest": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::ServiceQuotas::ServiceQuotaChangeRequest",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.ServiceQuotasServiceQuotaChangeRequest),
		GetDescriber:    nil,
	},

	"AWS::ServiceQuotas::Service": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::ServiceQuotas::Service",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.ServiceQuotasService),
		GetDescriber:    nil,
	},

	"AWS::EC2::VPCEndpointService": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::EC2::VPCEndpointService",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.EC2VPCEndpointService),
		GetDescriber:    nil,
	},

	"AWS::EC2::LaunchTemplate": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::EC2::LaunchTemplate",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.EC2LaunchTemplate),
		GetDescriber:    ParallelDescribeRegionalSingleResource(describer.GetEC2LaunchTemplate),
	},

	"AWS::EC2::LaunchTemplateVersion": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::EC2::LaunchTemplateVersion",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.EC2LaunchTemplateVersion),
		GetDescriber:    ParallelDescribeRegionalSingleResource(describer.GetEC2LaunchTemplateVersion),
	},

	"AWS::SNS::Subscription": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::SNS::Subscription",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.SNSSubscription),
		GetDescriber:    nil,
	},

	"AWS::S3::AccountSetting": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::S3::AccountSetting",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   SequentialDescribeGlobal(describer.S3AccountSetting),
		GetDescriber:    nil,
	},

	"AWS::SSM::ManagedInstanceCompliance": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::SSM::ManagedInstanceCompliance",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.SSMManagedInstanceCompliance),
		GetDescriber:    nil,
	},

	"AWS::SSM::ManagedInstancePatchState": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::SSM::ManagedInstancePatchState",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.SSMManagedInstancePatchState),
		GetDescriber:    nil,
	},

	"AWS::SSOAdmin::AccountAssignment": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::SSOAdmin::AccountAssignment",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.SSOAdminAccountAssignment),
		GetDescriber:    nil,
	},

	"AWS::SSOAdmin::UserEffectiveAccess": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::SSOAdmin::UserEffectiveAccess",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.UserEffectiveAccess),
		GetDescriber:    nil,
	},

	"AWS::SSOAdmin::Instance": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::SSOAdmin::Instance",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.SSOAdminInstance),
		GetDescriber:    nil,
	},

	"AWS::SSOAdmin::PermissionSet": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::SSOAdmin::PermissionSet",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.SSOAdminPermissionSet),
		GetDescriber:    nil,
	},

	"AWS::SSOAdmin::AttachedManagedPolicy": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::SSOAdmin::AttachedManagedPolicy",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.SSOAdminManagedPolicyAttachment),
		GetDescriber:    nil,
	},

	"AWS::ServiceDiscovery::Service": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::ServiceDiscovery::Service",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.ServiceDiscoveryService),
		GetDescriber:    nil,
	},

	"AWS::ServiceDiscovery::Namespace": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::ServiceDiscovery::Namespace",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.ServiceDiscoveryNamespace),
		GetDescriber:    nil,
	},

	"AWS::ServiceDiscovery::Instance": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::ServiceDiscovery::Instance",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.ServiceDiscoveryInstance),
		GetDescriber:    nil,
	},

	"AWS::ServiceCatalog::Portfolio": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::ServiceCatalog::Portfolio",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.ServiceCatalogPortfolio),
		GetDescriber:    nil,
	},

	"AWS::ServiceCatalog::Product": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::ServiceCatalog::Product",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.ServiceCatalogProduct),
		GetDescriber:    nil,
	},

	"AWS::IdentityStore::User": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::IdentityStore::User",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.IdentityStoreUser),
		GetDescriber:    nil,
	},

	"AWS::IdentityStore::Group": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::IdentityStore::Group",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.IdentityStoreGroup),
		GetDescriber:    nil,
	},

	"AWS::IdentityStore::GroupMembership": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "AWS::IdentityStore::GroupMembership",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   ParallelDescribeRegional(describer.IdentityStoreGroupMembership),
		GetDescriber:    nil,
	},
}
