package consts

import "os"

// JwtKey is JWT signing/verification key.
// Security note: do NOT hardcode real secrets in the repository.
// For local/dev use set WENCI_JWT_KEY; otherwise a placeholder key will be used.
var JwtKey = getJwtKey()

func getJwtKey() string {
	if v := os.Getenv("WENCI_JWT_KEY"); v != "" {
		return v
	}
	// Placeholder only for local development.
	return "dev-change-me"
}
