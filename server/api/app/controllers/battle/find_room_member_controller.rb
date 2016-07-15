class Battle::FindRoomMemberController < ApplicationController
  prepend TyphenApi::Controller::Submarine::Battle::FindRoomMember
  prepend TyphenApiRespondable

  def service
    render_response(room_member: room_member.try(:as_battle_room_member_api_type))
  end

  def room_member
    @room_member ||= RoomMember.find_by(room_key: params.room_key)
  end
end
