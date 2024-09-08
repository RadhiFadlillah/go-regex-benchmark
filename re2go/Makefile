build: generate
	@CGO_ENABLED=0 go build -o benchmark

generate:
	@for name in *.re; do \
		RE_IN=$$name; \
		RE_OUT=$$(echo $$name | sed 's/.re/.go/'); \
		re2go --input-encoding utf8 -i $$RE_IN -o $$RE_OUT; \
		gofmt -w $$RE_OUT; \
	done