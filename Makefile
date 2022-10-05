Go := go

.PHONY: generate
generate:
	@echo "  >  "$@"ing $(DESTDIR)"
	${MAKE} -C api generate

.PHONY: all
all:
	@echo "  >  "$@"ing $(DESTDIR)"
	@mkdir -p pack/bin
	${MAKE} -C cmd/backend
	${MAKE} -C cmd/recall
	${MAKE} -C cmd/predict
	${MAKE} -C cmd/userpref
	@cp -Rv cmd/backend/recsys-backend pack/bin/
	@cp -Rv cmd/predict/recsys-predict pack/bin/
	@cp -Rv cmd/recall/recsys-recall pack/bin/
	@cp -Rv cmd/userpref/recsys-userpref pack/bin/
