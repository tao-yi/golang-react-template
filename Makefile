.PHONY: dev
dev:
	air -d

.PHONY: buildClient
buildClient:
	cd client && yarn build && cd - 