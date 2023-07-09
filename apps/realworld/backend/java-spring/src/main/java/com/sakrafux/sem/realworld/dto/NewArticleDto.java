package com.sakrafux.sem.realworld.dto;

import jakarta.validation.constraints.NotNull;
import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@Builder
@NoArgsConstructor
@AllArgsConstructor
public class NewArticleDto {

    @NotNull
    private String title;
    @NotNull
    private String description;
    @NotNull
    private String body;
    private String[] tagList;

}
