//go:build tools
// +build tools

package tools

import (
	_ "github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen" // Tool import for 'go:generate' directives but not include in the final build
)
