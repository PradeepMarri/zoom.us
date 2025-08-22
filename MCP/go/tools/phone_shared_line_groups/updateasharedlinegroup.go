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

func UpdateasharedlinegroupHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		sharedLineGroupIdVal, ok := args["sharedLineGroupId"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: sharedLineGroupId"), nil
		}
		sharedLineGroupId, ok := sharedLineGroupIdVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: sharedLineGroupId"), nil
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
		url := fmt.Sprintf("%s/phone/shared_line_groups/%s%s", cfg.BaseURL, sharedLineGroupId, queryString)
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

func CreateUpdateasharedlinegroupTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("patch_phone_shared_line_groups_sharedLineGroupId",
		mcp.WithDescription("Update a shared line group"),
		mcp.WithString("sharedLineGroupId", mcp.Required(), mcp.Description("Unique identifier of the shared line group that is to be updated.")),
		mcp.WithString("display_name", mcp.Description("Input parameter: Display Name of the Shared Line Group.")),
		mcp.WithNumber("extension_number", mcp.Description("Input parameter: Extension number assigned to the Shared Line Group.")),
		mcp.WithObject("primary_number", mcp.Description("Input parameter: If you have multiple direct phone numbers assigned to the shared line group, select a number from those numbers as the primary number. The primary number shares the same line as the extension number. This means if a caller is routed to the shared line group through an auto receptionist, the line associated with the primary number will be used. A pending number cannot be used as a Primary Number.")),
		mcp.WithString("status", mcp.Description("Input parameter: Status of the Shared Line Group.")),
		mcp.WithString("timezone", mcp.Description("Input parameter: Timezone to be used for the Business Hours. A value should be provided from the IDs listed [here](https://marketplace.zoom.us/docs/api-reference/other-references/abbreviation-lists#timezones).")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    UpdateasharedlinegroupHandler(cfg),
	}
}
