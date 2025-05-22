package main

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "aic",
	Short: "AI Commit - Generate commit messages using AI",
	Long:  `AI Commit is a CLI tool that generates commit messages using AI (OpenAI or DeepSeek) based on git diff.`,
}

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Configure AI Commit settings",
}

var setOpenAICmd = &cobra.Command{
	Use:   "set-openai [api-key]",
	Short: "Set OpenAI API key and settings",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		config, err := LoadConfig()
		if err != nil {
			return fmt.Errorf("failed to load config: %w", err)
		}

		config.AIModel = OpenAIModel
		config.OpenAI.APIKey = args[0]

		model, _ := cmd.Flags().GetString("model")
		if model != "" {
			config.OpenAI.Model = model
		}

		temp, _ := cmd.Flags().GetFloat32("temperature")
		if temp > 0 {
			config.OpenAI.Temperature = temp
		}

		if err := SaveConfig(config); err != nil {
			return fmt.Errorf("failed to save config: %w", err)
		}

		fmt.Println("✅ OpenAI configuration saved successfully")
		return nil
	},
}

var setDeepSeekCmd = &cobra.Command{
	Use:   "set-deepseek [api-key]",
	Short: "Set DeepSeek API key and settings",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		config, err := LoadConfig()
		if err != nil {
			return fmt.Errorf("failed to load config: %w", err)
		}

		config.AIModel = DeepSeekModel
		config.DeepSeek.APIKey = args[0]

		model, _ := cmd.Flags().GetString("model")
		if model != "" {
			config.DeepSeek.Model = model
		}

		temp, _ := cmd.Flags().GetFloat32("temperature")
		if temp > 0 {
			config.DeepSeek.Temperature = temp
		}

		baseURL, _ := cmd.Flags().GetString("base-url")
		if baseURL != "" {
			config.DeepSeek.BaseURL = baseURL
		}

		if err := SaveConfig(config); err != nil {
			return fmt.Errorf("failed to save config: %w", err)
		}

		fmt.Println("✅ DeepSeek configuration saved successfully")
		return nil
	},
}

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate commit message from staged changes",
	RunE: func(cmd *cobra.Command, args []string) error {
		// Load config
		config, err := LoadConfig()
		if err != nil {
			return fmt.Errorf("failed to load config: %w", err)
		}

		// Get git diff
		diff, err := getGitDiff()
		if err != nil {
			return fmt.Errorf("error getting git diff: %w", err)
		}

		// Generate commit message
		var message string
		switch config.AIModel {
		case DeepSeekModel:
			message, err = generateWithDeepSeek(diff, config)
		default: // OpenAI
			message, err = generateWithOpenAI(diff, config)
		}

		if err != nil {
			return fmt.Errorf("error generating commit message: %w", err)
		}

		// Copy to clipboard
		if err := copyToClipboard(message); err != nil {
			return fmt.Errorf("error copying to clipboard: %w", err)
		}

		fmt.Printf("✅ Generated commit message (copied to clipboard):\n\n%s\n", message)
		return nil
	},
}

func init() {
	// Add commands
	rootCmd.AddCommand(generateCmd)
	rootCmd.AddCommand(configCmd)
	configCmd.AddCommand(setOpenAICmd)
	configCmd.AddCommand(setDeepSeekCmd)

	// Set OpenAI flags
	setOpenAICmd.Flags().String("model", "", "OpenAI model (default: gpt-3.5-turbo)")
	setOpenAICmd.Flags().Float32("temperature", 0.7, "Temperature for the model (0.0 to 2.0)")

	// Set DeepSeek flags
	setDeepSeekCmd.Flags().String("model", "", "DeepSeek model (default: deepseek-chat)")
	setDeepSeekCmd.Flags().Float32("temperature", 0.7, "Temperature for the model (0.0 to 2.0)")
	setDeepSeekCmd.Flags().String("base-url", "", "Base URL for DeepSeek API (default: https://api.deepseek.com/v1)")

	// Set generate as default command
	rootCmd.RunE = generateCmd.RunE
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("Error: %v", err)
	}
}

func main() {
	Execute()
}
