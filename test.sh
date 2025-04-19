go test -json -v ./... 2>&1 | tee /tmp/gotest.log | gotestfmt
