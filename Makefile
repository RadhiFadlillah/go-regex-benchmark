run-all:
	@echo "modernc/benchmark"
	@modernc/benchmark input-text.txt

	@echo
	@echo "grafana/benchmark"
	@grafana/benchmark input-text.txt

	@echo
	@echo "hyperscan/benchmark"
	@hyperscan/benchmark input-text.txt

	@echo
	@echo "regexp2/benchmark"
	@regexp2/benchmark input-text.txt

	@echo
	@echo "stdgo/benchmark"
	@stdgo/benchmark input-text.txt

	@echo
	@echo "pcre/benchmark"
	@pcre/benchmark input-text.txt

	@echo
	@echo "pcre-go/benchmark"
	@pcre-go/benchmark input-text.txt

	@echo
	@echo "re2/benchmark-wasm"
	@re2/benchmark-wasm input-text.txt

	@echo
	@echo "re2/benchmark-cgo"
	@re2/benchmark-cgo input-text.txt

build-all:
	@make -C grafana
	@make -C hyperscan
	@make -C modernc
	@make -C pcre
	@make -C pcre-go
	@make -C re2
	@make -C regexp2
	@make -C stdgo