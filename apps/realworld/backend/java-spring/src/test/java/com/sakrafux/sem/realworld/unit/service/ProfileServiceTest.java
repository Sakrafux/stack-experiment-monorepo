package com.sakrafux.sem.realworld.unit.service;

import static org.junit.jupiter.api.Assertions.assertAll;
import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertFalse;
import static org.junit.jupiter.api.Assertions.assertNull;
import static org.mockito.Mockito.mockStatic;
import static org.mockito.Mockito.verify;
import static org.mockito.Mockito.when;

import com.sakrafux.sem.realworld.entity.ApplicationUser;
import com.sakrafux.sem.realworld.exception.response.NotFoundResponseException;
import com.sakrafux.sem.realworld.mapper.ProfileMapper;
import com.sakrafux.sem.realworld.repository.ApplicationUserRepository;
import com.sakrafux.sem.realworld.service.ProfileService;
import com.sakrafux.sem.realworld.util.AuthUtil;
import java.util.Optional;
import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.boot.test.mock.mockito.MockBean;
import org.springframework.boot.test.mock.mockito.SpyBean;
import org.springframework.test.context.ActiveProfiles;

@SpringBootTest
@ActiveProfiles("test")
public class ProfileServiceTest {

    @MockBean
    private ApplicationUserRepository applicationUserRepository;

    @SpyBean
    private ProfileMapper profileMapper;

    @Autowired
    private ProfileService profileService;

    @Test
    void getProfileByUsername_givenUsername_whenFound_thenMappedProfileDto()
        throws NotFoundResponseException {
        var user = ApplicationUser.builder()
            .id(1L)
            .username("username")
            .email("email")
            .password("password")
            .bio("")
            .build();

        when(applicationUserRepository.findByUsername("username")).thenReturn(Optional.of(user));

        try (var authUtil = mockStatic(AuthUtil.class)) {
            authUtil.when(AuthUtil::getCurrentUser).thenReturn(null);

            var ret = profileService.getProfileByUsername("username");

            assertAll(
                () -> assertEquals("username", ret.getUsername()),
                () -> assertEquals("", ret.getBio()),
                () -> assertNull(ret.getImage()),
                () -> assertFalse(ret.isFollowing())
            );
            verify(profileMapper).entityToDto(user, false);
        }
    }

}
