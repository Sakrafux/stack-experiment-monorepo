package com.sakrafux.sem.realworld.service.impl;

import com.sakrafux.sem.realworld.dto.ProfileDto;
import com.sakrafux.sem.realworld.entity.ApplicationUser;
import com.sakrafux.sem.realworld.exception.response.NotFoundResponseException;
import com.sakrafux.sem.realworld.mapper.ProfileMapper;
import com.sakrafux.sem.realworld.repository.ApplicationUserRepository;
import com.sakrafux.sem.realworld.service.ProfileService;
import com.sakrafux.sem.realworld.util.AuthUtil;
import jakarta.transaction.Transactional;
import lombok.RequiredArgsConstructor;
import org.springframework.stereotype.Service;

@Service
@RequiredArgsConstructor
public class ProfileServiceImpl implements ProfileService {

    private final ApplicationUserRepository applicationUserRepository;

    private final ProfileMapper profileMapper;

    @Override
    @Transactional
    public ProfileDto getProfileByUsername(String username) throws NotFoundResponseException {
        var user = applicationUserRepository.findByUsername(username)
            .orElseThrow(() -> new NotFoundResponseException("Profile not found."));

        var currentUser = AuthUtil.getCurrentUser();

        if (currentUser == null) {
            return profileMapper.entityToDto(user, false);
        }

        currentUser = applicationUserRepository.findById(currentUser.getId()).orElseThrow(null);

        var following = currentUser.getFollowing().stream().map(ApplicationUser::getId)
            .anyMatch(id -> id.equals(user.getId()));

        return profileMapper.entityToDto(user, following);
    }

    @Override
    @Transactional
    public ProfileDto followUserByUsername(String username) throws NotFoundResponseException {
        var user = applicationUserRepository.findByUsername(username)
            .orElseThrow(() -> new NotFoundResponseException("Profile not found."));

        user.getFollowers().add(AuthUtil.getCurrentUser());

        applicationUserRepository.save(user);

        return profileMapper.entityToDto(user, true);
    }

    @Override
    @Transactional
    public ProfileDto unfollowUserByUsername(String username) throws NotFoundResponseException {
        var user = applicationUserRepository.findByUsername(username)
            .orElseThrow(() -> new NotFoundResponseException("Profile not found."));

        user.getFollowers().remove(AuthUtil.getCurrentUser());

        applicationUserRepository.save(user);

        return profileMapper.entityToDto(user, false);
    }

}
