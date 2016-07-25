class LoginController < ApplicationController
  prepend TyphenApi::Controller::Submarine::Login
  prepend TyphenApiRespondable

  def service
    user = User.find_by_auth_token(params.auth_token)
    raise GameError::LoginFailed.new('The auth token is incorrect') if user.nil?

    set_access_token(user.generate_access_token!)
    render(user: user.as_logged_in_user_api_type)
  end
end
