class GetRoomsController < ApplicationController
  prepend TyphenApi::Controller::Submarine::GetRooms
  prepend TyphenApiRespondable

  def service
    if current_user.room.present?
      rooms = []
    else
      rooms = Room.joinable.all
    end
    render(rooms: rooms.map { |r| r.as_room_api_type })
  end
end
