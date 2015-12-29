class CreateRoomController < ApplicationController
  include TyphenApi::Controller::Submarine::CreateRoom
  include TyphenApiRespondable

  def service
    render_response(room: new_room.to_joined_room_api_type(current_user))
  end

  def new_room
    @new_room ||= current_user.create_room(battle_server_base_uri: battle_server_base_uri)
  end

  def battle_server_base_uri
    # TODO: battle_server_base_uri is temporary.
    'ws://localhost:5000'
  end
end
