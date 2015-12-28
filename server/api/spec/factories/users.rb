FactoryGirl.define do

  factory :user do
    name { Faker::Name.first_name }
    password Faker::Internet.password(6, 20)
    lock_version 1

    trait :with_stupid_password do
      password 'secret'
    end
  end

end
