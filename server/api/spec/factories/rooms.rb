FactoryGirl.define do

  factory :room do
    battle_server_base_uri Faker::Internet.ip_v4_address

    trait :with_user do
      after(:create) do |room|
        user = create(:user)
        room.join_user!(user)
      end
    end

    trait :full do
      after(:create) do |room|
        Room.max_room_members_count.times do
          user = create(:user)
          room.join_user!(user)
        end
      end
    end
  end

end
