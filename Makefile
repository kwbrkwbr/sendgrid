GOCMD=go
GOVET=$(GOCMD) vet
GOGEN=$(GOCMD) generate
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOTOOL=$(GOCMD) tool
GOMOD=$(GOCMD) mod
GORUN=$(GOCMD) run

ensure:
	$(GOMOD) download

vet:
	$(GOVET) $$(go list ./... )

wire-echo:
	$(GORUN) github.com/google/wire/cmd/wire ./cmd/echo/...

wire-fiber:
	$(GORUN) github.com/google/wire/cmd/wire ./cmd/fiber/...

realize:
	realize s --build --run

newman:
	newman run ./var/postman/sendgrid.postman_collection.json

build: wire-echo
	$(GOBUILD) -ldflags="-w -s" -o main ./cmd/echo

build-mock: wire-echo
	go build -ldflags="-w -s" -tags "mock" -o main ./cmd/echo/

docker-sendgrid:
	docker build -f sendgrid.Dockerfile . -t sendgrid

cloud-build-sendgrid:
	gcloud builds submit --config cloudbuild-sendgrid.yaml .

install-realize:
	GO111MODULE=off go get -u github.com/oxequa/realize
	realize -v

install-delve:
	go get -u github.com/go-delve/delve/cmd/dlv@v1.5.0
	go build -o $(GOPATH)/bin/dlv $(GOPATH)/pkg/mod/github.com/go-delve/delve@v1.5.0/cmd/dlv

install-air:
	curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $$(go env GOPATH)/bin
	air -v # `air -v` が動かない場合は、自分で $$(GOPATH)/bin のairをpathに通してください。

air:
	air -c .air/.air.toml

install-task:
	sh -c "$$(curl -ssL https://taskfile.dev/install.sh)" -- -d -b $$(go env GOPATH)/bin
	task --version # `task --version` が動かない場合は、自分で $$(GOPATH)/bin のtaskをpathに通してください。

activate:
	gcloud config configurations activate $(GCP_CONFIG)