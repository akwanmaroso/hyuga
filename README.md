# Hyuga

Action for generate secret from vault into .env

## How to use ?

```yaml
name: "Test"

on:
  push:
    - master

jobs:
  test:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v2
      - name: "Generate Secret"
        uses: akwanmaroso/hyuga@main
        with:
          VAULT_ADDRESS: ${{ secrets.VAULT_ADDRESS }}
          VAULT_TOKEN: ${{ secrets.VAULT_TOKEN }}
          SECRET_PATH: ${{ secerts.SECRET_PATH }}
      - name: "Read .env"
        run: |
          cat .env
        shell: bash
```

## Next Feature ?

- [ ] Allow using vault v1
- [ ] Allow authentication in multiple method
