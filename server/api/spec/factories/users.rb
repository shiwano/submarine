FactoryGirl.define do

  factory :user do
    name { Faker::Name.first_name }
    auth_token { SecureRandom.hex(64) }
    lock_version 1

    trait :with_stupid_auth_token do
      auth_token 'secret'
    end

    trait :with_room do
      after(:create) do |user|
        user.create_room!
      end
    end
  end

end
