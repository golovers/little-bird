package stackdriver

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strings"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestHook(t *testing.T) {
	log := logrus.New()
	var service = "service"
	var version = "version"
	var errorMessage = "Error reported"
	var output bytes.Buffer
	hook := NewHook(service, version)
	hook.Output = &output
	log.Hooks.Add(hook)
	log.Error(errorMessage)
	var payload Payload
	err := json.Unmarshal(output.Bytes(), &payload)
	if err != nil {
		t.Fatal(err)
	}
	if payload.ServiceContext.Service != service {
		t.Errorf("Payload.ServiceContext.Service = %s; expected %s", payload.ServiceContext.Service, service)
	}
	if payload.ServiceContext.Version != version {
		t.Errorf("Payload.ServiceContext.Version = %s; expected %s", payload.ServiceContext.Version, version)
	}
	if !strings.HasPrefix(payload.Message, errorMessage) {
		t.Errorf("Payload.Message = %s; expected %s (rest left out)", payload.Message, errorMessage)
	}
}

func TestRequestHook(t *testing.T) {
	log := logrus.New()
	var service = "service"
	var version = "version"
	var errorMessage = "Error reported"
	var output bytes.Buffer
	hook := NewHook(service, version)
	hook.Output = &output
	log.Hooks.Add(hook)

	request, err := http.NewRequest("GET", "http://api.d21s.com/v1/api", bytes.NewReader([]byte("Request content")))
	if err != nil {
		t.Fatal(err)
	}
	log.WithField(HTTPRequestKey, request).Error(errorMessage)
	var payload Payload
	err = json.Unmarshal(output.Bytes(), &payload)
	if err != nil {
		t.Fatal(err)
	}
	if payload.ServiceContext.Service != service {
		t.Errorf("Payload.ServiceContext.Service = %s; expected %s", payload.ServiceContext.Service, service)
	}
	if payload.ServiceContext.Version != version {
		t.Errorf("Payload.ServiceContext.Version = %s; expected %s", payload.ServiceContext.Version, version)
	}
	if !strings.HasPrefix(payload.Message, errorMessage) {
		t.Errorf("Payload.Message = %s; expected %s (rest left out)", payload.Message, errorMessage)
	}
	if payload.Context.HTTPContext.Method != request.Method {
		t.Errorf("HTTPContext.Method = %s; expected %s", payload.Context.HTTPContext.Method, request.Method)
	}
	if payload.Context.HTTPContext.URL != request.Host+request.RequestURI {
		t.Errorf("HTTPContext.URL = %s; expected %s", payload.Context.HTTPContext.URL, request.Host+request.RequestURI)
	}
}
