package querier

import (
	"flag"
	"time"
)

// Config for a querier.
type Config struct {
	QueryTimeout    time.Duration `yaml:"query_timeout"`
	ExtraQueryDelay time.Duration `yaml:"extra_query_delay,omitempty"`
}

// RegisterFlagsAndApplyDefaults register flags.
func (cfg *Config) RegisterFlagsAndApplyDefaults(prefix string, f *flag.FlagSet) {
	cfg.QueryTimeout = 10 * time.Second
	cfg.ExtraQueryDelay = 0
}
