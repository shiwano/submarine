# == Schema Information
#
# Table name: rooms
#
#  id                     :integer          not null, primary key
#  battle_server_base_uri :string(255)
#  lock_version           :integer
#  created_at             :datetime         not null
#  updated_at             :datetime         not null
#  room_members_count     :integer
#

class Room < ActiveRecord::Base
  has_many :room_members, dependent: :delete_all
  has_many :users, through: :room_members

  def max_member_count
    4
  end

  def full?
    users.count < max_member_count
  end

  def join(user)
    with_lock do
      if full?
        RoomMember.create(room: self, user: user)
      else
        raise ApplicationError::RoomIsFull.new("room(#{id}) is full")
      end
    end
  end

  def to_api_type
    TyphenApi::Model::Submarine::Room.new(
      id: id,
      battle_server_base_uri: battle_server_base_uri,
      members: users.map { |u| u.to_api_type })
  end
end
