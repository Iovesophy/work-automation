.PHONY: setup
setup: build
	./setup.sh

.PHONY: build
build: clean
	./build.sh

.PHONY: clean
clean:
	rm -f genkey attach detach manhours

.PHONY: genkey
genkey:
	./genkey

.PHONY: attach
attach:
	if [ "$(shell cat config/touch.log)" = "1" ] ; then \
		./attach && printf "0" > config/touch.log; \
	fi

.PHONY: detach
detach:
	if [ "$(shell cat config/touch.log)" = "0" ] ; then \
		./detach && printf "1" > config/touch.log; \
	fi
	./manhours
