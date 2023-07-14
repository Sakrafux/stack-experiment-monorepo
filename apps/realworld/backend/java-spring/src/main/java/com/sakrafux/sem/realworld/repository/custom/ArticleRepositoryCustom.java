package com.sakrafux.sem.realworld.repository.custom;

import com.sakrafux.sem.realworld.entity.Article;
import org.springframework.data.domain.Page;

public interface ArticleRepositoryCustom {

    Page<Article> getArticlesByCriteria(String tag, String author, String favoritedBy, int offset,
                                        int limit);

}
