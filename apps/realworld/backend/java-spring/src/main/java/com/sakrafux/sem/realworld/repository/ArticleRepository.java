package com.sakrafux.sem.realworld.repository;

import com.sakrafux.sem.realworld.entity.Article;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

@Repository
public interface ArticleRepository extends JpaRepository<Article, Long> {

}
