# PonyClaw Features Documentation

This document provides a comprehensive overview of all features available in PonyClaw.

## Table of Contents

- [Core Features](#core-features)
- [AI Agent Capabilities](#ai-agent-capabilities)
- [Communication Channels](#communication-channels)
- [LLM Provider Support](#llm-provider-support)
- [Built-in Tools](#built-in-tools)
- [Advanced Features](#advanced-features)
- [Configuration & Management](#configuration--management)

---

## Core Features

### 1. Lightweight Architecture
- **Go-based Implementation**: Compiled binary with minimal dependencies
- **Low Memory Footprint**: Optimized for resource-constrained environments
- **Fast Startup**: Quick initialization and execution
- **Cross-Platform**: Supports x86_64, ARM64, RISC-V architectures

### 2. Multi-Agent System
- **Agent Registry**: Manage multiple AI agents with different configurations
- **Context Management**: Maintain conversation context and history
- **Session Management**: Handle multiple concurrent conversations
- **Workspace Isolation**: Each agent operates in its own workspace

### 3. Security Sandbox
- **Workspace Restriction**: Limit file and command access to configured workspace
- **Command Filtering**: Block dangerous system commands
- **Path Validation**: Prevent access to sensitive system directories
- **Configurable Security**: Enable/disable restrictions as needed

---

## AI Agent Capabilities

### Conversation Management
- **Multi-turn Conversations**: Maintain context across multiple exchanges
- **Session Persistence**: Save and restore conversation history
- **Memory System**: Long-term memory storage in MEMORY.md
- **Context Caching**: Optimize LLM API calls with context reuse

### Agent Customization
- **Identity Configuration**: Define agent personality via IDENTITY.md
- **Soul/Personality**: Customize agent behavior via SOUL.md
- **User Preferences**: Store user-specific settings in USER.md
- **Agent Behavior Guide**: Configure agent actions via AGENTS.md

### Autonomous Operations
- **Tool Execution**: Automatically use tools to accomplish tasks
- **Subagent Spawning**: Create child agents for parallel task execution
- **Async Task Handling**: Non-blocking execution for long-running operations
- **Error Recovery**: Automatic retry and fallback mechanisms

---

## Communication Channels

PonyClaw supports multiple chat platforms for user interaction:

### Instant Messaging
- **Telegram**: Full support with voice transcription, inline keyboards
- **Discord**: Socket mode, supports DMs and server channels
- **Slack**: Socket mode, enterprise-ready
- **WhatsApp**: Via bridge integration
- **LINE**: Webhook-based integration

### Enterprise Platforms
- **QQ**: Official bot API support
- **DingTalk (钉钉)**: Stream mode, no public IP required
- **Feishu (飞书)**: Enterprise collaboration features
- **WeCom (企业微信)**: 
  - Bot mode (group chats via webhook)
  - App mode (private chats with proactive messaging)

### Community Protocols
- **OneBot**: Compatible with NapCat, Go-CQHTTP implementations

### Channel Features
- **User Allowlists**: Restrict access to specific users
- **Message Routing**: Intelligent message distribution
- **Media Support**: Handle images, voice messages, files
- **Async Messaging**: Non-blocking message delivery

---

## LLM Provider Support

### Supported Providers

PonyClaw uses a model-centric configuration supporting multiple LLM providers:

#### Commercial Providers
- **OpenAI**: GPT-4, GPT-3.5, and other models
- **Anthropic**: Claude 3 family (Opus, Sonnet, Haiku)
- **Google Gemini**: Gemini Pro and Flash models
- **DeepSeek**: DeepSeek Chat and Coder models
- **Groq**: Fast inference with Llama, Mixtral models
- **Moonshot**: Moonshot AI models
- **Cerebras**: Ultra-fast inference

#### Chinese Providers
- **Zhipu AI (智谱)**: GLM-4 series models
- **Qwen (通义千问)**: Alibaba's Qwen models
- **Volcengine (火山引擎)**: ByteDance's LLM service
- **Shengsuanyun (神算云)**: Domestic LLM aggregator

#### Aggregators & Proxies
- **OpenRouter**: Access to 100+ models via single API
- **NVIDIA**: NVIDIA AI Foundation models

#### Local & Self-Hosted
- **Ollama**: Run models locally
- **VLLM**: Self-hosted model serving
- **GitHub Copilot**: Local Copilot integration

#### Special Integrations
- **Antigravity**: OAuth-based Google Cloud integration
- **Claude CLI**: Direct Claude desktop app integration
- **Codex CLI**: OpenAI Codex integration

### Provider Features
- **Load Balancing**: Distribute requests across multiple endpoints
- **Fallback Support**: Automatic failover to backup providers
- **OAuth Authentication**: Secure authentication for supported providers
- **Custom API Bases**: Use proxy or custom endpoints
- **Request Timeout**: Configurable timeout per model
- **Cooldown Management**: Rate limiting and retry logic

---

## Built-in Tools

PonyClaw provides a comprehensive set of tools for the AI agent:

### File System Operations
- **read_file**: Read file contents
- **write_file**: Create or overwrite files
- **edit_file**: Modify existing files
- **append_file**: Append content to files
- **list_dir**: List directory contents

### Command Execution
- **exec**: Execute shell commands
  - Timeout support
  - Working directory control
  - Environment variable management
  - Safety guards against dangerous commands

### Web & Search
- **web_search**: Search the internet
  - Brave Search API
  - DuckDuckGo (no API key required)
  - Perplexity API
  - Tavily API (AI-optimized)
- **web_fetch**: Fetch content from URLs
  - HTML parsing
  - Content extraction
  - Media download support

### Task Management
- **cron**: Schedule tasks and reminders
  - One-time reminders
  - Recurring tasks
  - Cron expression support
  - Task persistence

### Agent Operations
- **spawn**: Create child agents for async tasks
  - Independent context
  - Parallel execution
  - Result callback support
- **subagent**: Delegate tasks to specialized agents
  - Task isolation
  - Resource management
- **message**: Send messages to users
  - Multi-channel support
  - Async delivery

### Skills & Extensions
- **find_skills**: Search for available skills
  - ClawHub registry integration
  - Local skill discovery
- **install_skill**: Install new skills
  - Automatic dependency resolution
  - Skill validation

### Hardware Integration (Linux)
- **i2c**: I2C device communication
  - Read/write operations
  - Device scanning
- **spi**: SPI device communication
  - Data transfer
  - Device control

### Tool Features
- **Async Execution**: Non-blocking tool calls
- **Error Handling**: Graceful error recovery
- **Result Formatting**: Structured output for LLM consumption
- **Context Awareness**: Tools can access channel/chat context
- **Extensibility**: Easy to add custom tools

---

## Advanced Features

### 1. Heartbeat System
Periodic task execution without user intervention:
- **Configurable Interval**: Default 30 minutes, minimum 5 minutes
- **Task Definition**: Define tasks in HEARTBEAT.md
- **Async Task Support**: Use spawn for long-running tasks
- **Tool Access**: Full access to all agent tools
- **Enable/Disable**: Toggle via configuration

### 2. Skills System
Extend agent capabilities with custom skills:
- **Skill Registry**: Discover skills from ClawHub and local sources
- **Skill Installation**: Automatic download and setup
- **Skill Loading**: Dynamic skill loading at runtime
- **Skill Isolation**: Each skill runs in isolated context
- **Skill Search**: Find skills by keyword or category

### 3. Voice Transcription
Convert voice messages to text:
- **Groq Whisper**: Free voice transcription via Groq API
- **Telegram Integration**: Automatic voice message transcription
- **Multi-language Support**: Supports multiple languages

### 4. OAuth Authentication
Secure authentication for supported providers:
- **PKCE Flow**: Secure OAuth 2.0 with PKCE
- **Token Management**: Automatic token refresh
- **Provider Support**: Anthropic, Google Gemini
- **Credential Storage**: Secure local storage

### 5. Device Monitoring
Monitor and respond to hardware events:
- **USB Device Detection**: Detect USB device connections
- **Event Triggers**: Execute actions on device events
- **Device Information**: Query device details

### 6. Message Bus
Internal event system for component communication:
- **Inbound Messages**: User messages from channels
- **Outbound Messages**: Agent responses to channels
- **Event Broadcasting**: System-wide event distribution
- **Async Processing**: Non-blocking message handling

### 7. Migration System
Automatic configuration and workspace migration:
- **Config Migration**: Update old config formats
- **Workspace Migration**: Upgrade workspace structure
- **Backward Compatibility**: Support legacy configurations

---

## Configuration & Management

### Configuration Files
- **config.json**: Main configuration file
- **Workspace Files**: Agent-specific configurations
  - IDENTITY.md: Agent identity
  - SOUL.md: Agent personality
  - USER.md: User preferences
  - AGENTS.md: Agent behavior
  - HEARTBEAT.md: Periodic tasks
  - MEMORY.md: Long-term memory
  - TOOLS.md: Tool descriptions

### CLI Commands
- **onboard**: Initialize configuration and workspace
- **agent**: Interactive agent mode
- **gateway**: Start multi-channel gateway
- **status**: Show system status
- **auth**: Manage OAuth authentication
- **cron**: Manage scheduled tasks
- **skills**: Manage skills
- **migrate**: Run migrations

### Environment Variables
Override configuration via environment variables:
- `PONYCLAW_*`: All config options can be set via env vars
- `PONYCLAW_HEARTBEAT_ENABLED`: Enable/disable heartbeat
- `PONYCLAW_HEARTBEAT_INTERVAL`: Set heartbeat interval
- `PONYCLAW_GATEWAY_HOST`: Gateway listen address
- `PONYCLAW_GATEWAY_PORT`: Gateway listen port

### Logging
- **Structured Logging**: JSON-formatted logs
- **Log Levels**: Debug, Info, Warn, Error
- **Context Fields**: Rich contextual information
- **Component Tagging**: Identify log sources

### Health Monitoring
- **Health Endpoint**: HTTP health check endpoint
- **Status API**: Query system status
- **Channel Status**: Monitor channel health
- **Agent Status**: Track agent activity

---

## Docker Support

### Docker Compose
- **Gateway Mode**: Run as multi-channel gateway
- **Agent Mode**: One-shot agent execution
- **Volume Mounting**: Persist configuration and data
- **Network Configuration**: Expose ports as needed

### Container Features
- **Multi-architecture**: Support for amd64, arm64
- **Minimal Image**: Small container size
- **Health Checks**: Built-in health monitoring
- **Environment Config**: Configure via environment variables

---

## Performance Features

### Optimization
- **Context Caching**: Reduce LLM API costs
- **Connection Pooling**: Reuse HTTP connections
- **Concurrent Processing**: Handle multiple requests simultaneously
- **Memory Management**: Efficient memory usage

### Reliability
- **Automatic Retry**: Retry failed operations
- **Error Recovery**: Graceful error handling
- **Fallback Mechanisms**: Backup providers and strategies
- **Rate Limiting**: Respect API rate limits

### Monitoring
- **Execution Timing**: Track tool execution time
- **Token Usage**: Monitor LLM token consumption
- **Error Tracking**: Log and track errors
- **Performance Metrics**: System performance data

---

## Security Features

### Access Control
- **User Allowlists**: Restrict access per channel
- **Workspace Isolation**: Sandbox file system access
- **Command Filtering**: Block dangerous commands
- **Path Validation**: Prevent directory traversal

### Data Protection
- **Credential Storage**: Secure API key storage
- **Session Encryption**: Protect session data
- **Secure Communication**: HTTPS/WSS for channels
- **Token Management**: Secure OAuth token handling

### Safety Guards
- **Command Blacklist**: Block destructive commands
- **Path Restrictions**: Limit file system access
- **Resource Limits**: Prevent resource exhaustion
- **Timeout Protection**: Prevent infinite loops

---

## Extensibility

### Custom Tools
- **Tool Interface**: Simple tool creation API
- **Tool Registration**: Dynamic tool registration
- **Tool Discovery**: Automatic tool detection
- **Tool Documentation**: Self-documenting tools

### Custom Channels
- **Channel Interface**: Standard channel API
- **Channel Registration**: Dynamic channel registration
- **Message Routing**: Automatic message routing
- **Event Handling**: Standard event handling

### Custom Providers
- **Provider Interface**: Standard provider API
- **Provider Factory**: Dynamic provider creation
- **Protocol Support**: Multiple protocol support
- **Error Handling**: Standard error handling

### Plugin System
- **Skill Plugins**: Extend with custom skills
- **Tool Plugins**: Add custom tools
- **Provider Plugins**: Add custom LLM providers
- **Channel Plugins**: Add custom communication channels

---

## Future Features (Roadmap)

See [ROADMAP.md](ROADMAP.md) for planned features and development priorities.

---

## Getting Help

- **Documentation**: Check README.md and docs/ directory
- **Issues**: Report bugs on GitHub Issues
- **Discussions**: Ask questions on GitHub Discussions
- **Contributing**: See CONTRIBUTING.md for contribution guidelines
