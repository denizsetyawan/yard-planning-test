/*
 Navicat Premium Dump SQL

 Source Server         : localhost
 Source Server Type    : PostgreSQL
 Source Server Version : 170002 (170002)
 Source Host           : localhost:5432
 Source Catalog        : yard-planning
 Source Schema         : public

 Target Server Type    : PostgreSQL
 Target Server Version : 170002 (170002)
 File Encoding         : 65001

 Date: 16/11/2025 21:18:30
*/


-- ----------------------------
-- Table structure for yards
-- ----------------------------
DROP TABLE IF EXISTS "public"."yards";
CREATE TABLE "public"."yards" (
  "id" int8 NOT NULL DEFAULT nextval('yards_id_seq'::regclass),
  "code" varchar(10) COLLATE "pg_catalog"."default" NOT NULL,
  "name" varchar(100) COLLATE "pg_catalog"."default" NOT NULL,
  "description" text COLLATE "pg_catalog"."default",
  "created_at" timestamptz(6) NOT NULL DEFAULT now(),
  "updated_at" timestamptz(6) NOT NULL DEFAULT now()
)
;

-- ----------------------------
-- Records of yards
-- ----------------------------
INSERT INTO "public"."yards" VALUES (2, 'YRD2', '2', 'Yard 2', '2025-11-15 21:51:09.608265+07', '2025-11-16 17:47:02.173849+07');
INSERT INTO "public"."yards" VALUES (3, 'YRD3', '3', 'Yard 3', '2025-11-15 21:51:09.608265+07', '2025-11-16 17:47:08.223829+07');
INSERT INTO "public"."yards" VALUES (1, 'YRD1', '1', 'Yard 1', '2025-11-15 21:51:09.608265+07', '2025-11-16 17:47:13.09628+07');

-- ----------------------------
-- Triggers structure for table yards
-- ----------------------------
CREATE TRIGGER "public"."trg_yards_updated_at" BEFORE UPDATE ON "public"."yards"
FOR EACH ROW
EXECUTE PROCEDURE "public"."set_updated_at"();

-- ----------------------------
-- Uniques structure for table yards
-- ----------------------------
ALTER TABLE "public"."yards" ADD CONSTRAINT "yards_code_key" UNIQUE ("code");

-- ----------------------------
-- Primary Key structure for table yards
-- ----------------------------
ALTER TABLE "public"."yards" ADD CONSTRAINT "yards_pkey" PRIMARY KEY ("id");
