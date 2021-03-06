OSFLAG :=
ifeq ($(OS),Windows_NT)
	OSFLAG = windows
else
	UNAME_S := $(shell uname -s)
	ifeq ($(UNAME_S),Linux)
		OSFLAG = linux
	endif
	ifeq ($(UNAME_S),Darwin)
		OSFLAG = darwin
	endif
endif

build:
	GOOS=$(OSFLAG) GOARCH=amd64 CGO_ENABLED=0 go build -o letscloud ./cmd/letscloud