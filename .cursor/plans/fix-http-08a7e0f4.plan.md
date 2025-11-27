<!-- 08a7e0f4-b20d-478f-89a6-f1f80baf4fbf d4531d4d-bbed-43d9-bcfe-dc137f93e1a0 -->
# Fix HTTP Authentication Context Flow

## Problem

The MCP server is failing authentication in HTTP mode because:

1. API keys are correctly extracted from HTTP `Authorization` headers and stored in context
2. Tool handlers are registered with `context.Background()` instead of the request context
3. The `mcp-go` library's `AddTool` signature doesn't accept context: `func(arguments map[string]interface{}) (*mcp.CallToolResult, error)`
4. When handlers try to get API key from context, they fail and fall back to empty `cfg.PlantonAPIKey`

## Investigation Steps

### 1. Check mcp-go Library Support

Investigate whether the `mcp-go` library supports passing request context to tool handlers:

- Check if there's an alternative API that accepts context
- Look for context propagation in SSE transport
- Review library documentation/source code

### 2. Identify Workaround Options

If library doesn't support context:

**Option A: Store context in request metadata**

- Use HTTP headers or query parameters to pass API key to internal SSE server
- Internal SSE server retrieves and injects into context before calling handlers

**Option B: Custom middleware in mcp-go**

- Fork or extend mcp-go to support context in tool handlers
- Modify `AddTool` signature to accept context

**Option C: Global context storage**

- Store context in a thread-safe map keyed by request ID
- Tool handlers retrieve context from map
- Clean up after request completes

**Option D: Fix the tool registration**

- Check if the SSE server actually provides context to handlers
- The proxy correctly passes context to the internal SSE server
- Need to verify if internal SSE server passes it to tool handlers

## Implementation

### Primary Approach: Verify SSE Context Flow

1. **[`internal/mcp/http_server.go`](internal/mcp/http_server.go)** - Add debugging to confirm context is passed
2. **Check mcp-go SSEServer** - Does it provide context to tool handlers?
3. **Update tool registrations** - If context is available, use it instead of `context.Background()`

### Files to Modify

- [`internal/domains/infrahub/cloudresource/register.go`](internal/domains/infrahub/cloudresource/register.go) - All 8 tool registrations
- [`internal/domains/resourcemanager/environment/register.go`](internal/domains/resourcemanager/environment/register.go) - 1 tool registration
- Potentially [`internal/mcp/http_server.go`](internal/mcp/http_server.go) - If workaround needed

## Testing

1. Set `PLANTON_MCP_TRANSPORT=http` and `PLANTON_MCP_HTTP_AUTH_ENABLED=true`
2. Start server without `PLANTON_API_KEY` in environment
3. Make request with valid `Authorization: Bearer <api-key>` header
4. Should succeed if fix works

## Questions to Answer

1. Does the `mcp-go` library's SSE server pass request context to tool handlers?
2. If not, what's the best workaround that maintains security and multi-user support?
3. Should we consider contributing a fix upstream to `mcp-go`?