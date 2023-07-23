package com.sakrafux.sem.realworld.service;

import com.sakrafux.sem.realworld.dto.ArticleDto;
import com.sakrafux.sem.realworld.dto.CommentDto;
import com.sakrafux.sem.realworld.dto.NewArticleDto;
import com.sakrafux.sem.realworld.dto.NewCommentDto;
import com.sakrafux.sem.realworld.dto.UpdateArticleDto;
import com.sakrafux.sem.realworld.dto.request.PaginationParamDto;
import com.sakrafux.sem.realworld.exception.response.GenericErrorResponseException;
import com.sakrafux.sem.realworld.exception.response.NotFoundResponseException;
import java.util.List;
import org.springframework.data.util.Pair;

public interface ArticleService {

    Pair<List<ArticleDto>, Long> getArticles(PaginationParamDto params, String tag, String author,
                                             String favorited);

    Pair<List<ArticleDto>, Long> getArticlesFeed(PaginationParamDto params);

    ArticleDto getArticle(String slug) throws NotFoundResponseException;

    ArticleDto createArticle(NewArticleDto articleDto) throws GenericErrorResponseException;

    ArticleDto updateArticle(String slug, UpdateArticleDto articleDto)
        throws NotFoundResponseException;

    void deleteArticle(String slug);

    CommentDto createArticleComment(String slug, NewCommentDto commentDto)
        throws NotFoundResponseException;

    List<CommentDto> getArticleComments(String slug) throws NotFoundResponseException;

    void deleteArticleComment(String slug, Long id);

    ArticleDto favoriteArticle(String slug) throws NotFoundResponseException;

    ArticleDto unfavoriteArticle(String slug) throws NotFoundResponseException;

}
