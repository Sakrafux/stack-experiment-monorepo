package com.sakrafux.sem.realworld.util;

import com.sakrafux.sem.realworld.entity.ApplicationUser;
import org.springframework.security.core.context.SecurityContextHolder;

public class AuthUtil {

    public static ApplicationUser getCurrentUser() {
        var auth = SecurityContextHolder.getContext().getAuthentication().getPrincipal();
        if (auth instanceof String) {
            return null;
        }
        return (ApplicationUser) SecurityContextHolder.getContext().getAuthentication().getPrincipal();
    }

}
