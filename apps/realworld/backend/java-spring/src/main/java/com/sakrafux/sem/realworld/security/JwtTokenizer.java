package com.sakrafux.sem.realworld.security;

import io.jsonwebtoken.Jwts;
import io.jsonwebtoken.SignatureAlgorithm;
import io.jsonwebtoken.security.Keys;
import java.util.Date;
import java.util.List;
import lombok.RequiredArgsConstructor;
import org.springframework.stereotype.Component;

@Component
@RequiredArgsConstructor
public class JwtTokenizer {

    public String getAuthToken(String user) {
        return getAuthToken(user, List.of("ROLE_USER"));
    }

    public String getAuthToken(String user, List<String> roles) {
        byte[] signingKey = SecurityProperties.SECRET.getBytes();
        String token = Jwts.builder()
            .signWith(Keys.hmacShaKeyFor(signingKey), SignatureAlgorithm.HS512)
            .setHeaderParam("typ", SecurityProperties.TYPE)
            .setIssuer(SecurityProperties.ISSUER)
            .setAudience(SecurityProperties.AUDIENCE)
            .setSubject(user)
            .setExpiration(new Date(System.currentTimeMillis() + SecurityProperties.EXPIRATION_TIME))
            .claim("rol", roles)
            .compact();
        return SecurityProperties.PREFIX + token;
    }

}
