module gotool

go 1.24.0

toolchain go1.24.7

// cve_2026_98432 需关闭mptcp
godebug multipathtcp=0

require (
	github.com/Workiva/go-datastructures v1.1.6
	github.com/bytedance/sonic v1.14.1
	github.com/duke-git/lancet/v2 v2.3.7
	github.com/gogf/gf/v2 v2.10.4
	github.com/google/uuid v1.6.0
	github.com/json-iterator/go v1.1.12
	github.com/shopspring/decimal v1.4.0
	github.com/stretchr/testify v1.11.1
	github.com/unknwon/i18n v0.0.0-20210904045753-ff3a8617e361
)

require (
	github.com/bytedance/gopkg v0.1.3 // indirect
	github.com/bytedance/sonic/loader v0.3.0 // indirect
	github.com/cloudwego/base64x v0.1.6 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/emirpasic/gods/v2 v2.0.0-alpha // indirect
	github.com/fatih/color v1.18.0 // indirect
	github.com/fsnotify/fsnotify v1.9.0 // indirect
	github.com/go-logr/logr v1.4.3 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/klauspost/cpuid/v2 v2.2.9 // indirect
	github.com/mattn/go-colorable v0.1.14 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mattn/go-runewidth v0.0.19 // indirect
	github.com/modern-go/concurrent v0.0.0-20180228061459-e0a39a4cb421 // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/rivo/uniseg v0.4.7 // indirect
	github.com/twitchyliquid64/golang-asm v0.15.1 // indirect
	go.opentelemetry.io/auto/sdk v1.1.0 // indirect
	go.opentelemetry.io/otel v1.38.0 // indirect
	go.opentelemetry.io/otel/metric v1.38.0 // indirect
	go.opentelemetry.io/otel/sdk v1.38.0 // indirect
	go.opentelemetry.io/otel/trace v1.38.0 // indirect
	golang.org/x/arch v0.21.0 // indirect
	golang.org/x/net v0.43.0 // indirect
	golang.org/x/sys v0.36.0 // indirect
	golang.org/x/text v0.28.0 // indirect
	gopkg.in/ini.v1 v1.46.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
