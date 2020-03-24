IMAGE_VERSION = latest
REGISTRY = docker.io/joyme
IMAGE = ${REGISTRY}/cola-device-plugin:${IMAGE_VERSION}

.PHONY: build deploy

build:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o build/cola cmd/server/app.go

buildImage:
	docker build -t ${IMAGE} .

kindLoad:
	kind load docker-image ${IMAGE}

pushImage:
	docker push ${IMAGE}

deploy:
	helm install cola deploy/helm/cola

upgrade:
	helm upgrade cola deploy/helm/cola

dry-run:
	helm install cola deploy/helm/cola --dry-run