# Logger for go

### Install

```bash
go get github.com/nextmillenniummedia/logger-go
```

### Init

Init main logger:
```go
logger := loggergo.NewLogger().
    Params("trace_id", traceId).
    Level(loggergo.LOG_INFO).
```

Cloning the logger for child processes while preserving settings:
```go
logger = logger.Clone().From("Another service")
```

### Example

Example: in `json` format
```go
logger.Params("trace_id", trace_id).From("Service name")
logger.Info("You message", "param1": "value1")
```
Stdout:
```
{"level": 30, "from": "Service name", "message": "You message", "trace_id": "xxx-xxx-xxx", "param1": "value1"}
```

Example: for local development in `pretty` mode
```go
logger.Params("trace_id", trace_id).From("Service name")
logger.Pretty() // Enable pretty mode
logger.Info("You message")
```
Stdout:
```
12:31:33.988 [INFO]    [Service name] You message
    trace_id: xxx-xxx-xxx
    param1: value1
    param2: value2
```

### Levels with codes

| Level: | verbose | debug | info | warn | error | fatal | silent   |
|--------|---------|-------|------|------|-------|-------|----------|
| Value: | 10      | 20    | 30   | 40   | 50    | 60    | 100      |

Log message with level lower than set will be skipped

```go
logger := loggergo.NewLogger().Level(loggergo.LEVEL_ERROR)
logger.Error("Error message")
logger.Info("info message") // Will be skipped
```
