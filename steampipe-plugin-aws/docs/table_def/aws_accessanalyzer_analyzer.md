# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>name</td><td>The name of the Analyzer.</td></tr>
	<tr><td>arn</td><td>The ARN of the analyzer.</td></tr>
	<tr><td>status</td><td>The status of the analyzer.</td></tr>
	<tr><td>type</td><td>The type of analyzer, which corresponds to the zone of trust chosen for the analyzer.</td></tr>
	<tr><td>created_at</td><td>A timestamp for the time at which the analyzer was created.</td></tr>
	<tr><td>last_resource_analyzed</td><td>The resource that was most recently analyzed by the analyzer.</td></tr>
	<tr><td>last_resource_analyzed_at</td><td>The time at which the most recently analyzed resource was analyzed.</td></tr>
	<tr><td>status_reason</td><td>The statusReason provides more details about the current status of the analyzer.</td></tr>
	<tr><td>findings</td><td>A list of findings retrieved from the analyzer that match the filter criteria specified, if any.</td></tr>
	<tr><td>title</td><td>Title of the resource.</td></tr>
	<tr><td>tags</td><td>A map of tags for the resource.</td></tr>
	<tr><td>akas</td><td>Array of globally unique identifier strings (also known as) for the resource.</td></tr>
	<tr><td>partition</td><td>The AWS partition in which the resource is located (aws, aws-cn, or aws-us-gov).</td></tr>
	<tr><td>region</td><td>The AWS Region in which the resource is located.</td></tr>
	<tr><td>account_id</td><td>The AWS Account ID in which the resource is located.</td></tr>
	<tr><td>platform_account_id</td><td>The Platform Account ID in which the resource is located.</td></tr>
	<tr><td>platform_resource_id</td><td>The unique ID of the resource in opengovernance.</td></tr>
	<tr><td>platform_metadata</td><td>Platform Metadata of the AWS resource.</td></tr>
</table>