class LoginController < ApplicationController
  include TyphenApi::Controller::Submarine::Login
  include TyphenApiValidation

  def service
    if logged_in_user.blank?
      raise ApplicationError::Login.new('The user name or password is incorrect')
    end

    render_response(user: logged_in_user.to_api_type)
  end

  def logged_in_user
    @logged_in_user ||= login(params.name, params.password)
  end
end
