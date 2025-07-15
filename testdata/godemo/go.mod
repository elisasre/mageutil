module godemo

go 1.24.4

// Deps will be added in CI to avoid dependabot issues.

replace github.com/elisasre/mageutil => ../../

require github.com/elisasre/mageutil v1.10.3
