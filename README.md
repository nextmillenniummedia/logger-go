# Logger for go

### Format

Use `json` because it is easy to parse on the logging system side

Example: in `json` format
```
{"level": 1, "module": "predictor", "message": "Call predict response", "trace_id": "xxx-xxx-xxx", "param1": "value1", "param2": "value2"}
```

Example: for local development in pretty mode
```
12:31:33.988 [INFO] [predictor] Call predict response 
    trace_id xxx-xxx-xxx
    param1 value1
    param2 value2
```

### Levels with codes

| Level: | verbose | debug | info | warn | error | fatal | silent   |
|--------|---------|-------|------|------|-------|-------|----------|
| Value: | 10      | 20    | 30   | 40   | 50    | 60    | 100      |

Logs with a level which is smaller than in the settings will be skipped


### Setup logger

```go
logger := logging.NewLogger(...)
    Params("trace_id", traceId).
    Level(logLevel)
```


### Example of usage

In structure
```go
func NewService(ctx context.Context, ...) *IService {
    logger := logging.CloneFromContext(ctx, "service_name")
    return &Service{
        logger: logger 
    }
}

func (u *Service) BusinessMethod(param string) {
    u.logger.Info("BusinessMethod missed a step because ...", "param", param)
}
```

In simple function
```go
var LOG_MODULE = "module_name"

func SimpleFunction(ctx context.Context, param1 string) {
    logger := logging.CloneFromContext(ctx, LOG_MODULE)
    logger.Info("simpleFunction start", "param1", param1)
    // ...
    logger.Info("simpleFunction end")
}
```

### Helpers for context
```go
// Set logger to context
func SetToContext(ctx context.Context, logger ILogger) {...}
// Restore from context and clone
func CloneFromContext(ctx context.Context, module string) ILogger {...}
```

