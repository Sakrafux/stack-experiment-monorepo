package com.sakrafux.sem.realworld.endpoint;

import com.sakrafux.sem.realworld.dto.request.UpdateUserRequestDto;
import com.sakrafux.sem.realworld.dto.response.UserResponseDto;
import com.sakrafux.sem.realworld.exception.response.GenericErrorResponseException;
import com.sakrafux.sem.realworld.exception.response.UnauthorizedResponseException;
import jakarta.validation.Valid;
import lombok.RequiredArgsConstructor;
import org.springframework.http.HttpStatus;
import org.springframework.security.access.annotation.Secured;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PutMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.ResponseStatus;
import org.springframework.web.bind.annotation.RestController;

@RestController
@RequestMapping(value = "/api/user")
@RequiredArgsConstructor
public class UserEndpoint {

    @Secured("ROLE_USER")
    @GetMapping()
    @ResponseStatus(HttpStatus.OK)
    public UserResponseDto getCurrentUser() throws UnauthorizedResponseException,
        GenericErrorResponseException {
        return null;
    }

    @Secured("ROLE_USER")
    @PutMapping()
    @ResponseStatus(HttpStatus.OK)
    public UserResponseDto updateCurrentUser(@Valid @RequestBody UpdateUserRequestDto dto)
        throws UnauthorizedResponseException, GenericErrorResponseException {
        return null;
    }

}
