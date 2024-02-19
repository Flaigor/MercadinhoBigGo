GOALIAS ?= go
COVER_OUT := coverage.out
COVER_HTML := coverage.html

run:
	$(GOALIAS) run app\cmd\main.go

test:
	$(GOALIAS) test ./app/cmd -v
	$(GOALIAS) test ./domain/entities -v
	$(GOALIAS) test ./domain/services -v
	