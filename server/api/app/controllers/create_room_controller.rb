class CreateRoomController < ApplicationController
  include TyphenApi::Controller::Submarine::CreateRoom
  include TyphenApiRespondable

  def service
    render_response(room: new_room.to_api_type, room_key: current_user.room_member.room_key)
  end

  def new_room
    @new_room ||= current_user.create_room(battle_server_base_uri: battle_server_base_uri)
  end

  def battle_server_base_uri
    # TODO: battle_server_base_uri is temporary.
    'ws://localhost:5000'
  end
end
