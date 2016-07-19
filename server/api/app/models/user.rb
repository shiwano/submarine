# == Schema Information
#
# Table name: users
#
#  id                   :integer          not null, primary key
#  name                 :string(255)      not null
#  encrypted_auth_token :string(255)      not null
#  created_at           :datetime
#  updated_at           :datetime
#  lock_version         :integer
#
# Indexes
#
#  index_users_on_encrypted_auth_token  (encrypted_auth_token) UNIQUE
#  index_users_on_name                  (name)
#

class User < ApplicationRecord
  authenticates_with_sorcery!

  has_one :room_member
  has_one :room, through: :room_member

  validates :password, length: { minimum: 6 }
  validates :name, presence: true
  validates :name, uniqueness: true
  validates :name, format: { with: /\A[A-Z\d\-_.]+\z/i }

  def create_room!
    transaction do
      newRoom = Room.create!
      newRoom.join_user!(self)
      newRoom
    end
  end

  def as_user_api_type
    TyphenApi::Model::Submarine::User.new(id: id, name: name)
  end

  def as_logged_in_user_api_type
    TyphenApi::Model::Submarine::LoggedInUser.new(
      id: id,
      name: name,
      joined_room: room ? room.as_joined_room_api_type(self) : nil)
  end
end
