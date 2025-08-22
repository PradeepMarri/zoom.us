package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"bytes"

	"github.com/zoom-api/mcp-server/config"
	"github.com/zoom-api/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func AssignsipconfigHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
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
		// Create properly typed request body using the generated schema
		var requestBody map[string]interface{}
		
		// Optimized: Single marshal/unmarshal with JSON tags handling field mapping
		if argsJSON, err := json.Marshal(args); err == nil {
			if err := json.Unmarshal(argsJSON, &requestBody); err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("Failed to convert arguments to request type: %v", err)), nil
			}
		} else {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to marshal arguments: %v", err)), nil
		}
		
		bodyBytes, err := json.Marshal(requestBody)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to encode request body", err), nil
		}
		url := fmt.Sprintf("%s/accounts/%s/sip_trunk/settings%s", cfg.BaseURL, queryString)
		req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(bodyBytes))
		req.Header.Set("Content-Type", "application/json")
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

func CreateAssignsipconfigTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("patch_accounts_accountId_sip_trunk_settings",
		mcp.WithDescription("Assign SIP trunk configuration"),
		mcp.WithNumber("show_zoom_provided_numbers", mcp.Description("Input parameter: If the value of this option is set to `0`, the numbers provided by Zoom will be displayed in the account's list of available call-out and call-in numbers in the Zoom Web Portal and Zoom Client. \n\nIf the value of this option is set to `1`, the Zoom provided numbers will be shown in the Zoom Web Portal but will not be used unless specified by the user.<br> \n\nIf the value of this option is set to `2`, all Zoom provided numbers will be deleted and only internal numbers (provided by carrier partners) will be used.")),
		mcp.WithBoolean("enable", mcp.Description("Input parameter: Enable or delete the configuration.<br> The values can be one of the following:<br> `true`: Enable configuration.<br> `false`: Delete configuration")),
		mcp.WithBoolean("show_callout_internal_number", mcp.Description("Input parameter: If the value of this option is set to `true`, the call-out numbers provided by the Zoom carrier partners will be displayed in the account's list of available call-out numbers in the Zoom Web Portal and Zoom Client.")),
		mcp.WithNumber("show_zoom_provided_callout_countries", mcp.Description("Input parameter: If the value of this option is set to `0`, the call-out countries list provided by Zoom will be [displayed](https://support.zoom.us/hc/en-us/articles/200942859-Using-telephone-call-out) in the account's list of available call-out countries. \n\nIf the value of this option is set to `1`, the Zoom provided call-out countries will be hidden from the user's account.<br> \n\nIf the value of this option is set to `2`, all Zoom provided countries will be deleted and only internal countries (provided by carrier partners) will be used.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    AssignsipconfigHandler(cfg),
	}
}
