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

func DeleteusersettingHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		settingTypeVal, ok := args["settingType"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: settingType"), nil
		}
		settingType, ok := settingTypeVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: settingType"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["shared_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("shared_id=%v", val))
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
		url := fmt.Sprintf("%s/phone/users/%s/settings/%s%s", cfg.BaseURL, userId, settingType, queryString)
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

func CreateDeleteusersettingTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("delete_phone_users_userId_settings_settingType",
		mcp.WithDescription("Remove shared access"),
		mcp.WithString("userId", mcp.Required(), mcp.Description("Unique identifier of the user.")),
		mcp.WithString("settingType", mcp.Required(), mcp.Description("Corresponds to the setting item you wish to remove. Allowed values: `voice_mail`")),
		mcp.WithString("shared_id", mcp.Required(), mcp.Description("Required only for voicemail setting type.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    DeleteusersettingHandler(cfg),
	}
}
