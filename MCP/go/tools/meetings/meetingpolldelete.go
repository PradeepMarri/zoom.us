package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/zoom-api/mcp-server/config"
	"github.com/zoom-api/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func MeetingpolldeleteHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		meetingIdVal, ok := args["meetingId"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: meetingId"), nil
		}
		meetingId, ok := meetingIdVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: meetingId"), nil
		}
		pollIdVal, ok := args["pollId"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: pollId"), nil
		}
		pollId, ok := pollIdVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: pollId"), nil
		}
		queryParams := make([]string, 0)
		// Handle multiple authentication parameters
		if cfg.BearerToken != "" {
			queryParams = append(queryParams, fmt.Sprintf("next_page_token=%s", cfg.BearerToken))
		}
		if cfg.APIKey != "" {
			queryParams = append(queryParams, fmt.Sprintf("page_number=%s", cfg.APIKey))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/meetings/%s/polls/%s%s", cfg.BaseURL, meetingId, pollId, queryString)
		req, err := http.NewRequest("DELETE", url, nil)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// Set authentication based on auth type
		// Handle multiple authentication parameters
		// API key already added to query string
		// API key already added to query string
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
		var result map[string]interface{}
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

func CreateMeetingpolldeleteTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("delete_meetings_meetingId_polls_pollId",
		mcp.WithDescription("Delete a meeting poll"),
		mcp.WithNumber("meetingId", mcp.Required(), mcp.Description("The meeting ID in **long** format. The data type of this field is \"long\"(represented as int64 in JSON).\n\nWhile storing it in your database, store it as a **long** data type and **not as an integer**, as the Meeting IDs can be longer than 10 digits.")),
		mcp.WithString("pollId", mcp.Required(), mcp.Description("The poll ID")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    MeetingpolldeleteHandler(cfg),
	}
}
