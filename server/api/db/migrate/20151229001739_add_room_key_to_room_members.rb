class AddRoomKeyToRoomMembers < ActiveRecord::Migration
  def change
    add_column :room_members, :room_key, :string, null: false
  end
end
