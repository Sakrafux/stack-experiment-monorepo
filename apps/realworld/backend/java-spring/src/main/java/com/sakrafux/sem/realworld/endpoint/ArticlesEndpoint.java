package com.sakrafux.sem.realworld.endpoint;

import com.sakrafux.sem.realworld.dto.request.NewArticleRequestDto;
import com.sakrafux.sem.realworld.dto.request.PaginationParamDto;
import com.sakrafux.sem.realworld.dto.request.UpdateArticleRequestDto;
import com.sakrafux.sem.realworld.dto.response.MultipleArticlesResponseDto;
import com.sakrafux.sem.realworld.dto.response.MultipleCommentsResponseDto;
import com.sakrafux.sem.realworld.dto.response.SingleArticleResponseDto;
import com.sakrafux.sem.realworld.dto.response.SingleCommentResponseDto;
import com.sakrafux.sem.realworld.exception.response.GenericErrorResponseException;
import jakarta.validation.Valid;
import lombok.RequiredArgsConstructor;
import org.springframework.http.HttpStatus;
import org.springframework.security.access.annotation.Secured;
import org.springframework.web.bind.annotation.DeleteMapping;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.PutMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.ResponseStatus;
import org.springframework.web.bind.annotation.RestController;

@RestController
@RequestMapping("/api/articles")
@RequiredArgsConstructor
public class ArticlesEndpoint {

    @Secured("ROLE_USER")
    @GetMapping("/feed")
    @ResponseStatus(HttpStatus.OK)
    public MultipleArticlesResponseDto getArticlesFeed(@Valid PaginationParamDto params)
        throws GenericErrorResponseException {
        return null;
    }

    @GetMapping()
    @ResponseStatus(HttpStatus.OK)
    public MultipleArticlesResponseDto getArticles(@Valid PaginationParamDto params,
                                                   @RequestParam(required = false) String tag,
                                                   @RequestParam(required = false) String author,
                                                   @RequestParam(required = false) String favorited)
        throws GenericErrorResponseException {
        return null;
    }

    @Secured("ROLE_USER")
    @PostMapping()
    @ResponseStatus(HttpStatus.CREATED)
    public MultipleArticlesResponseDto createArticle(@Valid @RequestBody NewArticleRequestDto dto)
        throws GenericErrorResponseException {
        return null;
    }

    @GetMapping("/{slug}")
    @ResponseStatus(HttpStatus.OK)
    public SingleArticleResponseDto getArticle(@PathVariable String slug)
        throws GenericErrorResponseException {
        return null;
    }

    @Secured("ROLE_USER")
    @PutMapping("/{slug}")
    @ResponseStatus(HttpStatus.OK)
    public SingleArticleResponseDto updateArticle(@PathVariable String slug,
                                                  @Valid @RequestBody UpdateArticleRequestDto dto)
        throws GenericErrorResponseException {
        return null;
    }

    @Secured("ROLE_USER")
    @DeleteMapping("/{slug}")
    @ResponseStatus(HttpStatus.OK)
    public void deleteArticle(@PathVariable String slug)
        throws GenericErrorResponseException {
    }

    @GetMapping("/{slug}/comments")
    @ResponseStatus(HttpStatus.OK)
    public MultipleCommentsResponseDto getArticleComments(@PathVariable String slug)
        throws GenericErrorResponseException {
        return null;
    }

    @Secured("ROLE_USER")
    @PostMapping("/{slug}/comments")
    @ResponseStatus(HttpStatus.OK)
    public SingleCommentResponseDto createArticleComment(@PathVariable String slug,
                                                         @Valid @RequestBody SingleCommentResponseDto dto)
        throws GenericErrorResponseException {
        return null;
    }

    @Secured("ROLE_USER")
    @DeleteMapping("/{slug}/comments/{id}")
    @ResponseStatus(HttpStatus.OK)
    public void deleteArticleComment(@PathVariable String slug, @PathVariable Long id)
        throws GenericErrorResponseException {
    }

    @Secured("ROLE_USER")
    @PostMapping("/{slug}/favorite")
    @ResponseStatus(HttpStatus.OK)
    public SingleArticleResponseDto createArticleFavorite(@PathVariable String slug)
        throws GenericErrorResponseException {
        return null;
    }

    @Secured("ROLE_USER")
    @DeleteMapping("/{slug}/favorite")
    @ResponseStatus(HttpStatus.OK)
    public SingleArticleResponseDto deleteArticleFavorite(@PathVariable String slug)
        throws GenericErrorResponseException {
        return null;
    }

}
