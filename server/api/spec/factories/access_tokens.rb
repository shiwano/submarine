FactoryGirl.define do
  factory :access_token do
    user nil
    token "MyString"
    expires_at "2016-07-20 00:40:25"
  end
end
