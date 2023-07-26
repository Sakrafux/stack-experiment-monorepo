package com.sakrafux.sem.chat.service.impl;

import com.sakrafux.sem.chat.dto.MessageDto;
import com.sakrafux.sem.chat.dto.NewMessageDto;
import com.sakrafux.sem.chat.entity.Message;
import com.sakrafux.sem.chat.mapper.MessageMapper;
import com.sakrafux.sem.chat.repository.ApplicationUserRepository;
import com.sakrafux.sem.chat.repository.ChatRepository;
import com.sakrafux.sem.chat.repository.MessageRepository;
import com.sakrafux.sem.chat.service.MessageService;
import com.sakrafux.sem.chat.util.AuthUtil;
import java.util.List;
import lombok.RequiredArgsConstructor;
import org.springframework.stereotype.Service;

@Service
@RequiredArgsConstructor
public class MessageServiceImpl implements MessageService {

    private final ApplicationUserRepository applicationUserRepository;
    private final ChatRepository chatRepository;
    private final MessageRepository messageRepository;

    private final MessageMapper messageMapper;

    @Override
    public MessageDto sendMessage(NewMessageDto newMessageDto) {
        var userId = AuthUtil.getUserId();
        var user = applicationUserRepository.findByGid(userId).orElseThrow();

        var message = Message.builder()
                .chat(chatRepository.getReferenceById(newMessageDto.getChatId()))
                .user(user)
                .text(newMessageDto.getText())
                .build();

        var result = messageRepository.save(message);
        return messageMapper.toDto(result);
    }

    @Override
    public List<MessageDto> getMessagesByChatId(Long chatId) {
        return messageRepository.findAllByChatIdOrderByCreatedAtDesc(chatId).stream()
                .map(messageMapper::toDto)
                .toList();
    }
}
