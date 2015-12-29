class Battle::FindRoomController < ApplicationController
  include TyphenApi::Controller::Submarine::Battle::FindRoom
  include TyphenApiRespondable

  def service
    render_response(room: target_room.try(:to_battle_room_api_type))
  end

  def target_room
    @target_room ||= Room.find_by(id: params.room_id)
  end
end
