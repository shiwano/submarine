class Battle::CloseRoomController < ApplicationController
  prepend TyphenApi::Controller::Submarine::Battle::CloseRoom
  prepend TyphenApiRespondable

  def service
    room = Room.find_by(id: params.room_id)
    raise GameError::RoomNotFound.new("room(#{params.room_id}) not found") if room.nil?

    room.destroy!
    render_response({})
  end
end
