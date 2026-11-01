#!/bin/sh
set -e

# Remove any stale pidfile left over on persistent storage from previous container runs
rm -f /data/.agent-vault/agent-vault.pid /data/.agent-vault/*.tmp "$HOME/.agent-vault/agent-vault.pid" 2>/dev/null || true

# AGENT_VAULT_MASTER_PASSWORD is read natively by the agent-vault binary.
# This entrypoint just forwards arguments to agent-vault.
exec /usr/local/bin/agent-vault "$@"
