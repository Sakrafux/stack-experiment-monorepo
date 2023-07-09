package com.sakrafux.sem.realworld.dto.response;

import com.sakrafux.sem.realworld.dto.CommentDto;
import jakarta.validation.constraints.NotNull;
import java.util.List;
import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@Builder
@NoArgsConstructor
@AllArgsConstructor
public class MultipleCommentsResponseDto {

    @NotNull
    private List<CommentDto> comments;

}
