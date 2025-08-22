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

func GetphonerecordingsHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["page_size"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("page_size=%v", val))
		}
		if val, ok := args["next_page_token"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("next_page_token=%v", val))
		}
		if val, ok := args["from"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("from=%v", val))
		}
		if val, ok := args["to"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("to=%v", val))
		}
		if val, ok := args["owner_type"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("owner_type=%v", val))
		}
		if val, ok := args["recording_type"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("recording_type=%v", val))
		}
		if val, ok := args["site_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("site_id=%v", val))
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
		url := fmt.Sprintf("%s/phone/recordings%s", cfg.BaseURL, queryString)
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

func CreateGetphonerecordingsTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_phone_recordings",
		mcp.WithDescription("Get call recordings"),
		mcp.WithNumber("page_size", mcp.Description("The number of records returned within a single API call. The default is **30**, and the maximum is **100**.")),
		mcp.WithString("next_page_token", mcp.Description("The current page number of returned records.")),
		mcp.WithString("from", mcp.Description("Start date and time in **yyyy-mm-dd** format or **yyyy-MM-dd’T’HH:mm:ss’Z’** format. The date range defined by the from and to parameters should only be one month as the report includes only one month worth of data at once.\n")),
		mcp.WithString("to", mcp.Description("End date and time in **yyyy-mm-dd** format or **yyyy-MM-dd’T’HH:mm:ss’Z’** format, the same formats supported by the `from` parameter.\n\n")),
		mcp.WithString("owner_type", mcp.Description("The owner type. The allowed values are null, `user`, or `callQueue`. The default is null. If null, returns all owner types.\n")),
		mcp.WithString("recording_type", mcp.Description("The recording type. The allowed values are null, `OnDemand`, or `Automatic`. The default is null. If null, returns all recording types.\n")),
		mcp.WithString("site_id", mcp.Description("The site ID. The default is `All sites`.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    GetphonerecordingsHandler(cfg),
	}
}
