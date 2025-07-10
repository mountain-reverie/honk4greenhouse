# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

### Technical Overview

This is a Go web application called "honk4greenhouse" - a website for designing efficient and potentially passive greenhouses optimized for Canadian climate conditions. The project prioritizes maintainability and API stability, leveraging Go-native solutions wherever possible.

#### Core Technology Stack

**Backend & Framework**

- [Gin](https://gin-gonic.com/): HTTP web framework with middleware support
- [Goth](https://github.com/markbates/goth): Multi-provider OAuth authentication ([introduction blog](https://dizzy.zone/2018/06/01/OAuth-with-Gin-and-Goth/))

**Frontend & UI**

- [Templ](https://templ.guide/): Type-safe HTML templating system
- [Templ UI](https://templui.io/): Component library built on Templ and [HTMX](https://htmx.org/)
- [HTMX](https://htmx.org/): Dynamic web interactions without JavaScript

**Database & Infrastructure**

- [Turso](https://turso.tech/): SQLite-compatible database with local replicas on each host
- [Pulumi](https://www.pulumi.com/docs/iac/languages-sdks/go/): Infrastructure as Code for setup and maintenance

**Testing & Monitoring**

- [Playwright CI Go](https://github.com/mountain-reverie/playwright-ci-go): End-to-end web testing framework
- OpenTelemetry: Comprehensive observability (logs, metrics, traces)
- [SigNoz](https://signoz.io/docs/install/self-host/): Self-hosted observability platform
- [Beszel](https://www.youtube.com/watch?v=O_9wT-5LoHM): System monitoring solution

#### Development Philosophy

**Maintainability First**: The codebase is designed for long-term stability with automatic dependency updates via Dependabot. All dependencies are selected for their proven track record of API stability and backward compatibility.

**Testing Strategy**: Comprehensive integration and end-to-end testing with performance benchmarking. Go PGO (Profile-Guided Optimization) uses CI benchmark results for optimization. Turso's local replica capability enables full end-to-end testing with isolated database instances.

**CI/CD Pipeline**: GitHub Actions handle continuous integration and deployment with strict quality gates, including benchmark regression detection to prevent performance degradation and maintain security standards.

#### Deployment Architecture

**Infrastructure**: Deployed exclusively on European and Canadian cloud providers ([OVH](https://www.ovhcloud.com/en-ca/)) across 3 machines:

- 1 monitoring server (SigNoz + Beszel)
- 2 application servers (load balanced)

**Networking**:

- Cloudflare Tunnel for secure internet exposure
- Tailscale for VPN, administrative access, and core VPC management

**Reference Videos** (for understanding Tailscale deployment patterns):

- [How to use cloud-init and Tailscale | Infrastructure as Code Series Part 1](https://www.youtube.com/watch?v=e-X5FJwrkaA)
- [Automate your Tailscale cloud deployments with Terraform | Infrastructure as Code Series Part 2](https://www.youtube.com/watch?v=PEoMmZOj6Cg)
- [An Ansible primer for Devops | Infrastructure as Code Series Part 3](https://www.youtube.com/watch?v=k5Xgt31yK2U)
- [Static site deployments made easy with Github Actions and Tailscale](https://www.youtube.com/watch?v=OQJAX-Ce1YY)

**Security**: GitHub Action environments segregate secrets with progressive exposure based on CI validation results. Performance regression gates prevent potentially compromised code from reaching production environments.

## Build and Development Commands

Since this is a Go project, use standard Go commands:

- `go run cmd/service/main.go` - Run the main service
- `go build cmd/service/main.go` - Build the service binary
- `go test ./...` - Run all tests (when tests are added)
- `go mod tidy` - Clean up module dependencies

## Architecture

- **Entry Point**: `cmd/service/main.go` - Contains the main function and service entry point
- **Module**: `github.com/mountain-reverie/true-north-greenhouse` (Go 1.23.6)
- **Structure**: Simple Go application structure with cmd/ directory for executables

The codebase is minimal and appears to be a starter project ready for greenhouse design functionality to be implemented.
