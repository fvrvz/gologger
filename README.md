## 📦 go-logger

`go-logger` is a lightweight, reusable, and customizable logging library for Go that adds semantic log levels (`INFO`, `DEBUG`, `ERROR`, `WARN`) with **colored terminal output** for better readability during development.

Designed for simplicity and clarity, this library wraps Go's standard `log` package, making it easy to use across any project — from web servers to CLIs.

---

## ✨ Features

* ✅ Easy-to-use logging functions: `logger.INFO()`, `logger.ERROR()`, etc.
* 🎨 Colored log level prefixes in terminal output
* 📁 Log source file and line number included
* 🔁 Minimal external dependencies (uses only Go standard library)
* ⚙️ Extendable: add file logging, log filtering, or structured logging later

---

## 🚀 Usage Example

```go
import "github.com/fvrvz/go-logger"

func main() {
    logger.INFO("Server started on port %d", 8080)
    logger.DEBUG("Loaded config: %+v", config)
    logger.WARN("Using default fallback value")
    logger.ERROR("Database connection failed: %v", err)
}
```

---

## 📁 Installation

```bash
go get github.com/fvrvz/go-logger
```

---

## 🔧 Configuration (Optional)

You can easily add support for:

* Disabling colors in production via environment variables
* Logging to files or multiple outputs
* Integrating with Gin, Echo, or any Go web framework

Let me know if you want those features built-in via Issues or PRs.
