package logger

import (
	"github.com/kataras/iris/v12/middleware/accesslog"
	"os"
)

func NewRequestLogger() *accesslog.AccessLog {
	ac := accesslog.File("./access.log")
	ac.AddOutput(os.Stdout)

	// The default configuration:
	ac.Delim = '|'
	ac.TimeFormat = "2006-01-02 15:04:05"
	ac.Async = false
	ac.IP = true
	ac.BytesReceivedBody = true
	ac.BytesSentBody = true
	ac.BytesReceived = false
	ac.BytesSent = false
	ac.BodyMinify = true
	ac.RequestBody = true
	ac.ResponseBody = false
	ac.KeepMultiLineError = true
	ac.PanicLog = accesslog.LogCallers

	// Default line format if formatter is missing:
	// Time|Latency|Code|Method|Path|IP|Path Params Query Fields|Bytes Received|Bytes Sent|Request|Response|
	//
	// Set Custom Formatter:
	ac.SetFormatter(&accesslog.JSON{
		Indent:    " ",
		HumanTime: true,
	})
	//ac.SetFormatter(&accesslog.CSV{})
	//ac.SetFormatter(&accesslog.Template{Text: "{{.Code}}"})

	return ac
}
