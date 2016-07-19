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
  has_one :room_member
  has_one :room, through: :room_member

  validates :encrypted_auth_token, presence: true
  validates :encrypted_auth_token, uniqueness: true
  validates :name, presence: true
  validates :name, length: { maximum: 16 }

  def self.salt
    Rails.application.secrets[:secret_key_base]
  end

  def self.encrypt_auth_token(auth_token)
    Digest::SHA512.hexdigest(auth_token + salt)
  end

  def generate_auth_token
    self.auth_token = SecureRandom.hex(64)
  end

  def auth_token=(auth_token)
    self.encrypted_auth_token = self.class.encrypt_auth_token(auth_token)
    auth_token
  end

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
