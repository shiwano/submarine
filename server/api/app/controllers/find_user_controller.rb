class FindUserController < ApplicationController
  prepend TyphenApi::Controller::Submarine::FindUser
  prepend TyphenApiRespondable

  def service
    user = User.where(name: params.name).first
    render_response(user: user.try(:as_user_api_type))
  end
end
