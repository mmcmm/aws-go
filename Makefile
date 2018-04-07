BINARY := application
SOURCE_BUNDLE := photosapp.zip

.PHONY: bundle

bundle: build
	zip -r ${SOURCE_BUNDLE} application public sass templates Procfile

build: clean
#	@echo "Getting dep tool"
#	go get github.com/golang/dep/cmd/dep
#	@echo "Installing dependencies"
#	dep ensure
	@echo "Building 64-bit Linux binary for AWS: ${BINARY}"
	GOOS=linux GOARCH=amd64 go build -o ${BINARY}

clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi
	if [ -f ${SOURCE_BUNDLE} ] ; then rm ${SOURCE_BUNDLE} ; fi

assets:
	@echo "Building assets using gulp..."
	@type gulp >/dev/null 2>&1 || npm install gulp
	@./node_modules/.bin/gulp sass
