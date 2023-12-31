package com.sakrafux.sem.realworld.service.impl;

import com.sakrafux.sem.realworld.dto.UpdateUserDto;
import com.sakrafux.sem.realworld.dto.UserDto;
import com.sakrafux.sem.realworld.entity.ApplicationUser;
import com.sakrafux.sem.realworld.exception.response.GenericErrorResponseException;
import com.sakrafux.sem.realworld.exception.response.UnauthorizedResponseException;
import com.sakrafux.sem.realworld.mapper.UserMapper;
import com.sakrafux.sem.realworld.repository.ApplicationUserRepository;
import com.sakrafux.sem.realworld.security.JwtTokenizer;
import com.sakrafux.sem.realworld.service.UserService;
import com.sakrafux.sem.realworld.util.AuthUtil;
import lombok.RequiredArgsConstructor;
import org.springframework.security.crypto.password.PasswordEncoder;
import org.springframework.stereotype.Service;

@Service
@RequiredArgsConstructor
public class UserServiceImpl implements UserService {

    private final ApplicationUserRepository applicationUserRepository;

    private final JwtTokenizer jwtTokenizer;
    private final PasswordEncoder passwordEncoder;

    private final UserMapper userMapper;

    @Override
    public UserDto login(String email, String password) throws UnauthorizedResponseException {
        var user = applicationUserRepository.findByEmail(email)
            .orElseThrow(() -> new UnauthorizedResponseException("Invalid email or password."));

        if (!passwordEncoder.matches(password, user.getPassword())) {
            throw new UnauthorizedResponseException("Invalid email or password.");
        }

        return userMapper.entityToDto(user, jwtTokenizer.getAuthToken(user.getUsername()));
    }

    @Override
    public UserDto createUser(String username, String email, String password) throws
        GenericErrorResponseException {
        if (applicationUserRepository.existsByUsernameOrEmail(username, email)) {
            throw new GenericErrorResponseException("Email or username already exists.");
        }

        var user =
            applicationUserRepository.save(ApplicationUser.builder().username(username).email(email)
                .password(passwordEncoder.encode(password)).bio("").build());

        return userMapper.entityToDto(user, jwtTokenizer.getAuthToken(user.getUsername()));
    }

    @Override
    public UserDto getCurrentUser() {
        var user = AuthUtil.getCurrentUser();
        return userMapper.entityToDto(user, jwtTokenizer.getAuthToken(user.getUsername()));
    }

    @Override
    public UserDto updateUser(UpdateUserDto updateUserDto) throws GenericErrorResponseException {
        var user = AuthUtil.getCurrentUser();
        if (updateUserDto.getUsername() != null) {
            var entiyWithUsername = applicationUserRepository.findByUsername(updateUserDto.getUsername());
            if (entiyWithUsername.isPresent() && !entiyWithUsername.get().getId().equals(user.getId())) {
                throw new GenericErrorResponseException("Username already exists.");
            }
        }
        if (updateUserDto.getEmail() != null) {
            var entiyWithEmail = applicationUserRepository.findByEmail(updateUserDto.getEmail());
            if (entiyWithEmail.isPresent() && !entiyWithEmail.get().getId().equals(user.getId())) {
                throw new GenericErrorResponseException("Email already exists.");
            }
        }

        userMapper.updateUserDtoToEntity(updateUserDto, user);

        return userMapper.entityToDto(applicationUserRepository.save(user),
            jwtTokenizer.getAuthToken(user.getUsername()));
    }
}
