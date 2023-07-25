package com.sakrafux.sem.chat.endpoint;

import com.sakrafux.sem.chat.dto.NewMessageDto;
import com.sakrafux.sem.chat.service.MessageService;
import lombok.RequiredArgsConstructor;
import org.springframework.security.access.annotation.Secured;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
@RequestMapping("/api/message")
@RequiredArgsConstructor
public class MessageEndpoint {

    private final MessageService messageService;

    @Secured("ROLE_USER")
    @PostMapping
    public void sendMessage(@RequestBody NewMessageDto newMessageDto) {
        messageService.sendMessage(newMessageDto);
    }

}
