package com.sakrafux.sem.realworld.repository.custom.impl;

import com.sakrafux.sem.realworld.entity.Article;
import com.sakrafux.sem.realworld.repository.custom.ArticleRepositoryCustom;
import jakarta.persistence.EntityManager;
import jakarta.persistence.PersistenceContext;
import jakarta.persistence.criteria.CriteriaQuery;
import jakarta.persistence.criteria.JoinType;
import jakarta.persistence.criteria.Root;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.PageImpl;
import org.springframework.data.domain.PageRequest;
import org.springframework.data.util.Pair;

public class ArticleRepositoryCustomImpl implements ArticleRepositoryCustom {

    @PersistenceContext
    private EntityManager em;

    @Override
    public Page<Article> getArticlesByCriteria(String tag, String author, String favoritedBy,
                                               int offset, int limit) {
        var cb = em.getCriteriaBuilder();

        var query = cb.createQuery(Article.class);

        var pair = getArticlesByCriteriaQuery(query, tag, author, favoritedBy);
        query = pair.getFirst();
        var root = pair.getSecond();
        query.select(root);
        query.orderBy(cb.desc(root.get("createdAt")));

        var typedQuery = em.createQuery(query);

        typedQuery.setFirstResult(offset);
        typedQuery.setMaxResults(limit);

        var results = typedQuery.getResultList();

        var countQuery = cb.createQuery(Long.class);

        var countPair = getArticlesByCriteriaQuery(countQuery, tag, author, favoritedBy);
        countQuery = countPair.getFirst();
        root = countPair.getSecond();
        countQuery.select(cb.countDistinct(root));

        var count = em.createQuery(countQuery).getSingleResult();

        return new PageImpl<>(results, PageRequest.of(offset, limit), count);
    }

    private <T> Pair<CriteriaQuery<T>, Root<Article>> getArticlesByCriteriaQuery(
        CriteriaQuery<T> query, String tag, String author, String favoritedBy) {
        var cb = em.getCriteriaBuilder();

        var root = query.from(Article.class);

        var authorJoin = root.join("author", JoinType.LEFT);
        var tagJoin = root.join("tags", JoinType.LEFT);
        var favoritedJoin = root.join("favoritedBy", JoinType.LEFT);

        if (tag != null) {
            query.where(cb.equal(tagJoin.get("tag"), tag));
        }

        if (author != null) {
            query.where(cb.equal(authorJoin.get("username"), author));
        }

        if (favoritedBy != null) {
            query.where(cb.equal(favoritedJoin.get("username"), favoritedBy));
        }

        return Pair.of(query, root);
    }
}
