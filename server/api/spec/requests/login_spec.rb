require 'rails_helper'

RSpec.describe 'Login', type: :request do
  describe 'POST /login' do
    let(:user) { create(:user, :with_stupid_password) }
    let(:room) { create(:room) }
    let(:params) { { name: user.name, password: 'secret' } }

    context 'with a valid request' do
      before do
        room.join_user!(user)
        post(login_path, params)
      end

      it 'should work' do
        expect(response).to have_http_status(200)
      end
      it 'should return a response that includes the user' do
        expect(response_json[:user][:name]).to eq user.name
      end
      it 'should return a response that includes the joined room' do
        expect(response_json[:user][:joined_room][:id]).to eq room.id
      end
      it 'should return a response that includes the session cookie' do
        expect(response.headers['Set-Cookie']).to include('submarine_api_session')
      end
    end

    context 'with invalid params' do
      it 'should not work' do
        expect { post(login_path) }.to raise_error(Virtus::CoercionError)
      end
    end

    context 'with an incorrect password' do
      let(:params) { { name: user.name, password: 'incorrect' } }

      it 'should raise login error' do
        expect { post(login_path, params) }.to raise_error(GameError::LoginFailed)
      end
    end

    context 'with an incorrect user name' do
      let(:params) { { name: 'unknown', password: 'secret' } }

      it 'should raise login error' do
        expect { post(login_path, params) }.to raise_error(GameError::LoginFailed)
      end
    end
  end
end
