package config

import (
	_ "embed"
)

const schemaURL = "config-schema.json"

//go:embed config-schema.json
var schemaData string

func validateSchema(cfgFile string) error {
	_ = "STUB: not implemented"
	// ignore if file not exists
	return nil
}

// Config file does not exist or is not readable - ignore it
//nolint:nilerr

// Load YAML file

func validateYAML(yamlContent []byte) error { _ = "STUB: not implemented"; return nil }

// Convert YAML to JSON

// Load JSON schema

// validateSchema
