class JoinIntoRoomController < ApplicationController
  prepend TyphenApi::Controller::Submarine::JoinIntoRoom
  prepend TyphenApiRespondable

  def service
    if room.blank?
      raise GameError::RoomNotFound.new("room(#{params.room_id}) not found")
    end

    room.join_user!(current_user)
    render_response(room: room.as_joined_room_api_type(current_user))
  end

  def room
    @room ||= Room.joinable.find_by(id: params.room_id)
  end
end
