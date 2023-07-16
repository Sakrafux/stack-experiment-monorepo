package com.sakrafux.sem.realworld.unit.repository;

import static org.junit.jupiter.api.Assertions.assertTrue;
import static org.springframework.test.context.jdbc.Sql.ExecutionPhase.AFTER_TEST_METHOD;

import com.sakrafux.sem.realworld.repository.ApplicationUserRepository;
import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.autoconfigure.jdbc.AutoConfigureTestDatabase;
import org.springframework.boot.test.autoconfigure.orm.jpa.DataJpaTest;
import org.springframework.test.context.ActiveProfiles;
import org.springframework.test.context.jdbc.Sql;

@DataJpaTest
@ActiveProfiles("test")
// Using the normal database instead of an in-memory one
@AutoConfigureTestDatabase(replace = AutoConfigureTestDatabase.Replace.NONE)
public class ApplicationUserRepositoryTest {

    @Autowired
    private ApplicationUserRepository applicationUserRepository;

    @Test
    @Sql("/sql/single_user.sql")
    @Sql(value = "/sql/clean_up.sql", executionPhase = AFTER_TEST_METHOD)
    void findByUsername_givenDataExists_whenUserExists_thenFindUser() {
        var result = applicationUserRepository.findByUsername("user");
        assertTrue(result.isPresent());
    }

}
