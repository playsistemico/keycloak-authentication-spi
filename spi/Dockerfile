FROM quay.io/keycloak/keycloak:24.0.5

USER root
COPY plugins/*.jar /opt/keycloak/providers/
USER 1000

ENTRYPOINT ["/opt/keycloak/bin/kc.sh", "start-dev"]
