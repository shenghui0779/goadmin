IMAGE=uhub.service.ucloud.cn/leesin/goadmin:v0.1.4

build:
	systemctl stop goadmin
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o goadmin main.go
	systemctl start goadmin
git:
	bash shell/git.sh

backup:
	bash shell/backup.sh