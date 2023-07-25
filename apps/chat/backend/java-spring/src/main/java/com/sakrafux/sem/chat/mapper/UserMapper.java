package com.sakrafux.sem.chat.mapper;

import com.sakrafux.sem.chat.dto.UserDto;
import com.sakrafux.sem.chat.entity.ApplicationUser;
import org.mapstruct.Mapper;

@Mapper
public interface UserMapper {

    UserDto toDto(ApplicationUser user);

}
