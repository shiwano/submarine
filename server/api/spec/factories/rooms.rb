FactoryGirl.define do

  factory :room do
    battle_server_base_uri Faker::Internet.ip_v4_address
    lock_version 1
  end

end
