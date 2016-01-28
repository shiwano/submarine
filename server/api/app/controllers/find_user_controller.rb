class FindUserController < ApplicationController
  include TyphenApi::Controller::Submarine::FindUser
  include TyphenApiRespondable

  def service
    render_response(user: user.try(:as_user_api_type))
  end

  def user
    @user ||= User.where(name: params.name).first
  end
end
