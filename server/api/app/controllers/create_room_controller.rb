class CreateRoomController < ApplicationController
  include TyphenApi::Controller::Submarine::CreateRoom
  include TyphenApiRespondable

  def service
    render_response(room: new_room.as_joined_room_api_type(current_user))
  end

  def new_room
    @new_room ||= current_user.create_room!
  end
end
