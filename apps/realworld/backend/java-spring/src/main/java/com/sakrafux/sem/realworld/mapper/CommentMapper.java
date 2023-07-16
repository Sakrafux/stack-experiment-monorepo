package com.sakrafux.sem.realworld.mapper;

import com.sakrafux.sem.realworld.dto.CommentDto;
import com.sakrafux.sem.realworld.dto.NewCommentDto;
import com.sakrafux.sem.realworld.entity.Comment;
import org.mapstruct.Mapper;
import org.mapstruct.Mapping;

@Mapper
public interface CommentMapper {

    Comment newDtoToEntity(NewCommentDto comment);

    @Mapping(target = "author", ignore = true)
    CommentDto entityToDto(Comment comment);

}
