delete from tag_is_article_to_tag where article_id < 0 or tag_id < 0;
delete from favorite_is_article_to_user where article_id < 0 or user_id < 0;
delete from follow_is_user_to_user where following_user_id < 0 or followed_user_id < 0;
delete from tag where id < 0;
delete from comment where id < 0;
delete from article where id < 0;
delete from app_user where id < 0;