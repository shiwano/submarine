class AddUniqueToRoomMembersUser < ActiveRecord::Migration
  def change
    remove_foreign_key :room_members, :user
    remove_index :room_members, :user_id

    add_index :room_members, :user_id, unique: true
    add_foreign_key :room_members, :users
  end
end
