package cli

import (
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
	"github.com/adrg/xdg"
	"github.com/urfave/cli/v2"
)

type Config struct {
	API struct {
		Endpoint string `toml:"endpoint"`
		Username string `toml:"username"`
		Password string `toml:"password"`
		Timeout  int    `toml:"timeout"`
		TestMode bool   `toml:"test_mode"`
	} `toml:"api"`
	Output struct {
		Format string `toml:"format"`
		Colors bool   `toml:"colors"`
	} `toml:"output"`
	Logging struct {
		Level  string `toml:"level"`
		Colors bool   `toml:"colors"`
	} `toml:"logging"`
}

// locateConfigFile determines the configuration file path using XDG spec.
// It respects the --config flag, then checks the XDG user config directory,
// and finally the current working directory.
func locateConfigFile(c *cli.Context) (string, error) {
	// 1. Explicit flag
	if cfg := c.String("config"); cfg != "" {
		return cfg, nil
	}

	// 2. XDG user config directory
	xdgPath := filepath.Join(xdg.ConfigHome, "inwx", "inwx.toml")
	if _, err := os.Stat(xdgPath); err == nil {
		return xdgPath, nil
	}

	// 3. Current directory fallback
	if _, err := os.Stat("./inwx.toml"); err == nil {
		return "./inwx.toml", nil
	}

	// No config file found
	return "", nil
}

func LoadConfig(c *cli.Context) (*Config, error) {
	config := &Config{}

	// Set defaults
	config.API.Timeout = 30
	config.Output.Format = "table"
	config.Output.Colors = true
	config.Logging.Level = "warn"
	config.Logging.Colors = true

	// Try to load config file using XDG logic
	configPath, err := locateConfigFile(c)
	if err != nil {
		return nil, err
	}
	if configPath != "" {
		if _, err := toml.DecodeFile(configPath, config); err != nil {
			return nil, err
		}
	}

	// Override with CLI flags and environment variables
	if c.String("username") != "" {
		config.API.Username = c.String("username")
	}
	if c.String("password") != "" {
		config.API.Password = c.String("password")
	}
	if c.Bool("test") {
		config.API.TestMode = true
	}
	if c.String("output") != "" {
		config.Output.Format = c.String("output")
	}
	if c.Bool("no-colors") {
		config.Output.Colors = false
		config.Logging.Colors = false
	}
	if c.String("log-level") != "" {
		config.Logging.Level = c.String("log-level")
	}

	return config, nil
}
