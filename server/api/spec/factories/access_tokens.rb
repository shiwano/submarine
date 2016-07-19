FactoryGirl.define do
  factory :access_token do
    user { create(:user) }
    token { SecureRandom.hex(64) }
    expires_at '2016-07-20 00:40:25'
  end
end
