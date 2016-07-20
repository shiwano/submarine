class CreateRoomController < ApplicationController
  prepend TyphenApi::Controller::Submarine::CreateRoom
  prepend TyphenApiRespondable

  def service
    room = current_user.create_room!
    render_response(room: room.as_joined_room_api_type(current_user))
  end
end
