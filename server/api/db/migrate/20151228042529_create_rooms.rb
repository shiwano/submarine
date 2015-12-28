class CreateRooms < ActiveRecord::Migration
  def change
    create_table :rooms do |t|
      t.string :battle_server_base_uri
      t.integer :lock_version

      t.timestamps null: false
    end
  end
end
