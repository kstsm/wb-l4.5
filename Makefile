.PHONY: run linter bench-before bench-after bench-compare profile-cpu profile-mem profile-trace trace-analyze

BENCH_DIR=benchmarks

# Start
run:
	go run main.go

# Linter
linter:
	go run github.com/golangci/golangci-lint/v2/cmd/golangci-lint@v2.6.2 run

# BENCHMARKS
bench-before:
	mkdir -p $(BENCH_DIR)
	go test -bench=BenchmarkConcatenate -benchmem -count=6 ./internal/service > $(BENCH_DIR)/before.txt

bench-after:
	mkdir -p $(BENCH_DIR)
	go test -bench=BenchmarkConcatenate -benchmem -count=6 ./internal/service > $(BENCH_DIR)/after.txt

bench-compare:
	benchstat $(BENCH_DIR)/both.txt

# PPROF CPU & HEAP
profile-cpu:
	go tool pprof http://localhost:6060/debug/pprof/profile?seconds=10

profile-mem:
	go tool pprof http://localhost:6060/debug/pprof/heap

# TRACE
profile-trace:
	curl -o $(BENCH_DIR)/trace.out http://localhost:6060/debug/pprof/trace?seconds=10

trace-analyze:
	go tool trace $(BENCH_DIR)/trace.out
