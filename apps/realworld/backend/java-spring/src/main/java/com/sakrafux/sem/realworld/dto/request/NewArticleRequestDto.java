package com.sakrafux.sem.realworld.dto.request;

import com.sakrafux.sem.realworld.dto.NewArticleDto;
import jakarta.validation.constraints.NotNull;
import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@Builder
@NoArgsConstructor
@AllArgsConstructor
public class NewArticleRequestDto {

    @NotNull
    private NewArticleDto article;

}
