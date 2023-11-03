package http

import "github.com/temporalio/temporal-demo-infra/bff/pkg/clients/http"

type Doer = http.Doer
type Config = http.Config
type RequestParams = http.RequestParams

var NewClient = http.NewClient
var MustNewClient = http.MustNewClient
var NewRequest = http.NewRequest
var WithConfig = http.WithConfig
var WithBackoff = http.WithBackoff
var WithTimeout = http.WithTimeout
var WithRetryCount = http.WithRetryCount
var WithLogger = http.WithLogger
var WithPlugin = http.WithPlugin
var WithInnerHTTPClient = http.WithInnerHTTPClient
