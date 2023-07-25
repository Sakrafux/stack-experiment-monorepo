package com.sakrafux.sem.chat.service.impl;

import com.google.api.client.googleapis.auth.oauth2.GoogleIdToken;
import com.sakrafux.sem.chat.dto.UserDto;
import com.sakrafux.sem.chat.entity.ApplicationUser;
import com.sakrafux.sem.chat.mapper.UserMapper;
import com.sakrafux.sem.chat.repository.ApplicationUserRepository;
import com.sakrafux.sem.chat.service.UserService;
import com.sakrafux.sem.chat.util.AuthUtil;
import jakarta.transaction.Transactional;
import java.util.List;
import lombok.RequiredArgsConstructor;
import org.springframework.stereotype.Service;

@Service
@RequiredArgsConstructor
public class UserServiceImpl implements UserService {

    private final ApplicationUserRepository applicationUserRepository;

    private final UserMapper userMapper;

    @Override
    @Transactional
    public void createUserIfNotExists(GoogleIdToken.Payload payload) {
        var userId = payload.getSubject();
        var name = (String) payload.get("name");
        var pictureUrl = (String) payload.get("picture");

        if (applicationUserRepository.findByGid(userId).isEmpty()) {
            var user = new ApplicationUser();
            user.setGid(userId);
            user.setName(name);
            user.setPicture(pictureUrl);
            applicationUserRepository.save(user);
        }
    }

    @Override
    public List<UserDto> getAllUsers() {
        var userId = AuthUtil.getUserId();
        return applicationUserRepository.findAll().stream()
            .filter(user -> !user.getGid().equals(userId)).map(userMapper::toDto).toList();
    }
}
