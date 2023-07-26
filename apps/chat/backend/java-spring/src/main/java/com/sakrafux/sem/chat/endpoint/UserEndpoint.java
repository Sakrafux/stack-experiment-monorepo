package com.sakrafux.sem.chat.endpoint;

import com.google.api.client.googleapis.auth.oauth2.GoogleIdToken;
import com.google.api.client.googleapis.auth.oauth2.GoogleIdTokenVerifier;
import com.google.api.client.http.javanet.NetHttpTransport;
import com.google.api.client.json.gson.GsonFactory;
import com.sakrafux.sem.chat.dto.UserDto;
import com.sakrafux.sem.chat.security.SecurityProperties;
import com.sakrafux.sem.chat.service.UserService;
import com.sakrafux.sem.chat.websocket.LoginController;
import java.util.Collections;
import java.util.List;
import lombok.RequiredArgsConstructor;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.security.access.annotation.Secured;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestHeader;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
@RequestMapping("/api/user")
@RequiredArgsConstructor
public class UserEndpoint {

    private final UserService userService;

    private final LoginController loginController;

    @Value("${appl.google.client-id}")
    private String clientId;

    @Secured("ROLE_USER")
    @GetMapping("/login")
    public void login(@RequestHeader("Authorization") String token) {
        var verifier = new GoogleIdTokenVerifier.Builder(new NetHttpTransport(), new GsonFactory())
            .setAudience(Collections.singletonList(clientId))
            .build();

        GoogleIdToken idToken;
        try {
            idToken = verifier.verify(token.replace(SecurityProperties.PREFIX, ""));
        } catch (Exception e) {
            throw new IllegalArgumentException("Token is invalid");
        }

        if (idToken == null) {
            throw new IllegalArgumentException("Token is invalid");
        }

        var payload = idToken.getPayload();
        var name = (String) payload.get("name");

        userService.createUserIfNotExists(payload);

        loginController.sendLoginMessage(name);
    }

    @Secured("ROLE_USER")
    @GetMapping
    public List<UserDto> getAllUsers() {
        return userService.getAllUsers();
    }

}
