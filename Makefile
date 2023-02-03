all: TestSequential TestParallelBasic TestRPCStatus TestHTTPStatus TestParallelSumHTTP

TestSequential:
	cd sequential/; \
	export HOME="${CURDIR}"; \
    chmod 0777 ./test.sh; \
    ./test.sh


TestParallelBasic:
	cd parallel/; \
	export HOME="${CURDIR}"; \
	chmod 0777 ./test.sh; \
    ./test.sh

TestRPCStatus:
	cd http_parallel/; \
	export HOME="${CURDIR}"; \
	chmod 0777 ./test_rpc.sh; \
	./test_rpc.sh;

TestHTTPStatus:
	cd http_parallel/; \
	export HOME="${CURDIR}"; \
	chmod 0777 ./test_http.sh; \
	./test_http.sh

TestParallelSumHTTP:
	cd http_parallel/; \
	export HOME="${CURDIR}"; \
	chmod 0777 ./test.sh; \
    ./test.sh

