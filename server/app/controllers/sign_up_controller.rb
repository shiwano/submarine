class SignUpController < ApplicationController
  include TyphenApi::Controller::Submarine::SignUp
  include TyphenApiValidation

  def service
    new_user.save
    render_response(user: new_user.to_api_type)
  end

  def new_user
    @new_user ||= User.new(name: params.name, password: params.password)
  end
end
