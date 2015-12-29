FactoryGirl.define do

  factory :room_member do
    user nil
    room nil
    room_key Faker::Internet.password
  end

end
