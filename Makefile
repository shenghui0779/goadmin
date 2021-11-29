IMAGE=uhub.service.ucloud.cn/leesin/goadmin:v0.1.4

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 docker build -t ${IMAGE} .
	docker push ${IMAGE}