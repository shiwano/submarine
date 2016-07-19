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
end
