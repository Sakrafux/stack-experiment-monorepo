package com.sakrafux.sem.realworld.dto;

import com.fasterxml.jackson.annotation.JsonFormat;
import jakarta.validation.constraints.NotNull;
import java.time.LocalDateTime;
import java.util.List;
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
    private List<String> tagList;
    @NotNull
    @JsonFormat(pattern = "yyyy-MM-dd'T'HH:mm:ss.SSS'Z'")
    private LocalDateTime createdAt;
    @NotNull
    @JsonFormat(pattern = "yyyy-MM-dd'T'HH:mm:ss.SSS'Z'")
    private LocalDateTime updatedAt;
    @NotNull
    private boolean favorited;
    @NotNull
    private int favoritesCount;
    @NotNull
    private ProfileDto author;

}
