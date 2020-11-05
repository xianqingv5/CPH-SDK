
RM ?= rm -rf
GOBUILD = go build
GOTEST = go test
GOGET = go get -u
MASTER = bin/tmaster
WORKER = bin/tworker
GUARD = bin/tguard

VARS=vars.mk
$(shell ./build_config ${VARS})
include ${VARS}

.PHONY: main deps test bench clean

main:
	#${GOBUILD} -o ${GUARD} src/tworker_guard.go
	${GOBUILD} -o ${WORKER} src/worker.go

deps:
	${GOGET} github.com/aws/aws-sdk-go
	${GOGET} github.com/golang/snappy
	${GOGET} github.com/xuanxinhuiqing/gotools
	${GOGET} github.com/xuanxinhuiqing/godnf
	${GOGET} github.com/garyburd/redigo/redis
	${GOGET} gopkg.in/Shopify/sarama.v2

test:
	./auto_test.sh

bench:
	@pushd src/util > /dev/null && ${GOTEST} -bench=. && popd > /dev/null

clean:
	${RM} ${VARS} bin/*
