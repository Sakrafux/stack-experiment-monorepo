package com.sakrafux.sem.realworld.dto.response;

import com.sakrafux.sem.realworld.dto.CommentDto;
import jakarta.validation.constraints.NotNull;
import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@Builder
@NoArgsConstructor
@AllArgsConstructor
public class SingleCommentResponseDto {

    @NotNull
    private CommentDto comment;

}
