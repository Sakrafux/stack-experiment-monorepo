package com.sakrafux.sem.chat.endpoint;

import com.sakrafux.sem.chat.service.ChatService;
import lombok.RequiredArgsConstructor;
import org.springframework.security.access.annotation.Secured;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
@RequestMapping("/api/chat")
@RequiredArgsConstructor
public class ChatEndpoint {

    private final ChatService chatService;

    @Secured("ROLE_USER")
    @PostMapping("/{userId}")
    public Long establishChat(@PathVariable Long userId) {
        return chatService.establishChatIfNotExists(userId);
    }

}
