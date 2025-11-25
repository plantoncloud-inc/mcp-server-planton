package config

import (
	"fmt"
	"os"
)

// Config holds the MCP server configuration loaded from environment variables.
//
// Unlike agent-fleet-worker (which uses machine account), this server
// expects PLANTON_API_KEY to be passed via environment by LangGraph or other MCP clients.
type Config struct {
	// PlantonAPIKey is the user's API key for authentication with Planton Cloud APIs.
	// This can be either a JWT token or an API key from the Planton Cloud console.
	// This is passed by LangGraph via environment when spawning the MCP server.
	PlantonAPIKey string

	// PlantonAPIsGRPCEndpoint is the gRPC endpoint for Planton Cloud APIs.
	// Defaults to "localhost:8080" if not set.
	PlantonAPIsGRPCEndpoint string
}

// LoadFromEnv loads configuration from environment variables.
//
// Required environment variables:
//   - PLANTON_API_KEY: User's API key for authentication (can be JWT token or API key)
//
// Optional environment variables:
//   - PLANTON_APIS_GRPC_ENDPOINT: Planton Cloud APIs gRPC endpoint (default: localhost:8080)
//
// Returns an error if PLANTON_API_KEY is missing.
func LoadFromEnv() (*Config, error) {
	apiKey := os.Getenv("PLANTON_API_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf(
			"PLANTON_API_KEY environment variable required. " +
				"This should be set by LangGraph when spawning MCP server",
		)
	}

	endpoint := os.Getenv("PLANTON_APIS_GRPC_ENDPOINT")
	if endpoint == "" {
		endpoint = "localhost:8080"
	}

	return &Config{
		PlantonAPIKey:           apiKey,
		PlantonAPIsGRPCEndpoint: endpoint,
	}, nil
}
