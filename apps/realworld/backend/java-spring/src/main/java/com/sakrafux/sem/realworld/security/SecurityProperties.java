package com.sakrafux.sem.realworld.security;

import lombok.AccessLevel;
import lombok.NoArgsConstructor;

@NoArgsConstructor(access = AccessLevel.PRIVATE)
public class SecurityProperties {

    public static final String HEADER = "Authorization";
    public static final String PREFIX = "Bearer ";
    public static final String LOGIN_URI = "/api/users/login";

    public static final String SECRET = "PBORCqTNbx2+YrKmF9tjI/dAYWBnMh8LLPUXf8Gnm+aarCwWOjANi8YMOp9qLj0t";
    public static final String TYPE = "JWT";
    public static final String ISSUER = "realworld";
    public static final String AUDIENCE = "realworld";
    public static final Long EXPIRATION_TIME = 86_400_000L;

}
