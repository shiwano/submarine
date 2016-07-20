# == Schema Information
#
# Table name: room_members
#
#  id         :integer          not null, primary key
#  user_id    :integer
#  room_id    :integer
#  created_at :datetime         not null
#  updated_at :datetime         not null
#  room_key   :string(255)      not null
#
# Indexes
#
#  index_room_members_on_room_id   (room_id)
#  index_room_members_on_room_key  (room_key) UNIQUE
#  index_room_members_on_user_id   (user_id) UNIQUE
#

class RoomMember < ApplicationRecord
  belongs_to :user
  belongs_to :room, :counter_cache => true

  validates :room_key, presence: true

  def as_battle_room_member_api_type
    TyphenApi::Model::Submarine::Battle::RoomMember.new({
      name: user.name,
      id: user.id,
      room_id: room.id,
    })
  end
end
