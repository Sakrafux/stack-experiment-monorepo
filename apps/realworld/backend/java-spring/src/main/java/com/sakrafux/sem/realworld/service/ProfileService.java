package com.sakrafux.sem.realworld.service;

import com.sakrafux.sem.realworld.dto.ProfileDto;
import com.sakrafux.sem.realworld.exception.response.NotFoundResponseException;

public interface ProfileService {

    ProfileDto getProfileByUsername(String username) throws NotFoundResponseException;

    ProfileDto followUserByUsername(String username) throws NotFoundResponseException;

    ProfileDto unfollowUserByUsername(String username) throws NotFoundResponseException;

}
