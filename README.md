# ⬛ Kuro

**Kuro** (黒) is a fast, lightweight, and opinionated command-line task management tool written in Go. Built for developers who live in the terminal and value architectural simplicity over bloat.

---

## ✨ Features

- **Group-based Organization:** Categorize tasks into sub-groups seamlessly using `--group` / `-g`.
- **Fast Lookup:** Powered by Go's custom map structures for $O(1)$ group-level memory operations.
- **Clean JSON Persistence:** Human-readable data storage located directly in your terminal workspace.
- **Zero Global State:** Built with modern, idiomatic Cobra patterns to maintain thread safety and testability.

---

## 🏗️ Architecture & Design Philosophy

Kuro is built around the principle of **Separation of Concerns**:

- `internal/task/types.go` — Domain data structures (`Tasks`, `Group`, `Task`).
- `internal/task/ops.go` — Business logic and pure data mutations.
- `internal/task/files.go` — Isolated I/O logic responsible strictly for disk persistence.
- `cmd/task/` — Modular Cobra CLI commands and flag parsers with zero mutable package variables.

---

## 🚀 Quick Start

### Prerequisites
- [Go](https://go.dev/) 1.20 or higher.

### Installation

Clone the repository and build the binary:

```bash
git clone [https://github.com/YOUR_GITHUB_USERNAME/kuro.git](https://github.com/YOUR_GITHUB_USERNAME/kuro.git)
cd kuro
go build -o kuro main.go
