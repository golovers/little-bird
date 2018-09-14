package stackdriver

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

// Payload format:
// {
//   "eventTime": string,
//   "serviceContext": {
//     "service": string,     // Required.
//     "version": string
//   },
//   "message": string,       // Required. Should contain the full exception
//                            // message, including the stack trace.
//   "context": {
//     "httpRequest": {
//       "method": string,
//       "url": string,
//       "userAgent": string,
//       "referrer": string,
//       "responseStatusCode": number,
//       "remoteIp": string
//     },
//     "user": string,
//     "reportLocation": {    // Required if no stack trace in 'message'.
//       "filePath": string,
//       "lineNumber": number,
//       "functionName": string
//     }
//   }
// }
type Payload struct {
	EventTime      string `json:"eventTime"`
	ServiceContext struct {
		Service string `json:"service"` // Required
		Version string `json:"version,omitempty"`
	} `json:"serviceContext"`
	Message string `json:"message"` // Required. runtime.Stack()
	Context struct {
		HTTPContext struct {
			Method             string `json:"method,omitempty"`
			URL                string `json:"url,omitempty"` // attach some custom data to ctx from interceptor?
			UserAgent          string `json:"userAgent,omitempty"`
			Referrer           string `json:"referrer,omitempty"`
			RemoteIP           string `json:"remoteIp,omitempty"` // peer.FromContext()
			ResponseStatusCode int    `json:"responseStatusCode,omitempty"`
		} `json:"httpContext,omitempty"`
		User           string   `json:"user,omitempty"`
		ReportLocation struct { // Required if no stacktrace in 'message' (runtime.Caller(1))
			FilePath     string `json:"filePath,omitempty"`
			LineNumber   int    `json:"lineNumber,omitempty"`
			FunctionName string `json:"functionName,omitempty"`
		} `json:"reportLocation,omitempty"`
	} `json:"context,omitempty"`
}

var errorLevels = []logrus.Level{logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel}

// Hook should use https://godoc.org/cloud.google.com/go/errors#WithMessage as reference
// https://github.com/GoogleCloudPlatform/google-cloud-go/blob/master/errors/errors.go#L330
type Hook struct {
	Service string
	Version string
	Output  io.Writer
}

// NewHook create and returns a new instance of Hook reporting on StdErr.
func NewHook(service, version string) *Hook {
	return &Hook{
		Service: service,
		Version: version,
		Output:  os.Stderr,
	}
}

// Fire emits an error report based on the log entry.
func (h *Hook) Fire(e *logrus.Entry) error {
	data, err := h.formatError(e)
	if err != nil {
		return err
	}
	_, err = h.Output.Write(data)
	return err
}

// Levels provides which levels the hook reacts to.
func (h *Hook) Levels() []logrus.Level {
	return errorLevels
}

func (h *Hook) formatError(e *logrus.Entry) ([]byte, error) {
	var payload Payload
	payload.EventTime = e.Time.In(time.UTC).Format(time.RFC3339Nano)
	payload.ServiceContext.Service = h.Service
	payload.ServiceContext.Version = h.Version
	if r, ok := e.Data[HTTPRequestKey].(*http.Request); ok {
		payload.Context.HTTPContext.Method = r.Method
		payload.Context.HTTPContext.RemoteIP = r.RemoteAddr
		payload.Context.HTTPContext.URL = r.Host + r.RequestURI
		payload.Context.HTTPContext.Referrer = r.Referer()
		payload.Context.HTTPContext.UserAgent = r.UserAgent()
	}
	payload.Message = e.Message + "\n" + formatStack(e.Level)
	serialized, err := json.Marshal(&payload)
	if err != nil {
		return nil, err
	}
	return append(serialized, '\n'), nil
}

func formatStack(l logrus.Level) string {
	// limit the stack trace to 16k.
	var buf [16384]byte
	s := buf[0:runtime.Stack(buf[:], false)]
	lf := bytes.IndexByte(s, '\n')
	if lf == -1 {
		return string(s)
	}
	stack := s[lf:]
	var strips []string
	if l == logrus.PanicLevel {
		strips = []string{"panic("}
	} else {
		fn := strings.Title(l.String())
		strips = []string{
			fmt.Sprintf("github.com/sirupsen/logrus.%s", fn),
			fmt.Sprintf("github.com/sirupsen/logrus.(*Logger).%s", fn),
			fmt.Sprintf("github.com/sirupsen/logrus.(*Entry).%s", fn),
		}
	}
	var stripLine int
	for _, strip := range strips {
		stripLine = bytes.Index(stack, []byte(strip))
		if stripLine != -1 {
			break
		}
	}
	if stripLine == -1 {
		return string(s)
	}
	stack = stack[stripLine+1:]
	for i := 0; i < 2; i++ {
		nextLine := bytes.IndexByte(stack, '\n')
		if nextLine == -1 {
			return string(s)
		}
		stack = stack[nextLine+1:]
	}
	return string(s[:lf+1]) + string(stack)
}
