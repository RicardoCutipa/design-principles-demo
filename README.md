# Design Principles Demo (Go)

This repo is a small companion project for a Medium article about software design principles.

## What's inside

- `checkout/`: a tiny checkout pricing model using composition (`PricingRule`) instead of a growing conditional chain.
- GitHub Actions workflow that runs `go test ./...` on every push and pull request.

## Run locally

```bash
go test ./...
```

