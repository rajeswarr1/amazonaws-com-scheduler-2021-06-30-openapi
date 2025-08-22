package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/amazon-eventbridge-scheduler/mcp-server/config"
	"github.com/amazon-eventbridge-scheduler/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func DeletescheduleHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		NameVal, ok := args["Name"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: Name"), nil
		}
		Name, ok := NameVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: Name"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["clientToken"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("clientToken=%v", val))
		}
		if val, ok := args["groupName"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("groupName=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/schedules/%s%s", cfg.BaseURL, Name, queryString)
		req, err := http.NewRequest("DELETE", url, nil)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// Set authentication based on auth type
		// Handle multiple authentication parameters
		if cfg.BearerToken != "" {
			req.Header.Set("X-Amz-Security-Token", cfg.BearerToken)
		}
		req.Header.Set("Accept", "application/json")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Request failed", err), nil
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to read response body", err), nil
		}

		if resp.StatusCode >= 400 {
			return mcp.NewToolResultError(fmt.Sprintf("API error: %s", body)), nil
		}
		// Use properly typed response
		var result models.DeleteScheduleOutput
		if err := json.Unmarshal(body, &result); err != nil {
			// Fallback to raw text if unmarshaling fails
			return mcp.NewToolResultText(string(body)), nil
		}

		prettyJSON, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to format JSON", err), nil
		}

		return mcp.NewToolResultText(string(prettyJSON)), nil
	}
}

func CreateDeletescheduleTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("delete_schedules_Name",
		mcp.WithDescription("Deletes the specified schedule."),
		mcp.WithString("clientToken", mcp.Description(" Unique, case-sensitive identifier you provide to ensure the idempotency of the request. If you do not specify a client token, EventBridge Scheduler uses a randomly generated token for the request to ensure idempotency. ")),
		mcp.WithString("groupName", mcp.Description("The name of the schedule group associated with this schedule. If you omit this, the default schedule group is used.")),
		mcp.WithString("Name", mcp.Required(), mcp.Description("The name of the schedule to delete.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    DeletescheduleHandler(cfg),
	}
}
