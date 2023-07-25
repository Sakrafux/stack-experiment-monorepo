package com.sakrafux.sem.chat.mapper;

import com.sakrafux.sem.chat.dto.MessageDto;
import com.sakrafux.sem.chat.entity.Message;
import org.mapstruct.Mapper;
import org.mapstruct.Mapping;

@Mapper
public interface MessageMapper {

    @Mapping(target = "userId", source = "user.id")
    MessageDto toDto(Message message);

    @Mapping(target = "user.id", source = "userId")
    Message toEntity(MessageDto messageDto);

}
