# Color presets
GREEN=\033[1;32;40m
NC=\033[0m\n

PHONY: gitlab-token

# install dependencies: tidy for removing unused packages, vendor for installing packages in vendor directory
ref:
	@/bin/sh -c 'echo "${GREEN}[Install packages in vendor directory]${NC}"'
	@go mod tidy -v
	@go mod vendor -v
.PHONY: ref


# build
build:
	@/bin/sh -c 'echo "${GREEN}Start build process${NC}"'
	@mkdir -p bin
	@go mod download && go mod verify
	@go build -o bin/
.PHONY: build

# create changelog
current_changelog:
	@/bin/sh -c "echo \"${GREEN}[release version] $(shell npx standard-version --dry-run | grep tagging | cut -d ' ' -f4)${NC}\""
	@/bin/sh -c "echo \"${GREEN}[description] ${NC}\""
	@npx standard-version --dry-run --silent | grep -v Done | grep -v "\-\-\-" | grep -v standard-version
.PHONY: current_changelog

# static application security testing (SAST)
# go get -u github.com/securego/gosec/v2/cmd/gosec
sast:
	@/bin/sh -c 'echo "${GREEN}[Start SAST using gosec]${NC}"'
	@mkdir -p .public/sast
	@gosec -fmt=html -out=.public/sast/index.html ./...; gosec -fmt=json -out=.public/sast/results.json ./...; 
	@gosec ./...