class CreateRoomController < ApplicationController
  include TyphenApi::Controller::Submarine::CreateRoom
  include TyphenApiRespondable

  def service
    # Implement here.
    render_response(room: nil)
  end
end
