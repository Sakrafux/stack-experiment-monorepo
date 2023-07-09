package com.sakrafux.sem.realworld.security;

import com.sakrafux.sem.realworld.repository.ApplicationUserRepository;
import io.jsonwebtoken.Claims;
import io.jsonwebtoken.JwtException;
import io.jsonwebtoken.Jwts;
import jakarta.servlet.FilterChain;
import jakarta.servlet.ServletException;
import jakarta.servlet.http.HttpServletRequest;
import jakarta.servlet.http.HttpServletResponse;
import java.io.IOException;
import java.util.List;
import lombok.RequiredArgsConstructor;
import lombok.extern.slf4j.Slf4j;
import org.slf4j.MDC;
import org.springframework.security.authentication.UsernamePasswordAuthenticationToken;
import org.springframework.security.core.authority.SimpleGrantedAuthority;
import org.springframework.security.core.context.SecurityContextHolder;
import org.springframework.stereotype.Component;
import org.springframework.web.filter.OncePerRequestFilter;

@Component
@Slf4j
@RequiredArgsConstructor
public class JwtAuthorizationFilter extends OncePerRequestFilter {

    private final ApplicationUserRepository userRepository;

    @Override
    protected void doFilterInternal(HttpServletRequest request, HttpServletResponse response, FilterChain chain)
        throws IOException, ServletException {

        try {
            UsernamePasswordAuthenticationToken authToken = getAuthToken(request);
            if (authToken != null) {
                SecurityContextHolder.getContext().setAuthentication(authToken);
            }
        } catch (IllegalArgumentException | JwtException e) {
            log.debug("Invalid authorization attempt: {}", e.getMessage());
            response.setStatus(HttpServletResponse.SC_UNAUTHORIZED);
            response.getWriter().write("Invalid authorization header or token");
            return;
        }
        chain.doFilter(request, response);
    }

    private UsernamePasswordAuthenticationToken getAuthToken(HttpServletRequest request)
        throws JwtException, IllegalArgumentException {

        String token = request.getHeader(SecurityProperties.HEADER);
        if (token == null || token.isEmpty()) {
            return null;
        }

        if (!token.startsWith(SecurityProperties.PREFIX)) {
            throw new IllegalArgumentException("Authorization header is malformed or missing");
        }

        byte[] signingKey = SecurityProperties.SECRET.getBytes();

        Claims claims = Jwts.parserBuilder().setSigningKey(signingKey).build()
            .parseClaimsJws(token.replace(SecurityProperties.PREFIX, ""))
            .getBody();

        String username = claims.getSubject();

        List<SimpleGrantedAuthority> authorities = ((List<?>) claims
            .get("rol")).stream()
            .map(authority -> new SimpleGrantedAuthority((String) authority))
            .toList();

        if (username == null || username.isEmpty()) {
            throw new IllegalArgumentException("Token contains no user");
        }

        MDC.put("u", username);

        var user = userRepository.findByUsername(username);

        if (user.isEmpty()) {
            throw new IllegalArgumentException("User does not exist");
        }

        return new UsernamePasswordAuthenticationToken(user, null, authorities);
    }

}
