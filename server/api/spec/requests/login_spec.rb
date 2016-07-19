require 'rails_helper'

RSpec.describe 'Login', type: :request do
  describe 'POST /login' do
    let(:request_params) { { auth_token: 'secret' } }

    context 'with a valid request' do
      before do
        @user = create(:user, :with_stupid_auth_token, :with_room)
        post login_path, params: request_params
      end

      it 'should work' do
        expect(response).to have_http_status(200)
      end
      it 'should return the user' do
        expect(parsed_response.user.name).to eq @user.name
      end
      it 'should return the joined room' do
        expect(parsed_response.user.joined_room.id).to eq @user.room.id
      end
    end

    context 'with invalid params' do
      it 'should not work' do
        expect { post(login_path) }.to raise_error(Virtus::CoercionError)
      end
    end

    context 'with an incorrect auth_token' do
      let(:request_params) { { auth_token: 'incorrect' } }

      it 'should raise login error' do
        expect {
          post(login_path, params: request_params)
        }.to raise_error(GameError::LoginFailed)
      end
    end
  end
end
