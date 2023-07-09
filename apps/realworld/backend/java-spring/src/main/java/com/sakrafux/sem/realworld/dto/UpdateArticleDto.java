package com.sakrafux.sem.realworld.dto;

import lombok.Builder;
import lombok.Data;

@Data
@Builder
public class UpdateArticleDto {

    private String title;
    private String description;
    private String body;

}
