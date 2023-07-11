package com.sakrafux.sem.realworld.service;

import com.sakrafux.sem.realworld.dto.UpdateUserDto;
import com.sakrafux.sem.realworld.dto.UserDto;
import com.sakrafux.sem.realworld.exception.response.GenericErrorResponseException;

public interface UserService {

    UserDto getCurrentUser();

    UserDto updateUser(UpdateUserDto updateUserDto) throws GenericErrorResponseException;

}
