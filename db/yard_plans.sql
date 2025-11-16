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

 Date: 16/11/2025 21:18:18
*/


-- ----------------------------
-- Table structure for yard_plans
-- ----------------------------
DROP TABLE IF EXISTS "public"."yard_plans";
CREATE TABLE "public"."yard_plans" (
  "id" int4 NOT NULL DEFAULT nextval('yard_plans_id_seq'::regclass),
  "block_id" int4 NOT NULL,
  "slot" int4 NOT NULL,
  "row" int4 NOT NULL,
  "tier" int4 NOT NULL,
  "container_number" varchar(20) COLLATE "pg_catalog"."default",
  "container_size" int4,
  "container_height" numeric,
  "container_type" varchar(10) COLLATE "pg_catalog"."default",
  "created_at" timestamp(6) DEFAULT now(),
  "updated_at" timestamp(6) DEFAULT now(),
  "is_picked" bool DEFAULT false
)
;

-- ----------------------------
-- Records of yard_plans
-- ----------------------------
INSERT INTO "public"."yard_plans" VALUES (10, 1, 1, 2, 1, 'ALFI000001', 40, 0, '', '2025-11-16 17:14:06.417006', '2025-11-16 17:14:06.417006', 'f');
INSERT INTO "public"."yard_plans" VALUES (11, 1, 2, 2, 1, 'ALFI000001', 40, 0, '', '2025-11-16 17:14:06.417006', '2025-11-16 17:14:06.417006', 'f');
INSERT INTO "public"."yard_plans" VALUES (12, 1, 3, 1, 1, 'ALFI000003', 20, 0, '', '2025-11-16 17:17:00.432584', '2025-11-16 17:17:00.432584', 'f');
INSERT INTO "public"."yard_plans" VALUES (13, 1, 3, 2, 1, 'ALFI000004', 20, 0, '', '2025-11-16 17:20:45.224494', '2025-11-16 17:20:45.224494', 'f');
INSERT INTO "public"."yard_plans" VALUES (2, 2, 2, 1, 1, 'ALFI000012', 40, 0, '', '2025-11-16 17:06:07.663506', '2025-11-16 17:06:07.663506', 'f');
INSERT INTO "public"."yard_plans" VALUES (3, 2, 1, 2, 1, 'ALFI000011', 40, 0, '', '2025-11-16 17:14:06.417006', '2025-11-16 17:14:06.417006', 'f');
INSERT INTO "public"."yard_plans" VALUES (4, 2, 2, 2, 1, 'ALFI000011', 40, 0, '', '2025-11-16 17:14:06.417006', '2025-11-16 17:14:06.417006', 'f');
INSERT INTO "public"."yard_plans" VALUES (5, 2, 3, 1, 1, 'ALFI000013', 20, 0, '', '2025-11-16 17:17:00.432584', '2025-11-16 17:17:00.432584', 'f');
INSERT INTO "public"."yard_plans" VALUES (1, 2, 1, 1, 1, 'ALFI000012', 40, 0, '', '2025-11-16 17:06:07.663506', '2025-11-16 17:06:07.663506', 'f');
INSERT INTO "public"."yard_plans" VALUES (6, 2, 3, 2, 1, 'ALFI000014', 20, 0, '', '2025-11-16 17:20:45.224494', '2025-11-16 17:20:45.224494', 'f');
INSERT INTO "public"."yard_plans" VALUES (9, 1, 2, 1, 1, 'ALFI000002', 40, 0, '', '2025-11-16 17:06:07.663506', '2025-11-16 17:06:07.663506', 'f');
INSERT INTO "public"."yard_plans" VALUES (8, 1, 1, 1, 1, 'ALFI000002', 40, 0, '', '2025-11-16 17:06:07.663506', '2025-11-16 17:06:07.663506', 'f');

-- ----------------------------
-- Uniques structure for table yard_plans
-- ----------------------------
ALTER TABLE "public"."yard_plans" ADD CONSTRAINT "yard_plans_block_id_slot_row_tier_key" UNIQUE ("block_id", "slot", "row", "tier");

-- ----------------------------
-- Primary Key structure for table yard_plans
-- ----------------------------
ALTER TABLE "public"."yard_plans" ADD CONSTRAINT "yard_plans_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Foreign Keys structure for table yard_plans
-- ----------------------------
ALTER TABLE "public"."yard_plans" ADD CONSTRAINT "yard_plans_block_id_fkey" FOREIGN KEY ("block_id") REFERENCES "public"."blocks" ("id") ON DELETE CASCADE ON UPDATE NO ACTION;
