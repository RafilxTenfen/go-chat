package api

import (
	"encoding/json"
	"io"
	"strconv"
	"time"

	echo "github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	logr "github.com/rhizomplatform/log"
	"github.com/sirupsen/logrus"
)

type logrusToEchoLogger struct {
	*logrus.Logger
}

// LogrusAdapter returns the global logrus logger adapted to match the
// echo Logger interface.
func LogrusAdapter() echo.Logger {
	return &logrusToEchoLogger{Logger: logrus.StandardLogger()}
}

func (l *logrusToEchoLogger) Level() log.Lvl {
	switch l.Logger.Level {
	case logrus.DebugLevel:
		return log.DEBUG
	case logrus.InfoLevel:
		return log.INFO
	case logrus.WarnLevel:
		return log.WARN
	case logrus.ErrorLevel:
		return log.ERROR
	case logrus.FatalLevel:
		return log.ERROR
	case logrus.PanicLevel:
		return log.ERROR
	default:
		return log.OFF
	}
}

func (l *logrusToEchoLogger) SetLevel(v log.Lvl) {
	switch v {
	case log.DEBUG:
		l.Logger.SetLevel(logrus.DebugLevel)
	case log.INFO:
		l.Logger.SetLevel(logrus.InfoLevel)
	case log.WARN:
		l.Logger.SetLevel(logrus.WarnLevel)
	case log.ERROR:
		l.Logger.SetLevel(logrus.ErrorLevel)
	default:
		logrus.SetLevel(logrus.PanicLevel)
	}

}

func (l *logrusToEchoLogger) Output() io.Writer {
	return l.Logger.Out
}

func (l *logrusToEchoLogger) Prefix() string {
	return ""
}

func (l *logrusToEchoLogger) SetPrefix(p string) {
	//noop
}

func (l *logrusToEchoLogger) Printj(j log.JSON) {
	b, err := json.Marshal(j)
	if err != nil {
		panic(err)
	}

	l.Print(string(b))
}

func (l *logrusToEchoLogger) SetHeader(h string) {
	// noop
}

func (l *logrusToEchoLogger) Infoj(j log.JSON) {
	b, err := json.Marshal(j)
	if err != nil {
		panic(err)
	}

	l.Info(string(b))
}

func (l *logrusToEchoLogger) Warnj(j log.JSON) {
	b, err := json.Marshal(j)
	if err != nil {
		panic(err)
	}

	l.Warn(string(b))
}

func (l *logrusToEchoLogger) Debugj(j log.JSON) {
	b, err := json.Marshal(j)
	if err != nil {
		panic(err)
	}

	l.Debug(string(b))
}

func (l *logrusToEchoLogger) Errorj(j log.JSON) {
	b, err := json.Marshal(j)
	if err != nil {
		panic(err)
	}

	l.Error(string(b))
}

func (l *logrusToEchoLogger) Fatalj(j log.JSON) {
	b, err := json.Marshal(j)
	if err != nil {
		panic(err)
	}

	l.Fatal(string(b))
}

func (l *logrusToEchoLogger) Panicj(j log.JSON) {
	b, err := json.Marshal(j)
	if err != nil {
		panic(err)
	}

	l.Panic(string(b))
}

// Logger returns a middleware that logs HTTP requests.
func Logger() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()
			res := c.Response()
			start := time.Now()

			var err error
			if err = next(c); err != nil {
				c.Error(err)
			}
			stop := time.Now()

			id := req.Header.Get(echo.HeaderXRequestID)
			if id == "" {
				id = res.Header().Get(echo.HeaderXRequestID)
			}
			reqSize := req.Header.Get(echo.HeaderContentLength)
			if reqSize == "" {
				reqSize = "0"
			}

			logr.With(logr.F{
				"id":           id,
				"ip":           c.RealIP(),
				"finished":     stop.Format(time.RFC3339),
				"host":         req.Host,
				"method":       req.Method,
				"requestURI":   req.RequestURI,
				"status":       res.Status,
				"requestSize":  reqSize,
				"responseSize": strconv.FormatInt(res.Size, 10),
				"elapsed":      stop.Sub(start).String(),
				"referer":      req.Referer(),
				"userAgent":    req.UserAgent(),
			}).Info("completed handling request")
			return nil
		}
	}
}
