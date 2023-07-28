# Filepath or URL:
SPEC_INPUT = https://api.gitbook.com/v1/openapi.json
OUTPUT_DIR = .
PACKAGE_NAME = gitbook
GIT_USER_ID = GitbookIO
GIT_REPO_ID = go-gitbook

generate: clean
	openapi-generator generate \
		-i $(SPEC_INPUT) \
		-g go \
		-o $(OUTPUT_DIR) \
		--git-user-id=$(GIT_USER_ID) \
		--git-repo-id=$(GIT_REPO_ID) \
		--additional-properties=packageName=$(PACKAGE_NAME) \
		--additional-properties=structPrefix=true \
		--additional-properties=enumClassPrefix=true
	find . -type f -name "*.go" -print0 | xargs -0 sed -i "" "s/MapmapOfStringinterface{}/Properties/g"
	gofmt -w .
	go mod tidy

clean:
	rm -rf *.go go.mod go.sum


.PHONY: generate clean