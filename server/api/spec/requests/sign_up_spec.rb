require 'rails_helper'

RSpec.describe "SignUp", type: :request do
  describe "POST /sign_up" do
    let(:request_params) { { name: 'Kaga' } }

    context 'with a valid request' do
      before do
        post sign_up_path, params: request_params
      end

      it "should work" do
        expect(response).to have_http_status(200)
      end
      it "should return the signed up user" do
        expect(parsed_response.user.name).to eq 'Kaga'
      end
      it "should return the valid auth token" do
        expect(User.find_by_auth_token(parsed_response.auth_token)).to be_a_kind_of User
      end
      it "should return the valid access token" do
        expect(User.find_by_access_token(parsed_response.access_token)).to be_a_kind_of User
      end
    end

    context 'with invalid request_params' do
      it 'should not work' do
        expect { post(sign_up_path) }.to raise_error(Virtus::CoercionError)
      end
    end
  end
end
