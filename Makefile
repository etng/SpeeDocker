
test_client:
	go run main.go pull docker.elastic.co/logstash/logstash-oss:6.8.4
test_serve:
	go run main.go serve 50001 docker.elastic.co/logstash/logstash-oss:6.8.3 etng
init:
	go get github.com/spf13/cobra/cobra
	PACKAGE=github.com/etng/SpeeDocker
	cobra init --pkg-name ${PACKAGE} .
	go mod init ${PACKAGE}
	git init
	go run main.go
	go mod tidy
	touch LICENSE
	echo '.idea' >> .gitignore

	cat <<'EOT'>~/.cobra.yaml
	author: Bo Yi <etng2004@gmail.com>
	license: MIT
	EOT
	cat ~/.cobra.yaml


	cobra add pull
	cobra add serve
build:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags "-w -s" -o bin/speedocker main.go
	upx bin/speedocker
release:
	scp bin/speedocker vms_hen:~/speedocker
