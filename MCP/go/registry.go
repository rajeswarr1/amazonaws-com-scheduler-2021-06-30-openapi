package main

import (
	"github.com/amazon-eventbridge-scheduler/mcp-server/config"
	"github.com/amazon-eventbridge-scheduler/mcp-server/models"
	tools_schedule_groups "github.com/amazon-eventbridge-scheduler/mcp-server/tools/schedule_groups"
	tools_schedules "github.com/amazon-eventbridge-scheduler/mcp-server/tools/schedules"
	tools_tags "github.com/amazon-eventbridge-scheduler/mcp-server/tools/tags"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_schedule_groups.CreateListschedulegroupsTool(cfg),
		tools_schedule_groups.CreateCreateschedulegroupTool(cfg),
		tools_schedule_groups.CreateDeleteschedulegroupTool(cfg),
		tools_schedule_groups.CreateGetschedulegroupTool(cfg),
		tools_schedules.CreateListschedulesTool(cfg),
		tools_schedules.CreateUpdatescheduleTool(cfg),
		tools_schedules.CreateDeletescheduleTool(cfg),
		tools_schedules.CreateGetscheduleTool(cfg),
		tools_schedules.CreateCreatescheduleTool(cfg),
		tools_tags.CreateListtagsforresourceTool(cfg),
		tools_tags.CreateTagresourceTool(cfg),
		tools_tags.CreateUntagresourceTool(cfg),
	}
}
