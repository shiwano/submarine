class JoinIntoRoomController < ApplicationController
  include TyphenApi::Controller::Submarine::JoinIntoRoom
  include TyphenApiRespondable

  def service
    # Implement here.
    render_response()
  end
end
