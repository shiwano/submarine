# == Schema Information
#
# Table name: rooms
#
#  id                     :integer          not null, primary key
#  battle_server_base_uri :string(255)      not null
#  created_at             :datetime         not null
#  updated_at             :datetime         not null
#  room_members_count     :integer          default(0)
#
# Indexes
#
#  index_rooms_on_room_members_count  (room_members_count)
#

class Room < ApplicationRecord
  attribute :battle_server_base_uri, :string, default: -> { retrieve_battle_server_base_uri }

  has_many :room_members, dependent: :delete_all
  has_many :users, through: :room_members

  validates :battle_server_base_uri, presence: true
  scope :joinable, -> { where.has { room_members_count < Room.max_room_members_count } }

  def self.max_room_members_count
    4
  end

  def self.retrieve_battle_server_base_uri
    # TODO: battle_server_base_uri is temporary.
    GameConfig.battle_server_base_uri
  end

  def full?
    room_members_count >= Room.max_room_members_count
  end

  def random_room_key
    SecureRandom.urlsafe_base64
  end

  def join_user!(user)
    with_lock do
      if full?
        raise GameError::RoomIsFull.new("room(#{id}) is full")
      end

      begin
        RoomMember.create(room: self, user: user, room_key: random_room_key)
      rescue ActiveRecord::RecordNotUnique
        raise GameError::RoomAlreadyJoined.new("user(#{user.id}) has already joined")
      end
    end
  end

  def as_room_api_type
    TyphenApi::Model::Submarine::Room.new(
      id: id,
      members: users.map { |u| u.as_user_api_type })
  end

  def as_battle_room_api_type
    TyphenApi::Model::Submarine::Battle::Room.new(
      id: id,
      members: room_members.includes(:user).map { |m|
        m.as_battle_room_member_api_type
      })
  end

  def as_joined_room_api_type(user)
    TyphenApi::Model::Submarine::JoinedRoom.new(
      id: id,
      members: users.map { |u| u.as_user_api_type },
      room_key: room_members.includes(:user).find { |m| m.user == user }.room_key,
      battle_server_base_uri: battle_server_base_uri)
  end
end
