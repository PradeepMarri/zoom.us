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

func UpdatesipphoneHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		phoneIdVal, ok := args["phoneId"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: phoneId"), nil
		}
		phoneId, ok := phoneIdVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: phoneId"), nil
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
		url := fmt.Sprintf("%s/sip_phones/%s%s", cfg.BaseURL, phoneId, queryString)
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

func CreateUpdatesipphoneTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("patch_sip_phones_phoneId",
		mcp.WithDescription("Update SIP phone"),
		mcp.WithString("phoneId", mcp.Required(), mcp.Description("Unique Identifier of the SIP Phone. This can be retrieved from the List SIP Phones API.")),
		mcp.WithString("transport_protocol2", mcp.Description("Input parameter: Protocols supported by the SIP provider.<br> The value must be either `UDP`, `TCP`, `TLS`, `AUTO`.")),
		mcp.WithString("register_server2", mcp.Required(), mcp.Description("Input parameter: IP address of the server that accepts REGISTER requests. Note that if you are using the UDP transport protocol, the default port is 5060. If you are using UDP with a different port number, that port number must be included with the IP address.")),
		mcp.WithString("transport_protocol", mcp.Description("Input parameter: Protocols supported by the SIP provider.<br> The value must be either `UDP`, `TCP`, `TLS`, `AUTO`.")),
		mcp.WithString("password", mcp.Required(), mcp.Description("Input parameter: The password generated for the user in the SIP account.")),
		mcp.WithString("proxy_server", mcp.Required(), mcp.Description("Input parameter: IP address of the proxy server for SIP requests. Note that if you are using the UDP transport protocol, the default port is 5060. If you are using UDP with a different port number, that port number must be included with the IP address. If you are not using a proxy server, this value can be the same as the Register Server.")),
		mcp.WithNumber("registration_expire_time", mcp.Description("Input parameter: The number of minutes after which the SIP registration of the Zoom client user will expire, and the client will auto register to the SIP server.")),
		mcp.WithString("authorization_name", mcp.Required(), mcp.Description("Input parameter: Authorization name of the user  registered for SIP Phone.")),
		mcp.WithString("register_server", mcp.Required(), mcp.Description("Input parameter: IP address of the server that accepts REGISTER requests. Note that if you are using the UDP transport protocol, the default port is 5060. If you are using UDP with a different port number, that port number must be included with the IP address.")),
		mcp.WithString("domain", mcp.Required(), mcp.Description("Input parameter: The name or IP address of your provider’s SIP domain. (example: CDC.WEB).")),
		mcp.WithString("proxy_server2", mcp.Required(), mcp.Description("Input parameter: IP address of the proxy server for SIP requests. Note that if you are using the UDP transport protocol, the default port is 5060. If you are using UDP with a different port number, that port number must be included with the IP address. If you are not using a proxy server, this value can be the same as the Register Server.")),
		mcp.WithString("transport_protocol3", mcp.Description("Input parameter: Protocols supported by the SIP provider.<br> The value must be either `UDP`, `TCP`, `TLS`, `AUTO`.")),
		mcp.WithString("user_name", mcp.Required(), mcp.Description("Input parameter: The phone number associated with the user in the SIP account.")),
		mcp.WithString("register_server3", mcp.Required(), mcp.Description("Input parameter: IP address of the server that accepts REGISTER requests. Note that if you are using the UDP transport protocol, the default port is 5060. If you are using UDP with a different port number, that port number must be included with the IP address.")),
		mcp.WithString("voice_mail", mcp.Required(), mcp.Description("Input parameter: The number to dial for checking voicemail.")),
		mcp.WithString("proxy_server3", mcp.Required(), mcp.Description("Input parameter: IP address of the proxy server for SIP requests. Note that if you are using the UDP transport protocol, the default port is 5060. If you are using UDP with a different port number, that port number must be included with the IP address. If you are not using a proxy server, this value can be the same as the Register Server.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    UpdatesipphoneHandler(cfg),
	}
}
