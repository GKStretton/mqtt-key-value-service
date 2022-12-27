#!/bin/bash
# Generate mocks for interfaces

mockgen \
	-source "./storage/storage.go" \
	-destination "./storage/mock.storage.go" \
	-package "storage"

mockgen \
	-source "./pubsub/pubsub.go" \
	-destination "./pubsub/mock.pubsub.go" \
	-package "pubsub"
