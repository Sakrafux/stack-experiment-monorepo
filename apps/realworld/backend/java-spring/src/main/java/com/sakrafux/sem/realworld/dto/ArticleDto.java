package com.sakrafux.sem.realworld.dto;

import jakarta.validation.constraints.NotNull;
import java.time.LocalDateTime;
import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@Builder
@NoArgsConstructor
@AllArgsConstructor
public class ArticleDto {

    @NotNull
    private String slug;
    @NotNull
    private String title;
    @NotNull
    private String description;
    @NotNull
    private String body;
    @NotNull
    private String[] tagList;
    @NotNull
    private LocalDateTime createdAt;
    @NotNull
    private LocalDateTime updatedAt;
    @NotNull
    private boolean favorited;
    @NotNull
    private int favoritesCount;
    @NotNull
    private ProfileDto author;

}
