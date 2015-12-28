class GetRoomsController < ApplicationController
  include TyphenApi::Controller::Submarine::GetRooms
  include TyphenApiRespondable

  def service
    # Implement here.
    render_response(rooms: nil)
  end
end
