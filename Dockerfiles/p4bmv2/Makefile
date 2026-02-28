IMG=tinynetwork/p4bmv2:develop
build:
	docker build -t $(IMG) .
push:
	docker push $(IMG)
all: build push
run:
	docker run --rm -it --privileged $(IMG) bash
