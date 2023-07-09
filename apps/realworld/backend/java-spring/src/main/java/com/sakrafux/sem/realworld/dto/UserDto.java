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
public class UserDto {

    @NotNull
    private String username;
    @NotNull
    private String email;
    @NotNull
    private String token;
    @NotNull
    private String bio;
    @NotNull
    private String image;

}
