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

func RecordinggetHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		queryParams := make([]string, 0)
		if val, ok := args["include_fields"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("include_fields=%v", val))
		}
		if val, ok := args["ttl"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("ttl=%v", val))
		}
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
		url := fmt.Sprintf("%s/meetings/%s/recordings%s", cfg.BaseURL, meetingId, queryString)
		req, err := http.NewRequest("GET", url, nil)
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
		var result interface{}
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

func CreateRecordinggetTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_meetings_meetingId_recordings",
		mcp.WithDescription("Get meeting recordings"),
		mcp.WithString("meetingId", mcp.Required(), mcp.Description("To get Cloud Recordings of a meeting, provide the meeting ID or meeting UUID. If the meeting ID is provided instead of UUID,the response will be for the latest meeting instance. \n\nTo get Cloud Recordings of a webinar, provide the webinar ID or the webinar UUID. If the webinar ID is provided instead of UUID,the response will be for the latest webinar instance. \n\nIf a UUID starts with \"/\" or contains \"//\" (example: \"/ajXp112QmuoKj4854875==\"), you must **double encode** the UUID before making an API request. ")),
		mcp.WithString("include_fields", mcp.Description("Get the `download_access_token` field for downloading meeting recordings.")),
		mcp.WithNumber("ttl", mcp.Description("Time to live (TTL) of the `download_access_token`. This is only valid if the `include_fields` query parameter contains `download_access_token`. The range is between 0-604800.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    RecordinggetHandler(cfg),
	}
}
