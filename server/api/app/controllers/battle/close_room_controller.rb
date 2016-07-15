class Battle::CloseRoomController < ApplicationController
  prepend TyphenApi::Controller::Submarine::Battle::CloseRoom
  prepend TyphenApiRespondable

  def service
    if room.blank?
      raise GameError::RoomNotFound.new("room(#{params.room_id}) not found")
    end

    room.destroy!
    render_response({})
  end

  def room
    @room ||= Room.find_by(id: params.room_id)
  end
end
