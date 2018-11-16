EXEC=$(shell basename "$(shell pwd)")
all: $(EXEC) data/items.json data/locations.json

$(EXEC):
	go build

data/items.json data/locations.json:
	./convert

.PHONY: $(EXEC)
