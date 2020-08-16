# Go params
GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=hselect

all: build
build: 
	$(GOBUILD) -v
	@echo 'done'

test:
	$(GOTEST) -v -race ./... #run tests (TestXxx) excluding benchmarks
	@echo 'done'

bench:
	$(GOTEST) -v -race -run=XXX -bench=. ./... #run all benchmarks (BenchmarkXxx)
	@echo 'done'

test-bench:
	$(GOTEST) -v -race -bench=. ./... #run all tests and benchmarks (TestXxx and BenchmarkXxx)
	@echo 'done'

clean: 
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	@echo 'done'

install:
	if [ ! -d /usr/local/hselect ] ; then \
		sudo mkdir /usr/local/hselect; \
		sudo mkdir /usr/local/hselect/bin; \
	fi
	sudo cp hselect /usr/local/hselect/bin/
	sudo rm -f hselect
	@echo 'done'
