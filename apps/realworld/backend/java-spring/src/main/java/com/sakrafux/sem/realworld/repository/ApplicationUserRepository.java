package com.sakrafux.sem.realworld.repository;

import com.sakrafux.sem.realworld.entity.ApplicationUser;
import java.util.Optional;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

@Repository
public interface ApplicationUserRepository extends JpaRepository<ApplicationUser, Long> {

    Optional<ApplicationUser> findByUsername(String username);

    Optional<ApplicationUser> findByEmail(String email);

    boolean existsByUsernameOrEmail(String username, String email);

}
