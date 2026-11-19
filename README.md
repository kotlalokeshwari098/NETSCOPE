# NetScope

**A Layer-by-Layer Network Connectivity Diagnostic CLI built with Go.**

NetScope is a command-line network diagnostic tool designed to investigate connectivity issues by checking the different stages involved in establishing communication with a target host.

Instead of providing a simple success or failure message when a website or server is unreachable, NetScope aims to identify exactly which stage of the connection process fails — DNS resolution, TCP connection, TLS handshake, or HTTP communication.

The project is being built using Go, with a focus on understanding Linux networking, network protocols, sockets, and how real-world connectivity works under the hood.

## The Problem

When a website, API, or remote server becomes unreachable, identifying the actual cause can be difficult.

A connection failure can happen for several reasons:

- **DNS Resolution Failure:** The hostname cannot be resolved into an IP address.
- **TCP Connection Failure:** The destination server cannot be reached on the requested port.
- **TLS Handshake Failure:** A secure connection cannot be established.
- **HTTP Communication Failure:** The server is reachable, but the HTTP request fails or returns an unexpected response.

Developers often rely on multiple command-line utilities and manually interpret their outputs to investigate these problems.

This makes troubleshooting fragmented, especially when the actual failure occurs at a particular networking layer.

## The Solution

NetScope aims to bring these connectivity checks together into a single diagnostic workflow.

Given a hostname, NetScope will progressively investigate the connection process:

1. Resolve the hostname into its corresponding IP addresses.
2. Establish a TCP connection with the target host.
3. Perform a TLS handshake when using a secure connection.
4. Send an HTTP request and inspect the response.

Each stage will report its result independently.

A key design principle is **dependency-aware execution**: if an earlier stage fails, dependent stages should not execute.

For example, if DNS resolution fails, NetScope should stop further connection attempts and mark TCP, TLS, and HTTP as `NOT EXECUTED`, rather than incorrectly reporting them as failures.

## Diagnostic Workflow

```text
             NetScope
                 |
          Target Hostname
                 |
                 v
          DNS Resolution
                 |
           +-----+-----+
           |           |
          FAIL        PASS
           |           |
       Stop Here       v
                 TCP Connection
                       |
                 +-----+-----+
                 |           |
                FAIL        PASS
                 |           |
              Stop Here     v
                       TLS Handshake
                             |
                       +-----+-----+
                       |           |
                      FAIL        PASS
                       |           |
                    Stop Here     v
                             HTTP Request
                                   |
                                   v
                            Diagnostic Report
```

## Technology Stack

| Technology | Purpose |
|---|---|
| Go | Core programming language |
| Go `net` package | DNS resolution, IP handling, TCP connections |
| Go `crypto/tls` | TLS handshake implementation |
| Go `net/http` | HTTP communication |
| Linux / Ubuntu (WSL2) | Development and networking experiments |
| Git & GitHub | Version control and collaboration |

The project prioritizes Go's standard library to understand networking fundamentals instead of depending on large third-party networking frameworks.

## Project Philosophy

**Don't just check whether a connection works. Understand why it works or why it fails.**

NetScope is being developed incrementally, with each networking layer explored through Linux commands and practical experiments before being implemented in Go.

## Getting Started

Clone the repository:

```bash
git clone <YOUR_REPOSITORY_URL>
cd netscope
```

Run the project:

```bash
go run . google.com
```

## License
MIT
