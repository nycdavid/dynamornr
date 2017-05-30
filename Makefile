compile:
	docker build \
	-t dynamornr \
	.
list-tables:
	make compile \
	&& docker run \
	-v $(shell pwd):/go/src/github.com/nycdavid/dynamornr \
	-e ENV=test \
	-e AWS_SECRET_ACCESS_KEY=secretaccesskey \
	-e AWS_ACCESS_KEY_ID=accesskeyid \
	-e AWS_DEFAULT_REGION=us-east-1 \
	-it \
	--rm \
	--network=dynamornr-test \
	dynamornr \
	/bin/ash \
	-c "cd test && dynamornr tables:list"
dynamo:
	docker run \
	-it \
	-p 8001:8000 \
	--rm \
	--name ddb \
	--network=dynamornr-test \
	peopleperhour/dynamodb
tables:
	make compile \
	&& docker run \
	-v $(shell pwd):/go/src/github.com/nycdavid/dynamornr \
	-e ENV=test \
	-e AWS_SECRET_ACCESS_KEY=secretaccesskey \
	-e AWS_ACCESS_KEY_ID=accesskeyid \
	-e AWS_DEFAULT_REGION=us-east-1 \
	-it \
	--rm \
	--network=dynamornr-test \
	dynamornr \
	/bin/ash \
	-c "cd test && TABLENAME=users dynamornr tables:create"
