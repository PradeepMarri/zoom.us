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

func RecordingsettingsupdateHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		meetingIdVal, ok := args["meetingId"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: meetingId"), nil
		}
		meetingId, ok := meetingIdVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: meetingId"), nil
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
		url := fmt.Sprintf("%s/meetings/%s/recordings/settings%s", cfg.BaseURL, meetingId, queryString)
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

func CreateRecordingsettingsupdateTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("patch_meetings_meetingId_recordings_settings",
		mcp.WithDescription("Update meeting recording settings"),
		mcp.WithString("meetingId", mcp.Required(), mcp.Description("To get Cloud Recordings of a meeting, provide the meeting ID or meeting UUID. If the meeting ID is provided instead of UUID,the response will be for the latest meeting instance. \n\nTo get Cloud Recordings of a webinar, provide the webinar ID or the webinar UUID. If the webinar ID is provided instead of UUID,the response will be for the latest webinar instance. \n\nIf a UUID starts with \"/\" or contains \"//\" (example: \"/ajXp112QmuoKj4854875==\"), you must **double encode** the UUID before making an API request. ")),
		mcp.WithString("password", mcp.Description("Input parameter: Enable password protection for the recording by setting a password. The password must have a minimum of **eight** characters with a mix of numbers, letters and special characters.<br><br>\n**Note:** If the account owner or the admin has set minimum password strength requirements for recordings via Account Settings, the password value provided here must meet those requirements. <br><br>If the requirements are enabled, you can view those requirements by calling either the [Get User Settings API](https://marketplace.zoom.us/docs/api-reference/zoom-api/users/usersettings) or the [Get Account Settings](https://marketplace.zoom.us/docs/api-reference/zoom-api/accounts/accountsettings) API. ")),
		mcp.WithBoolean("send_email_to_host", mcp.Description("Input parameter: Send an email to host when someone registers to view the recording. This applies for On-demand recordings only.")),
		mcp.WithBoolean("show_social_share_buttons", mcp.Description("Input parameter: Show social share buttons on registration page. This applies for On-demand recordings only.")),
		mcp.WithNumber("approval_type", mcp.Description("Input parameter: Approval type for the registration.<br>\n`0`- Automatically approve the registration when a user registers.<br>\n`1` - Manually approve or deny the registration of a user.<br>\n`2` - No registration required to view the recording.")),
		mcp.WithString("authentication_option", mcp.Description("Input parameter: Authentication Options.")),
		mcp.WithString("share_recording", mcp.Description("Input parameter: Determine how the meeting recording is shared.")),
		mcp.WithString("topic", mcp.Description("Input parameter: Name of the recording.")),
		mcp.WithBoolean("recording_authentication", mcp.Description("Input parameter: Only authenticated users can view.")),
		mcp.WithBoolean("viewer_download", mcp.Description("Input parameter: Determine whether a viewer can download the recording file or not.")),
		mcp.WithString("authentication_domains", mcp.Description("Input parameter: Authentication domains.")),
		mcp.WithBoolean("on_demand", mcp.Description("Input parameter: Determine whether registration  isrequired to view the recording.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    RecordingsettingsupdateHandler(cfg),
	}
}
