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

func RecordingslistHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		userIdVal, ok := args["userId"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: userId"), nil
		}
		userId, ok := userIdVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: userId"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["page_size"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("page_size=%v", val))
		}
		if val, ok := args["next_page_token"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("next_page_token=%v", val))
		}
		if val, ok := args["mc"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("mc=%v", val))
		}
		if val, ok := args["trash"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("trash=%v", val))
		}
		if val, ok := args["from"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("from=%v", val))
		}
		if val, ok := args["to"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("to=%v", val))
		}
		if val, ok := args["trash_type"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("trash_type=%v", val))
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
		url := fmt.Sprintf("%s/users/%s/recordings%s", cfg.BaseURL, userId, queryString)
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

func CreateRecordingslistTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_users_userId_recordings",
		mcp.WithDescription("List all recordings"),
		mcp.WithString("userId", mcp.Required(), mcp.Description("The user ID or email address of the user. For user-level apps, pass `me` as the value for userId.")),
		mcp.WithNumber("page_size", mcp.Description("The number of records returned within a single API call.")),
		mcp.WithString("next_page_token", mcp.Description("The next page token is used to paginate through large result sets. A next page token will be returned whenever the set of available results exceeds the current page size. The expiration period for this token is 15 minutes.")),
		mcp.WithString("mc", mcp.Description("Query Metadata of Recording if an On-Premise Meeting Connector was used for the meeting.")),
		mcp.WithBoolean("trash", mcp.Description("Query trash.\n`true`: List recordings from trash.<br> `false`: Do not list recordings from the trash.<br> The default value is `false`. If you set it to `true`, you can use the `trash_type` property to indicate the type of Cloud recording that you need to retrieve. ")),
		mcp.WithString("from", mcp.Description("The start date in 'yyyy-mm-dd' UTC format for the date range for which you would like to retrieve recordings. The maximum range can be a month. If no value is provided for this field, the default will be current date. For example, if you make the API request on June 30, 2020, without providing the “from” and “to” parameters, by default the value of 'from' field will be “2020-06-30” and the value of the 'to' field will be “2020-07-01”. \n\n**Note**: The \"trash\" files cannot be filtered by date range and thus, the \"from\" and \"to\" fields should not be used for trash files.")),
		mcp.WithString("to", mcp.Description("End date in 'yyyy-mm-dd' 'yyyy-mm-dd' UTC format. ")),
		mcp.WithString("trash_type", mcp.Description("The type of Cloud recording that you would like to retrieve from the trash. The value can be one of the following:<br>\n    `meeting_recordings`: List all meeting recordings from the trash.<br>\n    `recording_file`: List all individual recording files from the trash. ")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    RecordingslistHandler(cfg),
	}
}
