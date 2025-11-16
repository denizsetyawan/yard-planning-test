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

 Date: 16/11/2025 21:17:55
*/


-- ----------------------------
-- Table structure for blocks
-- ----------------------------
DROP TABLE IF EXISTS "public"."blocks";
CREATE TABLE "public"."blocks" (
  "id" int4 NOT NULL DEFAULT nextval('blocks_id_seq'::regclass),
  "yard_id" int4 NOT NULL,
  "block_code" varchar(10) COLLATE "pg_catalog"."default" NOT NULL,
  "slots" int4 NOT NULL,
  "rows" int4 NOT NULL,
  "tiers" int4 NOT NULL,
  "created_at" timestamp(6) DEFAULT now(),
  "updated_at" timestamp(6) DEFAULT now()
)
;

-- ----------------------------
-- Records of blocks
-- ----------------------------
INSERT INTO "public"."blocks" VALUES (1, 1, 'LC01', 3, 2, 4, '2025-11-15 22:01:32.335081', '2025-11-15 22:01:32.335081');
INSERT INTO "public"."blocks" VALUES (2, 1, 'LC02', 3, 2, 4, '2025-11-15 22:01:32.335081', '2025-11-15 22:01:32.335081');

-- ----------------------------
-- Primary Key structure for table blocks
-- ----------------------------
ALTER TABLE "public"."blocks" ADD CONSTRAINT "blocks_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Foreign Keys structure for table blocks
-- ----------------------------
ALTER TABLE "public"."blocks" ADD CONSTRAINT "blocks_yard_id_fkey" FOREIGN KEY ("yard_id") REFERENCES "public"."yards" ("id") ON DELETE CASCADE ON UPDATE NO ACTION;
