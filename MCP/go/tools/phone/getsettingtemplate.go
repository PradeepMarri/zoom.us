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

func GetsettingtemplateHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		templateIdVal, ok := args["templateId"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: templateId"), nil
		}
		templateId, ok := templateIdVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: templateId"), nil
		}
		queryParams := make([]string, 0)
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
		url := fmt.Sprintf("%s/phone/setting_templates/%s%s", cfg.BaseURL, templateId, queryString)
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

func CreateGetsettingtemplateTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_phone_setting_templates_templateId",
		mcp.WithDescription("Get setting template details"),
		mcp.WithString("templateId", mcp.Required(), mcp.Description("Unique identifier of the template.")),
		mcp.WithString("custom_query_fields", mcp.Description("Provide the name of the field to use to filter the response. For example, if you provide \"description\" as the value of the field, you will get a response similar to the following: {“description”: “template description”}.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    GetsettingtemplateHandler(cfg),
	}
}
