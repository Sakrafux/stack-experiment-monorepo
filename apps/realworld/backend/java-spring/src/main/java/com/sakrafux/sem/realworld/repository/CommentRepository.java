package com.sakrafux.sem.realworld.repository;

import com.sakrafux.sem.realworld.entity.Comment;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

@Repository
public interface CommentRepository extends JpaRepository<Comment, Long> {

}
