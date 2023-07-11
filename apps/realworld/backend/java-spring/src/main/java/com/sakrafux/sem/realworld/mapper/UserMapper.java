package com.sakrafux.sem.realworld.mapper;

import com.sakrafux.sem.realworld.dto.UserDto;
import com.sakrafux.sem.realworld.entity.ApplicationUser;
import org.mapstruct.Mapper;

@Mapper
public interface UserMapper {

    UserDto entityToDto(ApplicationUser user, String token);

}
