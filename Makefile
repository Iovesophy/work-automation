.PHONY: setup
setup: build
	./tools/setup.sh
	./genkey

.PHONY: build
build: clean
	./tools/build.sh

.PHONY: clean
clean:
	rm -rf genkey attach detach manhours config/* ~/.ssh/work

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
