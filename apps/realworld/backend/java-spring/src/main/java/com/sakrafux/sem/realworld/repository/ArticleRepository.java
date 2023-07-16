package com.sakrafux.sem.realworld.repository;

import com.sakrafux.sem.realworld.entity.Article;
import com.sakrafux.sem.realworld.repository.custom.ArticleRepositoryCustom;
import java.util.Optional;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.Pageable;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Query;
import org.springframework.stereotype.Repository;

@Repository
public interface ArticleRepository extends JpaRepository<Article, Long>, ArticleRepositoryCustom {

    @Query("select a from ApplicationUser u join u.following fu join fu.articles a where u.id = :userId")
    Page<Article> findAllByFollowedUsers(Long userId, Pageable pageable);

    Optional<Article> findBySlug(String slug);

    void deleteBySlug(String slug);

}
