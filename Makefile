run:
	docker build \
	-t dynamornr \
	. \
	&& \
	docker run \
	-v $(shell pwd):/go/src/github.com/nycdavid/dynamornr \
	-it \
	--rm \
	dynamornr \
	/bin/ash \
	-c "cd test && dynamornr"
