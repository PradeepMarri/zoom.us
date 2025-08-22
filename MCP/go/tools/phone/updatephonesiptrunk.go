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

func UpdatephonesiptrunkHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		sipTrunkIdVal, ok := args["sipTrunkId"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: sipTrunkId"), nil
		}
		sipTrunkId, ok := sipTrunkIdVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: sipTrunkId"), nil
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
		url := fmt.Sprintf("%s/accounts/%s/phone/sip_trunk/trunks/%s%s", cfg.BaseURL, sipTrunkId, accountId, queryString)
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

func CreateUpdatephonesiptrunkTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("patch_accounts_accountId_phone_sip_trunk_trunks_sipTrunkId",
		mcp.WithDescription("Update SIP trunk details"),
		mcp.WithString("sipTrunkId", mcp.Required(), mcp.Description("Unique identifier of the SIP trunk.")),
		mcp.WithString("accountId", mcp.Required(), mcp.Description("Unique identifier of the sub account.")),
		mcp.WithString("name", mcp.Description("Input parameter: Name of the SIP Trunk.")),
		mcp.WithString("carrier_account", mcp.Description("Input parameter: Account associated with the carrier.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    UpdatephonesiptrunkHandler(cfg),
	}
}
