# AI Commit

A CLI tool that generates commit messages using AI (OpenAI or DeepSeek) based on git diff.

## Features

- 🤖 Supports both OpenAI and DeepSeek APIs
- ⚙️ Simple configuration management
- 📋 Copies the generated message to clipboard
- 🔒 Secure - Your API keys are stored locally

## Installation

### Using Go (recommended)

```bash
go install github.com/hoangneee/ai-commit@latest
```

### Manual Build with Alias (recommended for development)

1. Clone the repository:

   ```bash
   git clone https://github.com/hoangneee/ai-commit.git
   cd ai-commit
   ```

2. Make the build script executable and run it:

   ```bash
   chmod +x build.sh
   ./build.sh
   ```

3. Set up the alias (choose one method):

   For temporary use (current terminal session only):

   ```bash
   alias aic='$(pwd)/aic'
   ```

   For permanent use (add to your shell config):

   ```bash
   echo "alias aic='$(pwd)/aic'" >> ~/.zshrc  # or ~/.bashrc
   source ~/.zshrc  # or ~/.bashrc
   ```

4. Verify the installation:
   ```bash
   aic --help
   ```

### Manual Installation (legacy)

```bash
git clone https://github.com/hoangneee/ai-commit.git
cd ai-commit
go build -o aic .
sudo mv aic /usr/local/bin/
```

## Configuration

### Configure OpenAI

```bash
# Set OpenAI API key (required)
aic config set-openai your-openai-api-key

# Optional: Set model and temperature
aic config set-openai your-openai-api-key --model gpt-4 --temperature 0.7
```

### Configure DeepSeek

```bash
# Set DeepSeek API key (required)
aic config set-deepseek your-deepseek-api-key

# Optional: Set model, temperature, and base URL
aic config set-deepseek your-deepseek-api-key --model deepseek-chat --temperature 0.7 --base-url https://api.deepseek.com/v1
```

### Configure Google AI

```bash
# Set Google AI API key (required)
aic config set-googleai your-googleai-api-key

# Optional: Set model and temperature
aic config set-googleai your-googleai-api-key --model gemini-pro --temperature 0.7
```

## Usage

1. Stage your changes:

   ```bash
   git add .
   ```

2. Generate a commit message:

   ```bash
   aic generate
   # or simply
   aic
   ```

   The generated commit message will be copied to your clipboard and displayed in the terminal.

3. Paste and commit:
   ```bash
   git commit -m "$(pbpaste)"  # On macOS
   # or
   git commit -m "$(xclip -o)"  # On Linux with xclip
   # or just paste manually
   ```

## Configuration File

The configuration is stored in `~/.aicommit.yaml`. Here's an example:

```yaml
ai_model: openai # or "deepseek"
openai:
  api_key: your-openai-api-key
  model: gpt-3.5-turbo
  temperature: 0.7
deepseek:
  api_key: your-deepseek-api-key
  model: deepseek-chat
  temperature: 0.7
  base_url: https://api.deepseek.com/v1
googleai:
  api_key: your-googleai-api-key
  model: gemini-pro
  temperature: 0.7
```

## License

[License](LICENSE)
