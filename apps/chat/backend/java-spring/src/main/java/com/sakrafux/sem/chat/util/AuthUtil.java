package com.sakrafux.sem.chat.util;

import org.springframework.security.core.context.SecurityContextHolder;

public class AuthUtil {

    public static String getUserId() {
        return (String) SecurityContextHolder.getContext().getAuthentication().getPrincipal();
    }

}
