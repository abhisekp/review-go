# add_asm

Small example showing how to implement a Go function with an amd64 assembly backing.

What it is

- Declares `Add(a, b int64) int64` in Go and implements it in `add_amd64.s`.

Quick notes

- This is a package-level example (not a runnable `main`).
- Intended to demonstrate calling assembly from Go.

Files

- `add_asm.go` — Go declaration for `Add`.
- `add_amd64.s` — amd64 assembly implementation.

Usage

- Use the package from other Go code: `import "review-go/add_asm"` and call `add_asm.Add(...)`.

