# Extract X.Y.Z from the git-describe output. Works for v0.1.0,
# test-0.1.0-rc1, v1.2.3-4-gabcdef, etc. Upstream stripped only the
# first character which mangled non-'v' prefixes and broke windres.
VERSION=$(shell git describe --tags 2>/dev/null | grep -oE '[0-9]+\.[0-9]+\.[0-9]+' | head -1)

GO_LDFLAGS = -s -w -X 'main.Version=$(VERSION)'

export GOARCH      ?= amd64
export CGO_ENABLED  = 1

build_default:
	mkdir -p _output
	go build -tags "sqlite_foreign_keys sqlite_math_functions" -tags="" -ldflags="$(GO_LDFLAGS)" -o _output/assurcast-reader ./cmd/narr

build_macos:
	mkdir -p _output/macos
	GOOS=darwin go build -tags "sqlite_foreign_keys sqlite_math_functions macos" -ldflags="$(GO_LDFLAGS)" -o _output/macos/assurcast-reader ./cmd/narr
	cp src/platform/icon.png _output/macos/icon.png
	go run ./cmd/package_macos -outdir _output/macos -version "$(VERSION)"

build_linux:
	mkdir -p _output/linux
	GOOS=linux go build -tags "sqlite_foreign_keys sqlite_math_functions linux" -ldflags="$(GO_LDFLAGS)" -o _output/linux/assurcast-reader ./cmd/narr

build_windows:
	mkdir -p _output/windows
	go run ./cmd/generate_versioninfo -version "$(VERSION)" -outfile src/platform/versioninfo.rc
	windres -i src/platform/versioninfo.rc -O coff -o src/platform/versioninfo.syso
	GOOS=windows go build -tags "sqlite_foreign_keys sqlite_math_functions windows" -ldflags="$(GO_LDFLAGS) -H windowsgui" -o _output/windows/assurcast-reader.exe ./cmd/narr

serve:
	go run -tags "sqlite_foreign_keys sqlite_math_functions" ./cmd/narr -db local.db

install:
	go install -tags "sqlite_foreign_keys sqlite_math_functions" ./cmd/narr

test:
	go test -tags "sqlite_foreign_keys sqlite_math_functions" ./...
