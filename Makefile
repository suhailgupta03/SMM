srcdir := plugins
plugins := $(shell find $(srcdir) -type d -mindepth 1)


all: $(plugins)

$(plugins):
	$(MAKE) -C $@

.PHONY: all $(plugins)