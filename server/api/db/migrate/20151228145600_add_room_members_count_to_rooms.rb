class AddRoomMembersCountToRooms < ActiveRecord::Migration
  def change
    add_column :rooms, :room_members_count, :integer
  end
end
