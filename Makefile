.PHONY: build

polybar_hue:
	CGO_ENABLED=0 go build -o polybar_hue -tags netgo -ldflags '-extldflags "-static"'

build: polybar_hue

install: polybar_hue
	cp -f polybar_hue ~/.config/polybar/scripts/hue	
