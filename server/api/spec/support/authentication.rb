module AuthenticationTestHelpers
  module Integration
    attr_reader :current_user

    def login!
      user = create(:user, :with_stupid_auth_token)
      post login_path, params: { auth_token: 'secret' }
      @current_user = user.reload
    end

    def post(path, options=nil)
      if @current_user.present?
        options ||= {}
        options[:headers] ||= {}
        options[:headers]['X-Access-Token'] = @current_user.access_token.token
      end
      super(path, options)
    end
  end
end

shared_context 'with login', with_login: true do
  before do
    login!
  end
end

RSpec.configure do |config|
  config.include AuthenticationTestHelpers::Integration, type: :request
end
