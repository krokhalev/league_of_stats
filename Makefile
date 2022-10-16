.PHONY: init_web
init_web:
	docker run \
		--rm \
		-w /app \
		-v `pwd`:/app \
		node:16.15 bash -c "npm install -g pnpm@7.1.0 && pnpm create vite web -- --template vue-ts && cd web && pnpm install"
	sudo chmod -R 777 web

.PHONY: start_web
start_web:
	docker-compose up web

.PHONY: clean
clean:
	sudo rm -rf .pnpm-store/
	sudo rm -rf web/
	sudo rm .pnpm-debug.log
