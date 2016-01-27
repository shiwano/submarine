class FindUserController < ApplicationController
  include TyphenApi::Controller::Submarine::FindUser
  include TyphenApiRespondable

  def service
    render_response(user: target_user.try(:as_user_api_type))
  end

  def target_user
    @target_user ||= User.where(name: params.name).first
  end
end
