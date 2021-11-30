GOALIAS ?= go
COVER_OUT := coverage.out
COVER_HTML := coverage.html

run:
	$(GOALIAS) run app\cmd\main.go

test:
	$(GOALIAS) test ./app/cmd
	$(GOALIAS) test ./domain/entities
	$(GOALIAS) test ./domain/services
	