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

func UserupdateHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		queryParams := make([]string, 0)
		if val, ok := args["login_type"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("login_type=%v", val))
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
		url := fmt.Sprintf("%s/users/%s%s", cfg.BaseURL, userId, queryString)
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

func CreateUserupdateTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("patch_users_userId",
		mcp.WithDescription("Update a user"),
		mcp.WithString("userId", mcp.Required(), mcp.Description("The user ID or email address of the user. For user-level apps, pass `me` as the value for userId.")),
		mcp.WithString("login_type", mcp.Description("`0` - Facebook.<br>`1` - Google.<br>`99` - API.<br>`100` - Zoom.<br>`101` - SSO.")),
		mcp.WithNumber("pmi", mcp.Description("Input parameter: Personal meeting ID: length must be 10.")),
		mcp.WithString("phone_number", mcp.Description("Input parameter: **Note:** This field has been **deprecated** and will not be supported in the future. Use the **phone_numbers** field instead to assign phone number(s) to a user.\n\n\nPhone number of the user. To update a phone number, you must also provide the `phone_country` field.")),
		mcp.WithString("language", mcp.Description("Input parameter: language")),
		mcp.WithNumber("type", mcp.Description("Input parameter: User types:<br>`1` - Basic.<br>`2` - Licensed.<br>`3` - On-prem.<br>`99` - None (this can only be set with `ssoCreate`).")),
		mcp.WithString("last_name", mcp.Description("Input parameter: User's last name. Cannot contain more than 5 Chinese characters.")),
		mcp.WithString("location", mcp.Description("Input parameter: User's location.")),
		mcp.WithObject("custom_attributes", mcp.Description("Input parameter: Custom attribute(s) of the user.")),
		mcp.WithString("group_id", mcp.Description("Input parameter: Provide unique identifier of the group that you would like to add a [pending user](https://support.zoom.us/hc/en-us/articles/201363183-Managing-users#h_13c87a2a-ecd6-40ad-be61-a9935e660edb) to. The value of this field can be retrieved from [List Groups](https://marketplace.zoom.us/docs/api-reference/zoom-api/groups/groups) API.")),
		mcp.WithString("manager", mcp.Description("Input parameter: The manager for the user.")),
		mcp.WithString("phone_country", mcp.Description("Input parameter: **Note:** This field has been **deprecated** and will not be supported in the future. Use the **country** field of the **phone_numbers** object instead to select the country for the phone number.\n\n\n\n[Country ID](https://marketplace.zoom.us/docs/api-reference/other-references/abbreviation-lists#countries) of the phone number. For example, if the phone number provided in the `phone_number` field is a Brazil based number, the value of the `phone_country` field should be `BR`.")),
		mcp.WithObject("phone_numbers", mcp.Description("")),
		mcp.WithString("first_name", mcp.Description("Input parameter: User's first name. Cannot contain more than 5 Chinese characters.")),
		mcp.WithString("dept", mcp.Description("Input parameter: Department for user profile: use for report.")),
		mcp.WithString("timezone", mcp.Description("Input parameter: The time zone ID for a user profile. For this parameter value please refer to the ID value in the [timezone](https://marketplace.zoom.us/docs/api-reference/other-references/abbreviation-lists#timezones) list.")),
		mcp.WithString("company", mcp.Description("Input parameter: User's company.")),
		mcp.WithBoolean("use_pmi", mcp.Description("Input parameter: Use Personal Meeting ID for instant meetings.")),
		mcp.WithString("vanity_name", mcp.Description("Input parameter: Personal meeting room name.")),
		mcp.WithString("cms_user_id", mcp.Description("Input parameter: Kaltura user ID.")),
		mcp.WithString("host_key", mcp.Description("Input parameter: Host key. It should be a 6-10 digit number.")),
		mcp.WithString("job_title", mcp.Description("Input parameter: User's job title.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    UserupdateHandler(cfg),
	}
}
