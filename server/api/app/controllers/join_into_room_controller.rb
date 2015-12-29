class JoinIntoRoomController < ApplicationController
  include TyphenApi::Controller::Submarine::JoinIntoRoom
  include TyphenApiRespondable

  def service
    if target_room.blank?
      raise ApplicationError::RoomNotFound.new("room(#{params.room_id}) not found")
    end

    target_room.join_user!(current_user)
    render_response(room: target_room.to_joined_room_api_type(current_user))
  end

  def target_room
    @target_room ||= Room.joinable.find_by(id: params.room_id)
  end
end
