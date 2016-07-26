require 'rails_helper'

RSpec.describe ApplicationController, type: :controller do
  controller(ApplicationController) do
    def no_authentication_required?
      false
    end

    def service
      render body: 'OK'
    end
  end

  before do
    routes.draw { post 'test' => 'anonymous#service' }
  end

  context 'with login' do
    before do
      @current_user = create(:user, :with_stupid_auth_token)
      @access_token = @current_user.generate_access_token!
      request.headers['X-Access-Token'] = @access_token
    end

    describe '#generate_access_token_if_needed' do
      context 'when the current access token will not expires soon' do
        it 'should not generate access token' do
          post :service
          expect(response.headers['X-Set-Access-Token']).to be nil
        end
      end

      context 'when the current access token will expires soon' do
        before { Timecop.freeze(@current_user.access_token.expires_at - 10.seconds) }
        after { Timecop.return }

        it 'should generate access token' do
          expected_access_token = 'NEW-ACCESS-TOKEN'
          expect_any_instance_of(User).to \
            receive(:generate_access_token!).and_return expected_access_token
          post :service
          expect(response.headers['X-Set-Access-Token']).to eq expected_access_token
        end
      end
    end
  end

  context 'without login' do
    it 'should raise error' do
      expect { post :service }.to raise_error GameError::NotAuthenticated
    end
  end
end
