package com.sakrafux.sem.realworld.repository;

import com.sakrafux.sem.realworld.entity.Comment;
import java.util.List;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

@Repository
public interface CommentRepository extends JpaRepository<Comment, Long> {

    void deleteByArticleSlugAndId(String slug, Long id);

}
