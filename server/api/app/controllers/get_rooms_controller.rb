class GetRoomsController < ApplicationController
  prepend TyphenApi::Controller::Submarine::GetRooms
  prepend TyphenApiRespondable

  def service
    rooms = joinable_rooms.map { |r| r.as_room_api_type }
    render_response(rooms: rooms)
  end

  def joinable_rooms
    @joinable_rooms ||= current_user.room.present? ? [] : Room.joinable.all
  end
end
