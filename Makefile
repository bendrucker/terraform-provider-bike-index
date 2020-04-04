.PHONY: test testacc

test:
	go test $(TESTARGS) -timeout=30s -parallel=4
testacc:
	TF_ACC=1 go test $(TEST) -v $(TESTARGS) -timeout 15m


