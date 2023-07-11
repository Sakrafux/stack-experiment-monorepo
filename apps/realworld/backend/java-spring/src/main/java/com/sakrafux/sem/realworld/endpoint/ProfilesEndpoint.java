package com.sakrafux.sem.realworld.endpoint;

import com.sakrafux.sem.realworld.dto.response.ProfileResponseDto;
import com.sakrafux.sem.realworld.exception.response.GenericErrorResponseException;
import com.sakrafux.sem.realworld.exception.response.NotFoundResponseException;
import com.sakrafux.sem.realworld.service.ProfileService;
import lombok.RequiredArgsConstructor;
import org.springframework.http.HttpStatus;
import org.springframework.security.access.annotation.Secured;
import org.springframework.web.bind.annotation.DeleteMapping;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.ResponseStatus;
import org.springframework.web.bind.annotation.RestController;

@RestController
@RequestMapping("/api/profiles")
@RequiredArgsConstructor
public class ProfilesEndpoint {

    private final ProfileService profileService;

    @GetMapping("/{username}")
    @ResponseStatus(HttpStatus.OK)
    public ProfileResponseDto getProfileByUsername(@PathVariable("username") String username)
        throws NotFoundResponseException {
        return new ProfileResponseDto(profileService.getProfileByUsername(username));
    }

    @Secured("ROLE_USER")
    @PostMapping("/{username}/follow")
    @ResponseStatus(HttpStatus.OK)
    public ProfileResponseDto followUserByUsername(@PathVariable("username") String username)
        throws NotFoundResponseException {
        return new ProfileResponseDto(profileService.followUserByUsername(username));
    }

    @Secured("ROLE_USER")
    @DeleteMapping("/{username}/follow")
    @ResponseStatus(HttpStatus.OK)
    public ProfileResponseDto unfollowUserByUsername(@PathVariable("username") String username)
        throws NotFoundResponseException {
        return new ProfileResponseDto(profileService.unfollowUserByUsername(username));
    }

}
