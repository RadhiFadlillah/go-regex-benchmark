run-all:
	@for benchmark in **/benchmark*; do \
		echo $$benchmark; \
		./$$benchmark "input-text.txt"; \
		echo; \
	done

clean-all:
	@for benchmark in **/benchmark*; do \
		rm $$benchmark; \
	done

build-all:
	@make -C codesearch
	@make -C grafana
	@make -C hyperscan
	@make -C matloob
	@make -C modernc
	@make -C pcre
	@make -C re2
	@make -C re2go
	@make -C regexp2
	@make -C regexp2c
	@make -C regexp2go
	@make -C stdgo