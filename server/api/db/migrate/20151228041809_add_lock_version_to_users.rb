class AddLockVersionToUsers < ActiveRecord::Migration
  def change
    add_column :users, :lock_version, :integer
  end
end
