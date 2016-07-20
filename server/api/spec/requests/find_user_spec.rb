require 'rails_helper'

RSpec.describe 'FindUser', type: :request do
  describe 'POST /find_user', with_login: true do
    let(:request_params) { { name: 'Shimakaze' } }

    before do
      create(:user, name: 'Shimakaze')
      create(:user, name: 'Yamato')
    end

    context 'with a valid request' do
      before do
        post(find_user_path, params: request_params)
      end

      it 'should work' do
        expect(response).to have_http_status(200)
      end
      it 'should return a reasponse that includes a user' do
        expect(parsed_response.user.name).to eq 'Shimakaze'
      end
    end

    context 'with invalid request_params' do
      it 'should not work' do
        expect { post(find_user_path) }.to raise_error(Virtus::CoercionError)
      end
    end

    context 'with a no-existing user name' do
      let(:request_params) { { name: 'Kaga' } }

      it 'should return empty' do
        post(find_user_path, params: request_params)
        expect(parsed_response.user).to be_nil
      end
    end
  end
end
