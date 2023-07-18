package com.sakrafux.sem.realworld.mapper;

import com.sakrafux.sem.realworld.dto.ProfileDto;
import com.sakrafux.sem.realworld.entity.ApplicationUser;
import org.mapstruct.Mapper;
import org.mapstruct.Mapping;

@Mapper
public interface ProfileMapper {

    @Mapping(source = "following", target = "following")
    @Mapping(source = "user.image", target = "image", defaultValue = "")
    ProfileDto entityToDto(ApplicationUser user, boolean following);

}
