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

func WebinarregistrantcreateHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		webinarIdVal, ok := args["webinarId"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: webinarId"), nil
		}
		webinarId, ok := webinarIdVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: webinarId"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["occurrence_ids"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("occurrence_ids=%v", val))
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
		url := fmt.Sprintf("%s/webinars/%s/registrants%s", cfg.BaseURL, webinarId, queryString)
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyBytes))
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

func CreateWebinarregistrantcreateTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_webinars_webinarId_registrants",
		mcp.WithDescription("Add a webinar registrant"),
		mcp.WithNumber("webinarId", mcp.Required(), mcp.Description("The webinar ID in \"**long**\" format(represented as int64 data type in JSON). ")),
		mcp.WithString("occurrence_ids", mcp.Description("Occurrence ID. Get this value from the webinar get API. Multiple values separated by a comma.")),
		mcp.WithString("purchasing_time_frame", mcp.Description("Input parameter: This field can be included to gauge interest of webinar attendees towards buying your product or service.\n\nPurchasing Time Frame:<br>`Within a month`<br>`1-3 months`<br>`4-6 months`<br>`More than 6 months`<br>`No timeframe`")),
		mcp.WithString("email", mcp.Description("Input parameter: A valid email address of the registrant.")),
		mcp.WithString("first_name", mcp.Description("Input parameter: Registrant's first name.")),
		mcp.WithString("city", mcp.Description("Input parameter: Registrant's city.")),
		mcp.WithString("job_title", mcp.Description("Input parameter: Registrant's job title.")),
		mcp.WithString("country", mcp.Description("Input parameter: Registrant's country. The value of this field must be in two-letter abbreviated form and must match the ID field provided in the [Countries](https://marketplace.zoom.us/docs/api-reference/other-references/abbreviation-lists#countries) table.")),
		mcp.WithString("comments", mcp.Description("Input parameter: A field that allows registrants to provide any questions or comments that they might have.")),
		mcp.WithString("address", mcp.Description("Input parameter: Registrant's address.")),
		mcp.WithString("org", mcp.Description("Input parameter: Registrant's Organization.")),
		mcp.WithString("role_in_purchase_process", mcp.Description("Input parameter: Role in Purchase Process:<br>`Decision Maker`<br>`Evaluator/Recommender`<br>`Influencer`<br>`Not involved` ")),
		mcp.WithString("industry", mcp.Description("Input parameter: Registrant's Industry.")),
		mcp.WithString("phone", mcp.Description("Input parameter: Registrant's Phone number.")),
		mcp.WithString("state", mcp.Description("Input parameter: Registrant's State/Province.")),
		mcp.WithArray("custom_questions", mcp.Description("Input parameter: Custom questions.")),
		mcp.WithString("zip", mcp.Description("Input parameter: Registrant's Zip/Postal Code.")),
		mcp.WithString("last_name", mcp.Description("Input parameter: Registrant's last name.")),
		mcp.WithString("no_of_employees", mcp.Description("Input parameter: Number of Employees:<br>`1-20`<br>`21-50`<br>`51-100`<br>`101-500`<br>`500-1,000`<br>`1,001-5,000`<br>`5,001-10,000`<br>`More than 10,000`")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    WebinarregistrantcreateHandler(cfg),
	}
}
