package com.sakrafux.sem.realworld.service;

import com.sakrafux.sem.realworld.dto.UpdateUserDto;
import com.sakrafux.sem.realworld.dto.UserDto;
import com.sakrafux.sem.realworld.exception.response.GenericErrorResponseException;
import com.sakrafux.sem.realworld.exception.response.UnauthorizedResponseException;

public interface UserService {

    UserDto login(String email, String password) throws UnauthorizedResponseException;

    UserDto createUser(String username, String email, String password) throws
        GenericErrorResponseException;

    UserDto getCurrentUser();

    UserDto updateUser(UpdateUserDto updateUserDto) throws GenericErrorResponseException;

}
