class FindUserController < ApplicationController
  include TyphenApi::Controller::Submarine::FindUser
  include TyphenApiValidation

  def service
    render_response(user: target_user.try(:to_api_type))
  end

  def target_user
    @target_user ||= User.where(name: params.name).first
  end
end
