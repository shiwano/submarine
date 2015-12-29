FactoryGirl.define do

  factory :room_member do
    user { create(:user) }
    room { create(:room) }
    room_key Faker::Internet.password
  end

end
