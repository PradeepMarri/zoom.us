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

func ChangezoomroomsappversionHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		roomIdVal, ok := args["roomId"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: roomId"), nil
		}
		roomId, ok := roomIdVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: roomId"), nil
		}
		deviceIdVal, ok := args["deviceId"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: deviceId"), nil
		}
		deviceId, ok := deviceIdVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: deviceId"), nil
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
		url := fmt.Sprintf("%s/rooms/%s/devices/%s/app_version%s", cfg.BaseURL, roomId, deviceId, queryString)
		req, err := http.NewRequest("PUT", url, bytes.NewBuffer(bodyBytes))
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

func CreateChangezoomroomsappversionTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("put_rooms_roomId_devices_deviceId_app_version",
		mcp.WithDescription("Change Zoom Rooms' app version"),
		mcp.WithString("roomId", mcp.Required(), mcp.Description("Unique Identifier of the Zoom Room.")),
		mcp.WithString("deviceId", mcp.Required(), mcp.Description("Unique Identifier of the Mac or the Windows device. The value of this field can be retrieved from the [List Zoom Room Devices API](https://marketplace.zoom.us/docs/api-reference/zoom-api/rooms/listzrdevices).")),
		mcp.WithString("action", mcp.Description("Input parameter: Specify one of the following values for this field:\n\n`upgrade`: Upgrade to the latest Zoom Rooms App Version.<br>\n`downgrade`: Downgrade the Zoom Rooms App Version.<br>\n`cancel`: Cancel an ongoing upgrade or downgrade process.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    ChangezoomroomsappversionHandler(cfg),
	}
}
