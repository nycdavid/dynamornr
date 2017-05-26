image:
	docker build \
	-t dynamornr \
	.
run:
	docker run \
	-it \
	--rm \
	dynamornr \
	dynamornr
