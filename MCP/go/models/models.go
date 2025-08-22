package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// TagMap represents the TagMap schema from the OpenAPI specification
type TagMap struct {
}

// UpdateScheduleOutput represents the UpdateScheduleOutput schema from the OpenAPI specification
type UpdateScheduleOutput struct {
	Schedulearn interface{} `json:"ScheduleArn"`
}

// DeleteScheduleGroupOutput represents the DeleteScheduleGroupOutput schema from the OpenAPI specification
type DeleteScheduleGroupOutput struct {
}

// EcsParameters represents the EcsParameters schema from the OpenAPI specification
type EcsParameters struct {
	Platformversion interface{} `json:"PlatformVersion,omitempty"`
	Propagatetags interface{} `json:"PropagateTags,omitempty"`
	Placementstrategy interface{} `json:"PlacementStrategy,omitempty"`
	Tags interface{} `json:"Tags,omitempty"`
	Taskcount interface{} `json:"TaskCount,omitempty"`
	Capacityproviderstrategy interface{} `json:"CapacityProviderStrategy,omitempty"`
	Launchtype interface{} `json:"LaunchType,omitempty"`
	Placementconstraints interface{} `json:"PlacementConstraints,omitempty"`
	Referenceid interface{} `json:"ReferenceId,omitempty"`
	Group interface{} `json:"Group,omitempty"`
	Networkconfiguration interface{} `json:"NetworkConfiguration,omitempty"`
	Taskdefinitionarn interface{} `json:"TaskDefinitionArn"`
	Enableecsmanagedtags interface{} `json:"EnableECSManagedTags,omitempty"`
	Enableexecutecommand interface{} `json:"EnableExecuteCommand,omitempty"`
}

// SqsParameters represents the SqsParameters schema from the OpenAPI specification
type SqsParameters struct {
	Messagegroupid interface{} `json:"MessageGroupId,omitempty"`
}

// GetScheduleInput represents the GetScheduleInput schema from the OpenAPI specification
type GetScheduleInput struct {
}

// TagResourceOutput represents the TagResourceOutput schema from the OpenAPI specification
type TagResourceOutput struct {
}

// CreateScheduleGroupInput represents the CreateScheduleGroupInput schema from the OpenAPI specification
type CreateScheduleGroupInput struct {
	Tags interface{} `json:"Tags,omitempty"`
	Clienttoken interface{} `json:"ClientToken,omitempty"`
}

// ScheduleSummary represents the ScheduleSummary schema from the OpenAPI specification
type ScheduleSummary struct {
	Name interface{} `json:"Name,omitempty"`
	State interface{} `json:"State,omitempty"`
	Target interface{} `json:"Target,omitempty"`
	Arn interface{} `json:"Arn,omitempty"`
	Creationdate interface{} `json:"CreationDate,omitempty"`
	Groupname interface{} `json:"GroupName,omitempty"`
	Lastmodificationdate interface{} `json:"LastModificationDate,omitempty"`
}

// TagResourceInput represents the TagResourceInput schema from the OpenAPI specification
type TagResourceInput struct {
	Tags interface{} `json:"Tags"`
}

// PlacementConstraint represents the PlacementConstraint schema from the OpenAPI specification
type PlacementConstraint struct {
	Expression interface{} `json:"expression,omitempty"`
	TypeField interface{} `json:"type,omitempty"`
}

// SageMakerPipelineParameter represents the SageMakerPipelineParameter schema from the OpenAPI specification
type SageMakerPipelineParameter struct {
	Value interface{} `json:"Value"`
	Name interface{} `json:"Name"`
}

// CreateScheduleGroupOutput represents the CreateScheduleGroupOutput schema from the OpenAPI specification
type CreateScheduleGroupOutput struct {
	Schedulegrouparn interface{} `json:"ScheduleGroupArn"`
}

// GetScheduleGroupInput represents the GetScheduleGroupInput schema from the OpenAPI specification
type GetScheduleGroupInput struct {
}

// FlexibleTimeWindow represents the FlexibleTimeWindow schema from the OpenAPI specification
type FlexibleTimeWindow struct {
	Mode interface{} `json:"Mode"`
	Maximumwindowinminutes interface{} `json:"MaximumWindowInMinutes,omitempty"`
}

// UpdateScheduleInput represents the UpdateScheduleInput schema from the OpenAPI specification
type UpdateScheduleInput struct {
	Flexibletimewindow interface{} `json:"FlexibleTimeWindow"`
	State interface{} `json:"State,omitempty"`
	Target interface{} `json:"Target"`
	Clienttoken interface{} `json:"ClientToken,omitempty"`
	Enddate interface{} `json:"EndDate,omitempty"`
	Scheduleexpression interface{} `json:"ScheduleExpression"`
	Scheduleexpressiontimezone interface{} `json:"ScheduleExpressionTimezone,omitempty"`
	Startdate interface{} `json:"StartDate,omitempty"`
	Actionaftercompletion interface{} `json:"ActionAfterCompletion,omitempty"`
	Description interface{} `json:"Description,omitempty"`
	Groupname interface{} `json:"GroupName,omitempty"`
	Kmskeyarn interface{} `json:"KmsKeyArn,omitempty"`
}

// CreateScheduleInput represents the CreateScheduleInput schema from the OpenAPI specification
type CreateScheduleInput struct {
	Clienttoken interface{} `json:"ClientToken,omitempty"`
	Target interface{} `json:"Target"`
	Actionaftercompletion interface{} `json:"ActionAfterCompletion,omitempty"`
	Groupname interface{} `json:"GroupName,omitempty"`
	Scheduleexpressiontimezone interface{} `json:"ScheduleExpressionTimezone,omitempty"`
	State interface{} `json:"State,omitempty"`
	Description interface{} `json:"Description,omitempty"`
	Flexibletimewindow interface{} `json:"FlexibleTimeWindow"`
	Kmskeyarn interface{} `json:"KmsKeyArn,omitempty"`
	Scheduleexpression interface{} `json:"ScheduleExpression"`
	Startdate interface{} `json:"StartDate,omitempty"`
	Enddate interface{} `json:"EndDate,omitempty"`
}

// PlacementStrategy represents the PlacementStrategy schema from the OpenAPI specification
type PlacementStrategy struct {
	Field interface{} `json:"field,omitempty"`
	TypeField interface{} `json:"type,omitempty"`
}

// ListTagsForResourceOutput represents the ListTagsForResourceOutput schema from the OpenAPI specification
type ListTagsForResourceOutput struct {
	Tags interface{} `json:"Tags,omitempty"`
}

// DeleteScheduleInput represents the DeleteScheduleInput schema from the OpenAPI specification
type DeleteScheduleInput struct {
}

// SageMakerPipelineParameters represents the SageMakerPipelineParameters schema from the OpenAPI specification
type SageMakerPipelineParameters struct {
	Pipelineparameterlist interface{} `json:"PipelineParameterList,omitempty"`
}

// KinesisParameters represents the KinesisParameters schema from the OpenAPI specification
type KinesisParameters struct {
	Partitionkey interface{} `json:"PartitionKey"`
}

// CapacityProviderStrategyItem represents the CapacityProviderStrategyItem schema from the OpenAPI specification
type CapacityProviderStrategyItem struct {
	Base interface{} `json:"base,omitempty"`
	Capacityprovider interface{} `json:"capacityProvider"`
	Weight interface{} `json:"weight,omitempty"`
}

// ListScheduleGroupsInput represents the ListScheduleGroupsInput schema from the OpenAPI specification
type ListScheduleGroupsInput struct {
}

// EventBridgeParameters represents the EventBridgeParameters schema from the OpenAPI specification
type EventBridgeParameters struct {
	Detailtype interface{} `json:"DetailType"`
	Source interface{} `json:"Source"`
}

// GetScheduleOutput represents the GetScheduleOutput schema from the OpenAPI specification
type GetScheduleOutput struct {
	Name interface{} `json:"Name,omitempty"`
	Arn interface{} `json:"Arn,omitempty"`
	Creationdate interface{} `json:"CreationDate,omitempty"`
	Scheduleexpressiontimezone interface{} `json:"ScheduleExpressionTimezone,omitempty"`
	State interface{} `json:"State,omitempty"`
	Lastmodificationdate interface{} `json:"LastModificationDate,omitempty"`
	Startdate interface{} `json:"StartDate,omitempty"`
	Flexibletimewindow interface{} `json:"FlexibleTimeWindow,omitempty"`
	Groupname interface{} `json:"GroupName,omitempty"`
	Target interface{} `json:"Target,omitempty"`
	Enddate interface{} `json:"EndDate,omitempty"`
	Kmskeyarn interface{} `json:"KmsKeyArn,omitempty"`
	Description interface{} `json:"Description,omitempty"`
	Scheduleexpression interface{} `json:"ScheduleExpression,omitempty"`
	Actionaftercompletion interface{} `json:"ActionAfterCompletion,omitempty"`
}

// AwsVpcConfiguration represents the AwsVpcConfiguration schema from the OpenAPI specification
type AwsVpcConfiguration struct {
	Assignpublicip interface{} `json:"AssignPublicIp,omitempty"`
	Securitygroups interface{} `json:"SecurityGroups,omitempty"`
	Subnets interface{} `json:"Subnets"`
}

// DeadLetterConfig represents the DeadLetterConfig schema from the OpenAPI specification
type DeadLetterConfig struct {
	Arn interface{} `json:"Arn,omitempty"`
}

// ScheduleGroupSummary represents the ScheduleGroupSummary schema from the OpenAPI specification
type ScheduleGroupSummary struct {
	Arn interface{} `json:"Arn,omitempty"`
	Creationdate interface{} `json:"CreationDate,omitempty"`
	Lastmodificationdate interface{} `json:"LastModificationDate,omitempty"`
	Name interface{} `json:"Name,omitempty"`
	State interface{} `json:"State,omitempty"`
}

// CreateScheduleOutput represents the CreateScheduleOutput schema from the OpenAPI specification
type CreateScheduleOutput struct {
	Schedulearn interface{} `json:"ScheduleArn"`
}

// DeleteScheduleGroupInput represents the DeleteScheduleGroupInput schema from the OpenAPI specification
type DeleteScheduleGroupInput struct {
}

// ListSchedulesOutput represents the ListSchedulesOutput schema from the OpenAPI specification
type ListSchedulesOutput struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Schedules interface{} `json:"Schedules"`
}

// Tag represents the Tag schema from the OpenAPI specification
type Tag struct {
	Key interface{} `json:"Key"`
	Value interface{} `json:"Value"`
}

// TargetSummary represents the TargetSummary schema from the OpenAPI specification
type TargetSummary struct {
	Arn interface{} `json:"Arn"`
}

// UntagResourceOutput represents the UntagResourceOutput schema from the OpenAPI specification
type UntagResourceOutput struct {
}

// ListSchedulesInput represents the ListSchedulesInput schema from the OpenAPI specification
type ListSchedulesInput struct {
}

// GetScheduleGroupOutput represents the GetScheduleGroupOutput schema from the OpenAPI specification
type GetScheduleGroupOutput struct {
	Arn interface{} `json:"Arn,omitempty"`
	Creationdate interface{} `json:"CreationDate,omitempty"`
	Lastmodificationdate interface{} `json:"LastModificationDate,omitempty"`
	Name interface{} `json:"Name,omitempty"`
	State interface{} `json:"State,omitempty"`
}

// ListTagsForResourceInput represents the ListTagsForResourceInput schema from the OpenAPI specification
type ListTagsForResourceInput struct {
}

// ListScheduleGroupsOutput represents the ListScheduleGroupsOutput schema from the OpenAPI specification
type ListScheduleGroupsOutput struct {
	Schedulegroups interface{} `json:"ScheduleGroups"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// DeleteScheduleOutput represents the DeleteScheduleOutput schema from the OpenAPI specification
type DeleteScheduleOutput struct {
}

// NetworkConfiguration represents the NetworkConfiguration schema from the OpenAPI specification
type NetworkConfiguration struct {
	Awsvpcconfiguration interface{} `json:"awsvpcConfiguration,omitempty"`
}

// UntagResourceInput represents the UntagResourceInput schema from the OpenAPI specification
type UntagResourceInput struct {
}

// Target represents the Target schema from the OpenAPI specification
type Target struct {
	Ecsparameters interface{} `json:"EcsParameters,omitempty"`
	Kinesisparameters interface{} `json:"KinesisParameters,omitempty"`
	Retrypolicy interface{} `json:"RetryPolicy,omitempty"`
	Rolearn interface{} `json:"RoleArn"`
	Arn interface{} `json:"Arn"`
	Eventbridgeparameters interface{} `json:"EventBridgeParameters,omitempty"`
	Input interface{} `json:"Input,omitempty"`
	Sagemakerpipelineparameters interface{} `json:"SageMakerPipelineParameters,omitempty"`
	Sqsparameters interface{} `json:"SqsParameters,omitempty"`
	Deadletterconfig interface{} `json:"DeadLetterConfig,omitempty"`
}

// RetryPolicy represents the RetryPolicy schema from the OpenAPI specification
type RetryPolicy struct {
	Maximumretryattempts interface{} `json:"MaximumRetryAttempts,omitempty"`
	Maximumeventageinseconds interface{} `json:"MaximumEventAgeInSeconds,omitempty"`
}
