class Battle::CloseRoomController < ApplicationController
  include TyphenApi::Controller::Submarine::Battle::CloseRoom
  include TyphenApiRespondable

  def service
    # Implement here.
    render_response()
  end
end
