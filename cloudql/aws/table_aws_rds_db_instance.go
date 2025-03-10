package aws

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-aws/discovery/pkg/es"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/rds"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"

	"github.com/turbot/go-kit/helpers"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsRDSDBInstance(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_rds_db_instance",
		Description: "AWS RDS DB Instance",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("db_instance_identifier"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"DBInstanceNotFound"}),
			},
			Hydrate: opengovernance.GetRDSDBInstance,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListRDSDBInstance,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "db_cluster_identifier", Require: plugin.Optional},
				{Name: "resource_id", Require: plugin.Optional},
				{Name: "engine", Require: plugin.Optional},
			},
		},

		Columns: awsOgRegionalColumns([]*plugin.Column{
			{
				Name:        "db_instance_identifier",
				Description: "The friendly name to identify the DB Instance.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBInstance.DBInstanceIdentifier")},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) for the DB Instance.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBInstance.DBInstanceArn"),
			},
			{
				Name:        "db_cluster_identifier",
				Description: "The friendly name to identify the DB cluster, that the DB instance is a member of.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBInstance.DBClusterIdentifier")},
			{
				Name:        "status",
				Description: "Specifies the current state of this database.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBInstance.DBInstanceStatus")},
			{
				Name:        "class",
				Description: "Contains the name of the compute and memory capacity class of the DB instance.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBInstance.DBInstanceClass")},
			{
				Name:        "resource_id",
				Description: "The AWS Region-unique, immutable identifier for the DB instance.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBInstance.DbiResourceId")},
			{
				Name:        "allocated_storage",
				Description: "Specifies the allocated storage size specified in gibibytes(GiB).",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.DBInstance.AllocatedStorage")},
			{
				Name:        "auto_minor_version_upgrade",
				Description: "Specifies whether minor version patches are applied automatically, or not.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.DBInstance.AutoMinorVersionUpgrade")},
			{
				Name:        "availability_zone",
				Description: "Specifies the name of the Availability Zone the DB instance is located in.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBInstance.AvailabilityZone")},
			{
				Name:        "backup_retention_period",
				Description: "Specifies the number of days for which automatic DB snapshots are retained.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.DBInstance.BackupRetentionPeriod")},
			{
				Name:        "ca_certificate_identifier",
				Description: "The identifier of the CA certificate for this DB instance.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBInstance.CACertificateIdentifier")},
			{
				Name:        "character_set_name",
				Description: "Specifies the name of the character set that this instance is associated with.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBInstance.CharacterSetName")},
			{
				Name:        "copy_tags_to_snapshot",
				Description: "Specifies whether tags are copied from the DB instance to snapshots of the DB instance, or not.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.DBInstance.CopyTagsToSnapshot")},
			{
				Name:        "customer_owned_ip_enabled",
				Description: "Specifies whether a customer-owned IP address (CoIP) is enabled for an RDS on Outposts DB instance, or not.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.DBInstance.CustomerOwnedIpEnabled")},
			{
				Name:        "port",
				Description: "Specifies the port that the DB instance listens on.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.DBInstance.DbInstancePort")},
			{
				Name:        "db_name",
				Description: "Contains the name of the initial database of this instance that was provided at create time.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBInstance.DBName")},
			{
				Name:        "db_subnet_group_arn",
				Description: "The Amazon Resource Name (ARN) for the DB subnet group.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBInstance.DBSubnetGroup.DBSubnetGroupArn")},
			{
				Name:        "db_subnet_group_description",
				Description: "Provides the description of the DB subnet group.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBInstance.DBSubnetGroup.DBSubnetGroupDescription")},
			{
				Name:        "db_subnet_group_name",
				Description: "The name of the DB subnet group.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBInstance.DBSubnetGroup.DBSubnetGroupName")},
			{
				Name:        "db_subnet_group_status",
				Description: "Provides the status of the DB subnet group.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBInstance.DBSubnetGroup.SubnetGroupStatus")},
			{
				Name:        "deletion_protection",
				Description: "Specifies whether the DB instance has deletion protection enabled, or not.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.DBInstance.DeletionProtection")},
			{
				Name:        "endpoint_address",
				Description: "Specifies the DNS address of the DB instance.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBInstance.Endpoint.Address")},
			{
				Name:        "endpoint_hosted_zone_id",
				Description: "Specifies the ID that Amazon Route 53 assigns when you create a hosted zone.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBInstance.Endpoint.HostedZoneId")},
			{
				Name:        "endpoint_port",
				Description: "Specifies the port that the database engine is listening on.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.DBInstance.Endpoint.Port")},
			{
				Name:        "engine",
				Description: "The name of the database engine to be used for this DB instance.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBInstance.Engine")},
			{
				Name:        "engine_version",
				Description: "Indicates the database engine version.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBInstance.EngineVersion")},
			{
				Name:        "enhanced_monitoring_resource_arn",
				Description: "The ARN of the Amazon CloudWatch Logs log stream that receives the Enhanced Monitoring metrics data for the DB instance.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBInstance.EnhancedMonitoringResourceArn")},
			{
				Name:        "iam_database_authentication_enabled",
				Description: "Specifies whether the the mapping of AWS IAM accounts to database accounts is enabled, or not.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.DBInstance.IAMDatabaseAuthenticationEnabled")},
			{
				Name:        "create_time",
				Description: "Provides the date and time the DB instance was created.",
				Type:        proto.ColumnType_TIMESTAMP,

				Transform: transform.FromField("Description.DBInstance.InstanceCreateTime")},
			{
				Name:        "iops",
				Description: "Specifies the Provisioned IOPS (I/O operations per second) value.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.DBInstance.Iops")},
			{
				Name:        "kms_key_id",
				Description: "The AWS KMS key identifier for the encrypted DB instance.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBInstance.KmsKeyId")},
			{
				Name:        "latest_restorable_time",
				Description: "Specifies the latest time to which a database can be restored with point-in-time restore.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.DBInstance.LatestRestorableTime")},
			{
				Name:        "license_model",
				Description: "License model information for this DB instance.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBInstance.LicenseModel")},
			{
				Name:        "master_user_name",
				Description: "Contains the master username for the DB instance.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBInstance.MasterUsername")},
			{
				Name:        "max_allocated_storage",
				Description: "The upper limit to which Amazon RDS can automatically scale the storage of the DB instance.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.DBInstance.MaxAllocatedStorage")},
			{
				Name:        "monitoring_interval",
				Description: "The interval, in seconds, between points when Enhanced Monitoring metrics are collected for the DB instance.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.DBInstance.MonitoringInterval")},
			{
				Name:        "monitoring_role_arn",
				Description: "The ARN for the IAM role that permits RDS to send Enhanced Monitoring metrics to Amazon CloudWatch Logs.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBInstance.MonitoringRoleArn")},
			{
				Name:        "multi_az",
				Description: "Specifies if the DB instance is a Multi-AZ deployment.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.DBInstance.MultiAZ")},
			{
				Name:        "nchar_character_set_name",
				Description: "The name of the NCHAR character set for the Oracle DB instance.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBInstance.NcharCharacterSetName")},
			{
				Name:        "performance_insights_enabled",
				Description: "Specifies whether Performance Insights is enabled for the DB instance, or not.",
				Type:        proto.ColumnType_BOOL,
				Default:     false,
				Transform:   transform.FromField("Description.DBInstance.PerformanceInsightsEnabled")},
			{
				Name:        "performance_insights_kms_key_id",
				Description: "The AWS KMS key identifier for encryption of Performance Insights data.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBInstance.PerformanceInsightsKMSKeyId")},
			{
				Name:        "performance_insights_retention_period",
				Description: "The amount of time, in days, to retain Performance Insights data.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.DBInstance.PerformanceInsightsRetentionPeriod")},
			{
				Name:        "preferred_backup_window",
				Description: "Specifies the daily time range during which automated backups are created.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBInstance.PreferredBackupWindow")},
			{
				Name:        "preferred_maintenance_window",
				Description: "Specifies the weekly time range during which system maintenance can occur.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBInstance.PreferredMaintenanceWindow")},
			{
				Name:        "promotion_tier",
				Description: "Specifies the order in which an Aurora Replica is promoted to the primary instance after a failure of the existing primary instance.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.DBInstance.PromotionTier")},
			{
				Name:        "publicly_accessible",
				Description: "Specifies the accessibility options for the DB instance.",
				Type:        proto.ColumnType_BOOL,
				Default:     false,
				Transform:   transform.FromField("Description.DBInstance.PubliclyAccessible")},
			{
				Name:        "read_replica_source_db_instance_identifier",
				Description: "Contains the identifier of the source DB instance if this DB instance is a read replica.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBInstance.ReadReplicaSourceDBInstanceIdentifier")},
			{
				Name:        "replica_mode",
				Description: "The mode of an Oracle read replica.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBInstance.ReplicaMode")},
			{
				Name:        "secondary_availability_zone",
				Description: "Specifies the name of the secondary Availability Zone for a DB instance with multi-AZ support.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBInstance.SecondaryAvailabilityZone")},
			{
				Name:        "storage_encrypted",
				Description: "Specifies whether the DB instance is encrypted, or not.",
				Type:        proto.ColumnType_BOOL,
				Default:     false,
				Transform:   transform.FromField("Description.DBInstance.StorageEncrypted")},
			{
				Name:        "storage_throughput",
				Description: "Specifies the storage throughput for the DB instance. This setting applies only to the gp3 storage type.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.DBInstance.StorageThroughput")},
			{
				Name:        "storage_type",
				Description: "Specifies the storage type associated with DB instance.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBInstance.StorageType")},
			{
				Name:        "tde_credential_arn",
				Description: " The ARN from the key store with which the instance is associated for TDE encryption.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBInstance.TdeCredentialArn")},
			{
				Name:        "timezone",
				Description: "The time zone of the DB instance.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBInstance.Timezone")},
			{
				Name:        "vpc_id",
				Description: "Provides the VpcId of the DB subnet group.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBInstance.DBSubnetGroup.VpcId")},
			{
				Name:        "associated_roles",
				Description: "A list of AWS IAM roles that are associated with the DB instance.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DBInstance.AssociatedRoles")},
			{
				Name:        "certificate",
				Description: "The CA certificate associated with the DB instance.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Certificate")},
			{
				Name:        "db_parameter_groups",
				Description: "A list of DB parameter groups applied to this DB instance.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DBInstance.DBParameterGroups")},
			{
				Name:        "db_security_groups",
				Description: "A list of DB security group associated with the DB instance.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DBInstance.DBSecurityGroups")},
			{
				Name:        "domain_memberships",
				Description: "A list of Active Directory Domain membership records associated with the DB instance.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DBInstance.DomainMemberships")},
			{
				Name:        "enabled_cloudwatch_logs_exports",
				Description: "A list of log types that this DB instance is configured to export to CloudWatch Logs.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DBInstance.EnabledCloudwatchLogsExports")},
			{
				Name:        "option_group_memberships",
				Description: "A list of option group memberships for this DB instance",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DBInstance.OptionGroupMemberships")},
			{
				Name:        "pending_maintenance_actions",
				Description: "A list that provides details about the pending maintenance actions for the resource.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.PendingMaintenance")},
			{
				Name:        "processor_features",
				Description: "The number of CPU cores and the number of threads per core for the DB instance class of the DB instance.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DBInstance.ProcessorFeatures")},
			{
				Name:        "read_replica_db_cluster_identifiers",
				Description: "A list of identifiers of Aurora DB clusters to which the RDS DB instance is replicated as a read replica.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DBInstance.ReadReplicaDBClusterIdentifiers")},
			{
				Name:        "read_replica_db_instance_identifiers",
				Description: "A list of identifiers of the read replicas associated with this DB instance.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DBInstance.ReadReplicaDBInstanceIdentifiers")},
			{
				Name:        "status_infos",
				Description: "The status of a read replica.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DBInstance.StatusInfos")},
			{
				Name:        "subnets",
				Description: "A list of subnet elements.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DBInstance.DBSubnetGroup.Subnets")},
			{
				Name:        "vpc_security_groups",
				Description: "A list of VPC security group elements that the DB instance belongs to.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DBInstance.VpcSecurityGroups")},
			{
				Name:        "tags_src",
				Description: "A list of tags attached to the DB Instance.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DBInstance.TagList")},

			// Standard columns
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(getRDSDBInstanceTurbotTags),
			},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBInstance.DBInstanceIdentifier")},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DBInstance.DBInstanceArn").Transform(arnToAkas),
			},
		}),
	}
}

//// LIST FUNCTION

func listRDSDBInstances(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {

	// Create Session
	svc, err := RDSClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("aws_rds_db_instance.listRDSDBInstances", "connection_error", err)
		return nil, err
	}

	// Limiting the results
	maxLimit := int32(100)
	if d.QueryContext.Limit != nil {
		limit := int32(*d.QueryContext.Limit)
		if limit < maxLimit {
			if limit < 20 {
				maxLimit = 20
			} else {
				maxLimit = limit
			}
		}
	}

	input := &rds.DescribeDBInstancesInput{
		MaxRecords: aws.Int32(maxLimit),
	}

	filters := buildRdsDbInstanceFilter(d.Quals)
	if len(filters) > 0 {
		input.Filters = filters
	}

	paginator := rds.NewDescribeDBInstancesPaginator(svc, input, func(o *rds.DescribeDBInstancesPaginatorOptions) {
		o.Limit = maxLimit
		o.StopOnDuplicateToken = true
	})

	// List call
	for paginator.HasMorePages() {
		output, err := paginator.NextPage(ctx)
		if err != nil {
			plugin.Logger(ctx).Error("aws_rds_db_instance.listRDSDBInstances", "api_error", err)
			return nil, err
		}

		for _, items := range output.DBInstances {
			if helpers.StringSliceContains(
				[]string{
					"aurora",
					"aurora-mysql",
					"aurora-postgresql",
					"mysql",
					"postgres",
				},
				*items.Engine) {
				d.StreamListItem(ctx, items)
			}

			// Context can be cancelled due to manual cancellation or the limit has been hit
			if d.RowsRemaining(ctx) == 0 {
				return nil, nil
			}
		}
	}

	return nil, err
}

//// HYDRATE FUNCTIONS

//// TRANSFORM FUNCTIONS

func getRDSDBInstanceTurbotTags(_ context.Context, d *transform.TransformData) (interface{}, error) {
	dbInstance := d.HydrateItem.(opengovernance.RDSDBInstance).Description.DBInstance

	if dbInstance.TagList != nil {
		turbotTagsMap := map[string]string{}
		for _, i := range dbInstance.TagList {
			turbotTagsMap[*i.Key] = *i.Value
		}
		return turbotTagsMap, nil
	}
	return nil, nil
}

//// UTILITY FUNCTIONS

// build rds db instance list call input filter
func buildRdsDbInstanceFilter(quals plugin.KeyColumnQualMap) []types.Filter {
	filters := make([]types.Filter, 0)
	filterQuals := map[string]string{
		"db_cluster_identifier": "db-cluster-id",
		"resource_id":           "dbi-resource-id",
		"engine":                "engine",
	}

	for columnName, filterName := range filterQuals {
		if quals[columnName] != nil {
			filter := types.Filter{
				Name: aws.String(filterName),
			}
			value := getQualsValueByColumn(quals, columnName, "string")
			val, ok := value.(string)
			if ok {
				filter.Values = []string{val}
			}
			filters = append(filters, filter)
		}
	}
	return filters
}
