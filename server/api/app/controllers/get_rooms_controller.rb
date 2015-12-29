class GetRoomsController < ApplicationController
  include TyphenApi::Controller::Submarine::GetRooms
  include TyphenApiRespondable

  def service
    rooms = joinable_rooms.map { |r| r.to_room_api_type }
    render_response(rooms: rooms)
  end

  def joinable_rooms
    @joinable_rooms ||= Room.joinable.all
  end
end
