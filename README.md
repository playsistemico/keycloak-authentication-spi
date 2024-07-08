# keycloak-authorization-spi
Simple authorization SPI for Keycloak

## Run all
```
docker compose up
```

## Running Keycloak
```
docker run -d -p 9090:8080 \
 --name pcp-auth-server \
  -e KEYCLOAK_ADMIN=admin \
  -e KEYCLOAK_ADMIN_PASSWORD=admin \
  quay.io/keycloak/keycloak:24.0.5 start-dev
```
