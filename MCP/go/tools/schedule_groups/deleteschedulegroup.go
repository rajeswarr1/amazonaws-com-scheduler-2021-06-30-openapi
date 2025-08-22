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

func DeleteschedulegroupHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/schedule-groups/%s%s", cfg.BaseURL, Name, queryString)
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
		var result models.DeleteScheduleGroupOutput
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

func CreateDeleteschedulegroupTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("delete_schedule-groups_Name",
		mcp.WithDescription("<p>Deletes the specified schedule group. Deleting a schedule group results in EventBridge Scheduler deleting all schedules associated with the group. When you delete a group, it remains in a <code>DELETING</code> state until all of its associated schedules are deleted. Schedules associated with the group that are set to run while the schedule group is in the process of being deleted might continue to invoke their targets until the schedule group and its associated schedules are deleted.</p> <note> <p> This operation is eventually consistent. </p> </note>"),
		mcp.WithString("clientToken", mcp.Description(" Unique, case-sensitive identifier you provide to ensure the idempotency of the request. If you do not specify a client token, EventBridge Scheduler uses a randomly generated token for the request to ensure idempotency. ")),
		mcp.WithString("Name", mcp.Required(), mcp.Description("The name of the schedule group to delete.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    DeleteschedulegroupHandler(cfg),
	}
}
