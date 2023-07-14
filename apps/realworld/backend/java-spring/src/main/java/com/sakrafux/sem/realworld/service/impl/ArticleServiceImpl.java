package com.sakrafux.sem.realworld.service.impl;

import com.sakrafux.sem.realworld.dto.ArticleDto;
import com.sakrafux.sem.realworld.dto.CommentDto;
import com.sakrafux.sem.realworld.dto.NewArticleDto;
import com.sakrafux.sem.realworld.dto.NewCommentDto;
import com.sakrafux.sem.realworld.dto.UpdateArticleDto;
import com.sakrafux.sem.realworld.dto.request.PaginationParamDto;
import com.sakrafux.sem.realworld.entity.ApplicationUser;
import com.sakrafux.sem.realworld.entity.Article;
import com.sakrafux.sem.realworld.entity.Tag;
import com.sakrafux.sem.realworld.exception.response.NotFoundResponseException;
import com.sakrafux.sem.realworld.mapper.ArticleMapper;
import com.sakrafux.sem.realworld.repository.ApplicationUserRepository;
import com.sakrafux.sem.realworld.repository.ArticleRepository;
import com.sakrafux.sem.realworld.repository.TagRepository;
import com.sakrafux.sem.realworld.service.ArticleService;
import com.sakrafux.sem.realworld.service.ProfileService;
import com.sakrafux.sem.realworld.util.AuthUtil;
import jakarta.transaction.Transactional;
import java.time.LocalDateTime;
import java.util.List;
import lombok.RequiredArgsConstructor;
import org.springframework.data.util.Pair;
import org.springframework.stereotype.Service;

@Service
@RequiredArgsConstructor
public class ArticleServiceImpl implements ArticleService {

    private final ArticleRepository articleRepository;
    private final ApplicationUserRepository applicationUserRepository;
    private final TagRepository tagRepository;

    private final ProfileService profileService;

    private final ArticleMapper articleMapper;

    @Override
    @Transactional
    public Pair<List<ArticleDto>, Long> getArticles(PaginationParamDto params, String tag,
                                                    String author,
                                                    String favoritedBy) {
        var articles = articleRepository.getArticlesByCriteria(tag, author, favoritedBy,
            params.getOffset(), params.getLimit());

        var currentUser = AuthUtil.getCurrentUser();
        if (currentUser != null) {
            currentUser = applicationUserRepository.findById(currentUser.getId()).orElseThrow(null);
        }

        var finalCurrentUser = currentUser;
        var results = articles.stream().map(article -> {
            var articleDto = articleMapper.entityToDto(article);

            articleDto.setFavorited(isFavoritedByCurrentUser(article, finalCurrentUser));
            try {
                articleDto.setAuthor(
                    profileService.getProfileByUsername(article.getAuthor().getUsername()));
            } catch (NotFoundResponseException ignore) {
            }

            return articleDto;
        }).toList();

        return Pair.of(results, articles.getTotalElements());
    }

    @Override
    public List<ArticleDto> getArticlesFeed(PaginationParamDto params) {
        return null;
    }

    @Override
    public ArticleDto getArticle(String slug) {
        return null;
    }

    @Override
    @Transactional
    public ArticleDto createArticle(NewArticleDto articleDto) {
        var article = articleMapper.newDtoToEntity(articleDto);

        var tags = articleDto.getTagList().stream().map(tag -> {
            var existingTag = tagRepository.findByTag(tag);
            return existingTag.orElseGet(() -> tagRepository.save(Tag.builder().tag(tag).build()));
        }).toList();

        article.setTags(tags);
        article.setAuthor(AuthUtil.getCurrentUser());

        article = articleRepository.save(article);
        // Manually set timestamps, because Java doesn't synchronize triggered events
        article.setCreatedAt(LocalDateTime.now());
        article.setUpdatedAt(LocalDateTime.now());
        var resultDto = articleMapper.entityToDto(article);

        try {
            resultDto.setAuthor(profileService.getProfileByUsername(article.getAuthor().getUsername()));
        } catch (NotFoundResponseException ignored) {}

        return resultDto;
    }

    @Override
    public ArticleDto updateArticle(String slug, UpdateArticleDto articleDto) {
        return null;
    }

    @Override
    public void deleteArticle(String slug) {

    }

    @Override
    public CommentDto createArticleComment(String slug, NewCommentDto commentDto) {
        return null;
    }

    @Override
    public List<CommentDto> getArticleComments(String slug) {
        return null;
    }

    @Override
    public void deleteArticleComment(String slug, Long id) {

    }

    @Override
    public void favoriteArticle(String slug) {

    }

    @Override
    public void unfavoriteArticle(String slug) {

    }

    private boolean isFavoritedByCurrentUser(Article article, ApplicationUser currentUser) {
        if (currentUser == null) {
            return false;
        }
        return article.getFavoritedBy().stream()
            .anyMatch(user -> user.getId().equals(currentUser.getId()));
    }

}
