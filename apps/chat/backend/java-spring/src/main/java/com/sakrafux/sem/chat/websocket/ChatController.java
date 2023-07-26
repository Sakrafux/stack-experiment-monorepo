package com.sakrafux.sem.chat.websocket;

import com.google.api.client.googleapis.auth.oauth2.GoogleIdTokenVerifier;
import com.google.api.client.http.javanet.NetHttpTransport;
import com.google.api.client.json.gson.GsonFactory;
import com.sakrafux.sem.chat.dto.MessageDto;
import com.sakrafux.sem.chat.dto.NewMessageDto;
import com.sakrafux.sem.chat.security.SecurityProperties;
import com.sakrafux.sem.chat.service.MessageService;
import java.io.IOException;
import java.security.GeneralSecurityException;
import java.util.Collections;
import java.util.List;
import lombok.RequiredArgsConstructor;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.messaging.handler.annotation.Header;
import org.springframework.messaging.handler.annotation.MessageExceptionHandler;
import org.springframework.messaging.handler.annotation.MessageMapping;
import org.springframework.messaging.handler.annotation.Payload;
import org.springframework.messaging.handler.annotation.SendTo;
import org.springframework.messaging.simp.annotation.SendToUser;
import org.springframework.security.authentication.UsernamePasswordAuthenticationToken;
import org.springframework.security.core.authority.SimpleGrantedAuthority;
import org.springframework.security.core.context.SecurityContextHolder;
import org.springframework.stereotype.Controller;

@Controller
@RequiredArgsConstructor
public class ChatController {

    @Value("${appl.google.client-id}")
    private String clientId;

    private final MessageService messageService;

    // is prefixed with /ws in WebSocketConfig
    @MessageMapping("/chat/{chatId}")
    @SendTo("/topic/chat/{chatId}")
    // could extract {chatId} with @DestinationVariable
    public MessageDto sendMessage(@Payload NewMessageDto newMessageDto, @Header("Authorization") String token)
        throws GeneralSecurityException, IOException {

        var verifier = new GoogleIdTokenVerifier.Builder(new NetHttpTransport(), new GsonFactory())
            .setAudience(Collections.singletonList(clientId))
            .build();
        var idToken = verifier.verify(token.replace(SecurityProperties.PREFIX, ""));

        if (idToken == null) {
            throw new GeneralSecurityException("Invalid token");
        }

        var payload = idToken.getPayload();
        var userId = payload.getSubject();
        var authorities = List.of(new SimpleGrantedAuthority("ROLE_USER"));
        SecurityContextHolder.getContext().setAuthentication(new UsernamePasswordAuthenticationToken(userId, null, authorities));

        return messageService.sendMessage(newMessageDto);
    }

    @MessageExceptionHandler
    @SendToUser("/topic/errors/chat")
    // could handle specific exceptions here
    public String handleException(Throwable exception) {
        return exception.getMessage();
    }

}
