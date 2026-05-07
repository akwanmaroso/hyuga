# Hyuga

A GitHub Action that reads secrets from HashiCorp Vault (KV v2) and writes them to a `.env` file in your workspace.

## Inputs

| Input | Required | Default | Description |
|-------|----------|---------|-------------|
| `VAULT_ADDRESS` | Yes | — | URL of your Vault instance |
| `VAULT_TOKEN` | Yes | — | Vault token for authentication |
| `SECRET_PATH` | Yes | — | Path to the secret (e.g. `myapp/prod`) |
| `KV_MOUNT_PATH` | No | `kv` | KV engine mount path |

## Usage

```yaml
steps:
  - uses: actions/checkout@v4

  - name: Generate .env from Vault
    uses: akwanmaroso/hyuga@v-1.2.0
    with:
      VAULT_ADDRESS: ${{ secrets.VAULT_ADDRESS }}
      VAULT_TOKEN: ${{ secrets.VAULT_TOKEN }}
      SECRET_PATH: ${{ secrets.SECRET_PATH }}
      # KV_MOUNT_PATH: secrets  # optional, defaults to "kv"

  - name: Use .env
    run: cat .env
```

The action writes all key-value pairs from the specified Vault secret to `.env` in the current workspace. String values are quoted; non-string values (numbers, booleans) are written as-is.

## Roadmap

- [ ] Support KV v1
- [ ] Support multiple authentication methods (AppRole, JWT, etc.)
