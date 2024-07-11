package modo.keycloak.authentication.authenticator.browser;

import org.jboss.logging.Logger;
import org.keycloak.authentication.AuthenticationFlowContext;
import org.keycloak.authentication.Authenticator;
import org.keycloak.broker.provider.util.SimpleHttp;
import org.keycloak.events.Errors;
import org.keycloak.models.KeycloakSession;
import org.keycloak.models.RealmModel;
import org.keycloak.models.UserModel;
import org.keycloak.models.utils.KeycloakModelUtils;

import com.fasterxml.jackson.databind.JsonNode;

import jakarta.ws.rs.core.MultivaluedMap;

public class ValidateSession implements Authenticator {
    private static final Logger logger = Logger.getLogger(ValidateSession.class);

    @Override
    public void close() {
    }

    @Override
    public void authenticate(AuthenticationFlowContext context) {
        MultivaluedMap<String, String> inputData = context.getHttpRequest().getUri().getQueryParameters();
        String session = inputData.getFirst("session");
        if (session == null) {
            context.attempted();
            return;
        }
        String username = getUserFromSession(context, session);
        if (username == null) {
            context.attempted();
            return;
        }

        UserModel user = KeycloakModelUtils.findUserByNameOrEmail(context.getSession(), context.getRealm(), username);
        if (user == null){
            context.getEvent().error(Errors.USER_NOT_FOUND);
            context.attempted();
            return;
        }

        context.setUser(user);
        context.success();
    }

    @Override
    public void action(AuthenticationFlowContext context) {
    }

    @Override
    public boolean requiresUser() {
        return false;
    }

    @Override
    public boolean configuredFor(KeycloakSession session, RealmModel realm, UserModel user) {
        return true;
    }

    @Override
    public void setRequiredActions(KeycloakSession session, RealmModel realm, UserModel user) {
    }

    private String getUserFromSession(AuthenticationFlowContext context, String session){
        SimpleHttp httpResp = SimpleHttp.doGet("http://backend:8080/user?session="+session, context.getSession());

        try {
            JsonNode resp;
            resp = httpResp.asJson();
            logger.info("API RESPONSE: "+resp.toString());
            String username = resp.get("username").asText();
            return username;

        } catch (Exception e) {
            logger.info("API ERROR: "+e.toString());
            return null;
        }
    }
}
