VERSION ?= "dev"
LOCALDIR=$(PWD)
DOCKERIMG="docker.elastic.co/beats-dev/golang-crossbuild:1.18-main"
DARWIN_DOCKERIMG="docker.elastic.co/beats-dev/golang-crossbuild:1.18-darwin"
ALPINE_DOCKERIMG="golang:1.18-alpine"
REPO="github.com/axllent/golp"

linux-amd64:
	mkdir -p dist/${VERSION} && \
	docker run --rm -it -v "${LOCALDIR}":/src -w /src ${DOCKERIMG} \
	--build-cmd "go build -ldflags='-s -w -X ${REPO}/cmd.Version=${VERSION} -X ${REPO}/updater.ArchiveName=golp_linux_amd64.tar.gz' -o golp" \
	-p "linux/amd64" \
	&& tar zcvf dist/${VERSION}/golp_linux_amd64.tar.gz golp README.md LICENSE \
	&& rm -f golp

alpine-amd64:
	mkdir -p dist/${VERSION} && \
	docker run --rm -it -v "${LOCALDIR}":/src -w /src ${ALPINE_DOCKERIMG} sh - \
	"apk add --no-cache gcc g++ \ 
	&& go build -ldflags='-s -w -X ${REPO}/cmd.Version=${VERSION}' -o golp \
	&& tar zcvf dist/${VERSION}/golp_linux_alpine_amd64.tar.gz golp README.md LICENSE \
	&& rm -f golp.lol"

windows-386:
	mkdir -p dist/${VERSION} && \
	docker run --rm -it -v "${LOCALDIR}":/src -w /src ${DOCKERIMG} \
	--build-cmd "apt update && apt install -y libsass-dev && CGO_ENABLED=1 go build -ldflags='-s -w -linkmode external -extldflags -static -X ${REPO}/cmd.Version=${VERSION}' -o golp.exe" \
	-p "windows/386" \
	&& zip dist/${VERSION}/golp_windows_386.zip golp.exe README.md LICENSE \
	&& rm -f golp.exe

windows-amd64:
	mkdir -p dist/${VERSION} && \
	docker run --rm -it -v "${LOCALDIR}":/src -w /src ${DOCKERIMG} \
	--build-cmd "apt update && apt install -y libsass-dev && CGO_ENABLED=1 go build -ldflags='-s -w -linkmode external -extldflags -static -X ${REPO}/cmd.Version=${VERSION}' -o golp.exe" \
	-p "windows/amd64" \
	&& zip dist/${VERSION}/golp_windows_amd64.zip golp.exe README.md LICENSE \
	&& rm -f golp.exe

all: linux-amd64 windows-386 windows-amd64
