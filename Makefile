.PHONY: setup
setup: build
	./tools/setup.sh
	./genkey

.PHONY: build
build: clean
	./tools/build.sh

.PHONY: clean
clean:
	rm -rf genkey attach detach manhours config ~/.ssh/work

.PHONY: genkey
genkey:
	./genkey

.PHONY: attach
attach:
	./attach

.PHONY: detach
detach:
	./detach
	./manhours

.PHONY: shortcut
shortcut: setup
	rm -rf ~/config ~/Desktop/attach ~/Desktop/detach ~/Desktop/manhours
	cp -r config ~
	cp attach detach manhours ~/Desktop
