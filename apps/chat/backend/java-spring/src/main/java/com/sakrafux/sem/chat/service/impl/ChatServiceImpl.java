package com.sakrafux.sem.chat.service.impl;

import com.sakrafux.sem.chat.entity.Chat;
import com.sakrafux.sem.chat.repository.ApplicationUserRepository;
import com.sakrafux.sem.chat.repository.ChatRepository;
import com.sakrafux.sem.chat.service.ChatService;
import com.sakrafux.sem.chat.util.AuthUtil;
import lombok.RequiredArgsConstructor;
import org.springframework.stereotype.Service;

@Service
@RequiredArgsConstructor
public class ChatServiceImpl implements ChatService {

    private final ApplicationUserRepository applicationUserRepository;
    private final ChatRepository chatRepository;

    @Override
    public Long establishChatIfNotExists(Long userId) {
        var myUserId = AuthUtil.getUserId();
        var user = applicationUserRepository.findByGid(myUserId).orElseThrow();

        var chat = chatRepository.findByUser1IdAndUser2Id(user.getId(), userId)
            .orElseGet(
                () -> chatRepository.findByUser1IdAndUser2Id(userId, user.getId()).orElseGet(() -> {
                    var newChat = new Chat();
                    newChat.setUser1(user);
                    newChat.setUser2(applicationUserRepository.findById(userId).orElseThrow());
                    return chatRepository.save(newChat);
                }));
        return chat.getId();
    }
}
