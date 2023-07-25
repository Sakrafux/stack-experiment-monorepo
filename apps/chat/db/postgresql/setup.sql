CREATE TABLE "app_user" (
    "id" bigint PRIMARY KEY,
    "gid" varchar(50) UNIQUE NOT NULL,
    "name" varchar(100) NOT NULL,
    "picture" varchar(255) NOT NULL
);

CREATE TABLE "chat" (
    "id" bigint PRIMARY KEY,
    "fk_user_1" bigint NOT NULL,
    "fk_user_2" bigint NOT NULL
);

CREATE TABLE "message" (
    "id" bigint PRIMARY KEY,
    "fk_chat" bigint NOT NULL,
    "fk_user" bigint NOT NULL,
    "text" varchar(255) NOT NULL,
    "created_at" timestamp
);

CREATE INDEX "ix_user_gid" ON "app_user" ("gid");

CREATE INDEX "ix_message_fk_chat" ON "message" ("fk_chat");

ALTER TABLE "chat" ADD FOREIGN KEY ("fk_user_1") REFERENCES "app_user" ("id");

ALTER TABLE "chat" ADD FOREIGN KEY ("fk_user_2") REFERENCES "app_user" ("id");

ALTER TABLE "message" ADD FOREIGN KEY ("fk_chat") REFERENCES "chat" ("id");

ALTER TABLE "message" ADD FOREIGN KEY ("fk_user") REFERENCES "app_user" ("id");

CREATE SEQUENCE seq_user_id
START 1
INCREMENT 1
MINVALUE 1
OWNED BY app_user.id;

CREATE SEQUENCE seq_chat_id
START 1
INCREMENT 1
MINVALUE 1
OWNED BY chat.id;

CREATE SEQUENCE seq_message_id
START 1
INCREMENT 1
MINVALUE 1
OWNED BY message.id;