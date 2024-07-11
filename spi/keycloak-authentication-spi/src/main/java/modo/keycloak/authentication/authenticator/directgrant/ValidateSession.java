package modo.keycloak.authentication.authenticator.directgrant;

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
        MultivaluedMap<String, String> inputData = context.getHttpRequest().getDecodedFormParameters();
        String session = inputData.getFirst("session");
        logger.info("session from param: "+session);
        String username = getUserFromSession(context, session);
        if (username == "") {
            logger.info("user empty");
            context.getEvent().error(Errors.USER_NOT_FOUND);
            context.success();
            return;
        }

        UserModel user = KeycloakModelUtils.findUserByNameOrEmail(context.getSession(), context.getRealm(), username);
        logger.info("userID: "+user.getId());
        context.setUser(user);
        context.success();
    }

    @Override
    public void action(AuthenticationFlowContext context) {
        context.success();
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
        SimpleHttp httpResp = SimpleHttp.doGet("http://mockserver:8080/session?session="+session, context.getSession());
        JsonNode resp;
        try {
            resp = httpResp.asJson();
        } catch (Exception e) {
            logger.info("exception: "+e.toString());
            return "";
        }
        
        logger.info(resp.toString());
        return resp.get("username").asText();
    }
}