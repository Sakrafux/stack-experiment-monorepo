package com.sakrafux.sem.chat.endpoint;

import com.sakrafux.sem.chat.dto.MessageDto;
import com.sakrafux.sem.chat.service.MessageService;
import java.time.LocalDateTime;
import java.time.format.DateTimeFormatter;
import java.util.List;
import lombok.RequiredArgsConstructor;
import org.springframework.security.access.annotation.Secured;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
@RequestMapping("/api/message")
@RequiredArgsConstructor
public class MessageEndpoint {

    private final MessageService messageService;

    @Secured("ROLE_USER")
    @GetMapping("/{chatId}")
    public List<MessageDto> getMessages(@PathVariable Long chatId) {
        return messageService.getMessagesByChatId(chatId);
    }

    @Secured("ROLE_USER")
    @GetMapping("/{chatId}/{createdAt}")
    public List<MessageDto> getMoreMessages(@PathVariable Long chatId, @PathVariable String createdAt) {
        var parsedCreatedAt = LocalDateTime.parse(createdAt, DateTimeFormatter.ISO_LOCAL_DATE_TIME);
        parsedCreatedAt = parsedCreatedAt.plusHours(2);
        return messageService.getMoreMessagesByChatId(chatId, parsedCreatedAt);
    }

}
