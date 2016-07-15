class PingController < ApplicationController
  prepend TyphenApi::Controller::Submarine::Ping
  prepend TyphenApiRespondable

  def service
    render_response(message: params.message + ' PONG')
  end
end
