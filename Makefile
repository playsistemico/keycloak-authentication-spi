all: build-spi run

build-spi:
	cd spi/keycloak-authentication-spi && \
	mvn clean package
	mv ./spi/keycloak-authentication-spi/target/keycloak-authentication-spi-1.0-SNAPSHOT.jar ./spi/plugins/

run:
	docker compose up -d --build auth-server
