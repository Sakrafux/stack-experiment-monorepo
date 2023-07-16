package com.sakrafux.sem.realworld.integration.endpoint;

import static org.junit.jupiter.api.Assertions.assertAll;
import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertNull;
import static org.springframework.test.context.jdbc.Sql.ExecutionPhase.AFTER_TEST_METHOD;
import static org.springframework.test.web.servlet.request.MockMvcRequestBuilders.post;

import com.fasterxml.jackson.databind.ObjectMapper;
import com.sakrafux.sem.realworld.dto.LoginUserDto;
import com.sakrafux.sem.realworld.dto.request.LoginUserRequestDto;
import com.sakrafux.sem.realworld.dto.response.UserResponseDto;
import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.autoconfigure.jdbc.AutoConfigureTestDatabase;
import org.springframework.boot.test.autoconfigure.web.servlet.AutoConfigureMockMvc;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.http.HttpStatus;
import org.springframework.http.MediaType;
import org.springframework.test.context.ActiveProfiles;
import org.springframework.test.context.jdbc.Sql;
import org.springframework.test.web.servlet.MockMvc;

@SpringBootTest
@ActiveProfiles("test")
// Using the normal database instead of an in-memory one
@AutoConfigureTestDatabase(replace = AutoConfigureTestDatabase.Replace.NONE)
@AutoConfigureMockMvc
public class UsersEndpointTest {

    @SuppressWarnings("SpringJavaInjectionPointsAutowiringInspection")
    @Autowired
    private MockMvc mockMvc;

    @SuppressWarnings("SpringJavaInjectionPointsAutowiringInspection")
    @Autowired
    private ObjectMapper objectMapper;

    private static final String BASE_URI = "/api/users";

    @Test
    @Sql("/sql/single_user.sql")
    @Sql(value = "/sql/clean_up.sql", executionPhase = AFTER_TEST_METHOD)
    public void login_givenUserData_whenCredentialsAreValid_then200() throws Exception {
        String username = "user@email.com";
        String password = "password";
        var content = new LoginUserRequestDto(new LoginUserDto(username, password));
        String body = objectMapper.writeValueAsString(content);

        var mvcResult = mockMvc.perform(post(BASE_URI + "/login")
            .content(body)
            .contentType(MediaType.APPLICATION_JSON)).andReturn();

        var response = mvcResult.getResponse();
        var result =
            objectMapper.readValue(response.getContentAsString(), UserResponseDto.class).getUser();

        assertAll(
            () -> assertEquals(HttpStatus.OK.value(), response.getStatus()),
            () -> assertEquals(MediaType.APPLICATION_JSON_VALUE, response.getContentType()),
            () -> assertEquals("user", result.getUsername()),
            () -> assertEquals("user@email.com", result.getEmail()),
            () -> assertEquals("Just a simple user", result.getBio()),
            () -> assertNull(result.getImage())
        );
    }

}
