package com.sakrafux.sem.chat.security;

import lombok.AccessLevel;
import lombok.NoArgsConstructor;

@NoArgsConstructor(access = AccessLevel.PRIVATE)
public class SecurityProperties {

    public static final String HEADER = "Authorization";
    public static final String PREFIX = "Bearer ";

}
