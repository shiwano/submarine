FactoryGirl.define do
  factory :access_token do
    user { create(:user) }
    token { SecureRandom.hex(64) }
    expires_at { Time.now + 12.hours }

    trait :expired do
      expires_at { Time.now - 12.hours }
    end
  end
end
