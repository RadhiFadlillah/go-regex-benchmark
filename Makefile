run-all:
	@for benchmark in **/benchmark*; do \
		echo $$benchmark; \
		./$$benchmark "input-text.txt"; \
		echo; \
	done

build-all:
	@make -C grafana
	@make -C hyperscan
	@make -C modernc
	@make -C pcre
	@make -C re2
	@make -C re2go
	@make -C regexp2
	@make -C stdgo