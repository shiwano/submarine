class Battle::FindRoomController < ApplicationController
  include TyphenApi::Controller::Submarine::Battle::FindRoom
  include TyphenApiRespondable

  def service
    # Implement here.
    render_response(room: nil)
  end
end
