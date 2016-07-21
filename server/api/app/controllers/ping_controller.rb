class PingController < ApplicationController
  prepend TyphenApi::Controller::Submarine::Ping
  prepend TyphenApiRespondable

  def service
    render(message: params.message + ' PONG')
  end
end
