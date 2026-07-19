# AGENTS.md

## What This Is

WeiGo 2.0 is a **Go web framework library** (not an application). It wraps Echo v4 + GORM v2 to provide PHP-style MVC development. The module path is `github.com/wycto/weigo`.

**This repo has no `main.go` or executable.** It's imported by other projects. Do not look for `go run` or build targets.

## Quick Reference

```bash
# Run all tests
go test ./...

# Run tests for a single package
go test ./config/...
go test ./controller/...
go test ./model/...

# Run a specific test
go test -run TestLoadValidConfig ./config/

# Check for compilation errors
go build ./...
```

## Architecture (Key Insight: Reflection Routing)

Routes are **not** defined with explicit handler functions. Instead, `controller.Router()` uses reflection to auto-discover controller methods:

```go
// This automatically registers GET/POST/PUT/DELETE for /user/*
controller.Router(weigo.Echo, "/user/", "index", &index.UserController{})
```

**Action method signature requirements** (non-negotiable):
- Must be `func (c *XxxController) MethodName() error`
- Zero parameters (besides receiver)
- Returns exactly one `error`
- Cannot be defined on base `Controller` or `WebSocketController`
- Method name matching is **case-insensitive**, hyphens stripped

## URL Parameter Pattern (PHP-Style)

Parameters come as key/value pairs in the URL path:
```
/user/info/id/42/name/john
```
Maps to: `c.GetParam("id")` → `"42"`, `c.GetParam("name")` → `"john"`

## Response Format

All responses follow: `{code: int, msg: string, data: any}`

| Code | Meaning |
|------|---------|
| 0 | Success |
| 1 | Error (HTTP 200, business logic error) |
| 200101-200103 | Parameter errors |
| 200201-200205 | Data errors |
| 200501 | System error |

Standard HTTP status codes used: 400, 401, 403, 404, 500.

## Configuration

Loads from `config/config.json` (relative to working directory). Panics on missing file or invalid JSON.

**CORS behavior is non-obvious:**
- `Debug: true` + no `AllowOrigins` → allows ALL origins
- `Debug: true` + `AllowOrigins` configured → **only** allows listed origins
- `Debug: false` + no `AllowOrigins` → rejects ALL origins
- Wildcard `*` in `AllowOrigins` → allows all

## Database

Supports: MySQL, PostgreSQL, SQLite. Table prefix from config is **automatically prepended** by `DB.TableName()`.

```go
// This automatically adds prefix: "SELECT * FROM wei_users"
weigo.DB.Table("user").Find(&users)

// Chain query also adds prefix
model.NewChain(weigo.DB, "user").Find(&users)
```

SQLite skips connection pool configuration (in-memory for tests).

## Chain Query Builder

Custom wrapper over GORM (not standard GORM chaining):

```go
model.NewChain(weigo.DB, "user").
    Fields("id,name").
    Where("status = ?", 1).
    Where("age > ?", 18).
    Order("id desc").
    Page(1, 10).
    Find(&users)
```

**Note:** Each `Where()` call appends a condition. Parameters are matched to `?` placeholders by position.

## JSON Body Auto-Parsing

For POST/PUT/PATCH requests, the JSON body is automatically decoded into `c.body`. Access via:
- `c.GetBody()` → full map
- `c.GetInt("key", default)` → typed access
- `c.GetString("key", default)` → typed access
- `c.GetBool("key", default)` → typed access

Body is **not** parsed for GET/DELETE requests.

## Template Rendering

- Debug mode: templates re-parsed on every request (no caching)
- Production: templates cached after first parse
- Call `controller.ClearTemplateCache()` after modifying templates in production

## Testing Patterns

- Tests use white-box testing (same package, access unexported fields)
- Config tests: `t.TempDir()` for isolated file system
- Database tests: SQLite in-memory (`:memory:`)
- Controller/Response tests: `httptest.NewRecorder()` + `echo.NewContext()`

## Common Pitfalls

1. **Action methods must return `error`** — framework calls them via reflection, non-error returns cause silent failures
2. **CORS in debug mode** — don't assume "debug = permissive" when `AllowOrigins` is configured
3. **Table prefix** — never hardcode prefixed table names; use `weigo.DB.TableName("user")` or ChainQuery
4. **URL parameters** — key/value pairs, not query strings; `/user/info/id/42` not `/user/info?id=42`
5. **Response code 1 vs HTTP errors** — code 1 is business error with HTTP 200; use `BadRequest()`/`Unauthorized()` for proper HTTP status codes
