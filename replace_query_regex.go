// Package plugindemo a demo plugin.
package traefik_plugin_replace_query_regex

import (
	"context"
	"fmt"
	"net/http"
	"regexp"
	"strings"
)

// Config holds the plugin configuration.
type Config struct {
	Regex       string `json:"regex,omitempty"`
	Replacement string `json:"replacement,omitempty"`
}

// CreateConfig creates the default plugin configuration.
func CreateConfig() *Config {
	return &Config{}
}

type ReplaceQueryRegex struct {
	next        http.Handler
	regexp      *regexp.Regexp
	replacement string
}

func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	if len(config.Regex) == 0 {
		return nil, fmt.Errorf("regex cannot be empty")
	}

	if len(config.Replacement) == 0 {
		return nil, fmt.Errorf("replacement cannot be empty")
	}

	exp, err := regexp.Compile(strings.TrimSpace(config.Regex))
	if err != nil {
		return nil, fmt.Errorf("error compiling regular expression %s: %w", config.Regex, err)
	}

	return &ReplaceQueryRegex{
		next:        next,
		regexp:      exp,
		replacement: config.Replacement,
	}, nil
}

func (a *ReplaceQueryRegex) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	req.URL.RawQuery = a.regexp.ReplaceAllString(req.URL.RawQuery, a.replacement)
	a.next.ServeHTTP(rw, req)
}
