# keycloak-authorization-spi
Simple authorization SPI for Keycloak

- [Install dependencies]

### Useful links
- [Create Keycloak plugin](https://dev.to/yakovlev_alexey/how-to-create-a-keycloak-plugin-3acj)
- [Keycloak auth SPI docs](https://wjw465150.gitbooks.io/keycloak-documentation/content/server_development/topics/auth-spi.html)
- [Keycloak authenticators repo](https://github.com/keycloak/keycloak/tree/main/services/src/main/java/org/keycloak/authentication/authenticators)
- [Keycloak SPI example](https://github.com/zene22/keycloak-spi-example)



# Install dependencies
## General
1. Install Docker
2. Install make
## For SPI development
1. Install JavaJDK
2. Install Maven
    1. Download the `.tar.gz` from [Maven's download page](https://maven.apache.org/download.cgi)
    2. Extract with with `tar -xvf apache-maven-<version>.tar.gz`
    3. Add the env to `.zshrc` file
       ```
       M2_HOME=/Users/<user>/Downloads/apache-maven-<version>
       export PATH=$PATH:$M2_HOME/bin
       ```
    4. Refresh `.zshrc` with `source .zshrc`
3. Install dependencies
   ```
   cd spi/keycloak-authentication-spi
   mvn install
   ```
## For backend development
1. Install Go

## Build SPI
```
make build-spi
```

## Run all
```
docker compose up -d
```


