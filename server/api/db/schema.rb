# encoding: UTF-8
# This file is auto-generated from the current state of the database. Instead
# of editing this file, please use the migrations feature of Active Record to
# incrementally modify your database, and then regenerate this schema definition.
#
# Note that this schema.rb definition is the authoritative source for your
# database schema. If you need to create the application database on another
# system, you should be using db:schema:load, not running all the migrations
# from scratch. The latter is a flawed and unsustainable approach (the more migrations
# you'll amass, the slower it'll run and the greater likelihood for issues).
#
# It's strongly recommended that you check this file into your version control system.

ActiveRecord::Schema.define(version: 20151229001739) do

  create_table "room_members", force: :cascade do |t|
    t.integer  "user_id",    limit: 4
    t.integer  "room_id",    limit: 4
    t.datetime "created_at",             null: false
    t.datetime "updated_at",             null: false
    t.string   "room_key",   limit: 255, null: false
  end

  add_index "room_members", ["room_id"], name: "index_room_members_on_room_id", using: :btree
  add_index "room_members", ["user_id"], name: "index_room_members_on_user_id", unique: true, using: :btree

  create_table "rooms", force: :cascade do |t|
    t.string   "battle_server_base_uri", limit: 255
    t.integer  "lock_version",           limit: 4
    t.datetime "created_at",                                     null: false
    t.datetime "updated_at",                                     null: false
    t.integer  "room_members_count",     limit: 4,   default: 0
  end

  add_index "rooms", ["room_members_count"], name: "index_rooms_on_room_members_count", using: :btree

  create_table "users", force: :cascade do |t|
    t.string   "name",             limit: 255, null: false
    t.string   "crypted_password", limit: 255
    t.string   "salt",             limit: 255
    t.datetime "created_at"
    t.datetime "updated_at"
    t.integer  "lock_version",     limit: 4
  end

  add_index "users", ["name"], name: "index_users_on_name", unique: true, using: :btree

  add_foreign_key "room_members", "rooms"
  add_foreign_key "room_members", "users"
end
