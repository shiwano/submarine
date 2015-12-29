class Battle::CloseRoomController < ApplicationController
  include TyphenApi::Controller::Submarine::Battle::CloseRoom
  include TyphenApiRespondable

  def service
    if target_room.blank?
      raise ApplicationError::RoomNotFound.new("room(#{params.room_id}) not found")
    end

    target_room.destroy!
    render_response({})
  end

  def target_room
    @target_room ||= Room.find_by(id: params.room_id)
  end
end
