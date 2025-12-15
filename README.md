# review-go — Go samples & data structures

This repository is a personal collection of small Go programs, data-structure implementations, and algorithm examples used for learning and experimentation.

Table of Contents

- [add_asm](./add_asm) — Small example demonstrating an assembly-backed `Add` function (Go + amd64 assembly).
- [async-traverse](./async-traverse) — Binary tree serialization/deserialization using level-order traversal and utilities for async traversal.
- [bfs](./bfs) — Breadth-first search helper with a shortest-path helper and example `Person` type.
- [builder](./builder) — Example of the builder pattern and interface embedding / composition.
- [dfs](./dfs) — Skeleton for depth-first traversal utilities and a simple `FileNode` structure.
- [gcd](./gcd) — Greatest Common Divisor implementation (Euclidean algorithm) and tests.
- [hashtable](./hashtable) — Several simple hash functions for string keys.
- [lib](./lib) — Example interfaces and concrete implementations demonstrating embedding and construction patterns.
- [linkedlist](./linkedlist) — Thread-safe generic linked list implementation with common operations.
- [node](./node) — Generic node type used by linked-list and other structures.
- [pipeline](./pipeline) — Examples of composable pipeline/pipe functions with channels, fan-out/fan-in, and numeric ops.
- [queue](./queue) — Generic queue implemented on top of the linked list.
- [recursion](./recursion) — Recursive algorithms: sum, count, max, binary search, and quicksort (including concurrent variant).
- [set](./set) — Generic set implementation with union/intersection/difference helpers.
- [signal](./signal) — Signal handling examples using `os/signal` and `context` utilities.
- [stackFromQueue](./stackFromQueue) — Stack implemented using two queues.
- [ticker](./ticker) — `Tick` helper returning an iterator-style sequence of timed ticks.
- [utils](./utils) — Utility generators and helpers (fibonacci, primes, partition, selection sort helpers, etc.).

Quick usage

- Run the repository's root `main.go` (if it is a runnable program):

```cmd
> go run .\main.go
```

- Run an example package with a `Run` function (replace `async-traverse` with the folder you want):

```cmd
> go run .\async-traverse
```

- Run all tests in the repository:

```cmd
> go test .\...
```

Notes

- Many folders are library packages (no `main`) and are intended to be used by examples or imported into other packages.
- Some examples may rely on local import paths (this repo's module path) — using `go run` or `go test` from the repo root should resolve imports correctly if `go.mod` is present.

Contributions & Next steps

- If you'd like, I can add a short README inside each subfolder with usage examples, or add a single top-level command runner that demonstrates all examples in sequence.

