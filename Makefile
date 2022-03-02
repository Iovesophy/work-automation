.PHONY: startup
startup: setup genkey
	mkdir -p ~/.ssh/work
	cd cmd/genkey; go run main.go
	cd cmd/do-login; go run main.go

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
	cd cmd/attach; go run main.go

.PHONY: detach
detach:
	cd cmd/detach; go run main.go
	cd cmd/manhours; go run main.go
