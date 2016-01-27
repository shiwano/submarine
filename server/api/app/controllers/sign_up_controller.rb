class SignUpController < ApplicationController
  include TyphenApi::Controller::Submarine::SignUp
  include TyphenApiRespondable

  def service
    login(new_user.name, params.password)
    render_response(user: new_user.as_logged_in_user_api_type)
  end

  def new_user
    @new_user ||= User.create!(name: params.name, password: params.password)
  end
end
