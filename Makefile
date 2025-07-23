tests:
	go test -coverprofile=coverage.out -covermode atomic .
	go tool cover -func=coverage.out | grep total:
	go tool cover -html coverage.out -o coverage.html