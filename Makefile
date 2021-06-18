.PHONY: deps proto

deps:
	go mod download

proto:
	@protoc  --go_out=. \
	--go_opt=Mpb/crx3.proto=./pb \
	./pb/crx3.proto
