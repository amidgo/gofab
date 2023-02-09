build:
	go build -ldflags="-s -w"
move-to-ledger:
	make build
	cp gofab ../amidledger/test-network
