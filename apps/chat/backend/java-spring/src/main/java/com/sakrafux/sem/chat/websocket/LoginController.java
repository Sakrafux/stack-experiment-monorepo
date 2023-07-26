package com.sakrafux.sem.chat.websocket;

import lombok.RequiredArgsConstructor;
import org.springframework.messaging.simp.SimpMessagingTemplate;
import org.springframework.stereotype.Controller;

@Controller
@RequiredArgsConstructor
public class LoginController {

    private final SimpMessagingTemplate simpMessagingTemplate;

    public void sendLoginMessage(String username) {
        simpMessagingTemplate.convertAndSend("/topic/login", username);
    }

}
