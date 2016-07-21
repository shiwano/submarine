class JoinIntoRoomController < ApplicationController
  prepend TyphenApi::Controller::Submarine::JoinIntoRoom
  prepend TyphenApiRespondable

  def service
    room = Room.joinable.find_by(id: params.room_id)
    raise GameError::RoomNotFound.new("room(#{params.room_id}) not found") if room.nil?

    room.join_user!(current_user)
    render(room: room.as_joined_room_api_type(current_user))
  end
end
