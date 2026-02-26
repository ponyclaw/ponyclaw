<div align="center">
<img src="assets/logo.png" alt="PonyClaw" width="512">

<h1>PonyClaw: 基于 Go 语言的轻量级 AI 助手</h1>

<h3>低资源占用 · 快速启动 · 跨平台支持</h3>

  <p>
    <img src="https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go&logoColor=white" alt="Go">
    <img src="https://img.shields.io/badge/Arch-x86__64%2C%20ARM64%2C%20RISC--V-blue" alt="Hardware">
    <img src="https://img.shields.io/badge/license-MIT-green" alt="License">
    <br>
    <a href="https://github.com/lansely/ponyclaw"><img src="https://img.shields.io/badge/GitHub-lansely/ponyclaw-blue?style=flat&logo=github&logoColor=white" alt="GitHub"></a>
  </p>

**中文** | [English](README.md)

</div>

---

PonyClaw 是一个使用 Go 语言开发的轻量级个人 AI 助手，专注于低资源占用和高性能表现。

> [!NOTE]
> PonyClaw 目前处于活跃开发阶段，功能和 API 可能会有变化。

## ✨ 特性

🪶 **轻量级**: 使用 Go 语言开发，内存占用低，资源消耗小

⚡️ **快速启动**: 编译型语言带来的快速启动和执行速度

🌍 **跨平台**: 支持 x86_64、ARM64、RISC-V 等多种架构

🔧 **易于部署**: 单一二进制文件，无需复杂的依赖安装

🤖 **多模型支持**: 兼容多种 LLM 提供商（OpenAI、Anthropic、智谱等）

💬 **多渠道集成**: 支持 Telegram、Discord、Slack、QQ 等多个聊天平台

## 🦾 演示

### 🛠️ 标准助手工作流

<table align="center">
<tr align="center">
<th><p align="center">🧩 全栈工程师模式</p></th>
<th><p align="center">🗂️ 日志与规划管理</p></th>
<th><p align="center">🔎 网络搜索与学习</p></th>
</tr>
<tr>
<td align="center"><p align="center"><img src="assets/picoclaw_code.gif" width="240" height="180"></p></td>
<td align="center"><p align="center"><img src="assets/picoclaw_memory.gif" width="240" height="180"></p></td>
<td align="center"><p align="center"><img src="assets/picoclaw_search.gif" width="240" height="180"></p></td>
</tr>
<tr>
<td align="center">开发 • 部署 • 扩展</td>
<td align="center">日程 • 自动化 • 记忆</td>
<td align="center">发现 • 洞察 • 趋势</td>
</tr>
</table>

### 📱 在移动设备上运行

PonyClaw 可以在 Android 设备上通过 Termux 运行：

1. 从应用商店下载安装 Termux
2. 在 Termux 中执行：

```bash
# 下载最新版本（根据实际版本号调整）
wget https://github.com/lansely/ponyclaw/releases/latest/download/ponyclaw-linux-arm64
chmod +x ponyclaw-linux-arm64
pkg install proot
termux-chroot ./ponyclaw-linux-arm64 onboard
```

然后按照"快速开始"章节继续配置即可使用。

## 📦 安装

### 使用预编译二进制文件安装

从 [Release 页面](https://github.com/lansely/ponyclaw/releases) 下载适用于您平台的二进制文件。

### 从源码安装

```bash
git clone https://github.com/lansely/ponyclaw.git
cd ponyclaw
make deps

# 构建
make build

# 构建并安装
make install
```

## 🐳 Docker Compose

您也可以使用 Docker Compose 运行 PonyClaw，无需在本地安装任何环境。

```bash
# 1. 克隆仓库
git clone https://github.com/lansely/ponyclaw.git
cd ponyclaw

# 2. 设置 API Key
cp config/config.example.json config/config.json
vim config/config.json      # 设置 API keys 等

# 3. 构建并启动
docker compose --profile gateway up -d

# 4. 查看日志
docker compose logs -f ponyclaw-gateway

# 5. 停止
docker compose --profile gateway down
```

> [!TIP]
> **Docker 用户**: 默认情况下, Gateway 监听 `127.0.0.1`。如果需要通过端口映射访问，请在环境变量中设置 `PONYCLAW_GATEWAY_HOST=0.0.0.0` 或修改 `config.json`。

### Agent 模式 (一次性运行)

```bash
# 提问
docker compose run --rm ponyclaw-agent -m "2+2 等于几？"

# 交互模式
docker compose run --rm ponyclaw-agent
```

## 🚀 快速开始

> [!TIP]
> 在 `~/.ponyclaw/config.json` 中设置您的 API Key。
> 获取 API Key: [OpenRouter](https://openrouter.ai/keys) · [Zhipu (智谱)](https://open.bigmodel.cn/usercenter/proj-mgmt/apikeys)
> 网络搜索是可选的 - 获取免费的 [Tavily API](https://tavily.com) (每月 1000 次) 或 [Brave Search API](https://brave.com/search/api) (每月 2000 次)

**1. 初始化**

```bash
ponyclaw onboard
```

**2. 配置** (`~/.ponyclaw/config.json`)

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

**3. 对话**

```bash
ponyclaw agent -m "你好"
```

就是这样！您在几分钟内就拥有了一个可工作的 AI 助手。

---

## 💬 聊天应用集成

PonyClaw 支持多种聊天平台：

| 渠道                 | 设置难度    | 特性说明                                  |
| -------------------- | ----------- | ----------------------------------------- |
| **Telegram**         | ⭐ 简单     | 推荐，支持语音转文字，长轮询无需公网      |
| **Discord**          | ⭐ 简单     | Socket Mode，支持群组/私信，Bot 生态成熟  |
| **Slack**            | ⭐ 简单     | Socket Mode (无需公网 IP)，企业级支持     |
| **QQ**               | ⭐⭐ 中等   | 官方机器人 API，适合国内社群              |
| **钉钉 (DingTalk)**  | ⭐⭐ 中等   | Stream 模式无需公网，企业办公首选         |
| **企业微信 (WeCom)** | ⭐⭐⭐ 较难 | 支持群机器人(Webhook)和自建应用(API)      |
| **飞书 (Feishu)**    | ⭐⭐⭐ 较难 | 企业级协作，功能丰富                      |
| **Line**             | ⭐⭐⭐ 较难 | 需要 HTTPS Webhook                        |
| **OneBot**           | ⭐⭐ 中等   | 兼容 NapCat/Go-CQHTTP，社区生态丰富       |

## ⚙️ 配置详解

配置文件路径: `~/.ponyclaw/config.json`

### 工作区布局

PonyClaw 将数据存储在您配置的工作区中（默认：`~/.ponyclaw/workspace`）：

```
~/.ponyclaw/workspace/
├── sessions/          # 对话会话和历史
├── memory/           # 长期记忆
├── state/            # 持久化状态
├── cron/             # 定时任务数据库
├── skills/           # 自定义技能
├── AGENTS.md         # Agent 行为指南
├── HEARTBEAT.md      # 周期性任务提示词
├── IDENTITY.md       # Agent 身份设定
├── SOUL.md           # Agent 灵魂/性格
├── TOOLS.md          # 工具描述
└── USER.md           # 用户偏好
```

### 心跳 / 周期性任务

PonyClaw 可以自动执行周期性任务。在工作区创建 `HEARTBEAT.md` 文件：

```markdown
# Periodic Tasks

- Check my email for important messages
- Review my calendar for upcoming events
```

Agent 将每隔 30 分钟（可配置）读取此文件并执行任务。

### 模型配置

PonyClaw 采用以模型为中心的配置方式，支持多种 LLM 提供商：

| 厂商                | `model` 前缀  | 获取 API Key                                                      |
| ------------------- | ------------- | ----------------------------------------------------------------- |
| **OpenAI**          | `openai/`     | [获取密钥](https://platform.openai.com)                           |
| **Anthropic**       | `anthropic/`  | [获取密钥](https://console.anthropic.com)                         |
| **智谱 AI (GLM)**   | `zhipu/`      | [获取密钥](https://open.bigmodel.cn/usercenter/proj-mgmt/apikeys) |
| **DeepSeek**        | `deepseek/`   | [获取密钥](https://platform.deepseek.com)                         |
| **Google Gemini**   | `gemini/`     | [获取密钥](https://aistudio.google.com/api-keys)                  |
| **Groq**            | `groq/`       | [获取密钥](https://console.groq.com)                              |
| **通义千问 (Qwen)** | `qwen/`       | [获取密钥](https://dashscope.console.aliyun.com)                  |
| **Ollama**          | `ollama/`     | 本地（无需密钥）                                                  |
| **OpenRouter**      | `openrouter/` | [获取密钥](https://openrouter.ai/keys)                            |

配置示例：

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

## CLI 命令行参考

| 命令                       | 描述               |
| -------------------------- | ------------------ |
| `ponyclaw onboard`         | 初始化配置和工作区 |
| `ponyclaw agent -m "..."` | 与 Agent 对话      |
| `ponyclaw agent`           | 交互式聊天模式     |
| `ponyclaw gateway`         | 启动网关           |
| `ponyclaw status`          | 显示状态           |
| `ponyclaw cron list`       | 列出所有定时任务   |
| `ponyclaw cron add ...`    | 添加定时任务       |

### 定时任务 / 提醒

PonyClaw 通过 `cron` 工具支持定时提醒和重复任务：

- **一次性提醒**: "Remind me in 10 minutes" → 10分钟后触发一次
- **重复任务**: "Remind me every 2 hours" → 每2小时触发
- **Cron 表达式**: "Remind me at 9am daily" → 使用 cron 表达式

任务存储在 `~/.ponyclaw/workspace/cron/` 中并自动处理。

## 🤝 贡献

欢迎提交 PR 和 Issue！我们致力于保持代码库的简洁和可读性。

## 🐛 疑难解答

### 网络搜索提示 "API 配置问题"

如果您尚未配置搜索 API Key，这是正常的。启用网络搜索：

1. 在 [Tavily](https://tavily.com) 或 [Brave Search](https://brave.com/search/api) 获取免费 API Key
2. 添加到 `~/.ponyclaw/config.json` 的 `tools.web` 配置中

### 遇到内容过滤错误

某些提供商有严格的内容过滤。尝试改写您的问题或使用其他模型。

### Telegram bot 提示冲突

这表示有另一个机器人实例正在运行。请确保同一时间只有一个 `ponyclaw gateway` 进程在运行。

## 📄 许可证

本项目采用 MIT 许可证。详见 [LICENSE](LICENSE) 文件。
