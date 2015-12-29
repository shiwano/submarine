FactoryGirl.define do

  factory :room do
    battle_server_base_uri Faker::Internet.ip_v4_address
    lock_version 1

    trait :with_user do
      after(:create) do |room|
        user = create(:user)
        room.join_user!(user)
      end
    end
  end

end
