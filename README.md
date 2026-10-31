# anjanAI 🛡️

> **Zero-Trust Credential Broker & Transparent Security Proxy for AI Agents**

[![Go Version](https://img.shields.io/badge/Go-1.26+-00ADD8?style=flat&logo=go)](https://golang.org)
[![React](https://img.shields.io/badge/Frontend-React_19_+_Vite-61DAFB?style=flat&logo=react)](https://react.dev)
[![Database](https://img.shields.io/badge/Database-SQLite_%7C_PostgreSQL_%7C_YugabyteDB-336791?style=flat&logo=postgresql)](https://yugabyte.com)

---

## 💡 The Problem

Traditional secret management hands API keys directly to applications. When working with **autonomous AI agents** (Claude Code, Cursor, custom LangChain bots), this creates a massive security vulnerability:

> **Prompt Injection & Agent Hijacking:** An adversarial prompt can trick an AI agent into exfiltrating raw credentials (`ANTHROPIC_API_KEY`, `OPENAI_API_KEY`, `GITHUB_TOKEN`, Stripe secrets).

## 🚀 The anjanAI Solution

**anjanAI** acts as an isolated, intelligent security boundary between your AI agents and external APIs. 

* The agent **never sees or holds real credentials**.
* The agent makes outbound requests with dummy placeholder tokens (e.g., `__anthropic_api_key__`).
* **anjanAI intercepts the traffic**, verifies permissions, swaps the placeholders with real encrypted credentials on-the-fly, and securely relays the request to the target API.

---

## 🏗️ Architecture

```
┌────────────────────────────────────────────────────────┐
│ Target APIs (api.anthropic.com, api.github.com, etc.)  │
└───────────────────────────▲────────────────────────────┘
                            │ Outbound HTTPS (Real credentials injected)
┌───────────────────────────┴────────────────────────────┐
│ anjanAI Broker & Vault                                 │
│  - Port 14321: Management Web UI & Control-Plane API   │
│  - Port 14322: Transparent MITM Forward Proxy          │
└───────────────────────────▲────────────────────────────┘
                            │ Proxied via HTTPS_PROXY / HTTP_PROXY
┌───────────────────────────┴────────────────────────────┐
│ AI Agent (Claude, Cursor, Autonomous Sandboxes)        │
└────────────────────────────────────────────────────────┘
```

---

## ✨ Features

- 🔒 **Zero Credential Exposure:** Agents operate purely on dummy keys. Sensitive keys are stored encrypted at rest using AES-256-GCM.
- 🌐 **Transparent MITM Proxy (Port 14322):** Works out-of-the-box with standard `HTTPS_PROXY` and `HTTP_PROXY` environment variables.
- 🖥️ **Modern Web Management Dashboard (Port 14321):** Built with React 19, Tailwind CSS v4, and Vite for intuitive secret & service management.
- 🗄️ **Dual Database Engine:** Zero-config embedded SQLite for local development, with native support for **YugabyteDB Cloud** and PostgreSQL for production.
- ✍️ **Human-in-the-Loop Proposal Engine:** Agents cannot modify service rules or credentials on their own; they must submit change proposals that human operators approve via the UI or CLI.
- ⚡ **Cross-Platform:** Runs natively on Windows, macOS, Linux, and Docker containers.

---

## 🛠️ Quickstart

### 1. Clone & Configure

```bash
git clone https://github.com/sumiqverse/anjanAI.git
cd anjanAI
```

Create your `.env` file:

```env
# Encryption Password
AGENT_VAULT_MASTER_PASSWORD=your-super-secret-password

# Server Address & Port
AGENT_VAULT_ADDR=http://localhost:14321
PORT=14321

# Optional: YugabyteDB / PostgreSQL connection string
# DATABASE_URL=postgres://admin:password@host:5433/yugabyte?sslmode=require
```

---

### 2. Run the Server

#### Option A: Native Windows (Pre-compiled or Go)
```powershell
$env:AGENT_VAULT_MASTER_PASSWORD="your-super-secret-password"
.\agent-vault.exe server
```

Or run via Go directly:
```powershell
go run . server
```

#### Option B: Docker
```bash
docker run -it \
  -p 14321:14321 \
  -p 14322:14322 \
  -e AGENT_VAULT_MASTER_PASSWORD=your-super-secret-password \
  -v anjanai-data:/data \
  sumiqverse/anjanAI
```

---

### 3. Access the Dashboard

1. Open your browser at **`http://localhost:14321`**.
2. Create your instance **Owner/Admin account** on initial launch.
3. Configure your Vault, add credentials, and generate tokens for your AI agents.

---

### 4. Connect an AI Agent

Point your AI agent or runtime environment to **anjanAI**:

```bash
export AGENT_VAULT_ADDR=http://localhost:14321
export AGENT_VAULT_TOKEN=<your-agent-token>
export HTTPS_PROXY=http://localhost:14322
export HTTP_PROXY=http://localhost:14322
export ANTHROPIC_API_KEY=__anthropic_api_key__
```

Now, your AI agent can perform its tasks normally while your production credentials stay 100% safe!

---

## 📦 Tech Stack

- **Backend:** Go 1.26+, Cobra CLI, GORM, `modernc.org/sqlite`, `jackc/pgx`
- **Frontend:** React 19, TypeScript, Vite, Tailwind CSS v4, TanStack Router
- **Security:** AES-256-GCM, Argon2id, dynamic TLS CA certificate generator

---

