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
  belongs_to :user

  validates :user, presence: true
  validates :user, uniqueness: true
  validates :token, presence: true
  validates :token, uniqueness: true
  validates :expires_at, presence: true

  def generate_token
    self.token = SecureRandom.hex(64)
    self.expires_at = Time.now + 1.day
    self.token
  end

  def expired?
    Time.now >= expires_at
  end
end
