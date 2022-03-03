.PHONY: startup
startup: setup genkey
	mkdir -p ~/.ssh/work
	cd cmd/genkey; go run main.go
	cd cmd/do-login; go run main.go
	echo 1 > config/touch.log

.PHONY: setup
setup:
	./setup.sh

.PHONY: genkey
genkey:
	cd cmd/genkey; go run main.go

.PHONY: login
login:
	cd cmd/do-login; go run main.go

.PHONY: attach
attach:
	if [ "$(shell cat config/touch.log)" = "1" ] ; then \
		cd cmd/attach; go run main.go && echo 0 > config/touch.log \
	fi

.PHONY: detach
detach:
	if [ "$(shell cat config/touch.log)" = "0" ] ; then \
		cd cmd/detach; go run main.go && echo 1 > config/touch.log \
	fi
	cd cmd/manhours; go run main.go
