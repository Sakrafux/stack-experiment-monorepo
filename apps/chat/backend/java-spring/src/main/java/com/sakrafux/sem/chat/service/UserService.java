package com.sakrafux.sem.chat.service;

import com.google.api.client.googleapis.auth.oauth2.GoogleIdToken;
import com.sakrafux.sem.chat.dto.UserDto;
import java.util.List;

public interface UserService {

    void createUserIfNotExists(GoogleIdToken.Payload payload);

    List<UserDto> getAllUsers();

}
