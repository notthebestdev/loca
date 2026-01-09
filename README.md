# Loca

**Loca** is a fast and simple CLI tool written in Go to analyze how many lines of code you have written in a project.

It scans the current directory, respects `.gitignore`, ignores generated files, and provides a clean breakdown of your code by programming language.

---

## Features

- Count total lines of code  
- Breakdown by language  
- Respects `.gitignore`  
- Ignores generated files (`node_modules`, lock files, etc.)  
- Blazing fast, single binary

---

## Installation

Build the executable file:

```bash
task build
```

Move it into your PATH (optional):

```bash
sudo mv loca /usr/local/bin
```

## Usage

```bash
loca
```

## Developement

```bash
task build
task run
task lint
task test
```

## License

[MIT](LICENSE)
