#!/bin/bash
gocompat compare \
	--go1compat \
	--log-level=debug \
	--git-refs=bf95e1..e9400c \
	./...
