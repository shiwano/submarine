class FindUserController < ApplicationController
  include TyphenApi::Controller::Submarine::FindUser
  include TyphenApiValidation

  def service
    # Implement here.
    render_response(user: nil)
  end
end
