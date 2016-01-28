class Battle::CloseRoomController < ApplicationController
  include TyphenApi::Controller::Submarine::Battle::CloseRoom
  include TyphenApiRespondable

  def service
    if room.blank?
      raise ApplicationError::RoomNotFound.new("room(#{params.room_id}) not found")
    end

    room.destroy!
    render_response({})
  end

  def room
    @room ||= Room.find_by(id: params.room_id)
  end
end
