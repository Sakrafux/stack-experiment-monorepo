package com.sakrafux.sem.chat.service;

import com.sakrafux.sem.chat.dto.NewMessageDto;

public interface MessageService {

    void sendMessage(NewMessageDto newMessageDto);

}
