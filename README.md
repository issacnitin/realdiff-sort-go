# RealDiff Go sort-stability demo

A three-test fixture for enabling a deterministic code tie-break in a stable priority sort. The pull request edits only `config.go`; the behavior frontier remains in unedited pricing code, and only the exact tie-winner assertion reacts.

Run locally with `go test ./...`.
