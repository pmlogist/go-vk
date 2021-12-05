test:
	pwd
dev:
	@docker run -it --rm \
		-w "/go/src/github.com/hnktt/go-vk" \
		-v $(PWD):/go/src/github.com/hnktt/go-vk \
		-p 9090:9090 \
		cosmtrek/air

build:
	@echo "Not implemented";
