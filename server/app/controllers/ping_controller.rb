class PingController < ApplicationController
  include TyphenApi::Controller::Submarine::Ping
  include TyphenApiValidation

  def service
    # Implement here.
    render_response(message: nil)
  end
end
