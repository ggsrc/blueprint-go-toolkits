module github.com/ggsrc/blueprint-go-toolkits/grpc

go 1.23.0

replace (
	github.com/ggsrc/blueprint-go-toolkits/env => ../env
	github.com/ggsrc/blueprint-go-toolkits/interceptor => ../interceptor
	github.com/ggsrc/blueprint-go-toolkits/zerolog => ../zerolog
)

require (
	github.com/ggsrc/blueprint-go-toolkits/env v0.0.0-20250307074235-8cbb76b9e006
	github.com/ggsrc/blueprint-go-toolkits/interceptor v0.0.0-20250307074235-8cbb76b9e006
	github.com/grpc-ecosystem/go-grpc-middleware/v2 v2.3.1
	github.com/grpc-ecosystem/go-grpc-prometheus v1.2.0
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/rs/zerolog v1.34.0
	go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.60.0
	google.golang.org/grpc v1.72.0
)

require (
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/bytedance/sonic v1.13.2 // indirect
	github.com/bytedance/sonic/loader v0.2.4 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/cloudwego/base64x v0.1.5 // indirect
	github.com/getsentry/sentry-go v0.32.0 // indirect
	github.com/ggsrc/blueprint-go-toolkits/zerolog v0.0.0-20250307074235-8cbb76b9e006 // indirect
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/klauspost/cpuid/v2 v2.2.10 // indirect
	github.com/mattn/go-colorable v0.1.14 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/munnerz/goautoneg v0.0.0-20191010083416-a7dc8b61c822 // indirect
	github.com/prometheus/client_golang v1.22.0 // indirect
	github.com/prometheus/client_model v0.6.2 // indirect
	github.com/prometheus/common v0.63.0 // indirect
	github.com/prometheus/procfs v0.16.1 // indirect
	github.com/twitchyliquid64/golang-asm v0.15.1 // indirect
	go.opentelemetry.io/auto/sdk v1.1.0 // indirect
	go.opentelemetry.io/otel v1.35.0 // indirect
	go.opentelemetry.io/otel/metric v1.35.0 // indirect
	go.opentelemetry.io/otel/trace v1.35.0 // indirect
	golang.org/x/arch v0.16.0 // indirect
	golang.org/x/net v0.39.0 // indirect
	golang.org/x/sys v0.32.0 // indirect
	golang.org/x/text v0.24.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250425173222-7b384671a197 // indirect
	google.golang.org/protobuf v1.36.6 // indirect
)
