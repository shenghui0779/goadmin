IMAGE=uhub.service.ucloud.cn/leesin/goadmin:v0.1.4
TIME=$(date '+%Y-%m-%d %H:%M:%S')

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o goadmin main.go
	scp goadmin root@121.199.68.249:/data/goadmin/goadmin
git:
	bash shell/git.sh