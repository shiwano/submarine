class PingController < ApplicationController
  include TyphenApi::Controller::Submarine::Ping
  include TyphenApiValidation

  def service
    render_response(message: params.message + ' PONG')
  end
end
