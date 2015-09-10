class SorceryCore < ActiveRecord::Migration
  def change
    create_table :users do |t|
      t.string :name,            :null => false
      t.string :crypted_password
      t.string :salt

      t.timestamps
    end

    add_index :users, :name, unique: true
  end
end
