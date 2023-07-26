package com.sakrafux.sem.chat.service;

import com.sakrafux.sem.chat.dto.MessageDto;
import com.sakrafux.sem.chat.dto.NewMessageDto;
import java.util.List;

public interface MessageService {

    MessageDto sendMessage(NewMessageDto newMessageDto);

    List<MessageDto> getMessagesByChatId(Long chatId);

}
