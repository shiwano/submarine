# == Schema Information
#
# Table name: rooms
#
#  id                     :integer          not null, primary key
#  battle_server_base_uri :string(255)
#  lock_version           :integer
#  created_at             :datetime         not null
#  updated_at             :datetime         not null
#  room_members_count     :integer          default(0)
#
# Indexes
#
#  index_rooms_on_room_members_count  (room_members_count)
#

class Room < ActiveRecord::Base
  has_many :room_members, dependent: :delete_all
  has_many :users, through: :room_members

  validates :battle_server_base_uri, presence: true

  def max_member_count
    4
  end

  def full?
    users.count < max_member_count
  end

  def join_user(user)
    with_lock do
      unless full?
        raise ApplicationError::RoomIsFull.new("room(#{id}) is full")
      end

      begin
        RoomMember.create(room: self, user: user)
      rescue ActiveRecord::RecordNotUnique
        raise ApplicationError::RoomAlreadyJoined.new("user has already joined")
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
