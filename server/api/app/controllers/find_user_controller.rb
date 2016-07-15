class FindUserController < ApplicationController
  prepend TyphenApi::Controller::Submarine::FindUser
  prepend TyphenApiRespondable

  def service
    render_response(user: user.try(:as_user_api_type))
  end

  def user
    @user ||= User.where(name: params.name).first
  end
end
