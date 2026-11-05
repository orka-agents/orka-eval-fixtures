# orka-eval-fixtures

Small Go utility packages used as review fixtures by Orka's evals.

| Package | What it does |
|---|---|
| `safepath` | Joins an untrusted relative path onto a root without escaping it |
| `textutil` | Truncates text to a byte budget without splitting characters |
| `cache` | A size-bounded LRU cache with per-entry expiry |
| `retry` | Retries a function with capped exponential backoff |
| `config` | Parses timeout settings into a bounded duration |
| `store` | Writes files atomically |

The open pull requests labelled `do-not-merge` are review fixtures pinned by
Orka's PR review eval. Don't merge them, push to them or rewrite their history.

## Examples

```go
short := textutil.Truncate("héllo wörld", 9) // "héllo…"

path, err := safepath.Join("/srv/data", userPath)
if errors.Is(err, safepath.ErrEscapesRoot) {
	// reject the request
}
```

Run the tests with `go test ./...`.
