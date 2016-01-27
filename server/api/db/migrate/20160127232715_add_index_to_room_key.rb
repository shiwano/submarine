class AddIndexToRoomKey < ActiveRecord::Migration
  def change
    add_index :room_members, :room_key, unique: true
  end
end
