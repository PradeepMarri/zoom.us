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

func GetaccountlocksettingsHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		accountIdVal, ok := args["accountId"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: accountId"), nil
		}
		accountId, ok := accountIdVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: accountId"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["option"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("option=%v", val))
		}
		if val, ok := args["custom_query_fields"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("custom_query_fields=%v", val))
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
		url := fmt.Sprintf("%s/accounts/%s/lock_settings%s", cfg.BaseURL, accountId, queryString)
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

func CreateGetaccountlocksettingsTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_accounts_accountId_lock_settings",
		mcp.WithDescription("Get locked settings"),
		mcp.WithString("accountId", mcp.Required(), mcp.Description("Unique Identifier of the account. To retrieve locked settings of the master account or a regular account, provide \"me\" as the value of this field. <br> To retrieve locked settings of a sub account, provide the Account ID of the sub account in this field.")),
		mcp.WithString("option", mcp.Description("`meeting_security`: Use this query parameter to view meeting security settings applied on the account.<br>")),
		mcp.WithString("custom_query_fields", mcp.Description("Provide the name of the field by which you would like to filter the response. For example, if you provide \"host_video\" as the value of this field, you will get a response similar to the following:<br>\n{\n    \"schedule_meeting\": {\n        \"host_video\": false\n    }\n}\n<br>You can provide multiple values by separating them with commas(example: \"host_video,participant_video”).")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    GetaccountlocksettingsHandler(cfg),
	}
}
