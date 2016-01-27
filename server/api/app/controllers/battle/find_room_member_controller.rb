class Battle::FindRoomMemberController < ApplicationController
  include TyphenApi::Controller::Submarine::Battle::FindRoomMember
  include TyphenApiRespondable

  def service
    # Implement here.
    render_response(room_member: nil)
  end
end
