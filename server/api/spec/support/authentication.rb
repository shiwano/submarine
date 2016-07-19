module AuthenticationTestHelpers
  module Controller
    def login!
      user = create(:user, :with_stupid_auth_token)
      request.headers['X-Access-Token'] = user.generate_access_token!
      user
    end
  end

  module Integration
    def login!
      user = create(:user, :with_stupid_auth_token)
      post login_path, params: { auth_token: 'secret' }
      request.headers['X-Access-Token'] = user.access_token.token
      user
    end
  end
end

shared_context 'with login', with_login: true do
  let(:current_user) { @current_user }

  before do
    @current_user = login!
  end
end

RSpec.configure do |config|
  config.include AuthenticationTestHelpers::Controller, type: :controller
  config.include AuthenticationTestHelpers::Integration, type: :request
end
