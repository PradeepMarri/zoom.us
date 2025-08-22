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

func UpdatecallqueueHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		callQueueIdVal, ok := args["callQueueId"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: callQueueId"), nil
		}
		callQueueId, ok := callQueueIdVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: callQueueId"), nil
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
		url := fmt.Sprintf("%s/phone/call_queues/%s%s", cfg.BaseURL, callQueueId, queryString)
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

func CreateUpdatecallqueueTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("patch_phone_call_queues_callQueueId",
		mcp.WithDescription("Update call queue details"),
		mcp.WithString("callQueueId", mcp.Required(), mcp.Description("Unique Identifier of the Call Queue.")),
		mcp.WithString("site_id", mcp.Description("Input parameter: Unique identifier of the [site](https://support.zoom.us/hc/en-us/articles/360020809672-Managing-Multiple-Sites) where the Call Queue is assigned.")),
		mcp.WithString("status", mcp.Description("Input parameter: Status of the Call Queue. Allowed values:<br>\n`active`<br>\n`inactive`")),
		mcp.WithString("timezone", mcp.Description("Input parameter: [Timezone](https://marketplace.zoom.us/docs/api-reference/other-references/abbreviation-lists#timezones) of the Call Queue.")),
		mcp.WithString("description", mcp.Description("Input parameter: Description for the Call Queue.")),
		mcp.WithNumber("extension_number", mcp.Description("Input parameter: Phone extension number for the site.<br>\n\nIf a site code has been [assigned](https://support.zoom.us/hc/en-us/articles/360020809672-Managing-Multiple-Sites#h_79ca9c8f-c97b-4486-aa59-d0d9d31a525b) to the site, provide the short extension number instead of the original extension number.")),
		mcp.WithString("name", mcp.Description("Input parameter: Name of the Call Queue.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    UpdatecallqueueHandler(cfg),
	}
}
