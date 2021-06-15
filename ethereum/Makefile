tests:
	DOCKER_BUILDKIT=1 docker build --no-cache --rm -t qtum-contract-test .
	docker-compose up -d qtum_testchain janus tests