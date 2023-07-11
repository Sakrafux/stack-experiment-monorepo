package com.sakrafux.sem.realworld.mapper;

import com.sakrafux.sem.realworld.dto.UpdateUserDto;
import com.sakrafux.sem.realworld.dto.UserDto;
import com.sakrafux.sem.realworld.entity.ApplicationUser;
import org.mapstruct.Mapper;
import org.mapstruct.MappingTarget;
import org.mapstruct.NullValuePropertyMappingStrategy;

@Mapper(nullValuePropertyMappingStrategy = NullValuePropertyMappingStrategy.IGNORE)
public interface UserMapper {

    UserDto entityToDto(ApplicationUser user, String token);

    void updateUserDtoToEntity(UpdateUserDto userDto, @MappingTarget ApplicationUser user);

}
