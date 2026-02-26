<div align="center">
<img src="assets/logo.png" alt="PonyClaw" width="512">

<h1>PonyClaw: Lightweight AI Assistant in Go</h1>

<h3>Low Resource Usage · Fast Startup · Cross-Platform</h3>

  <p>
    <img src="https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go&logoColor=white" alt="Go">
    <img src="https://img.shields.io/badge/Arch-x86__64%2C%20ARM64%2C%20RISC--V-blue" alt="Hardware">
    <img src="https://img.shields.io/badge/license-MIT-green" alt="License">
    <br>
    <a href="https://github.com/lansely/ponyclaw"><img src="https://img.shields.io/badge/GitHub-lansely/ponyclaw-blue?style=flat&logo=github&logoColor=white" alt="GitHub"></a>
  </p>

[中文](README.zh.md) | **English**

</div>

---

PonyClaw is a lightweight personal AI assistant built with Go, focusing on low resource usage and high performance.

> [!NOTE]
> PonyClaw is under active development. Features and APIs may change.

## ✨ Features

🪶 **Lightweight**: Built with Go for low memory footprint and minimal resource consumption

⚡️ **Fast Startup**: Compiled language provides quick startup and execution speed

🌍 **Cross-Platform**: Supports x86_64, ARM64, RISC-V and more architectures

🔧 **Easy Deployment**: Single binary file, no complex dependencies

🤖 **Multi-Model Support**: Compatible with various LLM providers (OpenAI, Anthropic, Zhipu, etc.)

💬 **Multi-Channel Integration**: Supports Telegram, Discord, Slack, QQ and more chat platforms

## 🦾 Demonstration

### 🛠️ Standard Assistant Workflows

<table align="center">
  <tr align="center">
    <th><p align="center">🧩 Full-Stack Engineer</p></th>
    <th><p align="center">🗂️ Logging & Planning</p></th>
    <th><p align="center">🔎 Web Search & Learning</p></th>
  </tr>
  <tr>
    <td align="center"><p align="center"><img src="assets/picoclaw_code.gif" width="240" height="180"></p></td>
    <td align="center"><p align="center"><img src="assets/picoclaw_memory.gif" width="240" height="180"></p></td>
    <td align="center"><p align="center"><img src="assets/picoclaw_search.gif" width="240" height="180"></p></td>
  </tr>
  <tr>
    <td align="center">Develop • Deploy • Scale</td>
    <td align="center">Schedule • Automate • Memory</td>
    <td align="center">Discovery • Insights • Trends</td>
  </tr>
</table>

### 📱 Run on Mobile Devices

PonyClaw can run on Android devices via Termux:

1. Install Termux from app store
2. Execute in Termux:

```bash
# Download latest version (adjust version number as needed)
wget https://github.com/lansely/ponyclaw/releases/latest/download/ponyclaw-linux-arm64
chmod +x ponyclaw-linux-arm64
pkg install proot
termux-chroot ./ponyclaw-linux-arm64 onboard
```

Then follow the "Quick Start" section to complete configuration.

## 📦 Installation

### Install with Precompiled Binary

Download the binary for your platform from the [Release page](https://github.com/lansely/ponyclaw/releases).

### Install from Source

```bash
git clone https://github.com/lansely/ponyclaw.git
cd ponyclaw
make deps

# Build
make build

# Build and install
make install
```

## 🐳 Docker Compose

You can also run PonyClaw using Docker Compose without installing anything locally.

```bash
# 1. Clone repository
git clone https://github.com/lansely/ponyclaw.git
cd ponyclaw

# 2. Set API keys
cp config/config.example.json config/config.json
vim config/config.json      # Set API keys, etc.

# 3. Build & Start
docker compose --profile gateway up -d

# 4. Check logs
docker compose logs -f ponyclaw-gateway

# 5. Stop
docker compose --profile gateway down
```

> [!TIP]
> **Docker Users**: By default, Gateway listens on `127.0.0.1`. To access via port mapping, set `PONYCLAW_GATEWAY_HOST=0.0.0.0` in environment or update `config.json`.

### Agent Mode (One-shot)

```bash
# Ask a question
docker compose run --rm ponyclaw-agent -m "What is 2+2?"

# Interactive mode
docker compose run --rm ponyclaw-agent
```

## 🚀 Quick Start

> [!TIP]
> Set your API key in `~/.ponyclaw/config.json`.
> Get API keys: [OpenRouter](https://openrouter.ai/keys) · [Zhipu](https://open.bigmodel.cn/usercenter/proj-mgmt/apikeys)
> Web search is optional - get free [Tavily API](https://tavily.com) (1000/month) or [Brave Search API](https://brave.com/search/api) (2000/month)

**1. Initialize**

```bash
ponyclaw onboard
```

**2. Configure** (`~/.ponyclaw/config.json`)

```json
{
  "agents": {
    "defaults": {
      "workspace": "~/.ponyclaw/workspace",
      "model_name": "gpt4",
      "max_tokens": 8192,
      "temperature": 0.7,
      "max_tool_iterations": 20
    }
  },
  "model_list": [
    {
      "model_name": "gpt4",
      "model": "openai/gpt-4",
      "api_key": "your-api-key",
      "request_timeout": 300
    }
  ],
  "tools": {
    "web": {
      "tavily": {
        "enabled": false,
        "api_key": "YOUR_TAVILY_API_KEY",
        "max_results": 5
      }
    }
  }
}
```

**3. Chat**

```bash
ponyclaw agent -m "Hello"
```

That's it! You have a working AI assistant in minutes.

---

## 💬 Chat App Integration

PonyClaw supports multiple chat platforms:

| Channel              | Setup      | Features                                    |
| -------------------- | ---------- | ------------------------------------------- |
| **Telegram**         | ⭐ Easy    | Recommended, voice-to-text, no public IP    |
| **Discord**          | ⭐ Easy    | Socket Mode, groups/DMs, mature bot system  |
| **Slack**            | ⭐ Easy    | Socket Mode (no public IP), enterprise      |
| **QQ**               | ⭐⭐ Medium | Official bot API, good for Chinese users    |
| **DingTalk**         | ⭐⭐ Medium | Stream mode, no public IP, enterprise       |
| **WeCom**            | ⭐⭐⭐ Hard | Group bot (Webhook) and app (API)           |
| **Feishu**           | ⭐⭐⭐ Hard | Enterprise collaboration, feature-rich      |
| **Line**             | ⭐⭐⭐ Hard | Requires HTTPS webhook                      |
| **OneBot**           | ⭐⭐ Medium | Compatible with NapCat/Go-CQHTTP, community |

## ⚙️ Configuration

Config file: `~/.ponyclaw/config.json`

### Workspace Layout

PonyClaw stores data in your configured workspace (default: `~/.ponyclaw/workspace`):

```
~/.ponyclaw/workspace/
├── sessions/          # Conversation sessions and history
├── memory/           # Long-term memory
├── state/            # Persistent state
├── cron/             # Scheduled jobs database
├── skills/           # Custom skills
├── AGENTS.md         # Agent behavior guide
├── HEARTBEAT.md      # Periodic task prompts
├── IDENTITY.md       # Agent identity
├── SOUL.md           # Agent personality
├── TOOLS.md          # Tool descriptions
└── USER.md           # User preferences
```

### Heartbeat / Periodic Tasks

PonyClaw can perform periodic tasks automatically. Create a `HEARTBEAT.md` file in your workspace:

```markdown
# Periodic Tasks

- Check my email for important messages
- Review my calendar for upcoming events
```

The agent will read this file every 30 minutes (configurable) and execute tasks.

### Model Configuration

PonyClaw uses a model-centric configuration approach, supporting various LLM providers:

| Vendor              | `model` Prefix | Get API Key                                                      |
| ------------------- | -------------- | ---------------------------------------------------------------- |
| **OpenAI**          | `openai/`      | [Get Key](https://platform.openai.com)                           |
| **Anthropic**       | `anthropic/`   | [Get Key](https://console.anthropic.com)                         |
| **智谱 AI (GLM)**   | `zhipu/`       | [Get Key](https://open.bigmodel.cn/usercenter/proj-mgmt/apikeys) |
| **DeepSeek**        | `deepseek/`    | [Get Key](https://platform.deepseek.com)                         |
| **Google Gemini**   | `gemini/`      | [Get Key](https://aistudio.google.com/api-keys)                  |
| **Groq**            | `groq/`        | [Get Key](https://console.groq.com)                              |
| **通义千问 (Qwen)** | `qwen/`        | [Get Key](https://dashscope.console.aliyun.com)                  |
| **Ollama**          | `ollama/`      | Local (no key needed)                                            |
| **OpenRouter**      | `openrouter/`  | [Get Key](https://openrouter.ai/keys)                            |

Configuration example:

```json
{
  "model_list": [
    {
      "model_name": "gpt-4",
      "model": "openai/gpt-4",
      "api_key": "sk-..."
    },
    {
      "model_name": "glm-4",
      "model": "zhipu/glm-4",
      "api_key": "your-key"
    }
  ]
}
```

## CLI Command Reference

| Command                    | Description                  |
| -------------------------- | ---------------------------- |
| `ponyclaw onboard`         | Initialize config/workspace  |
| `ponyclaw agent -m "..."` | Chat with agent              |
| `ponyclaw agent`           | Interactive chat mode        |
| `ponyclaw gateway`         | Start gateway                |
| `ponyclaw status`          | Show status                  |
| `ponyclaw cron list`       | List all scheduled tasks     |
| `ponyclaw cron add ...`    | Add scheduled task           |

### Scheduled Tasks / Reminders

PonyClaw supports scheduled reminders and recurring tasks via the `cron` tool:

- **One-time reminder**: "Remind me in 10 minutes" → triggers once after 10 minutes
- **Recurring task**: "Remind me every 2 hours" → triggers every 2 hours
- **Cron expression**: "Remind me at 9am daily" → uses cron expression

Tasks are stored in `~/.ponyclaw/workspace/cron/` and handled automatically.

## 🤝 Contributing

PRs and Issues are welcome! We strive to keep the codebase clean and readable.

## 🐛 Troubleshooting

### Web Search Shows "API Configuration Issue"

If you haven't configured a search API key, this is normal. To enable web search:

1. Get a free API key from [Tavily](https://tavily.com) or [Brave Search](https://brave.com/search/api)
2. Add to `~/.ponyclaw/config.json` in the `tools.web` configuration

### Content Filtering Errors

Some providers have strict content filtering. Try rephrasing your question or using a different model.

### Telegram Bot Shows Conflict

This indicates another bot instance is running. Ensure only one `ponyclaw gateway` process is running at a time.

## 📄 License

This project is licensed under the MIT License. See [LICENSE](LICENSE) file for details.
