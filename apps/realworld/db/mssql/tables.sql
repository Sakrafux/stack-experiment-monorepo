CREATE TABLE [app_user] (
  [id] bigint PRIMARY KEY,
  [username] varchar(50) UNIQUE NOT NULL,
  [email] varchar(100) UNIQUE NOT NULL,
  [password] varchar(60) NOT NULL,
  [bio] varchar(255) NOT NULL,
  [image] varchar(255),
  [created_at] timestamp,
  [updated_at] timestamp,
  [version] integer
)
GO

CREATE TABLE [article] (
  [id] bigint PRIMARY KEY,
  [slug] varchar(100) UNIQUE NOT NULL,
  [title] varchar(100) UNIQUE NOT NULL,
  [description] varchar(255) NOT NULL,
  [body] text NOT NULL,
  [fk_author] bigint NOT NULL,
  [created_at] timestamp,
  [updated_at] timestamp,
  [version] integer
)
GO

CREATE TABLE [comment] (
  [id] bigint PRIMARY KEY,
  [body] text NOT NULL,
  [fk_article] bigint NOT NULL,
  [fk_author] bigint NOT NULL,
  [created_at] timestamp,
  [updated_at] timestamp,
  [version] integer
)
GO

CREATE TABLE [tag] (
  [id] bigint PRIMARY KEY,
  [tag] varchar(20) UNIQUE NOT NULL,
  [created_at] timestamp,
  [updated_at] timestamp,
  [version] integer
)
GO

CREATE TABLE [follow_is_user_to_user] (
  [following_user_id] bigint,
  [followed_user_id] bigint,
  PRIMARY KEY ([following_user_id], [followed_user_id])
)
GO

CREATE TABLE [tag_is_article_to_tag] (
  [article_id] bigint,
  [tag_id] bigint,
  PRIMARY KEY ([article_id], [tag_id])
)
GO

CREATE TABLE [favorite_is_article_to_user] (
  [article_id] bigint,
  [user_id] bigint,
  PRIMARY KEY ([article_id], [user_id])
)
GO

CREATE INDEX [ix_user_username] ON [app_user] ("username")
GO

CREATE INDEX [ix_user_email] ON [app_user] ("email")
GO

CREATE INDEX [ix_article_slug] ON [article] ("slug")
GO

CREATE INDEX [ix_article_fk_author] ON [article] ("fk_author")
GO

CREATE INDEX [ix_comment_fk_article] ON [comment] ("fk_article")
GO

CREATE INDEX [ix_comment_fk_author] ON [comment] ("fk_author")
GO

CREATE INDEX [ix_tag_tag] ON [tag] ("tag")
GO

ALTER TABLE [article] ADD FOREIGN KEY ([fk_author]) REFERENCES [app_user] ([id])
GO

ALTER TABLE [comment] ADD FOREIGN KEY ([fk_article]) REFERENCES [article] ([id])
GO

ALTER TABLE [comment] ADD FOREIGN KEY ([fk_author]) REFERENCES [app_user] ([id])
GO

ALTER TABLE [follow_is_user_to_user] ADD FOREIGN KEY ([following_user_id]) REFERENCES [app_user] ([id])
GO

ALTER TABLE [follow_is_user_to_user] ADD FOREIGN KEY ([followed_user_id]) REFERENCES [app_user] ([id])
GO

ALTER TABLE [tag_is_article_to_tag] ADD FOREIGN KEY ([article_id]) REFERENCES [article] ([id])
GO

ALTER TABLE [tag_is_article_to_tag] ADD FOREIGN KEY ([tag_id]) REFERENCES [tag] ([id])
GO

ALTER TABLE [favorite_is_article_to_user] ADD FOREIGN KEY ([article_id]) REFERENCES [article] ([id])
GO

ALTER TABLE [favorite_is_article_to_user] ADD FOREIGN KEY ([user_id]) REFERENCES [app_user] ([id])
GO
