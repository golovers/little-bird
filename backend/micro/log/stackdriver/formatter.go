package stackdriver

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

const (
	timestampFieldKey = "timestamp"
	severityFieldKey  = "severity"
	messageFieldKey   = "message"

	// HTTPRequestKey should be used to attach a http.Request to the
	// log statements.
	// TODO: Move out of this package and into a general package.
	HTTPRequestKey = "httpRequest"
)

// Formatter for logrus with the Stackdriver format.
type Formatter struct {
	DisableTimestamp bool
	// Due to a bug in k8s 1.6.6 json parsing got broken unless content
	// is placed within a container.
	UseContainer bool
}

// Format logrus.Entry objects for Stackdriver,
// based on the format located in the fluentd-google plugin.
// {
//   "timestamp": {"seconds": 0, "nanos": 0},
//   "severity": "DEFAULT DEBUG INFO NOTICE WARNING ERROR CRITICAL ALERT EMERGENCY",
//   "httpRequest": {
//     "httpMethod": "",
//     "requestUrl": "",
//     "requestSize": "",
//     "status": 0,
//     "responseSize": "",
//     "userAgent": "",
//     "remoteIp": "",
//   },
//   "message": "",
// }
func (f *Formatter) Format(entry *logrus.Entry) ([]byte, error) {
	data := make(logrus.Fields, len(entry.Data)+3)
	for k, v := range entry.Data {
		switch v := v.(type) {
		case error:
			data[k] = v.Error()
		case *http.Request:
			data[k] = parseHTTPRequest(v)
		default:
			data[k] = v
		}
	}
	if !f.DisableTimestamp {
		t := entry.Time.In(time.UTC)
		seconds := t.Unix()
		nanos := int32(t.Sub(time.Unix(seconds, 0).In(time.UTC)))
		data[timestampFieldKey] = map[string]interface{}{
			"seconds": seconds,
			"nanos":   nanos,
		}
	}
	data[messageFieldKey] = entry.Message
	data[severityFieldKey] = levelToSeverity(entry.Level)
	if f.UseContainer {
		data = map[string]interface{}{"log": data}
	}
	serialized, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("Failed to marshal fields to JSON, %v", err)
	}
	return append(serialized, '\n'), nil
}

func parseHTTPRequest(r *http.Request) map[string]interface{} {
	e := map[string]interface{}{
		"requestMethod":                  r.Method,
		"requestUrl":                     r.Host + r.RequestURI,
		"requestSize":                    r.ContentLength,
		"userAgent":                      r.UserAgent(),
		"remoteIp":                       r.RemoteAddr,
		"serverIp":                       "",
		"referer":                        r.Referer(),
		"cacheLookup":                    false,
		"cacheHit":                       false,
		"cacheValidatedWithOriginServer": false,
	}
	if m, ok := FromContext(r.Context()); ok {
		e["latency"] = fmt.Sprintf("%fs", m.Duration.Seconds())
		e["status"] = m.Code
		e["responseSize"] = m.Written
	}
	return e
}

// Convert the logrus.Level to a stackdriver supported level.
func levelToSeverity(l logrus.Level) string {
	switch l {
	case logrus.PanicLevel, logrus.FatalLevel:
		return "critical"
	}
	return l.String()
}
