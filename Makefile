TOOLS_IMG=league_of_stats/web

.PHONY: init_web
init_web:
	docker run \
		--rm \
		-w /app \
		-v `pwd`:/app \
		node:16.15 bash -c "npm install -g pnpm@7.1.0 && pnpm create vite web -- --template vue-ts"

.PHONY: start_web
start_web:
	docker-compose up web

.PHONY: los_start
los_start:
	docker-compose --profile email up -V

.PHONY: gen
gen:
	docker build -f Dockerfile.web_init -t ${TOOLS_IMG} .

.PHONY: init
init:
	docker build -f Dockerfile.web_init -t ${TOOLS_IMG} .
	docker run \
		--rm \
		-v `pwd`:/app \
		${TOOLS_IMG}
