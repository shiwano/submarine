FactoryGirl.define do

  factory :user do
    name { Faker::Internet.user_name(separators: %w(._-)) }
    password Faker::Internet.password(6, 20)
    lock_version 1

    trait :with_stupid_password do
      password 'secret'
    end
  end

end
