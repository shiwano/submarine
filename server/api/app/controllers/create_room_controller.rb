class CreateRoomController < ApplicationController
  prepend TyphenApi::Controller::Submarine::CreateRoom
  prepend TyphenApiRespondable

  def service
    render_response(room: new_room.as_joined_room_api_type(current_user))
  end

  def new_room
    @new_room ||= current_user.create_room!
  end
end
