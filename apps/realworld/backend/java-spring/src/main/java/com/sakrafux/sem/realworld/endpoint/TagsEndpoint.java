package com.sakrafux.sem.realworld.endpoint;

import com.sakrafux.sem.realworld.dto.response.TagsResponseDto;
import com.sakrafux.sem.realworld.exception.response.GenericErrorResponseException;
import lombok.RequiredArgsConstructor;
import org.springframework.http.HttpStatus;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.ResponseStatus;
import org.springframework.web.bind.annotation.RestController;

@RestController
@RequestMapping("/api/tags")
@RequiredArgsConstructor
public class TagsEndpoint {

    @GetMapping()
    @ResponseStatus(HttpStatus.OK)
    public TagsResponseDto getTags() throws GenericErrorResponseException {
        return null;
    }

}
