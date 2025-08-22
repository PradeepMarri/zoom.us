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

func PhoneusercalllogsHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		if val, ok := args["from"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("from=%v", val))
		}
		if val, ok := args["to"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("to=%v", val))
		}
		if val, ok := args["type"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("type=%v", val))
		}
		if val, ok := args["next_page_token"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("next_page_token=%v", val))
		}
		if val, ok := args["phone_number"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("phone_number=%v", val))
		}
		if val, ok := args["time_type"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("time_type=%v", val))
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
		url := fmt.Sprintf("%s/phone/users/%s/call_logs%s", cfg.BaseURL, userId, queryString)
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

func CreatePhoneusercalllogsTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_phone_users_userId_call_logs",
		mcp.WithDescription("Get user's call logs"),
		mcp.WithString("userId", mcp.Required(), mcp.Description("The user ID or email address of the user. For user-level apps, pass `me` as the value for userId.")),
		mcp.WithNumber("page_size", mcp.Description("The number of records returned within a single API call.")),
		mcp.WithString("from", mcp.Required(), mcp.Description("Start date in 'yyyy-mm-dd' format. The date range defined by the \"from\" and \"to\" parameters should only be one month as the report includes only one month worth of data at once.")),
		mcp.WithString("to", mcp.Required(), mcp.Description("End date.")),
		mcp.WithString("type", mcp.Description("")),
		mcp.WithString("next_page_token", mcp.Description("The next page token is used to paginate through large result sets. A next page token will be returned whenever the set of available results exceeds the current page size. The expiration period for this token is 15 minutes.")),
		mcp.WithString("phone_number", mcp.Description("Filter API responses to include call logs of only the phone number defined in this field.")),
		mcp.WithString("time_type", mcp.Description("Enables you to sort call logs by start or end time. Choose the sort time value. Values include `startTime` or `endTime`.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    PhoneusercalllogsHandler(cfg),
	}
}
