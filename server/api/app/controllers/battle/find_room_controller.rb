class Battle::FindRoomController < ApplicationController
  include TyphenApi::Controller::Submarine::Battle::FindRoom
  include TyphenApiRespondable

  def service
    render_response(room: room.try(:as_battle_room_api_type))
  end

  def room
    @room ||= Room.find_by(id: params.room_id)
  end
end
