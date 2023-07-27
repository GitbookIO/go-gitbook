SPEC_URL = https://api.gitbook.com/openapi.yaml
OUTPUT_DIR = .
PACKAGE_NAME = gitbook
GIT_REPO_ID = go-gitbook
GIT_USER_ID = GitbookIO

generate: clean
	export GO_POST_PROCESS_FILE="gofmt -w" && \
	openapi-generator generate \
		-i $(SPEC_URL) \
		-g go \
		-o $(OUTPUT_DIR) \
		--enable-post-process-file \
		--git-user-id=$(GIT_USER_ID) \
		--git-repo-id=$(GIT_REPO_ID) \
		--additional-properties=packageName=$(PACKAGE_NAME) \
		--additional-properties=structPrefix=true \
		--additional-properties=enumClassPrefix=true

clean:
	rm -rf *.go go.mod go.sum


.PHONY: generate clean