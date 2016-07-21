class Battle::FindRoomController < ApplicationController
  prepend TyphenApi::Controller::Submarine::Battle::FindRoom
  prepend TyphenApiRespondable

  def service
    room = Room.find_by(id: params.room_id)
    render(room: room.try(:as_battle_room_api_type))
  end
end
