/*
 Navicat Premium Dump SQL

 Source Server         : Localhost
 Source Server Type    : PostgreSQL
 Source Server Version : 140012 (140012)
 Source Host           : localhost:5432
 Source Catalog        : golang
 Source Schema         : mygram

 Target Server Type    : PostgreSQL
 Target Server Version : 140012 (140012)
 File Encoding         : 65001

 Date: 23/09/2024 23:11:24
*/


-- ----------------------------
-- Sequence structure for comment_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "mygram"."comment_id_seq";
CREATE SEQUENCE "mygram"."comment_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;
ALTER SEQUENCE "mygram"."comment_id_seq" OWNER TO "postgres";

-- ----------------------------
-- Sequence structure for photo_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "mygram"."photo_id_seq";
CREATE SEQUENCE "mygram"."photo_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;
ALTER SEQUENCE "mygram"."photo_id_seq" OWNER TO "postgres";

-- ----------------------------
-- Sequence structure for social_media_url_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "mygram"."social_media_url_id_seq";
CREATE SEQUENCE "mygram"."social_media_url_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;
ALTER SEQUENCE "mygram"."social_media_url_id_seq" OWNER TO "postgres";

-- ----------------------------
-- Sequence structure for user_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "mygram"."user_id_seq";
CREATE SEQUENCE "mygram"."user_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;
ALTER SEQUENCE "mygram"."user_id_seq" OWNER TO "postgres";

-- ----------------------------
-- Table structure for comment
-- ----------------------------
DROP TABLE IF EXISTS "mygram"."comment";
CREATE TABLE "mygram"."comment" (
  "id" int8 NOT NULL DEFAULT nextval('"mygram".comment_id_seq'::regclass),
  "message" varchar(255) COLLATE "pg_catalog"."default",
  "user_id" int4 NOT NULL,
  "photo_id" int4 NOT NULL,
  "created_at" timestamp(6),
  "updated_at" timestamp(6)
)
;
ALTER TABLE "mygram"."comment" OWNER TO "postgres";

-- ----------------------------
-- Table structure for photo
-- ----------------------------
DROP TABLE IF EXISTS "mygram"."photo";
CREATE TABLE "mygram"."photo" (
  "id" int8 NOT NULL DEFAULT nextval('"mygram".photo_id_seq'::regclass),
  "title" text COLLATE "pg_catalog"."default",
  "caption" text COLLATE "pg_catalog"."default",
  "photo_url" varchar(255) COLLATE "pg_catalog"."default",
  "user_id" int8 NOT NULL,
  "created_at" timestamp(6),
  "updated_at" timestamp(6)
)
;
ALTER TABLE "mygram"."photo" OWNER TO "postgres";

-- ----------------------------
-- Table structure for social_media
-- ----------------------------
DROP TABLE IF EXISTS "mygram"."social_media";
CREATE TABLE "mygram"."social_media" (
  "id" int8 NOT NULL DEFAULT nextval('"mygram".social_media_url_id_seq'::regclass),
  "name" varchar(255) COLLATE "pg_catalog"."default",
  "social_media_url" text COLLATE "pg_catalog"."default",
  "user_id" int8 NOT NULL,
  "created_at" timestamp(6),
  "updated_at" timestamp(6)
)
;
ALTER TABLE "mygram"."social_media" OWNER TO "postgres";

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS "mygram"."user";
CREATE TABLE "mygram"."user" (
  "id" int8 NOT NULL DEFAULT nextval('"mygram".user_id_seq'::regclass),
  "username" varchar(45) COLLATE "pg_catalog"."default" NOT NULL,
  "email" varchar(100) COLLATE "pg_catalog"."default" NOT NULL,
  "password" varchar(255) COLLATE "pg_catalog"."default",
  "age" int4,
  "created_at" timestamp(6),
  "updated_at" timestamp(6)
)
;
ALTER TABLE "mygram"."user" OWNER TO "postgres";

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "mygram"."comment_id_seq"
OWNED BY "mygram"."comment"."id";
SELECT setval('"mygram"."comment_id_seq"', 4, true);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "mygram"."photo_id_seq"
OWNED BY "mygram"."photo"."id";
SELECT setval('"mygram"."photo_id_seq"', 9, true);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "mygram"."social_media_url_id_seq"
OWNED BY "mygram"."social_media"."id";
SELECT setval('"mygram"."social_media_url_id_seq"', 3, true);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "mygram"."user_id_seq"
OWNED BY "mygram"."user"."id";
SELECT setval('"mygram"."user_id_seq"', 5, true);

-- ----------------------------
-- Primary Key structure for table comment
-- ----------------------------
ALTER TABLE "mygram"."comment" ADD CONSTRAINT "comment_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table photo
-- ----------------------------
ALTER TABLE "mygram"."photo" ADD CONSTRAINT "photo_pkey" PRIMARY KEY ("id", "user_id");

-- ----------------------------
-- Primary Key structure for table social_media
-- ----------------------------
ALTER TABLE "mygram"."social_media" ADD CONSTRAINT "social_media_url_pkey" PRIMARY KEY ("id", "user_id");

-- ----------------------------
-- Uniques structure for table user
-- ----------------------------
ALTER TABLE "mygram"."user" ADD CONSTRAINT "username" UNIQUE ("username");

-- ----------------------------
-- Primary Key structure for table user
-- ----------------------------
ALTER TABLE "mygram"."user" ADD CONSTRAINT "user_pkey" PRIMARY KEY ("id");
