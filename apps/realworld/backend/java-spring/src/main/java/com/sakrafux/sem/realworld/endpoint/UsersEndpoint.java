package com.sakrafux.sem.realworld.endpoint;

import com.sakrafux.sem.realworld.dto.NewUserDto;
import com.sakrafux.sem.realworld.dto.request.LoginUserRequestDto;
import com.sakrafux.sem.realworld.dto.response.UserResponseDto;
import com.sakrafux.sem.realworld.exception.response.GenericErrorResponseException;
import com.sakrafux.sem.realworld.exception.response.UnauthorizedResponseException;
import jakarta.validation.Valid;
import lombok.RequiredArgsConstructor;
import org.springframework.http.HttpStatus;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.ResponseStatus;
import org.springframework.web.bind.annotation.RestController;

@RestController
@RequestMapping("/api/users")
@RequiredArgsConstructor
public class UsersEndpoint {

    @PostMapping("/login")
    @ResponseStatus(HttpStatus.OK)
    public UserResponseDto login(@Valid @RequestBody LoginUserRequestDto dto) throws
        UnauthorizedResponseException, GenericErrorResponseException {
        return null;
    }

    @PostMapping()
    @ResponseStatus(HttpStatus.CREATED)
    public UserResponseDto createUser(@Valid @RequestBody NewUserDto dto)
        throws GenericErrorResponseException {
        return null;
    }

}
