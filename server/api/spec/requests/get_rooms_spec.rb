require 'rails_helper'

RSpec.describe 'GetRooms', type: :request do
  describe 'POST /get_rooms' do
    before do
      user = create :user, :with_stupid_password
      login_user(user)
    end

    context 'with a valid request' do
      before do
        create(:room)
        create(:room)
        create(:room, :full)
        post get_rooms_path
      end

      it 'should work' do
        expect(response).to have_http_status(200)
      end
      it 'should return joinable rooms' do
        expect(response_json[:rooms].length).to eq 2
      end
    end
  end
end
