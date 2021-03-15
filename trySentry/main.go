package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/getsentry/sentry-go"
)

func main() {
	err := sentry.Init(sentry.ClientOptions{
		Dsn: "http://6fcb967e6f684fe5898dddecc5633f0d@localhost:9000/1",
		BeforeSend: func(event *sentry.Event, hint *sentry.EventHint) *sentry.Event {
			if hint.Context != nil {
				if req, ok := hint.Context.Value(sentry.RequestContextKey).(*http.Request); ok {
					_ = req
					// You have access to the original Request here
				}
			}

			return event
		},
	})
	if err != nil {
		panic(fmt.Sprintf("sentry.Init: %s", err))
	}
	defer sentry.Flush(15 * time.Second)

	sentry.Logger.SetOutput(os.Stderr)
	hub := sentry.CurrentHub()
	hub.CaptureException(errors.New("CaptureException-1"))

	hub.CaptureMessage("8-try CaptureMessage-1")
	// hub.WithScope(func(scope *sentry.Scope) {
	// 	scope.SetExtra("extra-1", "extra-1 value")
	// 	scope.SetExtras(map[string]interface{}{"extra-2": "extra-2-value", "extra-3": "extra3-vaule"})
	// 	scope.SetContext("ctx-key-1", "ctx-value-1")
	// 	scope.SetFingerprint([]string{"finger1", "finger2"})
	// 	scope.SetLevel(sentry.LevelWarning)
	// 	scope.SetTag("tag-1", "tag-value-1")
	// 	scope.SetTransaction("try SetTransaction")
	// })
	// hub.CaptureException(errors.New("CaptureException-2"))
	hub.CaptureMessage("8-try CaptureMessage-2")
	sentry.CaptureMessage("8-try CaptureMessage-3")

	hub.ConfigureScope(func(scope *sentry.Scope) {
		scope.SetExtra("extra-2", "extra-2 value")
		scope.SetExtras(map[string]interface{}{"extra-2": "extra-2-value", "extra-3": "extra3-vaule"})
		scope.SetContext("ctx-key-2", map[string]interface{}{"trans1": "val2", "trans2": 2})
		scope.SetFingerprint([]string{"finger1", "finger2"})
		scope.SetTag("tag-2", "tag-value-2")
		scope.SetTransaction("try SetTransaction2")
	})
	// hub.CaptureException(errors.New("CaptureException-2"))
	hub.CaptureMessage("8-try CaptureMessage-4")
	hub.CaptureMessage("8-try CaptureMessage-4.5")
	sentry.CaptureMessage("8-try CaptureMessage-5")

	hub.PushScope()
	hub.CaptureMessage("8-try CaptureMessage-6")
	hub.PopScope()
	hub.CaptureMessage("8-try CaptureMessage-7")
	// hub.PopScope()
	// hub.CaptureMessage("8-try CaptureMessage-8")

}
