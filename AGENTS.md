# createpr

A Go CLI tool that opens GitHub's "New Pull Request" page in your browser directly from the terminal.

## Build Commands

```bash
make test     # Run all tests
make lint     # Run go vet
make build    # Build the binary
```

## Critical Rules

- Pin dependencies to exact versions (e.g., `"github.com/go-git/go-git/v5": "v5.16.4"`)
- Keep docs updated with every code change
- Keep Makefile updated - add new tasks as project evolves
- Keep the tool single-purpose: open PR pages, nothing else
- Support both SSH and HTTPS remote URL formats
