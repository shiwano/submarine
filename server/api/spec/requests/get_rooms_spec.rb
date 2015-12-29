require 'rails_helper'

RSpec.describe 'GetRooms', type: :request do
  describe 'POST /get_rooms', with_login: true do
    context 'with a valid request' do
      before do
        create(:room)
        create(:room)
        create(:room, :full)
        post(get_rooms_path)
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
