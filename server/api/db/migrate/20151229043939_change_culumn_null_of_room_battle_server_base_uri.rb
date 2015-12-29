class ChangeCulumnNullOfRoomBattleServerBaseUri < ActiveRecord::Migration
  def change
    change_column_null :rooms, :battle_server_base_uri, false
  end
end
