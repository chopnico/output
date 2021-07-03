fmt:
	go fmt
clean:
	rm -rf vendor
	go clean
vendor:
	go mod vendor
