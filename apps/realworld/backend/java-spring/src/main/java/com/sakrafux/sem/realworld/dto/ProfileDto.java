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
public class ProfileDto {

    @NotNull
    private String username;
    @NotNull
    private String bio;
    @NotNull
    private String image;
    @NotNull
    private boolean following;

}
