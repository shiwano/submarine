class LoginController < ApplicationController
  include TyphenApi::Controller::Submarine::Login
  include TyphenApiValidation

  def service
    # Implement here.
    render_response(user: nil)
  end
end
