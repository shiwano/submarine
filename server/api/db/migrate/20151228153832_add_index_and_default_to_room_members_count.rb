class AddIndexAndDefaultToRoomMembersCount < ActiveRecord::Migration
  def change
    change_column_default :rooms, :room_members_count, 0
    add_index :rooms, :room_members_count
  end
end
