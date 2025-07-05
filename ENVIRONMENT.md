ASYNC_SERVICE_DEBUG   enable debug level for logging
ASYNC_SERVICE_TRACING   enable tracing level for logging
ASYNC_SERVICE_NAME   set the name of the instance of the service
ASYNC_SERVICE_TITLE   set the title in the web page
ASYNC_SERVICE_PORT_NB   the local port of the web service (default=8080)
ASYNC_SERVICE_API_KEYS   space separated list of valid API keys
ASYNC_SERVICE_SESSION_KEY
ASYNC_SERVICE_POLICY   OPA policy for access control
ASYNC_SERVICE_OPASVC   OPA service port to get the OPA policy for access control
ASYNC_SERVICE_REALM   Basic authentication realm
ASYNC_SERVICE_STAFF_USER   username of the administrator
ASYNC_SERVICE_STAFF_PASSWORD   password of the administrator
ASYNC_SERVICE_PARTICIPANT_USER   username of the user
ASYNC_SERVICE_PARTICIPANT_PASSWORD   password of the user
ASYNC_SERVICE_CERT_PEM   certificate for TLS (HTTPS) communication
ASYNC_SERVICE_KEY_PEM   key for TLS (HTTPS) communication
ASYNC_SERVICE_LOG_FILE   filename of the logging file or if it is "-" log all messages to the console
ASYNC_SERVICE_LOKI_SERVER   URL of the Loki Server, e.g. the Grafana Cloud
ASYNC_SERVICE_LOKI_USER   user name as defined for the data source as basic authentication
ASYNC_SERVICE_LOKI_PASSWORD   password as defined for the data source as basic authentication
ASYNC_SERVICE_LOKI_KEY   key/token as defined for the data source
ASYNC_SERVICE_LABELS   set of labels for Grafana Loki, like "app:myapp,tenant:mycustomer"
ASYNC_SERVICE_BUFFERSIZE   number of log lines cached before sending to Grafana Loki
ASYNC_SERVICE_MAX_DELAY   max time in seconds after which logs are flushed from buffer
ASYNC_SERVICE_LANGUAGE
ASYNC_SERVICE_LANGUAGES
ASYNC_SERVICE_USE_SSE   enable support for _server side event_ communication (default=false)
ASYNC_SERVICE_PROGRESS_DURATION   default duration of the progress bar (default=100ms)
ASYNC_SERVICE_RAPIDOC_DOC   enable Rapidoc for the OpenAPI viewer (default=false)
ASYNC_SERVICE_ELEMENTS_DOC   enable Elements for the OpenAPI viewer (default=false)
