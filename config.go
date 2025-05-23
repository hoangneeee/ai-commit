package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

type AIModel string

const (
	OpenAIModel   AIModel = "openai"
	DeepSeekModel AIModel = "deepseek"
	GoogleAIModel AIModel = "googleai"
)

type OpenAIConfig struct {
	APIKey      string  `mapstructure:"apikey"`
	Model       string  `mapstructure:"model"`
	Temperature float32 `mapstructure:"temperature"`
}

type DeepSeekConfig struct {
	APIKey      string  `mapstructure:"apikey"`
	Model       string  `mapstructure:"model"`
	Temperature float32 `mapstructure:"temperature"`
	BaseURL     string  `mapstructure:"base_url"`
}

type GoogleAIConfig struct {
	APIKey      string  `mapstructure:"apikey"`
	Model       string  `mapstructure:"model"`
	Temperature float32 `mapstructure:"temperature"`
}

type Config struct {
	AIModel  AIModel        `mapstructure:"ai_model"`
	OpenAI   OpenAIConfig   `mapstructure:"openai"`
	DeepSeek DeepSeekConfig `mapstructure:"deepseek"`
	GoogleAI GoogleAIConfig `mapstructure:"googleai"`
}

func getConfigPath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("failed to get user home directory: %w", err)
	}
	return filepath.Join(homeDir, ".aicommit.yaml"), nil
}

func LoadConfig() (*Config, error) {
	configPath, err := getConfigPath()
	if err != nil {
		return nil, err
	}

	// Set default values
	viper.SetDefault("ai_model", string(OpenAIModel))
	viper.SetDefault("openai.model", "gpt-3.5-turbo")
	viper.SetDefault("openai.temperature", 0.7)
	viper.SetDefault("deepseek.model", "deepseek-chat")
	viper.SetDefault("deepseek.temperature", 0.7)
	viper.SetDefault("deepseek.base_url", "https://api.deepseek.com")
	viper.SetDefault("googleai.model", "gemma-3-27b-it")
	viper.SetDefault("googleai.temperature", 0.7)

	// Check if config file exists
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		// Create default config file if not exists
		if err := SaveConfig(&Config{
			AIModel: OpenAIModel,
			OpenAI: OpenAIConfig{
				APIKey:      "",
				Model:       "gpt-3.5-turbo",
				Temperature: 0.7,
			},
			DeepSeek: DeepSeekConfig{
				APIKey:      "",
				Model:       "deepseek-chat",
				Temperature: 0.7,
				BaseURL:     "https://api.deepseek.com",
			},
			GoogleAI: GoogleAIConfig{
				APIKey:      "",
				Model:       "googleai-model",
				Temperature: 0.7,
			},
		}); err != nil {
			return nil, fmt.Errorf("failed to create default config: %w", err)
		}
	}

	// Read config file
	viper.SetConfigFile(configPath)
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	switch config.AIModel {
	case OpenAIModel:
		return &config, nil
	case DeepSeekModel:
		return &config, nil
	case GoogleAIModel:
		return &config, nil
	default:
		return nil, fmt.Errorf("unsupported AI model: %s", config.AIModel)
	}
}

func SaveConfig(config *Config) error {
	configPath, err := getConfigPath()
	if err != nil {
		return err
	}

	// Create directory if not exists
	if err := os.MkdirAll(filepath.Dir(configPath), 0755); err != nil {
		return fmt.Errorf("failed to create config directory: %w", err)
	}

	// Set values to viper
	viper.Set("ai_model", config.AIModel)
	viper.Set("openai", config.OpenAI)
	viper.Set("deepseek", config.DeepSeek)
	viper.Set("googleai", config.GoogleAI)

	// Write config file
	if err := viper.WriteConfigAs(configPath); err != nil {
		return fmt.Errorf("failed to write config file: %w", err)
	}

	return nil
}

// GetActiveConfig trả về cấu hình đang được chọn (OpenAI hoặc DeepSeek)
func (c *Config) GetActiveConfig() (string, string, float32, string) {
	switch c.AIModel {
	case OpenAIModel:
		return c.OpenAI.APIKey, c.OpenAI.Model, c.OpenAI.Temperature, ""
	case DeepSeekModel:
		return c.DeepSeek.APIKey, c.DeepSeek.Model, c.DeepSeek.Temperature, c.DeepSeek.BaseURL
	case GoogleAIModel:
		return c.GoogleAI.APIKey, c.GoogleAI.Model, c.GoogleAI.Temperature, ""
	default:
		return "", "", 0, ""
	}
}
