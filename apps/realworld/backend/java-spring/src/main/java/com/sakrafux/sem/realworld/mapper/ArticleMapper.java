package com.sakrafux.sem.realworld.mapper;

import com.sakrafux.sem.realworld.dto.ArticleDto;
import com.sakrafux.sem.realworld.dto.NewArticleDto;
import com.sakrafux.sem.realworld.entity.ApplicationUser;
import com.sakrafux.sem.realworld.entity.Article;
import com.sakrafux.sem.realworld.entity.Tag;
import java.util.Comparator;
import java.util.List;
import org.mapstruct.Mapper;
import org.mapstruct.Mapping;
import org.mapstruct.Named;

@Mapper
public interface ArticleMapper {

    @Mapping(target = "tagList", source = "tags", qualifiedByName = "tagsToStrings")
    @Mapping(target = "favoritesCount", source = "favoritedBy", qualifiedByName = "countFavorites")
    @Mapping(target = "author", ignore = true)
    ArticleDto entityToDto(Article article);

    @Mapping(target = "slug", source = "title", qualifiedByName = "titleToSlug")
    Article newDtoToEntity(NewArticleDto articleDto);

    @Named("tagsToStrings")
    default List<String> tagsToStrings(List<Tag> tags) {
        return tags.stream().map(Tag::getTag).sorted(Comparator.naturalOrder()).toList();
    }

    @Named("countFavorites")
    default int countFavorites(List<ApplicationUser> favoritedBy) {
        if (favoritedBy == null) {
            return 0;
        }
        return favoritedBy.size();
    }

    @Named("titleToSlug")
    default String titleToSlug(String title) {
        return title.toLowerCase().replaceAll("\\s", "-");
    }

}
