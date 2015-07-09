class SignUpController < ApplicationController
  include TyphenApi::Controller::Submarine::SignUp
  include TyphenApiValidation

  def service
    # Implement here.
    render_response(user: nil)
  end
end
