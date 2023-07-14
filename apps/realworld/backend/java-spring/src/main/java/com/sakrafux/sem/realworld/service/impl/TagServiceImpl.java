package com.sakrafux.sem.realworld.service.impl;

import com.sakrafux.sem.realworld.entity.Tag;
import com.sakrafux.sem.realworld.repository.TagRepository;
import com.sakrafux.sem.realworld.service.TagService;
import java.util.List;
import lombok.RequiredArgsConstructor;
import org.springframework.stereotype.Service;

@Service
@RequiredArgsConstructor
public class TagServiceImpl implements TagService {

    private final TagRepository tagRepository;

    @Override
    public List<String> getTags() {
        return tagRepository.findAll().stream().map(Tag::getTag).toList();
    }

}
