class ChangeUsers < ActiveRecord::Migration[5.0]
  def change
    remove_column :users, :salt
    rename_column :users, :crypted_password, :encrypted_auth_token
    change_column_null :users, :encrypted_auth_token, false
    add_index :users, :encrypted_auth_token, unique: true
    remove_index :users, :name
    add_index :users, :name
  end
end
