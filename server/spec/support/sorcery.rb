module SorceryHelpers
  module Request
    def login_user(user)
      post login_path, name: user.name, password: 'secret'
    end
  end
end

RSpec.configure do |config|
  config.include Sorcery::TestHelpers::Rails::Controller, type: :controller
  config.include SorceryHelpers::Request, type: :request
end
