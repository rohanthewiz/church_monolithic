package config

import (
	"os"
	"strings"
)

func envOverride(env_cfg *EnvConfig) *EnvConfig {
	if log_level := strings.TrimSpace(os.Getenv("LOG_LEVEL")); len(log_level) > 0 {
		env_cfg.Log.Level = log_level
	}
	if log_format := strings.TrimSpace(os.Getenv("LOG_FORMAT")); len(log_format) > 0 {
		env_cfg.Log.Format = log_format
	}
	if pgUser := strings.TrimSpace(os.Getenv("PG_USER")); len(pgUser) > 0 {
		env_cfg.PG.User = pgUser
	}
	if pgWord := strings.TrimSpace(os.Getenv("PG_WORD")); len(pgWord) > 0 {
		env_cfg.PG.Word = pgWord
	}
	return env_cfg
}
