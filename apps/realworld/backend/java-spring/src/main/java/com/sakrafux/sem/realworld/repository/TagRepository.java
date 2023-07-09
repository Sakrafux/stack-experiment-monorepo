package com.sakrafux.sem.realworld.repository;

import com.sakrafux.sem.realworld.entity.Tag;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

@Repository
public interface TagRepository extends JpaRepository<Tag, Long> {

}
