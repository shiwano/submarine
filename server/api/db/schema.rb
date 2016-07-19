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

ActiveRecord::Schema.define(version: 20160719154025) do

  create_table "access_tokens", force: :cascade, options: "ENGINE=InnoDB DEFAULT CHARSET=utf8" do |t|
    t.integer  "user_id",    null: false
    t.string   "token",      null: false
    t.datetime "expires_at", null: false
    t.datetime "created_at", null: false
    t.datetime "updated_at", null: false
    t.index ["token"], name: "index_access_tokens_on_token", unique: true, using: :btree
    t.index ["user_id"], name: "index_access_tokens_on_user_id", unique: true, using: :btree
  end

  create_table "room_members", force: :cascade, options: "ENGINE=InnoDB DEFAULT CHARSET=utf8" do |t|
    t.integer  "user_id"
    t.integer  "room_id"
    t.datetime "created_at", null: false
    t.datetime "updated_at", null: false
    t.string   "room_key",   null: false
    t.index ["room_id"], name: "index_room_members_on_room_id", using: :btree
    t.index ["room_key"], name: "index_room_members_on_room_key", unique: true, using: :btree
    t.index ["user_id"], name: "index_room_members_on_user_id", unique: true, using: :btree
  end

  create_table "rooms", force: :cascade, options: "ENGINE=InnoDB DEFAULT CHARSET=utf8" do |t|
    t.string   "battle_server_base_uri",             null: false
    t.datetime "created_at",                         null: false
    t.datetime "updated_at",                         null: false
    t.integer  "room_members_count",     default: 0
    t.index ["room_members_count"], name: "index_rooms_on_room_members_count", using: :btree
  end

  create_table "users", force: :cascade, options: "ENGINE=InnoDB DEFAULT CHARSET=utf8" do |t|
    t.string   "name",                 null: false
    t.string   "encrypted_auth_token", null: false
    t.datetime "created_at"
    t.datetime "updated_at"
    t.integer  "lock_version"
    t.index ["encrypted_auth_token"], name: "index_users_on_encrypted_auth_token", unique: true, using: :btree
    t.index ["name"], name: "index_users_on_name", using: :btree
  end

  add_foreign_key "access_tokens", "users"
  add_foreign_key "room_members", "rooms"
  add_foreign_key "room_members", "users"
end
