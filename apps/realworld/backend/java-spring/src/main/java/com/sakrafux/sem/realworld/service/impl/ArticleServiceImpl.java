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
import com.sakrafux.sem.realworld.mapper.CommentMapper;
import com.sakrafux.sem.realworld.repository.ApplicationUserRepository;
import com.sakrafux.sem.realworld.repository.ArticleRepository;
import com.sakrafux.sem.realworld.repository.CommentRepository;
import com.sakrafux.sem.realworld.repository.TagRepository;
import com.sakrafux.sem.realworld.service.ArticleService;
import com.sakrafux.sem.realworld.service.ProfileService;
import com.sakrafux.sem.realworld.util.AuthUtil;
import jakarta.transaction.Transactional;
import java.time.LocalDateTime;
import java.util.List;
import lombok.RequiredArgsConstructor;
import org.springframework.data.domain.PageRequest;
import org.springframework.data.domain.Sort;
import org.springframework.data.util.Pair;
import org.springframework.stereotype.Service;

@Service
@RequiredArgsConstructor
public class ArticleServiceImpl implements ArticleService {

    private final ArticleRepository articleRepository;
    private final ApplicationUserRepository applicationUserRepository;
    private final TagRepository tagRepository;
    private final CommentRepository commentRepository;

    private final ProfileService profileService;

    private final ArticleMapper articleMapper;
    private final CommentMapper commentMapper;

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
        var results = articles.stream().map(article -> mapArticleWithAuthor(finalCurrentUser, article)).toList();

        return Pair.of(results, articles.getTotalElements());
    }

    @Override
    @Transactional
    public Pair<List<ArticleDto>, Long> getArticlesFeed(PaginationParamDto params) {
        var currentUser = AuthUtil.getCurrentUser();
        var articles = articleRepository.findAllByFollowedUsers(currentUser.getId(), PageRequest.of(
            params.getOffset(), params.getLimit(), Sort.by("createdAt").descending()));

        var results = articles.stream().map(article -> mapArticleWithAuthor(currentUser, article)).toList();

        return Pair.of(results, articles.getTotalElements());
    }

    @Override
    @Transactional
    public ArticleDto getArticle(String slug) throws NotFoundResponseException {
        var currentUser = AuthUtil.getCurrentUser();
        if (currentUser != null) {
            currentUser = applicationUserRepository.findById(currentUser.getId()).orElseThrow(null);
        }

        var article = articleRepository.findBySlug(slug).orElseThrow(
            () -> new NotFoundResponseException("Article not found"));

        return mapArticleWithAuthor(currentUser, article);
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
            resultDto.setAuthor(
                profileService.getProfileByUsername(article.getAuthor().getUsername()));
        } catch (NotFoundResponseException ignored) {
        }

        return resultDto;
    }

    @Override
    @Transactional
    public ArticleDto updateArticle(String slug, UpdateArticleDto articleDto)
        throws NotFoundResponseException {
        var article = articleRepository.findBySlug(slug).orElseThrow(
            () -> new NotFoundResponseException("Article not found"));

        article = articleMapper.updateDtoToEntity(articleDto, article);
        article = articleRepository.save(article);
        article.setUpdatedAt(LocalDateTime.now());
        var resultDto = articleMapper.entityToDto(article);

        try {
            resultDto.setAuthor(
                profileService.getProfileByUsername(article.getAuthor().getUsername()));
        } catch (NotFoundResponseException ignored) {
        }

        return resultDto;
    }

    @Override
    @Transactional
    public void deleteArticle(String slug) {
        articleRepository.deleteBySlug(slug);
    }

    @Override
    public CommentDto createArticleComment(String slug, NewCommentDto commentDto)
        throws NotFoundResponseException {
        var article = articleRepository.findBySlug(slug).orElseThrow(
            () -> new NotFoundResponseException("Article not found"));

        var currentUser = AuthUtil.getCurrentUser();

        var comment = commentMapper.newDtoToEntity(commentDto);
        comment.setArticle(article);
        comment.setAuthor(currentUser);
        comment.setCreatedAt(LocalDateTime.now());
        comment.setUpdatedAt(LocalDateTime.now());

        comment = commentRepository.save(comment);

        var resultDto = commentMapper.entityToDto(comment);
        resultDto.setAuthor(profileService.getProfileByUsername(currentUser.getUsername()));

        return resultDto;
    }

    @Override
    @Transactional
    public List<CommentDto> getArticleComments(String slug) throws NotFoundResponseException {
        var comments = articleRepository.findBySlug(slug).orElseThrow(
            () -> new NotFoundResponseException("Article not found")).getComments();

        return comments.stream().map(comment -> {
            var commentDto = commentMapper.entityToDto(comment);

            try {
                commentDto.setAuthor(profileService.getProfileByUsername(comment.getAuthor().getUsername()));
            } catch (NotFoundResponseException ignored) {
            }
            return commentDto;
        }).toList();
    }

    @Override
    @Transactional
    public void deleteArticleComment(String slug, Long id) {
        commentRepository.deleteByArticleSlugAndId(slug, id);
    }

    @Override
    @Transactional
    public ArticleDto favoriteArticle(String slug) throws NotFoundResponseException {
        var article = articleRepository.findBySlug(slug).orElseThrow(
            () -> new NotFoundResponseException("Article not found"));

        var currentUser = AuthUtil.getCurrentUser();

        article.getFavoritedBy().add(currentUser);
        articleRepository.save(article);

        return mapArticleWithAuthor(currentUser, article);
    }

    @Override
    @Transactional
    public ArticleDto unfavoriteArticle(String slug) throws NotFoundResponseException {
        var article = articleRepository.findBySlug(slug).orElseThrow(
            () -> new NotFoundResponseException("Article not found"));

        var currentUser = AuthUtil.getCurrentUser();

        article.getFavoritedBy().remove(currentUser);
        articleRepository.save(article);

        return mapArticleWithAuthor(currentUser, article);
    }

    private boolean isFavoritedByCurrentUser(Article article, ApplicationUser currentUser) {
        if (currentUser == null) {
            return false;
        }
        return article.getFavoritedBy().stream()
            .anyMatch(user -> user.getId().equals(currentUser.getId()));
    }

    private ArticleDto mapArticleWithAuthor(ApplicationUser currentUser, Article article) {
        var articleDto = articleMapper.entityToDto(article);

        articleDto.setFavorited(isFavoritedByCurrentUser(article, currentUser));
        try {
            articleDto.setAuthor(
                profileService.getProfileByUsername(article.getAuthor().getUsername()));
        } catch (NotFoundResponseException ignore) {
        }

        return articleDto;
    }

}
