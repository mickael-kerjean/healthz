build:
	docker build . --tag machines/healthz
	docker push machines/healthz

run:
	docker run -ti --name healthz --rm -p 5341:5341 machines/healthz
