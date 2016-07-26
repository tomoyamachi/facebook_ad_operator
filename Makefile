lint:
	@for dir in $$(glide novendor); do \
	golint $$dir; \
	done;
