class SignUpController < ApplicationController
  prepend TyphenApi::Controller::Submarine::SignUp
  prepend TyphenApiRespondable

  def service
    user = User.new(name: params.name)
    auth_token = user.generate_auth_token!
    access_token = user.generate_access_token!

    render({
      user: user.as_logged_in_user_api_type,
      auth_token: auth_token,
      access_token: access_token,
    })
  end
end
