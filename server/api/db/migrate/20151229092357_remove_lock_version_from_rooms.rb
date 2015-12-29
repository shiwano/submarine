class RemoveLockVersionFromRooms < ActiveRecord::Migration
  def change
    remove_column :rooms, :lock_version
  end
end
