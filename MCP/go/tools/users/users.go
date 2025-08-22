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

func UsersHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["status"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("status=%v", val))
		}
		if val, ok := args["page_size"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("page_size=%v", val))
		}
		if val, ok := args["role_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("role_id=%v", val))
		}
		if val, ok := args["page_number"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("page_number=%v", val))
		}
		if val, ok := args["include_fields"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("include_fields=%v", val))
		}
		if val, ok := args["next_page_token"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("next_page_token=%v", val))
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
		url := fmt.Sprintf("%s/users%s", cfg.BaseURL, queryString)
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

func CreateUsersTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_users",
		mcp.WithDescription("List users"),
		mcp.WithString("status", mcp.Description("User statuses:<br>`active` - Users with an active status.<br>`inactive` - Users who are deactivated.<br>`pending` - Users with a pending status.")),
		mcp.WithNumber("page_size", mcp.Description("The number of records returned within a single API call.")),
		mcp.WithString("role_id", mcp.Description("Unique identifier of the role. Provide this parameter if you would like to filter the response by a specific role. You can retrieve Role IDs from [List Roles](https://marketplace.zoom.us/docs/api-reference/zoom-api/roles/roles) API. ")),
		mcp.WithString("page_number", mcp.Description("The page number of the current page in the returned records.")),
		mcp.WithString("include_fields", mcp.Description("Provide a value for this field if you would like to see the following attribute in the response of this API call:<br>\n\n`custom_attributes`: Returns custom attributes that are associated with the user.<br>`host_key`: Returns [host key](https://support.zoom.us/hc/en-us/articles/205172555-Using-your-host-key) of the user.")),
		mcp.WithString("next_page_token", mcp.Description("The next page token is used to paginate through large result sets. A next page token will be returned whenever the set of available results exceeds the current page size. The expiration period for this token is 15 minutes.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    UsersHandler(cfg),
	}
}
