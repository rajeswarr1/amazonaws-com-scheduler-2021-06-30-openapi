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

func ListschedulesHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["ScheduleGroup"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("ScheduleGroup=%v", val))
		}
		if val, ok := args["MaxResults"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("MaxResults=%v", val))
		}
		if val, ok := args["NamePrefix"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("NamePrefix=%v", val))
		}
		if val, ok := args["NextToken"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("NextToken=%v", val))
		}
		if val, ok := args["State"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("State=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/schedules%s", cfg.BaseURL, queryString)
		req, err := http.NewRequest("GET", url, nil)
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
		var result models.ListSchedulesOutput
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

func CreateListschedulesTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_schedules",
		mcp.WithDescription("Returns a paginated list of your EventBridge Scheduler schedules."),
		mcp.WithString("ScheduleGroup", mcp.Description("If specified, only lists the schedules whose associated schedule group matches the given filter.")),
		mcp.WithNumber("MaxResults", mcp.Description("If specified, limits the number of results returned by this operation. The operation also returns a <code>NextToken</code> which you can use in a subsequent operation to retrieve the next set of results.")),
		mcp.WithString("NamePrefix", mcp.Description("Schedule name prefix to return the filtered list of resources.")),
		mcp.WithString("NextToken", mcp.Description("The token returned by a previous call to retrieve the next set of results.")),
		mcp.WithString("State", mcp.Description("If specified, only lists the schedules whose current state matches the given filter.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    ListschedulesHandler(cfg),
	}
}
