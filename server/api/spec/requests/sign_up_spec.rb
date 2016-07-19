require 'rails_helper'

RSpec.describe "SignUp", type: :request do
  describe "POST /sign_up" do
    context 'with a valid request' do
      before do
        post sign_up_path
      end

      it "should work" do
        expect(response).to have_http_status(200)
      end
      it "should return the signed up user" do
        expect(parsed_response.user.name).to eq 'NO NAME'
      end
      it "should return the valid auth token" do
        expect(User.find_by_auth_token(parsed_response.auth_token)).to be_a_kind_of User
      end
      it "should return the valid access token" do
        expect(User.find_by_access_token(parsed_response.access_token)).to be_a_kind_of User
      end
    end
  end
end
