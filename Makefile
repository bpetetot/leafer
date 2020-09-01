all: build-webapp build-server pack-final-build

build-webapp:
	(cd leafer-web && yarn build)

build-server:
	(cd leafer-app && go build)

pack-final-build:
	mkdir -p ./build
	mv leafer-web/build ./build/web
	mv leafer-app/leafer ./build
	cp leafer-app/.env ./build