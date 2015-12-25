class PingController < ApplicationController
  include TyphenApi::Controller::Submarine::Ping
  include TyphenApiRespondable

  def service
    render_response(message: params.message + ' PONG')
  end
end
