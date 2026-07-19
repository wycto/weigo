package helper

import (
	"testing"

	"github.com/wycto/weigo/config"
)

func TestMD5(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"", "d41d8cd98f00b204e9800998ecf8427e"},
		{"hello", "5d41402abc4b2a76b9719d911017c592"},
		{"test123", "cc03e747a6afbbcbf8be7668acfebee5"},
	}
	for _, tt := range tests {
		result := MD5(tt.input)
		if result != tt.expected {
			t.Errorf("MD5(%q) = %q, want %q", tt.input, result, tt.expected)
		}
	}
}

func TestSHA256(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"", "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"},
		{"hello", "2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824"},
	}
	for _, tt := range tests {
		result := SHA256(tt.input)
		if result != tt.expected {
			t.Errorf("SHA256(%q) = %q, want %q", tt.input, result, tt.expected)
		}
	}
}

func TestIsOriginAllowed(t *testing.T) {
	// nil config
	config.GlobalConfig = nil
	if IsOriginAllowed("http://example.com") {
		t.Error("nil config should reject all origins")
	}

	// Debug mode, no AllowOrigins
	config.GlobalConfig = &config.Config{
		App: config.AppConfig{Debug: true},
	}
	if !IsOriginAllowed("http://any.com") {
		t.Error("debug mode with no AllowOrigins should allow all")
	}

	// Non-debug, no AllowOrigins
	config.GlobalConfig = &config.Config{
		App: config.AppConfig{Debug: false},
	}
	if IsOriginAllowed("http://example.com") {
		t.Error("non-debug with no AllowOrigins should reject all")
	}

	// Configured whitelist
	config.GlobalConfig = &config.Config{
		App: config.AppConfig{
			AllowOrigins: []string{"http://a.com", "http://b.com"},
		},
	}
	if !IsOriginAllowed("http://a.com") {
		t.Error("should allow listed origin a.com")
	}
	if !IsOriginAllowed("http://b.com") {
		t.Error("should allow listed origin b.com")
	}
	if IsOriginAllowed("http://c.com") {
		t.Error("should reject unlisted origin c.com")
	}

	// Wildcard
	config.GlobalConfig = &config.Config{
		App: config.AppConfig{
			AllowOrigins: []string{"*"},
		},
	}
	if !IsOriginAllowed("http://anything.com") {
		t.Error("wildcard should allow any origin")
	}

	// Debug + configured whitelist (whitelist should restrict)
	config.GlobalConfig = &config.Config{
		App: config.AppConfig{
			Debug:        true,
			AllowOrigins: []string{"http://only.com"},
		},
	}
	if !IsOriginAllowed("http://only.com") {
		t.Error("should allow listed origin")
	}
	if IsOriginAllowed("http://other.com") {
		t.Error("debug + whitelist should still restrict to listed origins")
	}
}
