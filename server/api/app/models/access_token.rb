# == Schema Information
#
# Table name: access_tokens
#
#  id         :integer          not null, primary key
#  user_id    :integer          not null
#  token      :string(255)      not null
#  expires_at :datetime         not null
#  created_at :datetime         not null
#  updated_at :datetime         not null
#
# Indexes
#
#  index_access_tokens_on_token    (token) UNIQUE
#  index_access_tokens_on_user_id  (user_id) UNIQUE
#

class AccessToken < ApplicationRecord
  EXPIRATION_TIME = 1.day
  EXPIRATION_WARNING_TIME = 1.hour

  belongs_to :user

  validates :user, presence: true
  validates :user, uniqueness: true
  validates :token, presence: true
  validates :token, uniqueness: true
  validates :expires_at, presence: true

  scope :no_expired, -> { where.has { expires_at > Time.now } }

  def self.find_by_token(token)
    AccessToken.includes(:user).no_expired.where(token: token).limit(1).first
  end

  def generate_token
    self.token = SecureRandom.hex(32)
    self.expires_at = Time.now + EXPIRATION_TIME
    token
  end

  def expired?
    expires_at <= Time.now
  end

  def expires_soon?
    !expired? && expires_at <= Time.now + EXPIRATION_WARNING_TIME
  end
end
