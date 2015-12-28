class CreateRoomMembers < ActiveRecord::Migration
  def change
    create_table :room_members do |t|
      t.references :user, index: true, foreign_key: true
      t.references :room, index: true, foreign_key: true

      t.timestamps null: false
    end
  end
end
