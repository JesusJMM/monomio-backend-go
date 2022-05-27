build:
	go build -o server ./cmd/server/main.go

run-dev: build
	TOKEN_SECRET_KEY="75c4510a079123e3a0ae02afd3f30f71d4efe21e1e45c3af40b70627fdf0f62f155c3f16329237fd1a5987e76d77a137efb2" \
							 PORT=8080\
							 ./server
watch:
	ulimit -n 1000 #increase the file watch limit, might required on MacOS
	reflex -s -r '\.go$$' make run-dev

apiTest:
	go test ./api
