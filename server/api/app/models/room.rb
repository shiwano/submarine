# == Schema Information
#
# Table name: rooms
#
#  id                     :integer          not null, primary key
#  battle_server_base_uri :string(255)      not null
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

  def max_room_members_count
    4
  end

  def full?
    room_members_count < max_room_members_count
  end

  def random_room_key
    SecureRandom.urlsafe_base64
  end

  def join_user(user)
    with_lock do
      unless full?
        raise ApplicationError::RoomIsFull.new("room(#{id}) is full")
      end

      begin
        RoomMember.create(room: self, user: user, room_key: random_room_key)
      rescue ActiveRecord::RecordNotUnique
        raise ApplicationError::RoomAlreadyJoined.new("user(#{user.id}) has already joined")
      end
    end
  end

  def to_room_api_type
    TyphenApi::Model::Submarine::Room.new(
      id: id,
      members: users.map { |u| u.to_api_type })
  end

  def to_joined_room_api_type(user)
    TyphenApi::Model::Submarine::JoinedRoom.new(
      id: id,
      members: users.map { |u| u.to_user_api_type },
      room_key: room_members.find { |m| m.user == user }.room_key,
      battle_server_base_uri: battle_server_base_uri)
  end
end
