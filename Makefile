# Filepath or URL:
SPEC_INPUT = /Users/dstotijn/projects/gitbook-x/packages/api-client/static/openapi.json
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
	mkdir -p ./api
	mv *.go ./api
	find . -type f -name "*.go" -print0 | xargs -0 sed -i "" "s/MapmapOfStringinterface{}/Properties/g"
	gofmt -w ./api
	go mod tidy

clean:
	rm -rf api go.mod go.sum


.PHONY: generate clean