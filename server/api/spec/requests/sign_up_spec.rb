require 'rails_helper'

RSpec.describe "SignUp", type: :request do
  describe "POST /sign_up" do
    let(:params) { { name: 'Kongou', password: 'KantaiCollection' } }

    context 'with a valid request' do
      before do
        post(sign_up_path, params)
      end

      it "should work" do
        expect(response).to have_http_status(200)
      end
      it "should return a reasponse that includes a user" do
        expect(response_json[:user][:name]).to eq 'Kongou'
      end
    end

    context 'with invalid params' do
      it "should not work" do
        expect { post(sign_up_path) }.to raise_error(Virtus::CoercionError)
      end
    end
  end
end
