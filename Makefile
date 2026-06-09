.PHONY: test tidy check-deps build-examples

test:
	@for mod in $$(find . -name go.mod | xargs -n1 dirname | sort); do \
		echo "==> test $$mod"; (cd $$mod && go test ./...) || exit 1; \
	done

tidy:
	@for mod in $$(find . -name go.mod | xargs -n1 dirname | sort); do \
		echo "==> tidy $$mod"; (cd $$mod && go mod tidy); \
	done

check-deps:
	@bash scripts/check-deps.sh

build-examples:
	cd examples && go build ./...
